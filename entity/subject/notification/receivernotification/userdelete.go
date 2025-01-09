package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type UserDelete struct {
	User subject.User
	BaseNotification
}

func (u UserDelete) JSON() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (u UserDelete) Type() subject.SubjectType {
	return UserDeleteT
}

func (u UserDelete) Name() string {
	return fmt.Sprintf("Пользователь %s удалён %s(%s) в процессе %s (%s).",
		u.User.Username,
		u.Who.Username, u.Who.Uid,
		u.WhoProcess.PID, u.WhoProcess.NameProcess)
}

func (u UserDelete) Hash(hash hash.Hash) string {
	return hash(u.JSON())
}
