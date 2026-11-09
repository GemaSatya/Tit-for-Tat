package main

import "fmt"

func main() {
	
	poin := 0

	pilihan := 0
	fmt.Println("Game started")
	for true{
		lawan := pilihan
		fmt.Println("Masukkan pilihan:")
		fmt.Scanln(&pilihan)
		if poin == 0{
			lawan = pilihan
		}
		if pilihan == 1 && lawan == 1 {
			fmt.Println("Sama")
			poin += 3
			} else if pilihan == 2 && lawan == 2{
			fmt.Println("Sama")
			poin += 3
		}else{
			fmt.Println("Beda")
			poin -=1
		}
		
		fmt.Println("Poin =", poin)
		fmt.Println("pilihan anda:", pilihan)
		fmt.Println("pilihan lawan:", lawan)
	}
}