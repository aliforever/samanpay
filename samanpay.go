package samanpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	terminalID       string
	terminalPassword string
	redirectUrl      string

	httpClient *http.Client

	urlPrefix string
}

func NewClient(terminalID, terminalPassword string) *Client {
	return &Client{
		terminalID:       terminalID,
		terminalPassword: terminalPassword,

		httpClient: &http.Client{},

		urlPrefix: defaultPrefix,
	}
}

func NewClientWithReverseProxy(proxyURL string, terminalID, terminalPassword string) *Client {
	return &Client{
		terminalID:       terminalID,
		terminalPassword: terminalPassword,

		httpClient: &http.Client{},

		urlPrefix: proxyURL,
	}
}

func NewClientWithHttpProxy(terminalID, terminalPassword string, proxyURL string) (*Client, error) {
	proxyUrl, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	return &Client{
		terminalID:       terminalID,
		terminalPassword: terminalPassword,

		httpClient: &http.Client{Transport: transport},

		urlPrefix: defaultPrefix,
	}, nil
}

func (c *Client) GeneratePaymentToken(
	invoiceID string,
	amount float64,
	redirectURL string,
	cellNumber string,
) (*TokenResponse, error) {
	data := map[string]interface{}{
		"action":      "token",
		"TerminalId":  c.terminalID,
		"Amount":      amount,
		"ResNum":      invoiceID,
		"RedirectUrl": redirectURL,
		"CellNumber":  cellNumber,
	}

	j, _ := json.Marshal(data)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.urlPrefix, tokenEndpoint),
		bytes.NewReader(j),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf("unexpected status code: %d - %s", resp.StatusCode, string(b))
	}

	b, _ := io.ReadAll(resp.Body)

	var tokenResponse *TokenResponse

	err = json.Unmarshal(b, &tokenResponse)
	if err != nil {
		return nil, err
	}

	if tokenResponse.Status != 1 {
		return nil, fmt.Errorf("%d - %s", tokenResponse.Status, tokenResponse.ErrorDesc)
	}

	return tokenResponse, nil
}

func (c *Client) VerifyPayment(refNum string) (*VerifyResponse, error) {
	data := map[string]interface{}{
		"TerminalNumber": c.terminalID,
		"RefNum":         refNum,
	}

	j, _ := json.Marshal(data)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.urlPrefix, verificationEndpoint),
		bytes.NewReader(j),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf("unexpected status code: %d - %s", resp.StatusCode, string(b))
	}

	b, _ := io.ReadAll(resp.Body)

	var verifyResponse *VerifyResponse

	err = json.Unmarshal(b, &verifyResponse)
	if err != nil {
		return nil, err
	}

	if !verifyResponse.Success {
		return nil, fmt.Errorf("%d - %s", verifyResponse.ResultCode, verifyResponse.ResultDescription)
	}

	return verifyResponse, nil
}

func (c *Client) HttpCallback(
	callback func(c *Callback, r *http.Request, w http.ResponseWriter),
	onError func(err error, r *http.Request, w http.ResponseWriter),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		mid := r.Form.Get("MID")
		state := r.Form.Get("State")
		status := r.Form.Get("Status")
		rrn := r.Form.Get("RRN")
		refNum := r.Form.Get("RefNum")
		resNum := r.Form.Get("ResNum")
		terminalID := r.Form.Get("TerminalId")
		traceNo := r.Form.Get("TraceNo")
		amount := r.Form.Get("Amount")
		wage := r.Form.Get("Wage")
		securePen := r.Form.Get("SecurePen")
		hashedCardNumber := r.Form.Get("HashedCardNumber")

		if mid == "" || state == "" || resNum == "" {
			onError(MissingParams, r, w)

			return
		}

		cback := &Callback{
			MID:              mid,
			State:            state,
			Status:           status,
			RRN:              rrn,
			RefNum:           refNum,
			ResNum:           resNum,
			TerminalID:       terminalID,
			TraceNo:          traceNo,
			Wage:             wage,
			SecurePen:        securePen,
			HashedCardNumber: hashedCardNumber,
		}

		if cback.PaymentSuccessful() {
			transactionAmount, err := strconv.ParseInt(amount, 10, 64)
			if err != nil {
				onError(err, r, w)

				return
			}

			cback.Amount = transactionAmount

			callback(cback, r, w)
		}
	}
}
