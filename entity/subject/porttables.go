package subject

import (
	"encoding/json"
	"github.com/bastjan/netstat"
	"net"
	"server_siem/hash"
	"strconv"
)

type SocketAddress struct {
	IP   net.IP
	Port uint16
}

type Protocol struct {
	Name string
	Path string
}

type PortTables struct {
	LocalAddress  *SocketAddress
	RemoteAddress *SocketAddress
	State         netstat.TCPState
	UserId        string
	PID           string
	Protocol      Protocol
	TransmitQueue uint64
	ReceiveQueue  uint64
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
	return portTables.LocalAddress.IP.String() + strconv.Itoa(int(portTables.LocalAddress.Port))
}

func (portTables PortTables) Hash(hash hash.Hash) string {
	return hash(portTables.JSON())
}

func PortTablesFromJSON(jsoned string) (PortTables, error) {
	var portTables PortTables
	err := json.Unmarshal([]byte(jsoned), &portTables)
	return portTables, err
}
