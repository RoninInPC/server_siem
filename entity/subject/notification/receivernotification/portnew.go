package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type PortNew struct {
	Port subject.PortTables
	BaseNotification
}

func (p PortNew) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p PortNew) Type() subject.SubjectType {
	return PortNewT
}

func (p PortNew) Name() string {
	return fmt.Sprintf("Порт %s создан %s(%s) в процессе %s (%s).",
		p.Port.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p PortNew) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
