package main
import (
	"fmt"
)

func check_nilai(nilai int) int {
	if nilai >= 0 && nilai <= 100 {
		if nilai < 38 {
			return nilai
		}
		upper_limit := ((nilai / 5) + 1) * 5
		if (upper_limit - nilai) < 3 {
			return upper_limit
		} 
		return nilai
	}
	return -1
}

func print_nilai(result []int){
	for _, res := range result {
		if res < 0 {
			fmt.Println("Nilai di bawah 0 atau di atas 100")
		} else {
			fmt.Println(res)
		}
	}
}

func main(){
	var jumlah_mhs, nilai int
	var result []int
	fmt.Scanln(&jumlah_mhs)
	for jumlah_mhs > 0 && jumlah_mhs <= 60 {
		fmt.Scanln(&nilai)
		result = append(result, check_nilai(nilai))
		jumlah_mhs -= 1
	}
	print_nilai(result)
}
