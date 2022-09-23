package db

import (
	"net"
	"time"
)

type SniffRecord struct {
	Timestamp          time.Time `json:"t"`
	Version            int       `json:"v"`
	Segment            int       `json:"segment"`
	Opcode             *int      `json:"opcode"`
	SourceAddress      string    `json:"src_addr"`
	SourcePort         int       `json:"src_port"`
	DestinationAddress string    `json:"dst_addr"`
	DestinationPort    int       `json:"dst_port"`
	Data               []byte    `json:"data"`
}

func (s *SniffRecord) GetSourceAddress() net.IP {
	return net.ParseIP(s.SourceAddress)
}

func (s *SniffRecord) GetDestinationAddress() net.IP {
	return net.ParseIP(s.DestinationAddress)
}
