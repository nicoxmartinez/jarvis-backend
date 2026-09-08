package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"jarvis-backend/models"
	"jarvis-backend/services"
)

func ProcessAudioHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(&rw, http.StatusMethodNotAllowed, "Metodo no Permitido")
		http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var tempFilePath string
	userText := ""

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		err := r.ParseMultipartForm(10 << 20) // 10MB max
		if err == nil {
			file, header, errFile := r.FormFile("audio")
			if errFile == nil {
				defer file.Close()

				tempFilePath = filepath.Join("./uploads", header.Filename)
				out, errCreate := os.Create(tempFilePath)
				if errCreate == nil {
					defer out.Close()
					defer os.Remove(tempFilePath) // Se borra automáticamente al finalizar
					_, _ = io.Copy(out, file)
				}
			}
		}
	}

	//Transcribir el archivo
	if tempFilePath != "" {
		transcription, err := services.TranscribeAudio(tempFilePath)
		if err != nil {
			ErrorResponse(&rw, http.StatusInternalServerError, "Error en transcripción STT: "+err.Error())
			http.Error(rw, "Error en transcripción STT: "+err.Error(), http.StatusInternalServerError)
			return
		}
		userText = transcription
	}

	//Procesar la transcripción
	jarvisResp, err := services.QueryLLM(userText)
	if err != nil {
		ErrorResponse(&rw, http.StatusInternalServerError, "Error de LLM: "+err.Error())
		http.Error(rw, "Error de LLM: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.ProcessResponse{
		Transcription: userText,
		Correction:    jarvisResp.Correction,
		ResponseText:  jarvisResp.ReplyEnglish,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(resp)
}
