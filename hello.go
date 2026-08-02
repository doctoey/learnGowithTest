package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello รับค่า string มา 1 ค่า แล้ว return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
// เพิ่ม language เพื่อเลือกภาษาที่จะใช้
func Hello(name, language string) string {
	// return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
	// เพิ่มเงื่อนไขว่าถ้า name เป็น "" ให้ name = "World"
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	// พิมพ์ค่าที่ return จากฟังก์ชัน Hello
	fmt.Println(Hello("Toey", ""))
}
