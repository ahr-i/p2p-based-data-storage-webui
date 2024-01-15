package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahr-i/p2p-based-data-storage-webui/handler/webHandler"
	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/myLocal"
	"github.com/ahr-i/p2p-based-data-storage-webui/sourceCode/p2p"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	/* Node Setting */
	const bootstrap_node = "http://200.0.0.20:3000" // Bootstrap Node Address
	//const bootstrap_node = "http://localhost:3000" // Bootstrap Node Address
	const db_path = "./database/file_metadata.db" // Database Path
	const port = "2000"                           // Node Port

	/* ----- */
	mux := webHandler.CreateHandler(port, db_path, bootstrap_node)
	handler := negroni.Classic()

	defer mux.Close() // Database Close

	cors_ := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                            // 모든 출처의 Server 접근 허용
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // 허용된 HTTP Methods
		AllowedHeaders:   []string{"*"},                            // 모든 Header 허용
		AllowCredentials: true,                                     // 자격 증명 허용
	})
	handler.Use(cors_)
	handler.UseHandler(mux)

	resp, _ := http.Get(bootstrap_node + "/register?ip=" + myLocal.GetLocalIP() + "&port=" + port) // Bootstrap에 Node Register
	if resp == nil || resp.StatusCode != http.StatusOK {
		log.Println("There Is No Bootstrap.")

		os.Exit(1)
	}

	go p2p.ListenBroadcast()
	http.ListenAndServe(":"+port, handler)
}
