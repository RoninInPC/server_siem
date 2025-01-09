package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type ProcessNew struct {
	Process subject.Process
	BaseNotification
}

func (p ProcessNew) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p ProcessNew) Type() subject.SubjectType {
	return ProcessNewT
}

func (p ProcessNew) Name() string {
	return fmt.Sprintf("Процесс %s создан %s(%s) в процессе %s (%s).",
		p.Process.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p ProcessNew) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
