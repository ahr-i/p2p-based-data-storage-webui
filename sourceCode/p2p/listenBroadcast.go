package p2p

import (
	"net"

	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/myLocal"
)

func ListenBroadcast() {
	/* Message 수신 / Auth + 수신받을 Address */
	receiver_addr, err := net.ResolveUDPAddr("udp", ":"+listen_port)
	if err != nil {
		panic(err)
	}

	receiver_conn, err := net.ListenUDP("udp", receiver_addr)
	if err != nil {
		panic(err)
	}
	defer receiver_conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, err := receiver_conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		auth_str := string(buffer[:len(auth)])
		sender_addr := string(buffer[len(auth):n])
		if auth_str == auth {
			/* Message 송신 / Auth + 자신의 Address */
			sender_conn, err := net.Dial("udp", sender_addr)
			if err != nil {
				panic(err)
			}
			defer sender_conn.Close()

			_, err_ := sender_conn.Write([]byte(auth + myLocal.GetLocalIP() + ":2000"))
			if err_ != nil {
				panic(err)
			}
		}
	}
}
