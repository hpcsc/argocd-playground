package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hpcsc/argocd-playground-app/log"
	"github.com/hpcsc/argocd-playground-app/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
)

type infoHandler struct {
	versionFilePath string
}

func NewInfoHandler(filePath string) *infoHandler {
	return &infoHandler{versionFilePath: filePath}
}

func (h infoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := log.NewSugaredLoggerWithContext(logger, request).
		With(zap.String("file", h.versionFilePath))

	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	content, err := ioutil.ReadFile(h.versionFilePath)
	if err != nil {
		message := fmt.Sprintf("failed to read version: %v", err)
		sugar.Error(message)
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]string{"message": message})
		return
	}

	var version models.Version
	if err = json.Unmarshal(content, &version); err != nil {
		message := fmt.Sprintf("invalid file format: %v", err)
		sugar.Error(message)
		writer.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]string{"message": message})
		return
	}

	writer.WriteHeader(http.StatusOK)
	encoder.Encode(map[string]string{
		"env":     os.Getenv("ENVIRONMENT"),
		"version": version.Version,
		"commit":  version.Commit,
	})
}
