package types

type CaptchaAuthModel struct {
	SessionID        string `json:"sessionid"`
	UserName         string `json:"username"`
	CaptchaFailed    bool   `json:"captcha_failed"`
	LoginSucceeded   bool   `json:"login_succeeded"`
	CredentialsError bool   `json:"credentials_error"`
}
