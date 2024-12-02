package samanpay

type VerifyResponse struct {
	ResultCode        int    `json:"ResultCode"`
	ResultDescription string `json:"ResultDescription"`
	Success           bool

	TransactionDetail *VerifyResponseTransactionDetail `json:"TransactionDetail"`
}

type VerifyResponseTransactionDetail struct {
	RRN             string `json:"RRN"`
	RefNum          string `json:"RefNum"`
	MaskedPan       string `json:"MaskedPan"`
	HashedPan       string `json:"HashedPan"`
	TerminalNumber  int    `json:"TerminalNumber"`
	OriginalAmount  int    `json:"OriginalAmount"`
	AffectiveAmount int    `json:"AffectiveAmount"`
	StraceDate      string `json:"StraceDate"`
	StraceNo        string `json:"StraceNo"`
}
