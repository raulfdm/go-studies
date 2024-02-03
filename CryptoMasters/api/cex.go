package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"studies.com/cryptomasters/data"
)

const apiURL = "https://trade.cex.io/api/rest-public/get_ticker"

func GetRates(currency string) (*data.Rate, error) {
	upperCurrency := strings.ToUpper(currency)

	postBodyStr := fmt.Sprintf(`{"pairs":["BTC-%v"]}`, upperCurrency)

	postBodyAsBytes := []byte(postBodyStr)

	res, err := http.Post(apiURL, "application/json", bytes.NewBuffer(postBodyAsBytes))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		apiResponse := CEXResponse{}

		err = json.Unmarshal(bodyBytes, &apiResponse)

		if err != nil {
			return nil, err
		} else {

			priceAsFloat, err := strconv.ParseFloat(apiResponse.Data["BTC-"+upperCurrency].Last, 64)

			if err != nil {
				return nil, err
			}

			rate := data.Rate{
				Currency: upperCurrency,
				Price:    priceAsFloat,
			}

			return &rate, nil
		}

	} else {
		return nil, fmt.Errorf("API returned status code %v", res.StatusCode)
	}

}
