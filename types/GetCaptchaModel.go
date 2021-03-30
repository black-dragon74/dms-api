package types

type GetCaptchaModel struct {
	SessionID    string `json:"sessionid"`
	Generator    string `json:"generator"`
	EncodedImage string `json:"encoded_image"`
}
