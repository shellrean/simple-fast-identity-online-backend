package dto

type RegisterUser struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	DeviceId  string `json:"device_id"`
}
