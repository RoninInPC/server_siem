package subject

import "server_siem/hash"

type SubjectType int

const (
	FileT       SubjectType = 0
	ProcessT    SubjectType = 1
	PortTablesT SubjectType = 2
	UserT       SubjectType = 3
	ProcessEnd  SubjectType = 4
)

type Subject interface {
	JSON() string
	Type() SubjectType
	Name() string
	Hash(hash.Hash) string
}
