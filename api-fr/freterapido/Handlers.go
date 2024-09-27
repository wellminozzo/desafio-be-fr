package freterapido

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/wellminozzo/desafio-be-fr/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func HandlerStatusJson(c echo.Context) error {

	return c.JSON(200, "OK")
}

func consumeAPI(requestData APIRequest) (*http.Response, error) {
	// URL da API externa para obter a cotação
	url := "https://sp.freterapido.com/api/v3/quote/simulate"

	// Converte os dados da requisição para JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter os dados para JSON: %v", err)
	}

	// Cria a requisição HTTP
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Envia a requisição
	client := &http.Client{}
	return client.Do(req)
}

func handleQuote(c echo.Context) error {
	// Defina os dados obrigatórios (aqui você pode coletar os dados do corpo da requisição ou definir manualmente)
	var requestData APIRequest
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados de entrada inválidos"})
	}

	// Consome a API externa
	resp, err := consumeAPI(requestData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	// Lê a resposta da API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao ler a resposta"})
	}

	var responseBody map[string]interface{}

	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Debug("Error to unmarshal")
		return err
	}

	//log.Println("responseBody>>>>>>>>>>>>>>>>>>>>>>", responseBody)
	db, err = models.InitDB()
	if err != nil {
		logrus.Error("Conexão com o banco de dados não foi estabelecida>>>> ", err)
		return fmt.Errorf("conexão com o banco de dados não estabelecida")
	}

	for _, dispatcherData := range responseBody["dispatchers"].([]interface{}) {
		dispatcherMap := dispatcherData.(map[string]interface{})

		fmt.Println("dispatcherMap>>>>>>>>>>>>>>>>>>>>>>", dispatcherMap)

		// Cria o Dispatcher
		dispatcher := models.Dispatcher{
			RequestID:                  dispatcherMap["request_id"].(string),
			RegisteredNumberDispatcher: dispatcherMap["registered_number_dispatcher"].(string),
			RegisteredNumberShipper:    dispatcherMap["registered_number_shipper"].(string),
			ZipcodeOrigin:              int64(dispatcherMap["zipcode_origin"].(float64)),
		}

		// Salva o Dispatcher no banco de dados
		if err := db.Create(&dispatcher).Error; err != nil {
			logrus.Errorf("Erro ao salvar o Dispatcher: %v", err)
			return err
		}

		//var carrier models.Carrier

		// Processa as Offers dentro do Dispatcher
		for _, offerData := range dispatcherMap["offers"].([]interface{}) {
			offerMap := offerData.(map[string]interface{})

			fmt.Println("offerMap>>>>>>>>>>>>>>>>>>>>>>", offerMap)

			// Cria uma nova Offer
			offer := models.Offer{

				DispatcherID: dispatcher.ID, // Associa o DispatcherID
				//CarrierID:    carrier.ID,    // Não esqueça de atribuir CarrierID
				FinalPrice:   offerMap["final_price"].(float64),
				CostPrice:    offerMap["cost_price"].(float64),
				Expiration:   offerMap["expiration"].(string),
				Service:      offerMap["service"].(string),
				HomeDelivery: offerMap["home_delivery"].(bool),
				Modal:        offerMap["modal"].(string),
				//CompanyName:  offerMap["company_name"].(string),
			}

			if err := db.Create(&offer).Error; err != nil {
				logrus.Errorf("Erro ao salvar a Offer: %v", err)
				return err
			}

		}
	}

	// Retorna a resposta da API externa para o cliente
	return c.JSON(http.StatusOK, map[string]interface{}{
		"response": responseBody,
	})
}

func GetMetrics(c echo.Context) error {

	lastQuotes := c.QueryParam("last_quotes")

	var results []models.Offer
	log.Println("lastQuotes>>>>>>>>>>>>>>>>>>>>>>", lastQuotes)
	if err := db.Table("offers").
		Order("created_at DESC").
		Limit(atoi(lastQuotes)).
		Find(&results).Error; err != nil {
		logrus.Errorf("Erro ao obter últimas cotações: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter dados"})
	}

	return c.JSON(http.StatusOK, results)

}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0 // Retorna 0 se houver erro na conversão
	}
	return num
}
