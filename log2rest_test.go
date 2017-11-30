package log2rest

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/bestmethod/go-logger"
)

func logLine(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Body == nil {
		http.Error(w, "Request body missing!", 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var b Message
	err := decoder.Decode(&b)
	if err != nil {
		http.Error(w, "Invalid json provided in post content. Must be: {\"message\":\"Your Message To Log Here\"}", 400)
		return
	}
	fmt.Printf("%s\n",b.Message)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func DieNow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusInternalServerError)
}

func TestLog2Rest(t *testing.T) {
	router := httprouter.New()
	router.POST("/:logLevel/die/:me", DieNow)
	router.POST("/:logLevel", logLine)
	go http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", 5555), router)

	llogger := new(Logger.Logger)
	llogger.Init("SUBNAME", "SERVICENAME", Logger.LEVEL_DEBUG | Logger.LEVEL_INFO | Logger.LEVEL_WARN, Logger.LEVEL_ERROR | Logger.LEVEL_CRITICAL, Logger.LEVEL_NONE)

	logger := Log2Rest{Endpoint:"http://127.0.0.1:5555",LocalLogger:llogger}
	logger.Debug("Debug Test")
	logger.Info("Info Test")
	logger.Warn("Warn Test")
	logger.Error("Error Test")
	logger.Critical("Critical Test")
	logger = Log2Rest{Endpoint:"http://127.0.0.1:5555/"}
	logger.Debug("Debug Test")
	logger.Info("Info Test")
	logger.Warn("Warn Test")
	logger.Error("Error Test")
	logger.Critical("Critical Test")
	logger = Log2Rest{Endpoint:"http://127.0.0.1:5555/asdf"}
	err := logger.Debug("Debug Test")
	fmt.Printf("%s\n",err)
	logger = Log2Rest{Endpoint:"http://127.0.0.1:5555/die/die"}
	err = logger.Debug("Debug Test")
	fmt.Printf("%s\n",err)
}