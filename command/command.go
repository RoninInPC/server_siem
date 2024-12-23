package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server_siem/entity/subject"
	"server_siem/hostinfo"
	"server_siem/storageservers"
	"server_siem/storagesubject"
	"server_siem/token"
)

type Action interface {
	Action(*gin.Context)
}

type Post struct {
	Servers  storageservers.StorageServers
	Subjects storagesubject.StorageSubjects
}

func (a Post) Action(g *gin.Context) {
	m := ContextToMessage(g)

	h := MessageToHostInfo(m)
	for _, ip := range h.IPs {
		if g.RemoteIP() == ip {
			switch m.TypeMessage {
			case "init_server":
				token := token.MakeToken(&h)
				a.Servers.Add(h, storageservers.Server)
				g.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				break
			case "init_receiver":
				token := token.MakeToken(&h)
				a.Servers.Add(h, storageservers.Receiver)
				g.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				break
			case "init":
				a.Servers.Exists(h)
				break
			case "new":
				break
			}
		}
	}
}

type Update struct {
	Post
}

func (u Update) Action(g *gin.Context) {
	m := ContextToMessage(g)
	switch m.TypeMessage {
	case "update":
		break
	}
}

type Delete struct {
	Post
}

func (u Delete) Action(g *gin.Context) {
	m := ContextToMessage(g)
	switch m.TypeMessage {
	case "delete":
		break
	}
}

func ContextToMessage(g *gin.Context) subject.Message {
	j := g.Param("json")
	m := subject.Message{}
	json.Unmarshal([]byte(j), &m)
	return m
}

func MessageToHostInfo(s subject.Message) hostinfo.HostInfo {
	return hostinfo.HostInfo{
		HostName: s.HostName,
		HostOS:   s.SystemOS,
		IPs:      s.HostIP,
	}
}
