package App

import (
	"errors"
	"net/http"
)

type Init struct {
  cookie string
  server string `default:"cn"`
}

func (i *Init) App() (string, error) {
	if i.server != "os" && i.server != "cn" {
		return "", errors.New("not a valid server. valid options are `os` and `cn`")
	}

	return " ", nil
}

func getEndpoint(server string) string {
	if server == "os" {
		return "https://api-os-takumi.mihoyo.com/game_record/genshin/api"
	} else {
		return "https://api-takumi.mihoyo.com/game_record/app/genshin/api/"
	}
}

func getServer(id string) (string, error) {
	switch string(id[0]) {
	case "1":
		return "cn_gf01", nil
	case "2":
		return "cn_gf01", nil
	case "5":
		return "cn_qd01", nil
	case "6":
		return "os_usa", nil
	case "7":
		return "os_euro", nil
	case "8":
      return "os_asia", nil
    case "9":
      return "os_cht", nil
	default:
		return "", errors.New("invalid user id")
	}
}

func makeRequest(cookie string, url string, server string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
        return errors.New("got error " + err.Error())
    }

	// One more header to be added later
	req.Header.Set("Origin", "https://webstatic-sea.hoyolab.com")
	req.Header.Set("Referer", "https://webstatic-sea.hoyolab.com/")
	req.Header.Set("Accept", "application/json, text/plain, */*")
    req.Header.Set("Accept-Encoding", "gzip, deflate")
    req.Header.Set("Accept-Language", "zh-CN,en-US;q=0.8")
    req.Header.Set("x-rpc-language", "en-us")
    req.Header.Set("x-rpc-app_version", "1.5.0")
    req.Header.Set("x-rpc-client_type", "5")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
    req.Header.Set("Cookie", cookie)

	response, err := client.Do(req)
    if err != nil {
        return errors.New("got error " + err.Error())
    }
    defer response.Body.Close()

	return nil
}
