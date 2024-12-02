package samanpay

import (
	"html/template"
	"io"
)

type TokenResponse struct {
	Status    int    `json:"status"`
	Token     string `json:"token"`
	ErrCode   string `json:"errCode,omitempty"`
	ErrorDesc string `json:"errorDesc,omitempty"`
}

func (t *TokenResponse) HttpWriteSampleRedirectForm(rw io.Writer) error {
	data := struct {
		PaymentUrl   string
		PaymentToken string
	}{
		PaymentUrl:   paymentURL,
		PaymentToken: t.Token,
	}

	tmp, err := template.New("redirect").Parse(redirectionFormTemplate)
	if err != nil {
		return err
	}

	err = tmp.Execute(rw, data)
	if err != nil {
		return err
	}

	return nil
}
