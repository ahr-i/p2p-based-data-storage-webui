package p2p

import (
	"log"
	"time"
)

const listen_port = "50000"
const receiver_port = "50001"
const timeout = 1 * time.Second
const auth = "DISCOVERY"

type NodeList struct {
	Address string `json:"address"`
}

func GetNodeList(bootstrap string) []*NodeList {
	node_list := []*NodeList{}

	node_list = GetNodeListFromBroadcast()
	if len(node_list) <= 1 {
		if len(node_list) == 0 {
			log.Println("*(Node Discovery) Broadcast Error")
		}
		log.Println("*(Node Discovery) Use Bootstrap.")
		node_list = GetNodeListFromBootstrap(bootstrap)
	} else {
		log.Println("*(Node Discovery) Use BroadCast.")
	}

	return node_list
}
