package mapper

import (
	"server_siem/entity/subject"
	"server_siem/entity/subject/notification"
)

func JSONtoSubject(json string, subjectType subject.SubjectType) subject.Subject {
	switch subjectType {
	case subject.FileT:
		sub, _ := subject.FileFromJSON(json)
		return sub
	case subject.ProcessT:
		sub, _ := subject.ProcessFromJSON(json)
		return sub
	case subject.PortTablesT:
		sub, _ := subject.PortTablesFromJSON(json)
		return sub
	case subject.UserT:
		sub, _ := subject.UserFromJSON(json)
		return sub
	case subject.SyscallT:
		sub, _ := subject.SyscallFromJSON(json)
		return sub
	case subject.ProcessEnd:
		sub, _ := notification.NotificationProcessEndFromJSON(json)
		return sub
	default:
		return subject.Nope{}
	}
}
