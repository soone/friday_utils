package log

import (
	"fmt"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func New(name string) {
	rotateOptions := []rotatelogs.Option{
		rotatelogs.WithRotationTime(time.Hour * 24),
	}

	w, err := rotatelogs.New(path.Join("logs", fmt.Sprintf("%s-%%Y-%%m-%%d.log", name)), rotateOptions...)
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
