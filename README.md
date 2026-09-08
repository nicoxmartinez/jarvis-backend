# J.A.R.V.I.S. English Tutor - Backend (Go)

---

## 🛠️ Tecnologías Utilizadas

- **Lenguaje:** Go (v1.24+)
- **STT API:** OpenAI Whisper API (`whisper-1`)
- **LLM API:** OpenRouter API (Soporta `meta-llama/llama-3.3-70b-instruct:free`, `openai/gpt-4o-mini`, etc.)
- **Servidor HTTP:** `net/http` nativo de Go con middleware para soporte de CORS.

---

## ⚙️ Configuración e Instalación

### 1. Requisitos Previos

- Tener instalado **Go** (versión 1.20 o superior).
- Cuenta y clave de API en **OpenAI** (para Whisper STT).
- Cuenta y clave de API en **OpenRouter** (para el modelo de lenguaje).

### 2. Variables de Entorno

Exporta las claves de API en tu terminal antes de ejecutar el servidor:

```bash
# Clave para el servicio de Transcripción (OpenAI Whisper)
export OPENAI_API_KEY="tu-clave-de-openai"

# Clave para el modelo de lenguaje (OpenRouter)
export OPENROUTER_API_KEY="tu-clave-de-openrouter"

# (Opcional) Especificar modelo en OpenRouter
export OPENROUTER_MODEL="openrouter/free"
```

### 3. Ejecución del Servidor

```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`.

---