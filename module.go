package module

import "fmt"

type Menu struct {
	nama	string
	harga	int16
}

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

func Member(nama string, status bool) {
	if status {
		fmt.Println(nama, " merupakan seorang member")
	} else {
		fmt.Println(nama, " bukan member")
	}
}

func DaftarMenu(){
	menu1 := Menu{
		nama: "Kopi",
		harga: 30000,
	}

	menu2 := Menu{
		nama: "Air Mineral",
		harga: 10000,
	}

	menu3 := Menu{
		nama: "Susu",
		harga: 20000,
	}
	fmt.Println(menu1)
	fmt.Println(menu2)
	fmt.Println(menu3)
}
