package webHandler

import (
	"io"
	"net/http"
)

func (wh *WebHandler) GetNodeCntHandler(w http.ResponseWriter, r *http.Request) {
	/* Bootstrap에서 살아있는 Node의 개수 정보 받아오기 */
	response, err := http.Get(bootstrap + "/alive_node_cnt")
	if err != nil {
		http.Error(w, "Failed to get node count from bootstrap", http.StatusInternalServerError)

		return
	}
	defer response.Body.Close()

	io.Copy(w, response.Body)
}
