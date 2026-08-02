package interation

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersTosum ...[]int) []int {
	// lengthOfNumbers เก็บจำนวนของตัวแปร numbersTosum ที่ส่งเข้ามา
	// lengthOfNumbers := len(numbersTosum)
	// // make คือ การจองพื้นที่ในหน่วยความจำ ให้กับตัวแปร sums
	// // โดยกำหนดให้มีขนาดเท่ากับจำนวนของตัวแปร numbersTosum
	// // ไม่ต้องใส่ค่า default เพราะเราจะใช้ค่าเริ่มต้นคือ 0 อยู่แล้ว
	// // capacity คือ จำนวนครั้งที่เราสามารถเพิ่มค่าได้โดยไม่ต้องจองพื้นที่ใหม่
	// // ในที่นี้เราจองพื้นที่ไว้ 2 เท่าของจำนวนของตัวแปร numbersTosum (lengthOfNumbers*2)
	// // cap กับ length ต่างกันยังไง
	// // length คือ ขนาดของตัวแปร
	// // capacity คือ ขนาดของ array ที่จองไว้
	// // ถ้าเกินน่าจะ panic

	// sums := make([]int, lengthOfNumbers, lengthOfNumbers*2)

	// for i, numbers := range numbersTosum {
	// 	sums[i] = Sum(numbers)
	// }
	// return sums

	var sums []int
	for _, numbers := range numbersTosum {
		sums = append(sums, Sum(numbers))
	}

	return sums

}

func SumAllTails(numbersTosum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersTosum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// [1:] คือ การ slice array โดยการตัดตัวแรกออก
			// ถึง ตัวสุดท้าย เพราะไม่ระบุจำนวนสุดท้าย ก็ถึงตัวสุดท้าย
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
