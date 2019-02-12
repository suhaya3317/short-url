package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RequestJson struct {
	LongDynamicLink string `json:"longDynamicLink"`
	Suffix          `json:"suffix"`
}

type Suffix struct {
	Option string `json:"option"`
}

type ResponseJson struct {
	ShortLink   string
	PreviewLink string
}

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		panic(err)
	}

	url := "https://firebasedynamiclinks.googleapis.com/v1/shortLinks?key=" + os.Getenv("FIREBASE_KEY")

	longLink := os.Args[1]
	requestJson := &RequestJson{os.Getenv("DOMAIN") + "/?link=" + longLink, Suffix{"SHORT"}}

	requestJsonByte, err := json.Marshal(requestJson)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestJsonByte))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var responseJson ResponseJson

	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(responseJson.ShortLink)
}
