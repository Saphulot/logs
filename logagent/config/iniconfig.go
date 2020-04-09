package config

type Ini struct {
	Kafka   Kafka
	Taillog Taillog
}

type Kafka struct {
	Ip    string `ini:"ip"`
	Port  string `ini:"port"`
	Topic string `ini:"topic"`
}

type Taillog struct {
	FileName string `ini:"filename"`
}
