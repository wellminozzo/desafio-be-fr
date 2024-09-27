package freterapido

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	url := "https://sp.freterapido.com/api/v3/quote/simulate"

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter os dados para JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func handleQuote(c echo.Context) error {

	var requestData APIRequest
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados de entrada inválidos"})
	}

	resp, err := consumeAPI(requestData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao ler a resposta"})
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Debug("Erro ao deserializar resposta")
		return err
	}

	db, err = models.InitDB()
	if err != nil {
		logrus.Error("Conexão com o banco de dados não foi estabelecida>>>> ", err)
		return fmt.Errorf("conexão com o banco de dados não estabelecida")
	}

	for _, dispatcherData := range responseBody["dispatchers"].([]interface{}) {
		dispatcherMap, ok := dispatcherData.(map[string]interface{})
		if !ok {
			logrus.Error("Erro ao converter dispatcherData")
			continue
		}

		for _, offerData := range dispatcherMap["offers"].([]interface{}) {
			offerMap, ok := offerData.(map[string]interface{})
			if !ok {
				logrus.Error("Erro ao converter offerData")
				continue
			}

			carrierMap, ok := offerMap["carrier"].(map[string]interface{})
			if !ok {
				logrus.Error("Erro ao converter carrierData")
				continue
			}

			companyName, ok := carrierMap["company_name"].(string)
			if !ok {
				logrus.Error("company_name está ausente ou não é uma string")
				continue
			}

			finalPrice, ok := offerMap["final_price"].(float64)
			if !ok {
				logrus.Error("final_price está ausente ou não é um float64")
				continue
			}

			service, ok := offerMap["service"].(string)
			if !ok {
				logrus.Error("service está ausente ou não é uma string")
				continue
			}

			deadlineMap, ok := offerMap["delivery_time"].(map[string]interface{})
			if !ok || deadlineMap["estimated_date"] == nil {
				logrus.Error("delivery_time ou estimated_date está ausente")
				continue
			}

			deadline, ok := deadlineMap["estimated_date"].(string)
			if !ok {
				logrus.Error("estimated_date não é uma string")
				continue
			}

			quote := models.Quote{
				Name:     companyName,
				Price:    fmt.Sprintf("%.2f", finalPrice),
				Service:  service,
				Deadline: deadline,
			}

			if err := db.Create(&quote).Error; err != nil {
				logrus.Errorf("Erro ao salvar a cotação: %v", err)
				continue
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"response": responseBody,
	})
}

func GetMetrics(c echo.Context) error {

	lastQuotes := c.QueryParam("last_quotes")

	var results []models.Quote

	if err := db.Table("quotes").
		Order("id DESC").
		Limit(atoi(lastQuotes)).
		Find(&results).Error; err != nil {
		logrus.Errorf("Erro ao obter últimas cotações: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter dados"})
	}

	return c.JSON(http.StatusOK, results)

}

func GetMetricsByCarrier(c echo.Context) error {

	var stats []models.CompanyMetric

	if err := db.Table("quotes").
		Select("name, COUNT(*) as quotes_count, SUM(price) as total_price, AVG(price) as avg_price").
		Group("name").
		Find(&stats).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter dados"})
	}

	return c.JSON(http.StatusOK, stats)

}

func GetCheaperQuote(c echo.Context) error {
	cheaperQuote, err := models.GetCheaperQuote(db)
	if err != nil {
		logrus.Errorf("Erro ao obter cotação mais barata: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter dados"})
	}

	return c.JSON(http.StatusOK, cheaperQuote)
}

func GetExpensiveQuote(c echo.Context) error {
	expensiveQuote, err := models.GetExpensiveQuote(db)
	if err != nil {
		logrus.Errorf("Erro ao obter cotação mais cara: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao obter dados"})
	}

	return c.JSON(http.StatusOK, expensiveQuote)
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
