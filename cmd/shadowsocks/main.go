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

func start_goShadowsocks(ClientAddr *C. char, ServerAddr *C.char, Cipher *C.char, Password *C.char) {

	ctx, cancel = context.WithCancel(context.Background())

	var key []byte

	ciph, err := core.PickCipher(C.GoStrng(Cipher), key, C.GoString(Password))
	if err != nil {

	}

	go socksLocal(C.GoString(ClientAddr), C.GoString(ServerAddr), ciph.StreamConn, ctx)
	go udpSocksLocal(C.GoString(ClientAddr), C.GoString(ServerAddr), ciph.PacketConn, ctx)

}

func stop_goShadowsocks() {
	closeTcpLocal()
	closeUdpLocal()
}