# Learn Go with Tests 🧪🤖

โปรเจกต์สรุปการเรียนรู้ภาษา **Go** ผ่านแนวคิด **TDD (Test-Driven Development)** อ้างอิงจากบทเรียน [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

---

## 📌 โครงสร้างบทเรียน (Chapters & Packages)

| บทเรียน (Chapter) | แพ็กเกจ (Package) | หัวข้อสำคัญที่ได้เรียนรู้ (Key Concepts) |
| :--- | :--- | :--- |
| **Hello, World** | [`helloworld`](file:///Users/doctoey/Documents/code/learnGowithTest/helloworld) | การเขียนฟังก์ชันพื้นฐาน, Subtests (`t.Run`), Constants, HTTP Server (`net/http`) |
| **Integers** | [`integers`](file:///Users/doctoey/Documents/code/learnGowithTest/integers) | การเขียนฟังก์ชันบวกเลข, Examples (`// Output:`) สำหรับ Go Documentation |
| **Iteration** | [`repeat`](file:///Users/doctoey/Documents/code/learnGowithTest/repeat), [`interation`](file:///Users/doctoey/Documents/code/learnGowithTest/interation) | การวนลูป (`for`), Benchmark Testing (`testing.B`) |
| **Structs, Methods & Interfaces** | [`shape`](file:///Users/doctoey/Documents/code/learnGowithTest/shape) | Structs, Methods, Interfaces, Table Driven Tests |
| **Pointers & Errors** | [`wallet`](file:///Users/doctoey/Documents/code/learnGowithTest/wallet) | Pointers (`*`, `&`), Custom Error (`errors.New`), Nil check |
| **Maps** | [`dictionary`](file:///Users/doctoey/Documents/code/learnGowithTest/dictionary) | Maps, Custom Error Types (`const` errors), CRUD operations (Search, Add, Update, Delete) |
| **Dependency Injection** | [`di`](file:///Users/doctoey/Documents/code/learnGowithTest/di) | Dependency Injection, `io.Writer`, `bytes.Buffer`, `os.Stdout`, การทดสอบการ Output |
| **Mocking** | [`mocking`](file:///Users/doctoey/Documents/code/learnGowithTest/mocking) | Mocking, Spies (การสร้าง Spy เพื่อบันทึกลำดับการทำงานและการเรียกใช้งาน), Interface Isolation |

---

## 📁 Project Structure

```text
learnGowithTest/
├── di/
│   ├── di.go
│   └── di_test.go
├── dictionary/
│   ├── dictionary.go
│   └── dictionary_test.go
├── helloworld/
│   ├── hello.go
│   └── hello_test.go
├── integers/
│   ├── adder.go
│   └── adder_test.go
├── interation/
│   ├── repeat.go
│   └── repeat_test.go
├── mocking/
│   ├── countdown.go
│   └── countdown_test.go
├── repeat/
│   ├── repeat.go
│   └── repeat_test.go
├── shape/
│   ├── shape.go
│   └── shape_test.go
├── wallet/
│   ├── wallet.go
│   └── wallet_test.go
├── go.mod
└── README.md
```

---

## 🚀 การรัน Unit Tests

### รันทุก Test ทั้งโปรเจกต์
```bash
go test ./...
```

### รัน Test พร้อมแสดงรายละเอียด (Verbose)
```bash
go test -v ./...
```

### รันเฉพาะบาง Package (เช่น `mocking`)
```bash
go test -v ./mocking
```

---

## 🛠️ เครื่องมือช่วยเช็คโค้ด (Tools)

- **`errcheck`**: เครื่องมือตรวจเช็คการละทิ้งค่า error (Unchecked errors)
  ```bash
  go install github.com/kisielk/errcheck@latest
  errcheck .
  ```

---

## 📝 บันทึก TDD Workflow (วงจรพัฒนาด้วย Test)
1. 🔴 **RED**: เขียน Test ก่อนเสมอ และรันให้ล้มเหลวเพื่อยืนยันพฤติกรรมที่คาดหวัง
2. 🟢 **GREEN**: เขียนโค้ดขั้นต่ำที่สุดเพื่อให้ Test ผ่าน
3. 🔵 **REFACTOR**: ปรับแต่งโค้ดให้สะอาด ปลอดภัย และขยายผลได้ง่ายขึ้น โดยมี Test คอยการันตีความถูกต้อง
