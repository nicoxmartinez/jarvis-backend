package models

type ErrorSalida struct {
	Status  int    `json:"codigo_error"`
	Mensaje string `json:"mensaje"`
}

func Error(codigoError int, mensaje string) ErrorSalida {
	var e ErrorSalida
	e.Status = codigoError
	e.Mensaje = mensaje

	return e
}
