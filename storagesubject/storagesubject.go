package storagesubject

import "server_siem/entity/subject"

type StorageSubjects interface {
	ClearDatabase(host string) bool
	Add(string, subject.Subject) bool
	Update(string, subject.Subject) bool
	Get(string, subject.Subject) subject.Subject
	Delete(string, subject.Subject) bool
	GetHosts() []string
}
