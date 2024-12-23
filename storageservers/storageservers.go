package storageservers

import "server_siem/hostinfo"

type TypeHost string

var (
	Server   TypeHost = "server"
	Receiver TypeHost = "receiver"
	Nope     TypeHost = "nope"
)

type StorageServers interface {
	Add(hostinfo.HostInfo, TypeHost) bool
	Exists(hostinfo.HostInfo) (TypeHost, bool)
	Update(hostinfo.HostInfo) bool
	Delete(hostinfo.HostInfo) bool
	GetType(TypeHost) []string
}