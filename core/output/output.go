package output

import (
	"encoding/json"
	"net/http"
)

const (
	responseStatusOk    = "OK"
	responseStatusError = "ERROR"
)

// Response - Объект, возвращаемый в ответ на запрос клиенту
type Response struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Code    int                    `json:"code"`
}

// Ok Возвращает ответ со статусом OK если:
// - Запрос отработал нормально и запрашиваемые данные получены.
// - Запрос отработал нормально и запрашиваемые данные не найдены.
func Ok(w http.ResponseWriter, responsePayload map[string]interface{}) {

}

// Error Возвращает ответ со статусом ERROR если:
// - Параметры запроса не валидны, либо отсутствуют обязательные данные.
// - Не найден запрашиваемые контроллер или метод.
// - Произошла непредвиденная ошибка. Например ошибка при выполнении запроса к БД.
// - Отсутствуют необходимые права доступа.
func Error() {

}

// Формирует ответ в виде json структуры и отправляет клиенту
func send(
	w http.ResponseWriter,
	responseStatus string,
	responsePayload map[string]interface{},
	errorMessage string,
	errorCode int,
) {
	w.Header().Set("Content-Type", "application/json")

	responce := Response{
		Status:  responseStatus,
		Data:    responsePayload,
		Message: errorMessage,
		Code:    errorCode,
	}

	json.NewEncoder(w).Encode(responce)
}
