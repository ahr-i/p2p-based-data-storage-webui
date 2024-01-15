package webHandler

import (
	"fmt"
	"net/http"
)

func (wh *WebHandler) GetStorageInfoHandler(w http.ResponseWriter, r *http.Request) {
	storage_size := wh.Database.GetStorageInfo()

	w.Write([]byte(fmt.Sprintf("%d Byte", storage_size)))
}
