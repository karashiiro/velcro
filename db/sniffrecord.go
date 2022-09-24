package db

import (
	"net"
	"time"
)

type MessageHeader struct {
	Opcode    int    `json:"opcode"`
	Server    int    `json:"server"`
	Timestamp uint32 `json:"timestamp"`
}

type SegmentHeader struct {
	Size        uint32 `json:"size"`
	SourceActor uint32 `json:"source_actor"`
	TargetActor uint32 `json:"target_actor"`
	Type        int    `json:"type"`
}

type SniffRecord struct {
	Timestamp          time.Time      `json:"t"`
	Version            int            `json:"v"`
	SourceAddress      string         `json:"src_addr"`
	SourcePort         int            `json:"src_port"`
	DestinationAddress string         `json:"dst_addr"`
	DestinationPort    int            `json:"dst_port"`
	SegmentHeader      *SegmentHeader `json:"segment_header"`
	MessageHeader      *MessageHeader `json:"message_header"`
	MessageData        []byte         `json:"message_data"`
}

func (s *SniffRecord) GetSourceAddress() net.IP {
	return net.ParseIP(s.SourceAddress)
}

func (s *SniffRecord) GetDestinationAddress() net.IP {
	return net.ParseIP(s.DestinationAddress)
}
