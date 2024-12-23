package subject

import (
	"server_siem/hash"
)

type Nope struct {
}

func (n Nope) JSON() string {
	return ""
}

func (n Nope) Type() SubjectType {
	return NopeT
}

func (n Nope) Name() string {
	return ""
}

func (n Nope) Hash(hash hash.Hash) string {
	return ""
}
