package genshin

import "errors"

func getApiEndpoint(server string) string {
	switch server {
	case "os":
		return "https://bbs-api-os.mihoyo.com/game_record/genshin/api/"
	case "cn":
		return "https://api-takumi.mihoyo.com/game_record/app/genshin/api/"
	default:
		return "https://api-takumi.mihoyo.com/game_record/app/genshin/api/"
	}
}

func getServerLocale(uid string) (string, error) {
	switch string(uid[0]) {
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
		return "", errors.New("uid is invalid")
	}
}