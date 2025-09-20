package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body = append(lrw.body, b...) // сохраняем тело
	return lrw.ResponseWriter.Write(b)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// оборачиваем стандартный ResponseWriter
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// вызываем следующий обработчик
		next.ServeHTTP(lrw, r)

		duration := time.Since(start)

		log.Printf(
			"%s %s %s %d %v\nResponse: %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			lrw.statusCode, // статус ответа
			duration,
			string(lrw.body), // тело ответа
		)
	})
}
