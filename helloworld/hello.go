package main

import (
	// "log"
	"net/http"
	"os"

	"github.com/doctoey/learnGowithTest/di"
	"github.com/doctoey/learnGowithTest/mocking"
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

	return greetingPrefix(language) + name
}

// greetingPrefix รับค่า language string แล้ว return ค่า prefix string
// @param language string - ภาษาที่จะใช้
// @return prefix string - ค่า prefix string ที่ใช้
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// func main() {
// 	// พิมพ์ค่าที่ return จากฟังก์ชัน Hello
// 	fmt.Println(Hello("Toey", ""))
// }

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

// func main() {
// 	di.Greet(os.Stdout, "Elodie")
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

// }

func main() {
	mocking.Countdown(os.Stdout)
}
