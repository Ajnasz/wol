# WoL (Wake-on-LAN) Go Package

This Go package provides functionality to create and send Wake-on-LAN (WoL) magic packets to wake up computers over a network.

## Installation

To install the package, use the following command:

```sh
go get github.com/Ajnasz/wol
```

## Usage

### Creating a Magic Packet

To create a new Wake-on-LAN magic packet, use the `NewMagicPacket` function:

```go
import "github.com/Ajnasz/wol"

macAddr := "01:23:45:67:89:ab"
packet, err := wol.NewMagicPacket(macAddr)
if err != nil {
    // handle error
}
```

### Sending a Magic Packet

To send a magic packet to a specified broadcast address, use the `SendPacket` function:

```go
import "github.com/Ajnasz/wol"

macAddr := "01:23:45:67:89:ab"
broadcastAddr := "192.168.1.255"
err := wol.SendPacket(macAddr, broadcastAddr)
if err != nil {
    // handle error
}
```

## Functions

### `NewMagicPacket(macAddr string) (*MagicPacket, error)`

Creates a new Wake-on-LAN magic packet for the given MAC address.

- `macAddr`: The MAC address of the target device in string format.
- Returns a `MagicPacket` instance or an error if the MAC address is invalid.

### `(*MagicPacket) Bytes() []byte`

Returns the complete magic packet as a byte slice.

### `SendPacket(macAddr, broadcastAddr string) error`

Sends the magic packet to the specified broadcast address.

- `macAddr`: The MAC address of the target device in string format.
- `broadcastAddr`: The broadcast address to send the packet to.
- Returns an error if the packet could not be sent.

## Errors

- `ErrInvalidMACAddress`: Returned when the provided MAC address is invalid.

## Example

```go
package main

import (
    "github.com/Ajnasz/wol"
    "log"
)

func main() {
    macAddr := "01:23:45:67:89:ab"
    broadcastAddr := "192.168.1.255"
    
    err := wol.SendPacket(macAddr, broadcastAddr)
    if err != nil {
        log.Fatalf("Failed to send magic packet: %v", err)
    }
    
    log.Println("Magic packet sent successfully")
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

