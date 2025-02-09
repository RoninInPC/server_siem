package receivernotification

import (
	"encoding/json"
	"fmt"
	"server_siem/entity/subject"
	"server_siem/hash"
)

type FileNew struct {
	File subject.File
	BaseNotification
}

func (f FileNew) JSON() string {
	bytes, err := json.Marshal(f)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (f FileNew) Type() subject.SubjectType {
	return FileNewT
}

func (f FileNew) Name() string {
	return fmt.Sprintf("Файл %s создан %s(%s) в процессе %s (%s).",
		f.File.FullName,
		f.Who.Username, f.Who.Uid,
		f.WhoProcess.PID, f.WhoProcess.NameProcess)
}

func (f FileNew) Hash(hash hash.Hash) string {
	return hash(f.JSON())
}
