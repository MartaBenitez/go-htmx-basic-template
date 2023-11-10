package pages

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"htmx.try/m/v2/pkg/dbconn"
	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
	"htmx.try/m/v2/pkg/domain/results"
)

var conn = *dbconn.NewInMemoryDB()
var template = "index.html"

// var url = "http://manuelsanchez.sisnet360.com:8082/"
// var url = "http://127.0.0.1:8080/sisnet/api/v1/linea-negocio"
var url = "http://santiagobricio.sisnet360.com:8080/sisnet/api/v1/linea-negocio-mock"
var urlSantiago = "http://santiagobricio.sisnet360.com:42069/business_line"
var urlMarta = "http://martabenitez.sisnet360.com:8080"
var endpoint = "/sections_to_edit"
var endpointBusisnessLine = "/bussisnes-line"
var endpointBase = "/base"

//var endpoint = "/base"

func Index(c echo.Context) error {
	user := c.QueryParam("user")
	conversaciones := GetConversations(user)

	return c.Render(200, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
	})
}

func StartNewConversation(c echo.Context) error {
	user := c.QueryParam("user")
	conn.DeleteData(user)
	conn.DeleteBases(user)
	conversaciones := GetConversations(user)
	return c.Render(200, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
		Id:            "",
	})
}

func AddMessage(c echo.Context) error {
	user := c.FormValue("user")
	question := c.FormValue("question")
	module := c.FormValue("module")
	time := time.Now().Format("15:04:05")

	quest := domain.NewMessage(question, time)
	answ := domain.NewMessage("", "")
	conversacion := domain.NewConversation(quest, answ, false, "invisible", true)
	var conversaciones = GetConversations(user)
	if len(conversaciones) > 0 {
		indiceUltimo := len(conversaciones) - 1
		conversaciones[indiceUltimo].IsLast = false
	}

	conversaciones = append(conversaciones, conversacion)
	var cosas = GetFullConn(user)

	cosas.Conversations = conversaciones
	conn.SetData(user, cosas)

	go generateMessage(user, module)

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
		Id:            "",
	})
}

func CloseActions(c echo.Context) error {
	user := c.FormValue("user")
	conversaciones := GetConversations(user)
	indice := len(conversaciones) - 1
	conversaciones[indice].Actions = "invisible"

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
		Id:            "",
	})
}

func AcceptPrompt(c echo.Context) error {
	user := c.FormValue("user")

	base := GetLastBase(user)
	if base == nil {
		log.Println("No hay respuesta")
		return nil
	}
	mensajeServidor := GetLastConversation(user)
	if mensajeServidor == nil {
		log.Println("No hay mensaje del servidor")
		return nil
	}

	lineaNegocio, err := getBusinessLine(base.Business_line_data, base.SectionsToEdit, "business_line")

	if err != nil {
		log.Println(err)
		return nil
	}

	id := loadBussinessLine(*lineaNegocio)
	conversaciones := GetConversations(user)

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
		Id:            id,
	})
}

type RespRed struct {
	_id   string
	Texto string
}

var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func normalize(str string) (string, error) {
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), err
}

type DataIn struct {
	_id                primitive.ObjectID
	Business_line_data dto.BusinessLineData
}

func loadBussinessLine(base dto.BusinessLineResp) string {
	var introducir = DataIn{_id: primitive.NewObjectID(), Business_line_data: base.Result}
	val, err := SaveJSONData(NewMongoDB(), "SISnetAI", "Producto", introducir)
	if err != nil {
		return introducir._id.Hex()
	}
	return val
}

func SaveJSONData(client *mongo.Client, databaseName string, collectionName string, data DataIn) (string, error) {
	// Get a handle for your collection
	collection := client.Database(databaseName).Collection(collectionName)

	// Insert a single document
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	val, err := collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	id := strings.Split(fmt.Sprintf("%v", val.InsertedID), "\"")[1]
	return id, nil
}

