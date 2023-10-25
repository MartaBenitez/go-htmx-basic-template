package pages

import (
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
)

var conn = *dbconn.NewInMemoryDB()
var template = "index.html"
var url = "http://manuelsanchez.sisnet360.com:8082/"

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
	conn.DeleteResponses(user)
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

func GetBussinessLine(c echo.Context) error {
	user := c.FormValue("user")
	respuesta := GetLastResponse(user)
	if respuesta == nil {
		log.Println("No hay respuesta")
		return nil
	}
	mensajeServidor := GetLastConversation(user)
	if mensajeServidor == nil {
		log.Println("No hay mensaje del servidor")
		return nil
	}
	//string, _ := normalize(mensajeServidor.Answer.Text)
	id := loadBussinessLine(*respuesta)
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

func loadBussinessLine(respuesta dto.BusinessLineData) string {
	var introducir = DataIn{_id: primitive.NewObjectID(), Business_line_data: respuesta}
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

	if !checkStatus() {
		err := "Server disconnected"
		log.Println(err)
		return nil
	}

	messageNoSpaces := strings.Replace(message.Text, " ", "%20", -1)
	base, errBase := getBase(messageNoSpaces)

	sections, errSections := getSections(messageNoSpaces, module)

	if errBase != nil || errSections != nil {
		fmt.Println(errBase)
		fmt.Println(errSections)
		return nil
	}

	producto := base.Result.Business_line_data.Business_line.Producto

	props := AppendProps(sections.Result)
	mensaje := fmt.Sprintf("Si te he entendido correctamente, quieres que realice cambios sobre la linea de negocio %s, sobre las siguientes secciones:\n -%v", producto, props)
	//Guardamos respuesta en memoria

	conn.SetResponse(user, base.Result.Business_line_data)
	return &mensaje
}

func checkStatus() bool {
	res, err := http.Get(url + "health_check")
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}

func getBase(message string) (*dto.Base, error) {
	res, err := http.Get(url + "base?query=" + message)
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

	error := errors.New("Error: response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error

}

func getSections(message string, module string) (*dto.SectionsToEdit, error) {
	res, err := http.Get(url + "sections_to_edit?query=" + message + "&module=" + module)
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

	error := errors.New("Error: response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error
}
