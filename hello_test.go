package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		// t ย่อมาจาก testing
		// Errorf  ย่อมาจาก error format
		// got เก็บค่าที่ได้จากการ return
		// want เก็บค่าที่คาดหวัง
		// %q ใช้ในการแสดงผล string แบบมีเครื่องหมายคำพูดครอบ เช่น "Hello, world"
		// t.Errorf คือ การแจ้งเตือนว่า test fail พร้อมบอกรายละเอียด
		t.Errorf("got %q want %q", got, want)
	}
}
