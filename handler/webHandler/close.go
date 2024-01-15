package webHandler

func (wh *WebHandler) Close() {
	wh.Database.Close()
}
