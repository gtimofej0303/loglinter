package main

type Logger struct{}

func (Logger) Info(msg string, args ...any)  {}
func (Logger) Error(msg string, args ...any) {}
func (Logger) Warn(msg string, args ...any)  {}
func (Logger) Debug(msg string, args ...any) {}

var logger Logger

func main() {
	// Нарушение 1: заглавная буква в начале
	logger.Info("User logged in")

	// Нарушение 2: неанглийские буквы
	logger.Info("пользователь вошёл")

	// Нарушение 3: чувствительные данные
	logger.Info("user password reset")
	logger.Info("received token from client")
	logger.Info("auth failed for user")

	// Нарушение 4: спецсимволы
	logger.Info("success!")
	logger.Info("connection failed???")
	logger.Info("user admin logged in @ root")
	logger.Info("done ✅")

	// Кастомные нарушения
	logger.Debug("secreT")
	logger.Error("Credit Card")
}
