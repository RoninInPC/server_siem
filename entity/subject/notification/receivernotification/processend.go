package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type ProcessEnd struct {
	Process subject.Process
	BaseNotification
}

func (p ProcessEnd) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessEnd) Type() subject.SubjectType {
	return ProcessEndT
}

func (p ProcessEnd) Name() string {
	return fmt.Sprintf("Процесс %s завершился.",
		p.Process.Name())
}

func (p ProcessEnd) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
