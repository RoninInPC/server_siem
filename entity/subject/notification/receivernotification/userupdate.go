package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type UserUpdate struct {
	UserBefore subject.User
	UserAfter  subject.User
	BaseNotification
}

func (u UserUpdate) JSON() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (u UserUpdate) Type() subject.SubjectType {
	return UserUpdateT
}

func (u UserUpdate) Name() string {
	return fmt.Sprintf("Пользователь %s обновлён %s(%s) в процессе %s (%s).",
		u.UserBefore.Username,
		u.Who.Username, u.Who.Uid,
		u.WhoProcess.PID, u.WhoProcess.NameProcess)
}

func (u UserUpdate) Hash(hash hash.Hash) string {
	return hash(u.JSON())
}
