package sender

import (
	"server_siem/entity/subject"
)

type Sender interface {
	Send(address string, message subject.Message) bool
}
