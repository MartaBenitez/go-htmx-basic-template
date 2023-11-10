package pages

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"htmx.try/m/v2/pkg/dbconn"
	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
)

var conn = *dbconn.NewInMemoryDB()
var template = "index.html"

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

	go GenerateMessage(user, module)

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

	var err error
	var id *dto.IdMongo

	//switch del modulo
	switch(base.Module){
		case "business_line":
			id, err = GetBusinessLine(base.Business_line_data, base.SectionsToEdit, "business_line")
			break;
		case "coverage":
			//id, err = GetCoverage(base.Coverage_data, base.SectionsToEdit, "coverage")
			break;
		case "technical_product":
			//id, err = GetTechnicalProduct(base.Technical_product, base.SectionsToEdit, "technical_product")
			break;
	}
	
	
	if err != nil {
		log.Println(err)
		return nil
	}

	conversaciones := GetConversations(user)

	return c.Render(http.StatusOK, template, domain.InterfaceResponseFull{
		User:          user,
		Conversations: conversaciones,
		Id:            id.Result.Lob_id,
	})
}