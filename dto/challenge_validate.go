package dto

type ChallengeValidate struct {
	Id       string `json:"id"`
	Sign     string `json:"sign"`
	DeviceId string `json:"device_id"`
}
