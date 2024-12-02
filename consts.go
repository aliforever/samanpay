package samanpay

import "errors"

const (
	defaultPrefix           = "https://sep.shaparak.ir"
	tokenEndpoint           = "/onlinepg/onlinepg"
	verificationEndpoint    = "/verifyTxnRandomSessionkey/ipg/VerifyTransaction"
	paymentURL              = "https://sep.shaparak.ir/OnlinePG/OnlinePG"
	redirectionFormTemplate = `<!doctypehtml><html lang="en"><meta charset="UTF-8"><title>Pay...</title><style>.text-center{text-align:center}.mt-2{margin-top:2em}.spinner{margin:100px auto 0;width:70px;text-align:center}.spinner>div{width:18px;height:18px;background-color:#333;border-radius:100%;display:inline-block;-webkit-animation:sk-bouncedelay 1.4s infinite ease-in-out both;animation:sk-bouncedelay 1.4s infinite ease-in-out both}.spinner .bounce1{-webkit-animation-delay:-.32s;animation-delay:-.32s}.spinner .bounce2{-webkit-animation-delay:-.16s;animation-delay:-.16s}@-webkit-keyframes sk-bouncedelay{0%,100%,80%{-webkit-transform:scale(0)}40%{-webkit-transform:scale(1)}}@keyframes sk-bouncedelay{0%,100%,80%{-webkit-transform:scale(0);transform:scale(0)}40%{-webkit-transform:scale(1);transform:scale(1)}}</style><body onload="submitForm()"><div class="spinner"><div class="bounce1"></div><div class="bounce2"></div><div class="bounce3"></div></div><form action="{{.PaymentUrl}}" class="mt-2 text-center" method="POST"><p>در حال انتقال به درگاه پرداخت<p>در صورتیکه بعد از<span id="countdown">5</span>ثانیه... وارد درگاه پرداخت نشدید کلیک کنید</p><input name="token" type="hidden" value="{{.PaymentToken}}"> <input name="language" type="hidden" value="fa"><button type="submit">ورود به درگاه پرداخت</button></form><script>var seconds=5;function submitForm(){document.forms[0].submit()}function countdown(){(seconds-=1)<=0?submitForm():(document.getElementById("countdown").innerHTML=seconds,window.setTimeout("countdown()",5e3))}countdown()</script>`
)

var (
	MissingParams error = errors.New("missing_parameters")
	MissingRefNum error = errors.New("missing_ref_num")
)
