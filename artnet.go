package artnet

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

const (
	OpPoll             = 0x2000
	OpPollReply        = 0x2100
	OpDiagData         = 0x2300
	OpCommand          = 0x2400
	OpOutput           = 0x5000
	OpDmx              = 0x5000
	OpNzs              = 0x5100
	OpSync             = 0x5200
	OpAddress          = 0x6000
	OpInput            = 0x7000
	OpTodRequest       = 0x8000
	OpTodData          = 0x8100
	OpTodControl       = 0x8200
	OpRdm              = 0x8300
	OpRdmSub           = 0x8400
	OpVideoSetup       = 0xa010
	OpVideoPalette     = 0xa020
	OpVideoData        = 0xa040
	OpMacMaster        = 0xf000
	OpMacSlave         = 0xf100
	OpFirmwareMaster   = 0xf200
	OpFirmwareReply    = 0xf300
	OpFileTnMaster     = 0xf400
	OpFileFnMaster     = 0xf500
	OpFileFnReply      = 0xf600
	OpIpProg           = 0xf800
	OpIpProgReply      = 0xf900
	OpMedia            = 0x9000
	OpMediaPatch       = 0x9100
	OpMediaControl     = 0x9200
	OpMediaContrlReply = 0x9300
	OpTimeCode         = 0x9700
	OpTimeSync         = 0x9800
	OpTrigger          = 0x9900
	OpDirectory        = 0x9a00
	OpDirectoryReply   = 0x9b00
)

var ARTNET = [8]uint8{65, 114, 116, 45, 78, 101, 116, 0}

const ProtVerHi = 0
const ProtVerLo = 14

type OpCode int16

type ArtDmx struct {
	Id        [8]uint8
	Opcode    OpCode
	ProtVerHi uint8
	ProtVerLo uint8
	Sequence  uint8
	Physical  uint8
	SubUni    uint8
	Net       uint8
	LengthHi  uint8
	Length    uint8
	Data      [512]uint8
}

func opByteSwap(opcode OpCode) OpCode {
	return (opcode>>8)&0xff | opcode&0xff
}

type Node struct {
	conn *net.UDPConn
}

func NewNode(service string) (*Node, error) {
	node := &Node{}

	addr, err := net.ResolveUDPAddr("udp", service)
	if err != nil {
		return nil, err
	}

	node.conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", node.conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", node.conn.LocalAddr().String())

	return node, nil
}

func (node *Node) Close() error {
	return node.conn.Close()
}

func (node *Node) Dmx(universe uint8, data [512]uint8) error {

	// Currently only support sending all 512 channels
	length := 512

	p := ArtDmx{
		Id:        ARTNET,
		Opcode:    opByteSwap(OpDmx),
		ProtVerHi: ProtVerHi,
		ProtVerLo: ProtVerLo,
		Sequence:  0x00,
		Physical:  0x00,
		SubUni:    (universe & 0xff),
		Net:       (universe >> 8) & 0xff,
		LengthHi:  uint8((length >> 8) & 0xff),
		Length:    uint8((length & 0xff)),
		Data:      data,
	}

	b := &bytes.Buffer{}
	binary.Write(b, binary.BigEndian, p)

	_, err := node.conn.Write(b.Bytes())

	return err
}
