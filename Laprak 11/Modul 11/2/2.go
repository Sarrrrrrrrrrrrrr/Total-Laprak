package main

import "fmt"

func main() {
	var jam, harga float64
	var jenis_kendaraan string

	fmt.Print("Masukkan jenis kendaraan dan durasi parkir (jam): ")
	fmt.Scan(&jenis_kendaraan, &jam)

	switch jenis_kendaraan {
	case "Motor", "motor":
		if jam > 0 {
			if jam <= 1 {
				jam = 1
			}
			harga = jam * 2000
			fmt.Printf("Rp. %.0f\n", harga)
		} else {
			fmt.Println("Masukkan jam yang sesuai.")
		}
	case "Mobil", "mobil":
		if jam > 0 {
			if jam <= 1 {
				jam = 1
			}
			harga = jam * 5000
			fmt.Printf("Rp. %.0f\n", harga)
		} else {
			fmt.Println("Masukkan jam yang sesuai.")
		}
	case "Truk", "truk":
		if jam > 0 {
			if jam <= 1 {
				jam = 1
			}
			harga = jam * 8000
			fmt.Printf("Rp. %.0f\n", harga)
		} else {
			fmt.Println("Masukkan jam yang sesuai.")
		}
	default:
		fmt.Println("Masukkan jenis kendaraan dengan benar.")
	}
}
