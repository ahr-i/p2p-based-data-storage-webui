package webHandler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/fileToHash"

	"github.com/gorilla/mux"
)

func (wh *WebHandler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	/* Download File ID값 추출 */
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)

		return
	}

	/* DB에서 File Metadata 가져오기 */
	file_metadata := wh.Database.GetFileMetadata(id)
	if file_metadata == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "File Not Found")

		return
	}

	/* 무결성 Check */
	hash, err := fileToHash.CalcSHA256(file_metadata.Path)
	if err != nil || hash != file_metadata.Hash {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("File Error :", id)

		return
	}

	/* File 전송 */
	http.ServeFile(w, r, file_metadata.Path)
}
