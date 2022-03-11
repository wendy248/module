package module

import "fmt"

type Status func(string) bool

func Intro() {
	fmt.Println("\nProgram coffee shop sedang dijalankan...")
}

func Salam(a string) string {
	return "Hai " + a + ", selamat datang"
}

func Recent(favorit string) (float32, float32) {
	var a, b float32

	if favorit == "kopi" {
		a = 30000
	} else if favorit == "air" {
		a = 10000
	}

	b = a + (a * 10 / 100)
	return a, b
}

func Member(nama string, status Status) {
	if status(nama) {
		fmt.Println(nama, " merupakan seorang member")
	} else {
		fmt.Println(nama, " bukan member")
	}
}
