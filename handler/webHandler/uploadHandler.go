package webHandler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ahr-i/p2p-based-data-storage-webui/handler/dbHandler"
	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/fileToHash"
	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/p2p"

	"github.com/gorilla/mux"
)

type NodeList struct {
	Address string `json:"address"`
}

/* 단일 Upload */
func (wh *WebHandler) Upload(w http.ResponseWriter, r *http.Request) {
	/* ----- Storage 저장 ----- */
	r.ParseMultipartForm(32 << 20)

	/* Upload할 File Data 가져오기 */
	upload_file, info, err := r.FormFile("user_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)

		return
	}
	defer upload_file.Close()

	/* 저장할 경로에 미리 File 생성 */
	dir_name := "./upload"
	file_path := fmt.Sprintf("%s/%s", dir_name, info.Filename)

	os.MkdirAll(dir_name, 0777)
	file, err := os.Create(file_path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)

		return
	}
	defer file.Close()

	/* File 저장 */
	io.Copy(file, upload_file)

	/* ----- Database 저장 ----- */
	/* File Hasing */
	hash, err := fileToHash.CalcSHA256(file_path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)

		return
	}

	/* File 존재 여부 확인 */
	if wh.Database.IsExistData(hash) {
		return
	}

	/* File Metadata Block 생성 및 Database 저장 */
	file_metadata := dbHandler.FileMetadata{
		Name:     info.Filename,
		Path:     file_path,
		Size:     int(info.Size),
		Hash:     hash,
		CreateAt: time.Now(),
	}
	id := wh.Database.SaveFileMetadata(file_metadata)
	file_metadata.ID = id

	rend.JSON(w, http.StatusOK, file_metadata)
}

/* 전체 Upload */
func (wh *WebHandler) AllUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("* (All Upload) All Upload Is Working.")
	/* Node List 가져오기 */
	node_list := p2p.GetNodeList(bootstrap)

	/* Upload할 File Data 읽어오기 */
	file, info, err := r.FormFile("user_file")
	file_name := info.Filename
	if err != nil {
		log.Println("Error reading file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	/* 순서대로 모든 Node에게 File 전달 */
	for _, node := range node_list {
		node_url := "http://" + node.Address + "/upload/0"

		/* Request Body 작성 */
		var request_body bytes.Buffer

		writer := multipart.NewWriter(&request_body)
		part, err := writer.CreateFormFile("user_file", file_name)
		if err != nil {
			log.Println("Error creating form file:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = io.Copy(part, file)
		if err != nil {
			log.Println("Error copying file:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.Close()
		file.Seek(0, 0)

		req, err := http.NewRequest("POST", node_url, &request_body)
		if err != nil {
			log.Println("Error creating request:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending request:", err)
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("* (All Upload) Received Response: %d %s", resp.StatusCode, resp.Status)
		log.Println("* (All Upload) Response body:", string(bodyBytes))
		resp.Body.Close()
	}

	rend.JSON(w, http.StatusOK, nil)
}

func (wh *WebHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	/* Type Check */
	vars := mux.Vars(r)
	upload_type, err := strconv.Atoi(vars["type"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)

		return
	}

	/*
		0 : 단일 Upload
		1 : 전체 Upload
	*/
	switch int(upload_type) {
	case 0:
		wh.Upload(w, r)
	case 1:
		wh.AllUpload(w, r)
	}
}
