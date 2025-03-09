package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"cryptomasters.com/go/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency code should be 3")
	}

	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		jsonErr := json.Unmarshal(bodyBytes, &response)
		if jsonErr != nil {
			return nil, jsonErr
		}
	} else {
		return nil, fmt.Errorf("status code recieved: %v", res.StatusCode)
	}
	floatPrice, err := strconv.ParseFloat(response.Last, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting string price to float: %v", err)
	}
	rate := datatypes.Rate{Currency: currency, Price: floatPrice}
	return &rate, nil
}
