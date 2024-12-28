package wol

import (
	"errors"
	"net"
)

var ErrInvalidMACAddress = errors.New("Invalid MAC address")

// MagicPacket represents the Wake-on-LAN packet structure
type MagicPacket struct {
	header  []byte
	macAddr []byte
}

// NewMagicPacket creates a new Wake-on-LAN magic packet
func NewMagicPacket(macAddr string) (*MagicPacket, error) {
	// Parse MAC address
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		if err.(*net.AddrError).Err == "invalid MAC address" {
			return nil, errors.Join(ErrInvalidMACAddress, err)
		}

		return nil, err
	}

	if len(mac) != 6 {
		return nil, ErrInvalidMACAddress
	}

	// Create magic packet structure
	packet := &MagicPacket{
		header:  make([]byte, 6),  // 6 bytes of 0xFF
		macAddr: make([]byte, 96), // MAC address repeated 16 times
	}

	// Fill header with 0xFF
	for i := 0; i < 6; i++ {
		packet.header[i] = 0xFF
	}

	// Fill MAC address portion (16 repetitions)
	for i := 0; i < 16; i++ {
		copy(packet.macAddr[i*6:(i+1)*6], mac)
	}

	return packet, nil
}

// Bytes returns the complete magic packet as a byte slice
func (mp *MagicPacket) Bytes() []byte {
	packet := make([]byte, 102) // 6 + 96 bytes total
	copy(packet[0:6], mp.header)
	copy(packet[6:], mp.macAddr)
	return packet
}

// SendPacket sends the magic packet to the specified address
func SendPacket(macAddr, broadcastAddr string) error {
	packet, err := NewMagicPacket(macAddr)
	if err != nil {
		return err
	}

	// Create UDP address
	addr, err := net.ResolveUDPAddr("udp", broadcastAddr+":0")
	if err != nil {
		return err
	}

	// Create UDP connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send packet
	_, err = conn.Write(packet.Bytes())
	return err
}
