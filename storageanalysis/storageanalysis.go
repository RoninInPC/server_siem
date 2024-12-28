package storageanalysis

import "server_siem/entity/subject"

type AddAnalysis struct {
	HostName string
	PID      string
	Subject  subject.Subject
}

type StorageAnalysis interface {
	Add(string, string, AddAnalysis) bool
	Get(string, string) AddAnalysis
	Del(string, string) bool
}
