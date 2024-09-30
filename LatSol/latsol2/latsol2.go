package main
import "fmt"
func main() {
	//input
	var name, class, region, nim, word string
	fmt.Print("masukkan nama panggilan anda: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan kelas anda: ")
	fmt.Scan(&class)
	fmt.Print("Masukkan NIM anda: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan tempat lahir anda: ")
	fmt.Scan(&region)
	fmt.Print("Satu kata untuk TELKOM UNIVERSITY: ")
	fmt.Scan(&word)

	//output
	fmt.Println()
	fmt.Println("---RESUME SINGKAT---")
	fmt.Printf("Halo semuanya, perkenalkan nama saya %s dengan NIM %s.\n", name, nim)
	fmt.Printf("Saya berasal dari kelas %s, Lalu saya lahir di %s. Satu kata untuk telkom adalah %s", class, region, word)
}