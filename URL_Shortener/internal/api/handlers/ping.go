package handlers

import (
	"net/http"
	"urlshortener/utils"
)

func PingServer(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Server was pinged"))
	if err != nil {
		http.Error(w, utils.EncodingMessageError.Error(), utils.EncodingMessageError.GetStatusCode())
		return
	}
}
