package subject

import "server_siem/hash"

type SubjectType int

const (
	FileT       SubjectType = 0
	ProcessT    SubjectType = 1
	PortTablesT SubjectType = 2
	UserT       SubjectType = 3
	ProcessEnd  SubjectType = 4
	MessageT    SubjectType = 5
	SyscallT    SubjectType = 6
	NopeT       SubjectType = -1
	HostT       SubjectType = 7
)

type Subject interface {
	JSON() string
	Type() SubjectType
	Name() string
	Hash(hash.Hash) string
}
