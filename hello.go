package main

import "fmt"

// Hello รับค่า string มา 1 ค่า แล้ว return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
func Hello(name string) string {
	// return ค่า string ที่ต่อท้ายด้วย "Hello, " + name
	return "Hello, " + name
}

func main() {
	// พิมพ์ค่าที่ return จากฟังก์ชัน Hello
	fmt.Println(Hello("Toey"))
}
