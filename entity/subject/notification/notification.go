package notification

import "server_siem/entity/subject"

type Notification interface {
	subject.Subject
}
