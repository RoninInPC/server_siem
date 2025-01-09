package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type FileDelete struct {
	File subject.File
	BaseNotification
}

func (f FileDelete) JSON() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (f FileDelete) Type() subject.SubjectType {
	return FileDeleteT
}

func (f FileDelete) Name() string {
	return fmt.Sprintf("Файл %s удалён %s(%s) в процессе %s (%s).",
		f.File.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileDelete) Hash(hash hash.Hash) string {
	return hash(f.JSON())
}
