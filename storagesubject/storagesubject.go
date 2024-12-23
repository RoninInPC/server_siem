package storagesubject

import "server_siem/entity/subject"

type StorageSubjects interface {
	Add(subject.Subject) bool
	Update(subject.Subject) bool
	Get(subject.Subject) subject.Subject
	Delete(subject.Subject) bool
}
