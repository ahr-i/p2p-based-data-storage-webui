package webHandler

import "net/http"

func (wh *WebHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
