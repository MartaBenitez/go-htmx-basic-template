package pages

import (
	"encoding/json"
	"fmt"
	"os"

	"htmx.try/m/v2/pkg/domain"
	"htmx.try/m/v2/pkg/domain/dto"
)

//var mongoUri = "mongodb://root:example@20.56.93.5:27017"

func GetLastBase(user string) *domain.BaseToSave {
	vals, ok := conn.GetBases(user)
	if !ok {
		return nil
	}
	val := vals[len(vals)-1]
	return &val
}

func GetLastConversation(user string) *domain.Conversation {
	vals := GetConversations(user)
	for _, val := range vals {
		if val.IsLast {
			return &val
		}
	}
	return nil
}

func GetFullConn(user string) domain.InterfaceResponseFull {
	//get the user conversations from the database
	val, ok := conn.GetData(user)
	if !ok {
		return domain.InterfaceResponseFull{}
	}
	return val
}

func GetConversations(user string) []domain.Conversation {
	val, ok := conn.GetData(user)
	if !ok {
		return []domain.Conversation{}
	}
	return val.Conversations
}

/*func NewMongoDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoUri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}*/

func RecoverExample() *dto.BusinessLineData {
	var respuesta dto.BusinessLineData
	raw, err := os.ReadFile("/home/usuario/Escritorio/ejemploCob.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(raw, &respuesta)
	return &respuesta
}

/*func AppendProps(prop dto.ResultSections) []string {
	var props []string
	if prop.BusinessLine != "" {
		props = append(props, "Business Line")
	}
	if prop.CommercialNetworkAttribute != "" {
		props = append(props, "Commercial Network Attribute")
	}
	if prop.ProductPaymentMethod != "" {
		props = append(props, "Product Payment Method")
	}
	if prop.ProductRenewalCycle != "" {
		props = append(props, "Product Renewal Cycle")
	}
	if prop.RenewalParameter != "" {
		props = append(props, "Renewal Parameter")
	}

	return props
}*/
