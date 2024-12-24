package notification

import (
	"encoding/json"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type NotificationProcessEnd struct {
	PID string
}

func (n NotificationProcessEnd) JSON() string {
	bytes, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (n NotificationProcessEnd) Name() string {
	return n.PID
}

func (n NotificationProcessEnd) Type() subject.SubjectType {
	return subject.ProcessEnd
}

func (n NotificationProcessEnd) Hash(hash hash.Hash) string {
	return hash(n.JSON())
}

func NotificationProcessEndFromJSON(jsoned string) (NotificationProcessEnd, error) {
	var end NotificationProcessEnd
	err := json.Unmarshal([]byte(jsoned), &end)
	return end, err
}
