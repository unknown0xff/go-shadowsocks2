package shadowsocks

import "C"
import (
	"context"

	"github.com/Trojan-Qt5/go-shadowsocks2/core"
	"github.com/Trojan-Qt5/go-shadowsocks2/stat"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func StartShadowsocks(ClientAddr string, ServerAddr string, Cipher string, Password string) {

	var err error
	addr := ServerAddr

	ctx, cancel = context.WithCancel(context.Background())

	var key []byte

	ciph, err := core.PickCipher(Cipher, key, Password)
	if err != nil {
	}

	meter := &stat.MemoryTrafficMeter{}

	go socksLocal(ClientAddr, addr, meter, ciph.StreamConn, ctx)
	go udpSocksLocal(ClientAddr, addr, ciph.PacketConn, ctx)

}

func StopShadowsocks() {

	cancel()

	closeTcpLocal()
	closeUdpLocal()
}
