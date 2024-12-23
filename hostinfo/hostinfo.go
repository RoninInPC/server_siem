package hostinfo

import (
	"encoding/json"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type HostInfo struct {
	HostName string
	HostOS   string
	Token    string
	IPs      []string
}

func (h HostInfo) JSON() string {
	bytes, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (h HostInfo) Type() subject.SubjectType {
	return subject.HostT
}

func (h HostInfo) Name() string {
	return h.HostName
}

func (h HostInfo) Hash(hash hash.Hash) string {
	return hash(h.JSON())
}
