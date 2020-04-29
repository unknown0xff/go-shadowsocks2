package shadowsocks

import (
	"C"

	"context"

	"github.com/TheWanderingCoel/go-shadowsocks2/core"

)

var (
	ctx     context.Context
	cancel  context.CancelFunc
)

func start_goShadowsocks(ClientAddr *C. char, ServerAddr *C.char, Cipher *C.char, Password *C.char, Plugin *C.char, PluginOptions *C.char) {

	var err error
	addr := C.GoString(ServerAddr)

	ctx, cancel = context.WithCancel(context.Background())

	var key []byte

	ciph, err := core.PickCipher(C.GoStrng(Cipher), key, C.GoString(Password))
	if err != nil {
	}

	if C.GoString(Plugin) != "" {
		addr, err = startPlugin(C.GoString(Plugin), C.GoString(PluginOptions), addr, false)
		if err != nil {
		}
	}

	go socksLocal(C.GoString(ClientAddr), addr, ciph.StreamConn, ctx)
	go udpSocksLocal(C.GoString(ClientAddr), addr, ciph.PacketConn, ctx)

}

func stop_goShadowsocks() {
	killPlugin()

	closeTcpLocal()
	closeUdpLocal()
}