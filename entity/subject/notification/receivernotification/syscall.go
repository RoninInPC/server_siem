package receivernotification

import "server_siem/entity/subject"

type Syscall struct {
	subject.Syscall
	BaseNotification
}

func (s Syscall) GetProcessPID() string {
	return s.PID
}

func (s Syscall) GetUsername() string {
	return s.Username
}
