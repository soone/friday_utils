package log

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func New() {
	rotateOptions := []rotatelogs.Option{
		rotatelogs.WithRotationTime(time.Hour * 24),
	}

	w, err := rotatelogs.New(path.Join("logs", "notion-%Y-%m-%d.log"), rotateOptions...)
	if err != nil {
		logrus.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}

	logFormatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	}

	logLevels := GetLogLevel(viper.GetString("output.log-level"))
	logrus.SetLevel(logLevels[0])

	logrus.AddHook(NewLocalHook(w, logFormatter, logLevels...))
}
