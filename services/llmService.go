package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jarvis-backend/data"
	"jarvis-backend/models"
	"net/http"
	"os"
)

func QueryLLM(userTranscript string) (*models.JarvisResponse, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	modelName := os.Getenv("OPENROUTER_MODEL")
	apiURL := data.LLM_API_URL

	if apiKey == "" {
		return &models.JarvisResponse{
			Correction:   "API_KEY del LLM no configurada.",
			ReplyEnglish: "",
		}, nil
	}

	payload := map[string]interface{}{
		"model": modelName,
		"messages": []map[string]string{
			{"role": "system", "content": data.SYSTEM_PROMPT},
			{"role": "user", "content": userTranscript},
		},
		"response_format": map[string]string{"type": "json_object"},
		"temperature":     0.7,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error serializando payload: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creando petición a: %s Error: %w", apiURL, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("X-Title", "J.A.R.V.I.S. English Tutor") // Opcional (para identficar la app)
	//req.Header.Set("HTTP-Referer", "http://localhost:8080") // Opcional

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error enviando petición a: %s: Error: %w", apiURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta de: %s error: %w", apiURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LLM: error de API (status %d): %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, &models.ChatResponse); err != nil {
		return nil, fmt.Errorf("error parseando respuesta: %w", err)
	}

	if len(models.ChatResponse.Choices) == 0 {
		return nil, fmt.Errorf("respuesta recibida de %s vacía", apiURL)
	}

	rawContent := models.ChatResponse.Choices[0].Message.Content

	var jarvisResp models.JarvisResponse
	if err := json.Unmarshal([]byte(rawContent), &jarvisResp); err != nil {
		return nil, fmt.Errorf("error parseando JSON de Jarvis (%v). Contenido: %s", err, rawContent)
	}

	return &jarvisResp, nil
}
