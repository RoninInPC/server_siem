package sender

import (
	"server_siem/entity/subject"
)

type JWTSender struct {
	HostServer string
	methods    map[string]CommandJWT
}

func InitJWTSender(hostSubject string) *JWTSender {
	return &JWTSender{HostServer: hostSubject,
		methods: map[string]CommandJWT{
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
