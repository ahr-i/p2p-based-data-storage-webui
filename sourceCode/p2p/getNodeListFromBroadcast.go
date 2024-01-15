package p2p

import (
	"log"
	"net"
	"time"

	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/myLocal"
)

var address = myLocal.GetLocalIP() + ":" + receiver_port

func GetNodeListFromBroadcast() []*NodeList {
	node_list := []*NodeList{}
	Broadcast_address := myLocal.GetBroadcastAddress() + ":" + listen_port // Broadcast Address 계산

	/* Message 송신 / Auth + 수신받을 Address */
	sender_conn, err := net.Dial("udp", Broadcast_address)
	if err != nil {
		panic(err)
	}
	defer sender_conn.Close()

	_, err = sender_conn.Write([]byte(auth + address))
	if err != nil {
		panic(err)
	}

	/* Message 수신 / Auth + Node의 Address */
	buffer := make([]byte, 1024)
	receiver_addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}

	receiver_conn, err := net.ListenUDP("udp", receiver_addr)
	if err != nil {
		panic(err)
	}
	defer receiver_conn.Close()

	for {
		/* Time Out Setting */
		receiver_conn.SetReadDeadline(time.Now().Add(timeout))

		/* 모든 Node에게 Message를 수신받음 */
		n, _, err := receiver_conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err, "/ Time Out")

			break
		}

		auth_str := string(buffer[:len(auth)])
		node_addr := string(buffer[len(auth):n])
		if auth_str == auth {
			node_list = append(node_list, &NodeList{Address: node_addr})
		}
	}

	return node_list
}
