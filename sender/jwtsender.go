package sender

import (
	"server_siem/entity/subject"
)

type JWTSender struct {
	methods map[string]CommandJWT
}

func InitJWTSender() *JWTSender {
	return &JWTSender{methods: map[string]CommandJWT{
		"send_receiver": CommandJWTPostForm{},
	}}
}

func (j *JWTSender) Send(address string, message subject.Message) bool {
	resp, err := j.methods[message.TypeMessage].Command(address, message.JSON())
	if err != nil {
		return false
	}

	return resp.StatusCode == 200
}
