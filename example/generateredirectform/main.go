package main

import (
	"bytes"

	"github.com/aliforever/samanpay"
)

func main() {
	tr := samanpay.TokenResponse{
		Status: 1,
		Token:  "",
	}

	buf := &bytes.Buffer{}

	err := tr.HttpWriteSampleRedirectForm(buf)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}
