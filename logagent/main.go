package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/config"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)

var iniConfig = new(config.Ini)

func main() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v/n", err)
		return
	}
	err = cfg.MapTo(iniConfig)

	err = taillog.Init(iniConfig.Taillog.FileName)
	if err != nil {
		fmt.Printf("init taillog failed :%v/n", err)
		return
	}

	ip := iniConfig.Kafka.Ip
	port := iniConfig.Kafka.Port

	address := []string{ip + ":" + port}

	err = kafka.Init(address)
	if err != nil {
		fmt.Printf("init kafka failed :%v/n", err)
		return
	}

	for {
		select {
		case data := <-taillog.GetLine():
			kafka.SendMsg(iniConfig.Kafka.Topic, data.Text)
		default:
			time.Sleep(time.Second)
		}

	}

}
