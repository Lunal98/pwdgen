package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Setup() {
	multi := zerolog.MultiLevelWriter(os.Stdout, rotate("pwdgen.log"))
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

}
func rotate(path string) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    200,
		MaxAge:     90,
		MaxBackups: 10,
		Compress:   true,
	}
}
