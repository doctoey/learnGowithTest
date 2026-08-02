package main

import "testing"

func TestHello(t *testing.T) {
	// เพิ่มตัวแปร name เข้าไปในฟังก์ชัน Hello
	// name เก็บค่า "Toey" เพื่อส่งเข้าไปในฟังก์ชัน Hello
	t.Run("saying Hello to Toey", func(t *testing.T) {
		got := Hello("Toey", "")
		want := "Hello, Toey"

		// เรียกใช้ฟังก์ชัน assertCorrectMessage เพื่อตรวจสอบค่าที่ได้จากการ return
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Santi", "French")
		want := "Bonjour, Santi"
		assertCorrectMessage(t, got, want)
	})
}

// ตัวเล็ก private func
// t testing.TB - ใช้สำหรับ test
// @got - ค่าที่ได้จากการ return
// @want - ค่าที่คาดหวัง
// %q ใช้ในการแสดงผล string แบบมีเครื่องหมายคำพูดครอบ เช่น "Hello, world"
// t.Helper() - บอกว่าฟังก์ชันนี้เป็น helper function ทำให้ test ที่ล้มเหลวจะรายงานผลไปยังฟังก์ชันที่เรียกใช้ฟังก์ชันนี้
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
