package analysis

import "server_siem/entity/subject"

type AnalysisInput struct {
	HostName string
	PID      string
	Subject  subject.Subject
}

type AnalysisService struct {
	channel chan AnalysisInput
}
