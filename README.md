# J.A.R.V.I.S. English Tutor - Backend (Go)

---

## 🛠️ Tecnologías Utilizadas

- **Lenguaje:** Go (v1.24+)
- **Servidor HTTP:** `net/http` (nativo de Go)
- **STT API**
- **LLM API**

---

## ⚙️ Configuración e Instalación

### 1. Requisitos Previos

- Tener instalado **Go** (versión 1.20 o superior).
- Cuenta y clave de API para la conversión de voz a texto.
- Cuenta y clave de API para el modelo de lenguaje.

### 2. Variables de Entorno

Exporta las claves de API en tu terminal antes de ejecutar el servidor:

```bash
# URL de la API
export SST_API_URL="sst-api-url"

# Clave para el servicio de transcripción 
export SST_API_KEY="tu-clave"

# Especificar modelo
export SST_MODEL="nombre-del-modelo"

# URL de la API
export LLM_API_URL="llm-api-url"

# Clave para el modelo de lenguaje 
export LLM_API_KEY="tu-clave"

# Especificar modelo
export LLM_MODEL="nombre-del-modelo"
```

### 3. Ejecución del Servidor

```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`.

---
