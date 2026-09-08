package models

type ProcessResponse struct {
	Transcription string `json:"transcription"`
	Correction    string `json:"correction"`
	ResponseText  string `json:"response_text"`
	AudioBase64   string `json:"audio_base64,omitempty"`
}
