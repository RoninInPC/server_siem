package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type ProcessDelete struct {
	Process subject.Process
	BaseNotification
}

func (p ProcessDelete) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessDelete) Type() subject.SubjectType {
	return ProcessDeleteT
}

func (p ProcessDelete) Name() string {
	return fmt.Sprintf("Процесс %s удалён %s(%s) в процессе %s (%s).",
		p.Process.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p ProcessDelete) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
