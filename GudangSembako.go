package main
import "fmt"
const NMAX = 1000

type barang struct {
	kode, nama string
	stok, harga int
}

type gudang [NMAX]barang

func main() {
	var data gudang
	var n, pilihan int

	for {
		menu()
		fmt.Print("  Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("\n")
			tambahBarang(&data, &n)
		case 2:
			tampilBarang(data, n)
		case 3:
			fmt.Print("\n")
			cariBarangSequential(data, n)
		case 4:
			fmt.Print("\n")
			insertionSortKodeAsc(&data, n)
			cariBarangBinary(data, n)
		case 5:
			editBarang(&data, n)
		case 6:
			hapusBarang(&data, &n)
		case 7:
			barangMasuk(&data, n)
		case 8:
			barangKeluar(&data, n)
		case 9:
			menuSorting(&data, n)
		case 0:
			fmt.Print("\n")
			fmt.Println("Program selesai.\n")
			return
		default:
			fmt.Print("\n")
			fmt.Println("Menu tidak tersedia\n")
		}
	}
}

func menu() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║	    GUDANG SEMBAKO             ║")
	fmt.Println("╟══════════════════════════════════════╢")
	fmt.Println("║ 1. Tambah Barang                     ║")
	fmt.Println("║ 2. Tampilkan Barang                  ║")
	fmt.Println("║ 3. Cari Barang Sequential            ║")
	fmt.Println("║ 4. Cari Barang Binary                ║")
	fmt.Println("║ 5. Edit Barang                       ║")
	fmt.Println("║ 6. Hapus Barang                      ║")
	fmt.Println("║ 7. Barang Masuk                      ║")
	fmt.Println("║ 8. Barang Keluar                     ║")
	fmt.Println("║ 9. Sorting                           ║")
	fmt.Println("║ 0. Keluar                            ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func tambahBarang(g *gudang, n *int) {
	var jumlah int

	fmt.Print("Berapa barang yang ingin ditambahkan? ")
	fmt.Scan(&jumlah)

	for i := 1; i <= jumlah; i++ {
		fmt.Printf("\nBarang ke-%d\n", i)

		fmt.Print("Kode  : ")
		fmt.Scan(&g[*n].kode)

		fmt.Print("Nama  : ")
		fmt.Scan(&g[*n].nama)

		fmt.Print("Stok  : ")
		fmt.Scan(&g[*n].stok)

		fmt.Print("Harga : ")
		fmt.Scan(&g[*n].harga)

		*n++

		fmt.Print("\n")
	}
}

func tampilBarang(g gudang, n int) {
	fmt.Println("\n============== DATA BARANG ===============")

	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s | %s | Stok: %d | Harga: %d\n",
			i+1,
			g[i].kode,
			g[i].nama,
			g[i].stok,
			g[i].harga)
	}
	fmt.Println("\n")
}

func cariBarangSequential(g gudang, n int) {
	var kode string

	fmt.Print("Masukkan kode barang: ")
	fmt.Scan(&kode)

	idx := sequentialSearch(g, n, kode)

	if idx != -1 {
		fmt.Println("Barang ditemukan")
		fmt.Println(g[idx],"\n")
	} else {
		fmt.Println("Barang tidak ditemukan\n")
	}
}

func cariBarangBinary(g gudang, n int) {
	var kode string

	fmt.Print("Masukkan kode barang: ")
	fmt.Scan(&kode)

	idx := binarySearch(g, n, kode)

	if idx != -1 {
		fmt.Println("Barang ditemukan")
		fmt.Println(g[idx],"\n")
	} else {
		fmt.Println("Barang tidak ditemukan\n")
	}
}

