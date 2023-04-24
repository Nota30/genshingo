package genshin

import (
	"errors"
	"io"
	"net/http"
)

func makeOSRequest(uid string, endpoint string, locale string, ds string, headers http.Header) (string, error) {
	client := &http.Client{}

	site := endpoint + "index" + "?role_id=" + uid + "&server=" + locale

	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
        return "", errors.New("got error " + err.Error())
    }

	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
        return "", errors.New("got error " + err.Error())
    }
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
    	bodyString := string(bodyBytes)
		return bodyString,  nil
	}

	return "", errors.New("request failed")
}