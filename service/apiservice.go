package service

import (
	"server_siem/api"
	"server_siem/command"
)

type Method int

var (
	POST    Method = 0
	GET     Method = 1
	HEAD    Method = 2
	OPTIONS Method = 3
	PATCH   Method = 4
	PUT     Method = 5
	DELETE  Method = 6
)

type PathWork struct {
	Path   string
	Action command.Action
}

type ApiService struct {
	Address  string
	API      api.Api
	Commands map[Method]PathWork
}

func (a ApiService) Work() {
	if &a != nil {
		a.API = api.InitApi()
	}
	for method, work := range a.Commands {
		switch method {
		case POST:
			a.API.Post(work.Path, work.Action)
		case GET:
			a.API.Get(work.Path, work.Action)
		case HEAD:
			a.API.Head(work.Path, work.Action)
		case OPTIONS:
			a.API.Options(work.Path, work.Action)
		case PATCH:
			a.API.Patch(work.Path, work.Action)
		case PUT:
			a.API.Patch(work.Path, work.Action)
		case DELETE:
			a.API.Delete(work.Path, work.Action)
		default:
			continue
		}
	}
	a.API.Run(a.Address)
}
