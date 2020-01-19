package main

import (
	"flag"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

func main() {
	flag.Parse()
	valueStr := flag.Arg(0)
	valueLen := len(valueStr)

	if (valueLen == 8 || valueLen == 12) && strings.HasSuffix(valueStr, "=") {
		base64string := valueStr
		decodeValue, _ := base64.StdEncoding.DecodeString(base64string)
		if valueLen == 8 {
			// handle float32 bits
			bits := binary.LittleEndian.Uint32(decodeValue)
			float := math.Float32frombits(bits)
			fmt.Println(float)
		} else {
			// handle float64 bits
			bits := binary.LittleEndian.Uint64(decodeValue)
			float := math.Float64frombits(bits)
			fmt.Println(float)
		}
	}
}