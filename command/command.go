package command

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"server_siem/entity/subject"
	"server_siem/entity/subject/notification/receivernotification"
	"server_siem/hostinfo"
	"server_siem/mapper"
	"server_siem/storagepids"
	"server_siem/storageservers"
	"server_siem/storagesubject"
	"server_siem/token"
	"strings"
	"time"
)

type Action interface {
	Action(*gin.Context)
}

type Post struct {
	PIDs     storagepids.StoragePIDs
	Servers  storageservers.StorageServers
	Subjects storagesubject.StorageSubjects
	Channel  chan receivernotification.Notification
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
					a.PostInChannel("new", sub, h, m.Username, m.PID, m.Time)
					a.Subjects.Add(h.HostName, sub)
				}
				break
			case "syscall":
				sub := mapper.JSONtoSubject(m.Json, m.TypeSubject)
				if _, compare := a.Servers.Compare(h); compare {
					if a.PIDs.Contains(h.HostName, m.PID) {
						a.PostInChannel("syscall", sub, h, m.Username, m.PID, m.Time)
					}
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
					sub := mapper.JSONtoSubject(m.Json, m.TypeSubject)
					if u.PIDs.Contains(h.HostName, m.PID) {
						u.PostInChannel("update", sub, h, m.Username, m.PID, m.Time)
					}
					u.Subjects.Update(h.HostName, sub)
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

					u.PostInChannel("delete", sub, h, m.Username, m.PID, m.Time)

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
	request := g.Request
	body, _ := io.ReadAll(request.Body)
	defer request.Body.Close()
	log.Println(string(body))
	jsoned := string(body)
	jsoned = strings.Replace(jsoned, "%7B", "{", -1)
	jsoned = strings.Replace(jsoned, "%22", "\"", -1)
	jsoned = strings.Replace(jsoned, "%3A", ":", -1)
	jsoned = strings.Replace(jsoned, "%2C", ",", -1)
	jsoned = strings.Replace(jsoned, "%7D", "}", -1)
	jsoned = strings.Replace(jsoned, "json=", "", -1)
	m := subject.Message{}
	json.Unmarshal([]byte(jsoned), &m)
	return m
}

func MessageToHostInfo(s subject.Message) hostinfo.HostInfo {
	return hostinfo.HostInfo{
		HostName: s.HostName,
		HostOS:   s.SystemOS,
		IPs:      s.HostIP,
	}
}

func (p Post) PostInChannel(tag string, sub subject.Subject, info hostinfo.HostInfo, username, pid string, t time.Time) bool {
	who := p.Subjects.Get(info.HostName, subject.User{Username: username}).(subject.User)
	whoProcess := p.Subjects.Get(info.HostName, subject.Process{PID: pid}).(subject.Process)
	base := receivernotification.BaseNotification{Host: info, Time: t, Who: &who, WhoProcess: &whoProcess}
	switch tag {
	case "new":
		switch sub.Type() {
		case subject.ProcessT:
			p.Channel <- receivernotification.ProcessNew{
				Process:          sub.(subject.Process),
				BaseNotification: base}
			return true
		case subject.FileT:
			p.Channel <- receivernotification.FileNew{
				File:             sub.(subject.File),
				BaseNotification: base,
			}
			return true
		case subject.UserT:
			p.Channel <- receivernotification.UserNew{
				User:             sub.(subject.User),
				BaseNotification: base,
			}
			return true
		case subject.PortTablesT:
			p.Channel <- receivernotification.PortNew{
				Port:             sub.(subject.PortTables),
				BaseNotification: base,
			}
			return true
		}
		break
	case "delete":
		switch sub.Type() {
		case subject.ProcessT:
			p.Channel <- receivernotification.ProcessDelete{
				Process:          sub.(subject.Process),
				BaseNotification: base}
			return true
		case subject.FileT:
			p.Channel <- receivernotification.FileDelete{
				File:             sub.(subject.File),
				BaseNotification: base,
			}
			return true
		case subject.UserT:
			p.Channel <- receivernotification.UserDelete{
				User:             sub.(subject.User),
				BaseNotification: base,
			}
			return true
		case subject.PortTablesT:
			p.Channel <- receivernotification.PortDelete{
				Port:             sub.(subject.PortTables),
				BaseNotification: base,
			}
			return true
		}
		break
	case "update":
		subBefore := p.Subjects.Get(info.HostName, sub)
		switch sub.Type() {
		case subject.ProcessT:
			p.Channel <- receivernotification.ProcessUpdate{
				ProcessBefore:    subBefore.(subject.Process),
				ProcessAfter:     sub.(subject.Process),
				BaseNotification: base}
			return true
		case subject.FileT:
			p.Channel <- receivernotification.FileUpdate{
				FileBefore:       subBefore.(subject.File),
				FileAfter:        sub.(subject.File),
				BaseNotification: base,
			}
			return true
		case subject.UserT:
			p.Channel <- receivernotification.UserUpdate{
				UserBefore:       subBefore.(subject.User),
				UserAfter:        sub.(subject.User),
				BaseNotification: base,
			}
			return true
		case subject.PortTablesT:
			p.Channel <- receivernotification.PortUpdate{
				PortBefore:       subBefore.(subject.PortTables),
				PortAfter:        sub.(subject.PortTables),
				BaseNotification: base,
			}
			return true
		}
		break
	case "syscall":
		p.Channel <- receivernotification.Syscall{
			Syscall:          sub.(subject.Syscall),
			BaseNotification: base,
		}
		return true
	}
	return false
}
