package api

import (
	"cryptomasters/datatypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (*datatypes.Rate, error) {
	upperCurrency := strings.ToUpper(currency)

	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("response error %v", err)
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}

