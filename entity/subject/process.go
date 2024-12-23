package subject

import (
	"encoding/json"
	"server_siem/hash"
	"time"
)

type Process struct {
	PID           string    `bson:"pid"`
	UID           string    `bson:"uid"`
	Nice          int32     `bson:"nice"`
	IsRunning     bool      `bson:"is_running"`
	IsBackGround  bool      `bson:"is_background"`
	CreateTime    time.Time `bson:"create_time"`
	Status        []string  `bson:"statuses"`
	NameProcess   string    `bson:"name_process"`
	CMDLine       string    `bson:"cmdline"`
	PercentCPU    float64   `bson:"percent_cpu"`
	PercentMemory float32   `bson:"percent_memory"`
}

func (process Process) JSON() string {
	bytes, err := json.Marshal(process)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (process Process) Type() SubjectType {
	return ProcessT
}

func (process Process) Name() string {
	return process.PID
}

func (process Process) Hash(hash hash.Hash) string {
	return hash(process.JSON())
}

func ProcessFromJSON(jsoned string) (Process, error) {
	var process Process
	err := json.Unmarshal([]byte(jsoned), &process)
	return process, err
}