func generateMessage(user string, module string) {
	time.Sleep(1 * time.Second)
	var conversaciones = GetConversations(user)

	for pos, val := range conversaciones {
		if !val.IsAnswered {
			resp := requestAnswer(conversaciones[pos].Question, user, module)
			var response string

			if resp == nil {
				response = "Ha ocurrido un error"
			} else {
				response = *resp
				conversaciones[pos].Actions = "visible"
			}

			conversaciones[pos].Answer = domain.Message{Text: response, Time: time.Now().Format("15:04:05")}
			conversaciones[pos].IsAnswered = true
			var updatedConv = GetFullConn(user)
			updatedConv.Conversations = conversaciones
			conn.SetData(user, updatedConv)
			return
		}

	}
}

func requestAnswer(message domain.Message, user string, module string) *string {

	/*if !checkStatus() {
		err := "Server disconnected"
		log.Println(err)
		return nil
	}*/

	base, errBase := getBase(message.Text)

	sections, errSections := getSections(message.Text, module)

	if errBase != nil || errSections != nil {
		fmt.Println(errSections)
		fmt.Println(errBase)
		return nil
	}

	producto := base.Result.Business_line_data.Business_line.Producto
	var secciones []string
	for _, value := range sections.Body.MAPISECTOEDIT.Sections {
		secciones = append(secciones, value.Section_code)
	}

	mensaje := fmt.Sprintf("Si te he entendido correctamente, quieres que realice cambios sobre la linea de negocio %s, sobre las siguientes secciones:\n -%v", producto, secciones)

	//Guardamos respuesta en memoria

	basetosave := results.NewBaseToSave(sections.Body.MAPISECTOEDIT, base.Result.Business_line_data, base.Result.Coverage_data)

	conn.SetBase(user, basetosave)
	return &mensaje
}

/*func checkStatus() bool {
	res, err := http.Get(url + "health_check")
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}*/

func getBase(message string) (*dto.Base, error) {
	/*---- REVISAD LAS PETICIONES QUE SE ESTÁN FORMANDO ----
	  ---- Y LUEGO HAY QUE ENGANCHARLAS CON LAS API QUE ESTÁ HACIENDO RAFA, SANTI, Y LA QUE MAÑANA VA A EMPEZAR ADRI ------*/

	var messajeProcessed = "{ \"query\":" + message + "}"

	log.Println("0")

	body, err := json.Marshal(messajeProcessed)
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return nil, err
	}

	res, err := http.Post("http://santiagobricio.sisnet360.com:42069"+endpointBase, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return nil, err
	}

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
			return nil, err
		}

		var response dto.Base
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		return &response, nil
	}

	error := errors.New("Error, POST request to " + url + endpoint + ": response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error
}

func getSections(message string, module string) (*dto.SectionsToEdit, error) {
	var messageProcessed = `{"query":"` + message + `","module_query": "` + module + `"}`

	res, err := http.Post("http://santiagobricio.sisnet360.com:42069/sections_to_edit", "application/json", bytes.NewBuffer([]byte(messageProcessed)))
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return nil, err
	}

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
			return nil, err
		}
		var response dto.SectionsToEdit
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		return &response, nil
	}

	error := errors.New("Error, POST request to " + url + "/sections_to_edit: response received with status code " + res.Status)
	return nil, error
}

func getBusinessLine(base dto.BusinessLineData, section dto.Sections, module_query string) (*dto.BusinessLineResp, error) {

	// Create a HTTP post request
	fmt.Println()
	vara, err := json.Marshal(base)
	if err != nil {
		log.Println("No ha podido hacer el marshal bien")
	}
	var body = "{ \"module_query\":" + module_query + ",\"sections_to_edit\":" + section.ToStringGuay() + ",\"base\":" + string(vara) + "}"
	if err != nil {
		print(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", urlSantiago, bytes.NewReader([]byte(body)))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("impossible to send request: %s", err)
	}

	defer res.Body.Close()

	//res, err := http.Post(urlMarta, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		log.Println("Impossible to build bussisnes line " + err.Error())
		return nil, err
	}

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
		}
		var response dto.BusinessLineResp
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		return &response, nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Impossible to read all body of response " + err.Error())
	}

	error := errors.New("Error, POST request to " + url + endpointBusisnessLine + ": response received with status code " + res.Status + "respuesta" + string(resBody))
	return nil, error
}