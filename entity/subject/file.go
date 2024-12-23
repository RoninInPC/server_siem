package subject

import (
	"encoding/json"
	"server_siem/hash"
	"time"
)

type File struct {
	FullName string    `bson:"filename"`
	Content  []byte    `bson:"content"`
	Size     int64     `bson:"filesize"`
	Mode     string    `bson:"mode"`
	Modified time.Time `bson:"mod_time"`
}

func (file File) JSON() string {
	bytes, err := json.Marshal(file)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (file File) Type() SubjectType {
	return FileT
}

func (file File) Name() string {
	return file.FullName
}

func (file File) Hash(hash hash.Hash) string {
	return hash(file.JSON())
}

func FileFromJSON(jsoned string) (File, error) {
	var file File
	err := json.Unmarshal([]byte(jsoned), &file)
	return file, err
}
