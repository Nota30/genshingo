package genshin

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func makeOSRequest(uid string, endpoint string, locale string, ds string, headers http.Header) (User, error) {
	client := &http.Client{}

	site := endpoint + "index" + "?role_id=" + uid + "&server=" + locale

	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
        return User{}, errors.New("got error " + err.Error())
    }

	req.Header = headers

	response, err := client.Do(req)
	if err != nil {
        return User{}, errors.New("got error " + err.Error())
    }
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		// bodyBytes, err := io.ReadAll(response.Body)
		// if err != nil {
		// 	return "", err
		// }
    	// bodyString := string(bodyBytes)
		// return bodyString,  nil
		var user User
		data, _ := io.ReadAll(response.Body)
		json.Unmarshal(data, &user)
		return user, nil
	}

	return User{}, errors.New("request failed")
}