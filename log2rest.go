package log2rest

//TODO: unittests

import (
	"github.com/bestmethod/go-logger"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"errors"
)

type Log2Rest struct {
	Endpoint string
	LocalLogger *Logger.Logger
	sep *string
}

type Message struct {
	Message string
}

func (l *Log2Rest) Debug(m string) error {
	if l.LocalLogger != nil { l.LocalLogger.Debug(m) }
	return l.call("DEBUG",m)
}

func (l *Log2Rest) Info(m string) error {
	if l.LocalLogger != nil { l.LocalLogger.Info(m) }
	return l.call("INFO",m)
}

func (l *Log2Rest) Warn(m string) error {
	if l.LocalLogger != nil { l.LocalLogger.Warn(m) }
	return l.call("WARN",m)
}

func (l *Log2Rest) Error(m string) error {
	if l.LocalLogger != nil { l.LocalLogger.Error(m) }
	return l.call("ERROR",m)
}

func (l *Log2Rest) Critical(m string) error {
	if l.LocalLogger != nil { l.LocalLogger.Critical(m) }
	return l.call("CRITICAL",m)
}

func (l *Log2Rest) call(level string, m string) error {
	//json encode
	j, err := json.Marshal(&Message{Message:m})
	if err != nil { return err }

	//if we didn't work out separator yet, do it now. Only doing it once for speed.
	if l.sep == nil {
		var sep string
		if l.Endpoint[len(l.Endpoint)-1:] == "/" {
			sep = ""
		} else {
			sep = "/"
		}
		l.sep = &sep
	}

	//perform request, handle errors
	req, err := http.NewRequest("POST", strings.Join([]string{l.Endpoint,level},*l.sep), bytes.NewBuffer(j))
	if err != nil { return err }
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil { return err }
	defer resp.Body.Close()

	//handle non-200 response
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.New(fmt.Sprintf("REST Upload failed. StatusCode: %d, Status: %s, Header: %s", resp.StatusCode, resp.Status, resp.Header))
		} else {
			return errors.New(fmt.Sprintf("REST Upload failed. StatusCode: %d, Status: %s, Header: %s, Body: %s", resp.StatusCode, resp.Status, resp.Header, string(body)))
		}
	}
	return nil
}
