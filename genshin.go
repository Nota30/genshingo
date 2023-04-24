package genshin

import "errors"

type Init struct {
  Cookie string
  Server string `default:"cn"`
}

func (init Init) GetUserData(uid string) (string, error) {
	if init.Server != "os" && init.Server != "cn" {
		return "", errors.New("not a valid server. valid options are `os` and `cn`")
	}

	locale, err := getServerLocale(uid)
	if err != nil {
		return "", err
	}

	ds, err := getDS(init.Server)
	if err != nil {
		return "", err
	}

	headers := getHeaders(init.Server, ds, init.Cookie)

	endpoint := getApiEndpoint(init.Server)

	req, err := makeOSRequest(uid, endpoint,locale, ds, headers)
	if err != nil {
		return "", err
	}

	return req, nil
}