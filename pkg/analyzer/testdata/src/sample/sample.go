package sample

import (
	"fmt"
)

type Logger struct{}

func (Logger) Info(msg string, args ...any)  {}
func (Logger) Error(msg string, args ...any) {}
func (Logger) Warn(msg string, args ...any)  {}
func (Logger) Debug(msg string, args ...any) {}

var logger Logger

func valid() {
	logger.Info("user logged in")
	logger.Info("connection established")
	logger.Debug("processing request")
}

func invalidLowercase() {
	logger.Info("User logged in")     // want `log message must start with a lowercase letter`
	logger.Error("Connection failed") // want `log message must start with a lowercase letter`
}

func invalidEnglish() {
	logger.Info("пользователь вошёл") // want `log message must contain only English letters`
	logger.Warn("ошибка соединения")  // want `log message must contain only English letters`
}

func invalidSensitive() {
	logger.Info("user password reset")        // want `log message must not contain sensitive data`
	logger.Info("received token from server") // want `log message must not contain sensitive data`
	logger.Info("auth failed")                // want `log message must not contain sensitive data`
}

func invalidSpecialChars() {
	logger.Info("success!")              // want `log message must not contain special characters`
	logger.Info("user @admin logged in") // want `log message must not contain special characters`
	logger.Info("done 🎉")                // want `log message must not contain special characters`
}

func notLogCall() {
	fmt.Println("User logged in")
}
