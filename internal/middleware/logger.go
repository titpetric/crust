package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/crusttech/crust/internal/logger"
)

// ContextLogger middleware binds logger to request's context.
//
// This allows us to use logger from context (with requestID)
// inside our (generated) handers and controllers
func ContextLogger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var requestID = middleware.GetReqID(req.Context())

			w.Header().Add("X-Request-Id", requestID)

			req = req.WithContext(logger.ContextWithValue(
				req.Context(),
				log.With(zap.String("requestID", requestID)).Named("rest"),
			))

			next.ServeHTTP(w, req)
		})
	}
}

// LogRequest middleware logs request details
//
// It uses logger from context, see ContextLogger()
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.ContextValue(req.Context()).Info(
			"HTTP request "+req.Method+" "+req.URL.Path,
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.Int64("size", req.ContentLength),
			zap.String("remote", req.RemoteAddr[:strings.Index(req.RemoteAddr, ":")]),
		)
		next.ServeHTTP(w, req)
	})
}

// LogResponse middleware logs response details
//
// It uses logger from context, see ContextLogger()
func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		wrapped := middleware.NewWrapResponseWriter(w, req.ProtoMajor)
		t := time.Now()

		defer func() {
			logger.ContextValue(req.Context()).Info(
				"HTTP response "+req.Method+" "+req.URL.Path,
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", wrapped.Status()),
				zap.Int("size", wrapped.BytesWritten()),
				zap.Float64("duration", time.Now().Sub(t).Seconds()),
			)
		}()

		next.ServeHTTP(wrapped, req)

	})

}
