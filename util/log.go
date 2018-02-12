package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// UseJSONLogFormat sets up the JSON log formatter
func UseJSONLogFormat() {
	env := GetEnv("GIN_ENV", "development")
	program := GetEnv("SERVICE_NAME", "gin-api-demo")

	log.SetFormatter(&JSONFormatter{
		Program: program,
		Env:     env,
	})

	// so our debug entries appear!
	log.SetLevel(log.DebugLevel)
}

// Timestamps in microsecond resolution (like time.RFC3339Nano but microseconds)
var timeStampFormat = "2006-01-02T15:04:05.000000Z07:00"

// JSONFormatter is a logger for use with Logrus
type JSONFormatter struct {
	Program string
	Env     string
}

// Format includes the program, environment, and a custom time format: microsecond resolution
func (f *JSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["time"] = entry.Time.UTC().Format(timeStampFormat)
	data["msg"] = entry.Message
	data["level"] = strings.ToUpper(entry.Level.String())
	data["program"] = f.Program
	data["env"] = f.Env

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

// GetClientIP gets the correct IP for the end client instead of the proxy
func GetClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

// GetUserID gets the current_user ID as a string
func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if exists {
		return userID.(string)
	}
	return ""
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
