package main
import "fmt"
// DEKLARASI TIPE BENTUKAN BERNAMA BOOK BERSERTA VARIABEL-VARIABELNYA
// Declaration of a user-defined type 'Book' and its variables
type book struct {
judul_buku string
genre string
tahun_rilis int
rating float64
}

func main() {

// DEKLARASI DUA VARIABEL BERTIPE BNETUKAN
// Declaration of two variables of a user-defined type
var buku1, buku2 book
// INPUT
fmt.Scan(&buku1.judul_buku, &buku2.judul_buku)
fmt.Scan(&buku1.genre, &buku2.genre)
fmt.Scan(&buku1.tahun_rilis, &buku2.tahun_rilis)
fmt.Scan(&buku1.rating, &buku2.rating)
// OUTPUT
fmt.Println("--- Book 1 ---")
fmt.Println(buku1.judul_buku)
fmt.Println(buku1.genre)
fmt.Println(buku1.tahun_rilis)
fmt.Println(buku1.rating)
fmt.Println("--- Book 2 ---")
fmt.Println(buku2.judul_buku)
fmt.Println(buku2.genre)
fmt.Println(buku2.tahun_rilis)
fmt.Println(buku2.rating)
}