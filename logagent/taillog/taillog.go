package taillog

import (
	"github.com/hpcloud/tail"
)

var tails *tail.Tail

func Init(fileName string) (err error) {
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}
	tails, err = tail.TailFile(fileName, config)
	return err
}

func GetLine() chan *tail.Line {
	return tails.Lines
}
