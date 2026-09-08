package routers

import (
	"jarvis-backend/handlers"
	"net/http"
)

func SetServer() {
	http.HandleFunc("/api/process-audio", CORS(handlers.ProcessAudioHandler))
	http.HandleFunc("/api/health", CORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"online","system":"J.A.R.V.I.S."}`))
	}))
}
