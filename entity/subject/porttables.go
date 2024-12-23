package subject

import (
	"encoding/json"
	"github.com/bastjan/netstat"
	"net"
	"server_siem/hash"
	"strconv"
)

type SocketAddress struct {
	IP   net.IP `bson:"socket_ip"`
	Port uint16 `bson:"socket_port"`
}

type Protocol struct {
	Name string `bson:"protocol_name"`
	Path string `bson:"protocol_path"`
}

type LocalRemote struct {
	LocalAddress  string           `bson:"local_address"`
	RemoteAddress *SocketAddress   `bson:"remote_address"`
	State         netstat.TCPState `bson:"state"`
	UserId        string           `bson:"user_id"`
	PID           string           `bson:"pid"`
	Protocol      Protocol         `bson:"protocol"`
	TransmitQueue uint64
	ReceiveQueue  uint64
}

type PortTables struct {
	Port         uint64        `bson:"port"`
	LocalRemotes []LocalRemote `bson:"remotes_config"`
}

func (portTables PortTables) JSON() string {
	bytes, err := json.Marshal(portTables)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (portTables PortTables) Type() SubjectType {
	return PortTablesT
}

func (portTables PortTables) Name() string {
	return strconv.Itoa(int(portTables.Port))
}

func (portTables PortTables) Hash(hash hash.Hash) string {
	return hash(portTables.JSON())
}

func PortTablesFromJSON(jsoned string) (PortTables, error) {
	var portTables PortTables
	err := json.Unmarshal([]byte(jsoned), &portTables)
	return portTables, err
}