func editBarang(g *gudang, n int) {
	var kode string

	fmt.Print("Kode barang: ")
	fmt.Scan(&kode)
	idx := sequentialSearch(*g, n, kode)

	if idx != -1 {
		fmt.Print("Kode baru : ")
		fmt.Scan(&g[idx].kode)

		fmt.Print("Nama baru : ")
		fmt.Scan(&g[idx].nama)

		fmt.Print("Stok baru : ")
		fmt.Scan(&g[idx].stok)

		fmt.Print("Harga baru : ")
		fmt.Scan(&g[idx].harga)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusBarang(g *gudang, n *int) {
	var kode string

	fmt.Print("Kode barang yang dihapus: ")
	fmt.Scan(&kode)
	idx := sequentialSearch(*g, *n, kode)

	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			g[i] = g[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func barangMasuk(g *gudang, n int) {
	var kode string
	var jumlah int

	fmt.Print("Kode barang : ")
	fmt.Scan(&kode)
	fmt.Print("Jumlah masuk : ")
	fmt.Scan(&jumlah)
	idx := sequentialSearch(*g, n, kode)

	if idx != -1 {
		g[idx].stok += jumlah
		fmt.Println("Stok berhasil ditambah")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func barangKeluar(g *gudang, n int) {
	var kode string
	var jumlah int

	fmt.Print("Kode barang : ")
	fmt.Scan(&kode)
	fmt.Print("Jumlah keluar : ")
	fmt.Scan(&jumlah)
	idx := sequentialSearch(*g, n, kode)

	if idx != -1 {
		if jumlah <= g[idx].stok {
			g[idx].stok -= jumlah
			fmt.Println("Stok berhasil dikurangi")
		} else {
			fmt.Println("Stok tidak cukup")
		}
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func menuSorting(g *gudang, n int) {
	var pilih int

	fmt.Println("\n===== SORTING =====")
	fmt.Println("1. Stok Ascending")
	fmt.Println("2. Harga Descending")
	fmt.Println("3. Nama Ascending")
	fmt.Println("4. Nama Descending")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		selectionSortStokAsc(g, n)
	case 2:
		selectionSortHargaDesc(g, n)
	case 3:
		insertionSortNamaAsc(g, n)
	case 4:
		insertionSortNamaDesc(g, n)
	}
}

func sequentialSearch(g gudang, n int, x string) int {
	for i := 0; i < n; i++ {
		if g[i].kode == x {
			return i
		}
	}
	return -1
}

func binarySearch(g gudang, n int, x string) int {

	left := 0
	right := n - 1

	for left <= right {

		mid := (left + right) / 2

		if g[mid].kode == x {
			return mid
		}

		if x < g[mid].kode {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func selectionSortStokAsc(g *gudang, n int) {
	var pass, idx, i int

	for pass = 0; pass < n-1; pass++ {
		idx = pass

		for i = pass + 1; i < n; i++ {
			if g[i].stok < g[idx].stok {
				idx = i
			}
		}

		g[pass], g[idx] = g[idx], g[pass]
	}
}

func selectionSortHargaDesc(g *gudang, n int) {
	var pass, idx, i int

	for pass = 0; pass < n-1; pass++ {
		idx = pass

		for i = pass + 1; i < n; i++ {
			if g[i].harga > g[idx].harga {
				idx = i
			}
		}

		g[pass], g[idx] = g[idx], g[pass]
	}
}

func insertionSortNamaAsc(g *gudang, n int) {
	var pass int
	var temp barang

	for pass = 1; pass < n; pass++ {

		temp = g[pass]
		i := pass - 1

		for i >= 0 && g[i].nama > temp.nama {
			g[i+1] = g[i]
			i--
		}

		g[i+1] = temp
	}
}

func insertionSortNamaDesc(g *gudang, n int) {
	var pass int
	var temp barang

	for pass = 1; pass < n; pass++ {

		temp = g[pass]
		i := pass - 1

		for i >= 0 && g[i].nama < temp.nama {
			g[i+1] = g[i]
			i--
		}

		g[i+1] = temp
	}
}

func insertionSortKodeAsc(g *gudang, n int) {
	var pass int
	var temp barang

	for pass = 1; pass < n; pass++ {
		temp = g[pass]
		i := pass - 1
		for i >= 0 && g[i].kode > temp.kode {
			g[i+1] = g[i]
			i--
		}
		g[i+1] = temp
	}
}
