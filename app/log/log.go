package log

import (
	"github.com/hpcsc/argocd-playground-app/middlewares"
	"go.uber.org/zap"
	"net/http"
)

func toZapStrings(data map[string]string) []interface{} {
	var output []interface{}
	for k, v := range data {
		output = append(output, zap.String(k, v))
	}
	return output
}

func NewSugaredLoggerWithContext(logger *zap.Logger, request *http.Request) *zap.SugaredLogger {
	data := request.Context().Value(middlewares.ContextDataKey)
	if data == nil {
		return logger.Sugar()
	}

	return logger.Sugar().With(toZapStrings(data.(map[string]string))...)
}
