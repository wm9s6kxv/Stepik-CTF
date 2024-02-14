package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetServer() {
	url := "http://172.21.0.30:8080/auth"

	creds := Credentials{
		Username: "asdfsddmin",
		Password: "fgdfgdfg",
	}

	jsonData, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("Ошибка при маршалинге JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при отправке POST запроса:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	if string(body) != "NO OK" {

		fmt.Println("Good")
	} else {
		fmt.Println("Error")
	}

}

func main() {
	for {
		GetServer()
		time.Sleep(time.Second * 10)
	}
}
