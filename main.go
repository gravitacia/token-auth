package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenAuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenAuthResponse struct {
	Token string `json:"token"`
}

func main() {
	email := "mail"
	password := "pass"

	authRequest := TokenAuthRequest{
		Email:    email,
		Password: password,
	}

	requestBody, err := json.Marshal(authRequest)
	if err != nil {
		fmt.Printf("Error marshalling the request: %s\n", err)
		return
	}

	resp, err := http.Post("https://lk.biacorp.ru", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("Error sending the request: %s\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Authentication successful.")
	} else {
		fmt.Printf("Authentication failed. Status code: %d\n", resp.StatusCode)
	}
}
