package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server_siem/entity/subject"
	"server_siem/hostinfo"
	"server_siem/mapper"
	"server_siem/storagepids"
	"server_siem/storageservers"
	"server_siem/storagesubject"
	"server_siem/token"
	"time"
)

type Action interface {
	Action(*gin.Context)
}

type Post struct {
	PIDs     storagepids.StoragePIDs
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
				a.Subjects.ClearDatabase(h.HostName)
				g.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				break
			case "init_receiver":
				token := token.MakeToken(&h)
				h.CodeName = m.Message
				a.Servers.Add(h, storageservers.Receiver)
				g.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				break
			case "init":
				if _, compare := a.Servers.Compare(h); compare {
					a.Subjects.Add(h.HostName, mapper.JSONtoSubject(m.Json, m.TypeSubject))
				}
				break
			case "new":
				sub := mapper.JSONtoSubject(m.Json, m.TypeSubject)
				if _, compare := a.Servers.Compare(h); compare {
					if h.Type() == subject.ProcessT {
						pr := sub.(subject.Process)
						a.PIDs.AddTemporalPID(m.HostName, pr.PID, time.Minute*10)
					}
					a.Subjects.Add(h.HostName, sub)
				}
				break
			case "syscall":
				if _, compare := a.Servers.Compare(h); compare {
				}
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
	h := MessageToHostInfo(m)
	for _, ip := range h.IPs {
		if g.RemoteIP() == ip {
			switch m.TypeMessage {
			case "update":
				if _, compare := u.Servers.Compare(h); compare {
					u.Subjects.Update(h.HostName, mapper.JSONtoSubject(m.Json, m.TypeSubject))
				}
				break
			}
		}
	}
}

type Delete struct {
	Post
}

func (u Delete) Action(g *gin.Context) {
	m := ContextToMessage(g)
	h := MessageToHostInfo(m)
	for _, ip := range h.IPs {
		if g.RemoteIP() == ip {
			switch m.TypeMessage {
			case "delete":
				sub := mapper.JSONtoSubject(m.Json, m.TypeSubject)
				if _, compare := u.Servers.Compare(h); compare {
					if h.Type() == subject.ProcessT {
						pr := sub.(subject.Process)
						u.PIDs.DeletePID(m.HostName, pr.PID)
					}
					u.Subjects.Delete(h.HostName, mapper.JSONtoSubject(m.Json, m.TypeSubject))
				}
				break
			}
		}
	}
}

type PostCommand struct {
	Post
}

func (p PostCommand) Action(g *gin.Context) {
	m := ContextToMessage(g)
	h := MessageToHostInfo(m)
	for _, ip := range h.IPs {
		if g.RemoteIP() == ip {
			switch m.TypeMessage {
			case "new_pid":
				break
			case "delete_pid":
				break
			case "new_temporal_pid":
				break
			}
		}
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
