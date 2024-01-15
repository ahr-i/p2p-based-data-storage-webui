package webHandler

import "net/http"

func (wh WebHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "index.html", http.StatusTemporaryRedirect)
}
