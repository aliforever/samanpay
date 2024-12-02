package main

import (
	"encoding/json"
	"os"

	"github.com/aliforever/samanpay"
)

func main() {
	terminalID := os.Getenv("terminal_id")
	terminalPassword := os.Getenv("terminal_password")

	client, err := samanpay.NewClientWithHttpProxy(terminalID, terminalPassword, "http://localhost:11809")
	if err != nil {
		panic(err)
	}

	result, err := client.GeneratePaymentToken(
		"1",
		20000,
		"https://pay.gocademy.ir/sep",
		"09120000000",
	)
	if err != nil {
		panic(err)
	}

	j, _ := json.Marshal(result)

	println(string(j))
}
