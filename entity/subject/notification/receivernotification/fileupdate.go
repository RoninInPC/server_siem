package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type FileUpdate struct {
	FileBefore subject.File
	FileAfter  subject.File
	BaseNotification
}

func (f FileUpdate) JSON() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (f FileUpdate) Type() subject.SubjectType {
	return FileChangeT
}

func (f FileUpdate) Name() string {
	return fmt.Sprintf("Файл %s изменён пользователем %s(%s) в процессе %s (%s).",
		f.FileBefore.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileUpdate) Hash(hash hash.Hash) string {
	return hash(f.JSON())
}
