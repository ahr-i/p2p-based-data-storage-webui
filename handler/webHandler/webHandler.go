package webHandler

import (
	"net/http"

	"github.com/ahr-i/p2p-based-data-storage-webui/handler/dbHandler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rend *render.Render = render.New()
var bootstrap string
var port string

type WebHandler struct {
	http.Handler
	Database dbHandler.DBHandler
}

func CreateHandler(my_port string, db_path string, bootstrap_node string) *WebHandler {
	bootstrap = bootstrap_node // Bootstrap Address
	port = my_port             // 자신의 Port

	mux := mux.NewRouter()
	web_handler := &WebHandler{
		Handler:  mux,
		Database: dbHandler.CreateDBHandler(db_path),
	}

	mux.HandleFunc("/", web_handler.HomeHandler).Methods("GET")                         // Home URL
	mux.HandleFunc("/files", web_handler.GetFileListHandler).Methods("GET")             // File List Return URL
	mux.HandleFunc("/download/{id:[0-9]+}", web_handler.DownloadHandler).Methods("GET") // File Download URL
	mux.HandleFunc("/upload/{type:[0-9]+}", web_handler.UploadHandler).Methods("POST")  // File Upload URL

	mux.HandleFunc("/storage", web_handler.GetStorageInfoHandler).Methods("GET") // Storage Size Return URL
	mux.HandleFunc("/node_cnt", web_handler.GetNodeCntHandler).Methods("GET")    // 살아있는 Node 개수 Return URL
	mux.HandleFunc("/ping", web_handler.PingHandler).Methods("GET")              // Node Check 용 URL

	return web_handler
}
