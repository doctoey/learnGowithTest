package shape

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// interface ชี้ไปที่ method ใด ๆ ก็ได้ที่มี Area() float64
// เหมือน abstract class ใน java
// Golang ไม่มี abstract class
// Golang ไม่มี inheritance
// Golang ใช้ interface แทน
type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return (rectangle.Height * rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

// method ของ circle
// () คือ receiver คือตัวแปร c ชี้ไปที่ Circle
// Area() คือ method
// ถ้าเป็นชื่อเดียวกับ function ตัวบนจะเรียก method ไม่ได้
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}
