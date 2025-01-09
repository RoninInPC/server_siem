package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type PortDelete struct {
	Port subject.PortTables
	BaseNotification
}

func (p PortDelete) JSON() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (p PortDelete) Type() subject.SubjectType {
	return PortDeleteT
}

func (p PortDelete) Name() string {
	return fmt.Sprintf("Порт %s удалён %s(%s) в процессе %s (%s).",
		p.Port.Name(),
		p.Who.Username, p.Who.Uid,
		p.WhoProcess.PID, p.WhoProcess.NameProcess)
}

func (p PortDelete) Hash(hash hash.Hash) string {
	return hash(p.JSON())
}
