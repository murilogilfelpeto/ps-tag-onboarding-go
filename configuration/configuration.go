package configuration

var (
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}
