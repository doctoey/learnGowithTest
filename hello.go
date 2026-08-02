package main

import (
	"fmt"
)

const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello รับค่า string มา 1 ค่า แล้ว return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
// เพิ่ม language เพื่อเลือกภาษาที่จะใช้
func Hello(name, language string) string {
	// return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
	// เพิ่มเงื่อนไขว่าถ้า name เป็น "" ให้ name = "World"
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	// พิมพ์ค่าที่ return จากฟังก์ชัน Hello
	fmt.Println(Hello("Toey", ""))
}
