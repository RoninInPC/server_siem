package storage

import "server_siem/entity/subject"

type Storage[sub subject.Subject] interface {
	Add(sub) bool
	Update(sub) bool
	Get(sub) sub
	Delete(sub) bool
}
