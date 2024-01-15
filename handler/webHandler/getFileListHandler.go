package webHandler

import (
	"net/http"
)

func (wh *WebHandler) GetFileListHandler(w http.ResponseWriter, r *http.Request) {
	file_list := wh.Database.GetFileList()

	rend.JSON(w, http.StatusOK, file_list)
}
