package verify

type GEtheader struct {
	Jwt string `json:"jwt"`
	OTP string `json:"otp"`
}
type DATA struct {
	Username, Tag, UserId, Time, Email string

	Subdata struct {
		Password string
	}
}
