package subject

import (
	"encoding/json"
	"server_siem/hash"
)

type User struct {
	Uid        string `bson:"uid"`
	Gid        string `bson:"gid"`
	Username   string `bson:"username"`
	SimpleName string `bson:"simple_name"`
	HomeDir    string `bson:"home_dir"`
}

func (user User) JSON() string {
	bytes, err := json.Marshal(user)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (user User) Type() SubjectType {
	return UserT
}

func (user User) Name() string {
	return user.Username
}

func (user User) Hash(hash hash.Hash) string {
	return hash(user.JSON())
}

func UserFromJSON(jsoned string) (User, error) {
	var user User
	err := json.Unmarshal([]byte(jsoned), &user)
	return user, err
}
