package subject

import (
	"encoding/json"
	"github.com/RoninInPC/gosyscalltrace"
	"server_siem/hash"
)

type Syscall struct {
	gosyscalltrace.TraceInfo
}

func (m Syscall) JSON() string {
	bytes, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (m Syscall) Type() SubjectType {
	return SyscallT
}

func (m Syscall) Name() string {
	return m.SyscallName
}

func (m Syscall) Hash(hash hash.Hash) string {
	return hash(m.JSON())
}

func SyscallFromJSON(jsoned string) (Syscall, error) {
	var syscall Syscall
	err := json.Unmarshal([]byte(jsoned), &syscall)
	return syscall, err
}
