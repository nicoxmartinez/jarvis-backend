# J.A.R.V.I.S. English Tutor - Backend (Go)
backend diseñado como un asistente de procesamiento de voz e inteligencia artificial en tiempo real. Su objetivo principal es recibir entradas de voz del usuario, convertirlas a texto mediante reconocimiento de voz, procesar el contexto con un modelo de lenguaje (LLM) y devolver una respuesta estructurada.

NOTA: Almacenamiento temporal para archivos de audio entrantes (.wav) en la carpeta ``/uploads``
---

## 🛠️ Tecnologías Utilizadas

- **Lenguaje:** Go (v1.24+)
- **Servidor HTTP:** `net/http` (nativo de Go)
- **Environment:** godotenv (v1.5.1) 
- **STT API**
- **LLM API**

---

## ⚙️ Configuración e Instalación

### Requisitos Previos

- Tener instalado **Go** (versión 1.20 o superior).
- Cuenta y clave de API para el modelo de lenguaje.
- Cuenta y clave de API para la conversión de voz a texto.

### Instalar dependencias

```bash
go mod tidy
```

### Ejecución del Servidor

```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`.

---
