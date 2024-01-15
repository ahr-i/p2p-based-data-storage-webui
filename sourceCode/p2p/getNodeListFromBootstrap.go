package p2p

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetNodeListFromBootstrap(bootstrap string) []*NodeList {
	node_list := []*NodeList{}
	response, err := http.Get(bootstrap + "/nodes")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&node_list)
	if err != nil {
		log.Println("Error decoding response:", err)
		return nil
	}

	return node_list
}
