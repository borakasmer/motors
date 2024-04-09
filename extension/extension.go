package extension

import "strings"

type String struct {
	Value string
}

func (str *String) Slice() string {
	cleaned := strings.ReplaceAll(str.Value, "\n", "")
	// Replace spaces with an empty string
	cleaned = strings.TrimSpace(cleaned)
	return cleaned
}

var MotorsMap = map[string]string{
	"honda-cbr-650-r":      "Cbr650R",
	"kawasaki-ninja-zx-6r": "Zx6R",
	"honda-cbr-600-rr":     "Cbr600RR",
	"yamaha-yzf-r25-abs":   "R25",
}
var Motors = newMotorRegistry()

func newMotorRegistry() *motorsEnum {
	return &motorsEnum{
		Cbr650R:  "honda-cbr-650-r",
		Zx6R:     "kawasaki-ninja-zx-6r",
		Cbr600RR: "honda-cbr-600-rr",
		R25:      "yamaha-yzf-r25-abs",
	}
}

type motorsEnum struct {
	Cbr650R  string
	Zx6R     string
	Cbr600RR string
	R25      string
}
