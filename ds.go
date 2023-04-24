package genshin

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func getDS(server string) (string, error) {
	switch server {
	case "os":
		return osDS("6s25p5ox5y14umn1p61aqyyvbvvl3lrt")
	default:
		return osDS("6s25p5ox5y14umn1p61aqyyvbvvl3lrt")
	}
}

func osDS(salt string) (string, error) {
	date := time.Now().UTC().UnixNano() / 1e6
	timeNow := math.Floor(float64(date) / 1000)
    if timeNow >= math.MaxInt64 || timeNow <= math.MinInt64 {
        return "", errors.New("64 is out of int64 range.")
    }
    t := int64(timeNow)
	random := randomString(6)

	hash := md5.Sum([]byte("salt=" + salt + "&t=" + strconv.FormatInt(t, 10) + "&r=" + random))

	return strconv.FormatInt(t, 10) + "," + random + "," + hex.EncodeToString(hash[:]), nil
}

func randomString(n int) string {
	rand.Seed(time.Now().Unix())
	runes := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)

    for i := range b {
        b[i] = runes[rand.Intn(len(runes))]
    }
	
    return string(b)
}