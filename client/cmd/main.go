package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	authData := map[string]string{
		"Username": "Пользователь11",
		"Password": "Password",
		"Email":    "example@example.ru",
	}
	jsonData, _ := json.Marshal(authData)

	resp, err := http.Post(
		"http://localhost:63000/login",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	defer resp.Body.Close()

	head := resp.Header.Get("Authorization")
	fmt.Println("Токен JWT: ", head)
}
