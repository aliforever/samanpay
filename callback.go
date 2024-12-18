package samanpay

var (
	paymentFailedTranslations = map[string]string{
		"CanceledByUser":             "کاربر انصراف داده است",
		"Failed":                     "پرداخت انجام نشد.",
		"SessionIsNull":              " کاربر در بازه زمانی تعیین شده پاسخی ارسال نکرده است.",
		"InvalidParameters":          "پارامترهای ارسالی نامعتبر است.",
		"MerchantIpAddressIsInvalid": "آدرس سرور پذیرنده نامعتبر است )در پرداخت های بر پایه توکن(",
		"TokenNotFound":              "توکن ارسال شده یافت نشد.",
		"TokenRequired":              "با این شماره ترمینال فقط تراکنش های توکنی قابل پرداخت هستند.",
		"TerminalNotFound":           "شماره ترمینال ارسال شده یافت نشد.",
		"MultisettlePolicyErrors":    "محدودیت های مدل چند حسابی رعایت نشده",
	}
)

type Callback struct {
	MID              string `json:"MID"`
	State            string `json:"State"`
	Status           string `json:"Status"`
	RRN              string `json:"RRN"`
	RefNum           string `json:"RefNum"`
	ResNum           string `json:"ResNum"`
	TerminalID       string `json:"TerminalId"`
	TraceNo          string `json:"TraceNo"`
	Amount           int64  `json:"Amount"`
	Wage             string `json:"Wage"`
	SecurePen        string `json:"SecurePen"`
	HashedCardNumber string `json:"HashedCardNumber"`
	Token            string `json:"Token"`
}

func (c *Callback) PaymentSuccessful() bool {
	return c.State == "OK"
}

func (c *Callback) PaymentFailedTranslation() string {
	if translation, ok := paymentFailedTranslations[c.State]; ok {
		return translation
	}

	return c.State
}
