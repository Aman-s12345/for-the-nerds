package stringQues

import (
	"net/url"
	"strings"
)

func ParseURLParams(inputURL string) map[string]string {



	urlParams := make(map[string]string)

	parts := strings.SplitN(inputURL, "?", 2) // [arr1,arr2]

	// No query params
	if len(parts) < 2 {
		return urlParams
	}

	paramsURL := parts[1]

	tokenParams := strings.Split(paramsURL, "&")

	for _, param := range tokenParams {

		keyVal := strings.SplitN(param, "=", 2)

		// Invalid param
		if len(keyVal) < 2 {
			continue
		}

		key := keyVal[0]
		value := keyVal[1]

		decodedValue, err := url.QueryUnescape(value)
		if err != nil {
			continue
		}

		urlParams[key] = decodedValue
	}

	return urlParams



}
