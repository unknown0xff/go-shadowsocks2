package shadowsocks

import "C"
import (
	"context"

	"github.com/TheWanderingCoel/go-shadowsocks2/core"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func StartGoShadowsocks(ClientAddr string, ServerAddr string, Cipher string, Password string, Plugin string, PluginOptions string) {

	var err error
	addr := ServerAddr

	ctx, cancel = context.WithCancel(context.Background())

	var key []byte

	ciph, err := core.PickCipher(Cipher, key, Password)
	if err != nil {
	}

	if Plugin != "" {
		addr, err = startPlugin(Plugin, PluginOptions, addr, false)
		if err != nil {
		}
	}

	go socksLocal(ClientAddr, addr, ciph.StreamConn, ctx)
	go udpSocksLocal(ClientAddr, addr, ciph.PacketConn, ctx)

}

func StopGoShadowsocks() {
	killPlugin()

	cancel()

	closeTcpLocal()
	closeUdpLocal()
}
