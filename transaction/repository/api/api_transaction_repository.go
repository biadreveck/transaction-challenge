package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"stone/transaction-challenge/domain"

	"github.com/spf13/viper"
)

const (
	endpointName = "transactions"
)

type apiTransactionRepo struct {
	Endpoint string
}

func NewApiTransactionRepository(baseUrl string) domain.TransactionRepository {
	return &apiTransactionRepo{
		Endpoint: fmt.Sprintf("%s/%s", baseUrl, endpointName),
	}
}

func (r *apiTransactionRepo) Insert(ctx context.Context, authType string, t *domain.Transaction) (result map[string]interface{}, err error) {
	apiKey := viper.GetString("pagarmeapi.key")
	endpoint := r.Endpoint

	if authType == domain.AUTH_TYPE_URL {
		endpoint = fmt.Sprintf("%s?api_key=%s", endpoint, apiKey)
	} else if authType == domain.AUTH_TYPE_BODY {
		t.ApiKey = apiKey
	}

	t.Async = false
	// log.Printf("Request data: %+v\n", t)
	requestBody, err := json.Marshal(t)
	if err != nil {
		return
	}

	// log.Print(endpoint)
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return
	}
	request.Header.Set("Content-type", "application/json")

	if authType == domain.AUTH_TYPE_BASIC {
		request.SetBasicAuth(apiKey, "x")
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		// log.Printf("Error on API request: %s\n", err.Error())
		return
	}

	defer response.Body.Close()

	// log.Printf("Response status: %d\n", response.StatusCode)

	result = make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return
	}
	// log.Printf("Response data: %+v\n", result)

	if response.StatusCode != 200 {
		switch response.StatusCode {
		case 400:
			err = domain.ErrBadParamInput
			break
		case 401:
			err = domain.ErrUnauthorized
			break
		case 404:
			err = domain.ErrNotFound
			break
		default:
			err = domain.ErrInternal
			break
		}
	}

	return
}
