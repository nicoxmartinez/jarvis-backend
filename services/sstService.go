package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jarvis-backend/models"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// TranscribeAudio envía un archivo de audio a la API de SST y devuelve el texto transcrito
func TranscribeAudio(filePath string) (string, error) {
	apiURL := os.Getenv("SST_API_URL")
	apiKey := os.Getenv("SST_API_KEY")
	modelName := os.Getenv("SST_MODEL")

	if apiKey == "" {
		return "API_KEY del SST no configurada.", nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error abriendo archivo de audio para STT: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("error creando form-file: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("error copiando bytes de audio: %w", err)
	}

	err = writer.WriteField("model", modelName)
	if err != nil {
		return "", fmt.Errorf("error escribiendo campo model: %s Error: %w", modelName, err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return "", fmt.Errorf("error creando petición a %s Error: %w", apiURL, err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error enviando petición a: %s Error: %w", apiURL, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error leyendo respuesta de: %s Error: %w", apiURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("SST: error API  (status %d): %s", resp.StatusCode, string(respBody))
	}

	var sstResponse models.SstResponse
	if err := json.Unmarshal(respBody, &sstResponse); err != nil {
		return "", fmt.Errorf("error parseando respuesta JSON : %w", err)
	}

	return sstResponse.Text, nil
}
