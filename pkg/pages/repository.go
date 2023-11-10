package pages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
)

type DataIn struct {
	_id                primitive.ObjectID
	Business_line_data dto.BusinessLineData
}

var urlMock = "http://santiagobricio.sisnet360.com:42069"
var urlSisnet = ""
var endpointBase = "/base"
var endpointSections = "/sections_to_edit"
var endpointBusiness = "/business_line"

func GenerateMessage(user string, module string) {
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
	base, errBase := getBase(message.Text)

	sections, errSections := getSections(message.Text, module)

	if errBase != nil || errSections != nil {
		fmt.Println(errSections)
		fmt.Println(errBase)
		return nil
	}

	producto := base.Result.Business_line_data.Business_line.Producto
	var secciones []string
	for _, value := range sections.NAPISALSECTOEDIT.Sections_to_edit.Sections{
		secciones = append(secciones, value.Section_code)
	}

	mensaje := fmt.Sprintf("Si te he entendido correctamente, quieres que realice cambios sobre la linea de negocio %s, sobre las siguientes secciones:\n -%v", producto, secciones)

	//Guardamos respuesta en memoria
	basetosave := domain.NewBaseToSave("business_line", sections.NAPISALSECTOEDIT.Sections_to_edit, base.Result.Business_line_data, base.Result.Coverage_data, base.Result.Tecnical_product_data)
	conn.SetBase(user, basetosave)
	return &mensaje
}

func getBase(message string) (*dto.Base, error) {
	var messajeProcessed = "{ \"query\":" + message + "}"

	body, err := json.Marshal(messajeProcessed)
	if err != nil {
		log.Println("Impossible to build request: " + err.Error())
		return nil, err
	}
	fmt.Println(urlMock)
	res, err := http.Post(urlMock+endpointBase, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Impossible to do request: " + err.Error())
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
		log.Println("------------Base----------")
		log.Println(response)
		return &response, nil
	}

	error := errors.New("Error, POST request to " + urlMock + endpointBase + ": response received with status code " + res.Status)
	log.Println(error.Error())
	return nil, error
}

func getSections(message string, module string) (*dto.SectionsToEdit, error) {
	var messageProcessed = `{"query":"` + message + `","module_query": "` + module + `"}`

	res, err := http.Post(urlMock+endpointSections, "application/json", bytes.NewBuffer([]byte(messageProcessed)))
	if err != nil {
		log.Println("Impossible to do request: " + err.Error())
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
		log.Println("------------Secciones----------")
		log.Println(response)
		return &response, nil
	}

	error := errors.New("Error, POST request to " + urlMock + endpointSections + ": response received with status code " + res.Status)
	return nil, error
}

/*func LoadBussinessLine(base dto.BusinessLineResp) string {
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
}*/

func GetBusinessLine(base dto.BusinessLineData, section dto.Sections, module string) (*dto.IdMongo, error) {
	vara, err := json.Marshal(base)
	if err != nil {
		log.Println("No ha podido hacer el marshal bien")
	}
	var body = "{ \"module_query\":" + module + ",\"sections_to_edit\":" + section.ToStringGuay() + ",\"base\":" + string(vara) + "}"
	if err != nil {
		print(err)
		return nil, err
	}

	res, err := http.Post(urlMock+endpointBusiness, "application/json", bytes.NewReader([]byte(body)))
	if err != nil {
		log.Println("Impossible to do request " + err.Error())
		return nil, err
	}

	if res.StatusCode == 200 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Impossible to read all body of response " + err.Error())
		}
		var response dto.IdMongo
		err = json.Unmarshal(resBody, &response)
		if err != nil {
			log.Println("Impossible to parse the response " + err.Error())
			return nil, err
		}
		log.Println("------------Linea Negocio----------")
		log.Println(response)
		return &response, nil
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Impossible to read all body of response " + err.Error())
	}

	error := errors.New("Error, POST request to " + urlMock + endpointBusiness + ": response received with status code " + res.Status + "respuesta" + string(resBody))
	return nil, error
}

/*func getCoverage(base dto.Coverage, section dto.Sections, module_query string) (*dto.CoverageResp, error) {
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
		var response dto.CoverageResp
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
}*/
