package service

import (
	"server_siem/api"
	"server_siem/command"
	"server_siem/storagepids"
	"server_siem/storageservers"
	"server_siem/storagesubject"
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
	Method Method
	Path   string
	Action command.Action
}

type ApiService struct {
	Address  string
	API      api.Api
	Commands []PathWork
}

func InitApiService(address string,
	ds storagepids.StoragePIDs,
	servers storageservers.StorageServers,
	subjects storagesubject.StorageSubjects) ApiService {
	post := command.Post{ds, servers, subjects}
	return ApiService{API: api.InitApi(), Address: address, Commands: []PathWork{
		{POST, "/api/command", command.PostCommand{post}},
		{POST, "/api/server", post},
		{PATCH, "/api/server", command.Update{post}},
		{DELETE, "/api/server", command.Delete{post}},
	}}
}

func (a ApiService) Work() {
	if &a != nil {
		a.API = api.InitApi()
	}
	for _, c := range a.Commands {
		switch c.Method {
		case POST:
			a.API.Post(c.Path, c.Action)
		case GET:
			a.API.Get(c.Path, c.Action)
		case HEAD:
			a.API.Head(c.Path, c.Action)
		case OPTIONS:
			a.API.Options(c.Path, c.Action)
		case PATCH:
			a.API.Patch(c.Path, c.Action)
		case PUT:
			a.API.Patch(c.Path, c.Action)
		case DELETE:
			a.API.Delete(c.Path, c.Action)
		default:
			continue
		}
	}
	a.API.Run(a.Address)
}
