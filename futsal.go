package main

import "fmt"

//konstanta array
const NMAX = 100

//deklarasi tipe data	
type Lapangan struct {
	ID 		  	  int
    Nama          string
    Ukuran        string
    Jenis         string
    Harga         int
    Status        string
    Durasi        string
    JamBuka       string
    JamTutup      string
}
type Penyewa struct {
	ID 		int
	nama 	string
	noHP 	string
}
type Transaksi struct {
    NamaPenyewa 	string
	NamaLapangan 	string
	TanggalSewa 	Tanggal
	JamMulai 		string
	TotalHarga 		int
	Status 			string
}
type Tanggal struct {
	Hari 	int
	Bulan 	int
	Tahun 	int
}

type daftarLapangan [NMAX]Lapangan
type tabPenyewa 	[NMAX]Penyewa
type tabTransaksi 	[NMAX]Transaksi	
type TanggalSewa 	[NMAX]Tanggal

var nTransaksi, nLapangan, nPenyewa int

//data dummy
var daftarPenyewa = tabPenyewa{
    {1, "Nolan", "081234567890"},
    {2, "Kagura", "081345678901"},
    {3, "Suyou", "082156789012"},
    {4, "Gusion", "085767890123"},
    {5, "Aamon", "087878901234"},
	{6, "Benedetta", "089989012345"},
	{7, "Beatrix", "081234567890"},
}

func tambahPenyewa(T *tabPenyewa, n *int) {
	if *n < NMAX {
		// Input data
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&T[*n].ID)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&T[*n].nama)
		fmt.Print("Masukkan No HP: ")
		fmt.Scan(&T[*n].noHP)
		*n++

		fmt.Println("✅ Data penyewa berhasil ditambahkan!")
	} else {
		fmt.Println("❌ Kapasitas data penuh!")
	}
}

var lapangan daftarLapangan = daftarLapangan{
	{1, "LapanganA", "25x15", "Sintetis", 100000, "Tersedia", "2", "8AM", "10PM"},
	{2, "LapanganB", "25x15", "Vinyl", 110000, "Tersedia", "2", "8AM", "10PM"},
	{3, "LapanganC", "25x15", "Semen", 90000, "Tersedia", "1", "8AM", "10PM"},
	{4, "LapanganD", "25x15", "Sintetis", 120000, "Tidak_Tersedia", "3", "8AM", "10PM"},
	{5, "LapanganE", "25x15", "Vinyl", 130000, "Tersedia", "2", "8AM", "10PM"},
	{6, "LapanganF", "25x15", "Semen", 95000, "Tersedia", "1", "8AM", "10PM"},
	{7, "LapanganG", "25x15", "Sintetis", 140000, "Tersedia", "3", "8AM", "11PM"},
	{8, "LapanganH", "25x15", "Vinyl", 115000, "Tidak_Tersedia", "2", "8AM", "10PM"},
	{9, "LapanganI", "25x15", "Sintetis", 125000, "Tersedia", "2", "9AM", "11PM"},
	{10, "LapanganJ", "25x15", "Semen", 85000, "Tersedia", "1", "8AM", "9PM"},
}

func tambahLapangan(L *daftarLapangan, n *int) {
	if *n < NMAX {
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&L[*n].ID)
		fmt.Print("Nama Lapangan: ")
		fmt.Scan(&L[*n].Nama)
		fmt.Print("Ukuran: ")
		fmt.Scan(&L[*n].Ukuran)
		fmt.Print("Jenis: ")
		fmt.Scan(&L[*n].Jenis)
		fmt.Print("Harga: ")
		fmt.Scan(&L[*n].Harga)
		fmt.Print("Status: ")
		fmt.Scan(&L[*n].Status)
		fmt.Print("Durasi: ")
		fmt.Scan(&L[*n].Durasi)
		fmt.Print("Jam Buka: ")
		fmt.Scan(&L[*n].JamBuka)
		fmt.Print("Jam Tutup: ")
		fmt.Scan(&L[*n].JamTutup)
		*n++
		fmt.Println("✅ Data lapangan berhasil ditambahkan!")
	} else {
		fmt.Println("❌ Kapasitas data penuh!")
	}
}

var daftarTransaksi = tabTransaksi{
    {"Nolan", "Lapangan A", Tanggal{1, 6, 2026}, "8AM", 200000, "Lunas"},
    {"Kagura", "Lapangan B", Tanggal{1, 6, 2026}, "10AM", 220000, "Lunas"},
    {"Suyou", "Lapangan C", Tanggal{2, 6, 2026}, "9AM", 90000, "DP"},
    {"Gusion", "Lapangan A", Tanggal{2, 6, 2026}, "1PM", 360000, "Lunas"},
    {"Aamon", "Lapangan B", Tanggal{3, 6, 2026}, "8AM", 260000, "DP"},
    {"Benedetta", "Lapangan C", Tanggal{3, 6, 2026}, "11AM", 95000, "Lunas"},
    {"Beatrix", "Lapangan A", Tanggal{4, 6, 2026}, "2PM", 420000, "Lunas"},
}

func tambahTransaksi(T *tabTransaksi, n *int) {
	if *n < NMAX {
		// Input data
		fmt.Print("Masukkan Nama Penyewa: ")
		fmt.Scan(&T[*n].NamaPenyewa)
		fmt.Print("Masukkan Nama Lapangan: ")
		fmt.Scan(&T[*n].NamaLapangan)
		fmt.Print("Masukkan Tanggal Sewa (hari bulan tahun): ")
		fmt.Scan(&T[*n].TanggalSewa.Hari, &T[*n].TanggalSewa.Bulan, &T[*n].TanggalSewa.Tahun)
		fmt.Print("Masukkan Jam Mulai: ")
		fmt.Scan(&T[*n].JamMulai)
		fmt.Print("Masukkan Total Harga: ")
		fmt.Scan(&T[*n].TotalHarga)
		fmt.Print("Masukkan Status: ")
		fmt.Scan(&T[*n].Status)
		*n++

		fmt.Println("✅ Data transaksi berhasil ditambahkan!")
	} else {
		fmt.Println("❌ Kapasitas data penuh!")
	}
}

//tampil data lapangan
func tampilLapangan(lapangan [NMAX]Lapangan, n int) {
    fmt.Println("\n=== DATA LAPANGAN FUTSAL ===")

    // HEADER
    fmt.Printf("%-3s | %-12s | %-15s | %-10s | %-15s | %-15s | %-15s | %-8s | %-8s\n",
        "ID",
        "Nama",
        "Ukuran (Meter)",
        "Jenis",
        "Harga (Rp)",
        "Status",
        "Durasi (Jam)",
        "Buka",
        "Tutup")

    fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

    // DATA
    for i := 0; i < n; i++ {
        fmt.Printf("%-3d | %-12s | %-15s | %-10s | %-15d | %-15s | %-15s | %-8s | %-8s\n",
            lapangan[i].ID,
            lapangan[i].Nama,
            lapangan[i].Ukuran,
            lapangan[i].Jenis,
            lapangan[i].Harga,
            lapangan[i].Status,
            lapangan[i].Durasi,
            lapangan[i].JamBuka,
            lapangan[i].JamTutup)
    }
}

func tampilDaftarPenyewa(T [NMAX]Penyewa, n int) {
    fmt.Println("\n=== DATA PENYEWA FUTSAL ===")

    fmt.Printf("%-5s | %-15s | %-15s\n",
        "ID", "Nama Penyewa", "No HP")

    fmt.Println("-------------------------------------------")

    if n == 0 {
        fmt.Println("Belum ada data penyewa.")
        return
    }

    for i := 0; i < n; i++ {
        fmt.Printf("%-5d | %-15s | %-15s\n",
            T[i].ID,
            T[i].nama,
            T[i].noHP)
    }
}

func tampilDaftarTransaksi(T [NMAX]Transaksi, n int) {
	fmt.Println("\n==========================================================================================")
	fmt.Println("||                                  DATA TRANSAKSI FUTSAL                               ||")
	fmt.Println("==========================================================================================")
	
	// Header Tabel
	fmt.Printf("%-15s | %-15s | %-12s | %-8s | %-12s | %-10s\n",
		"Penyewa", "Lapangan", "Tanggal", "Jam", "Total (Rp)", "Status")
	fmt.Println("------------------------------------------------------------------------------------------")

	if n == 0 {
		fmt.Println("                            Belum ada data transaksi.                            ")
	} else {
		for i := 0; i < n; i++ {
			// Format tanggal
			tanggal := fmt.Sprintf("%02d/%02d/%d",
				T[i].TanggalSewa.Hari,
				T[i].TanggalSewa.Bulan,
				T[i].TanggalSewa.Tahun)

			// Menampilkan data
			fmt.Printf("%-15s | %-15s | %-12s | %-8s | %-12d | %-10s\n",
				T[i].NamaPenyewa,
				T[i].NamaLapangan,
				tanggal,
				T[i].JamMulai,
				T[i].TotalHarga,
				T[i].Status)
		}
	}
	fmt.Println("==========================================================================================")
}

func menuUtama() {
	fmt.Println("\n==================================================")
	fmt.Println("||        APLIKASI PEMESANAN LAPANGAN FUTSAL    ||")
	fmt.Println("==================================================")
	fmt.Println("1. Menu Lapangan")
	fmt.Println("2. Menu Data Penyewa")
	fmt.Println("3. Menu Transaksi")
	fmt.Println("4. Keluar")
	fmt.Println("--------------------------------------------------")
}

func menuLapangan() {
	fmt.Println("\n===============================================")
	fmt.Println("||           MENU DATA LAPANGAN FUTSAL       ||")
	fmt.Println("===============================================")
	fmt.Println("1. Tambah Data Lapangan")
	fmt.Println("2. Ubah Data Lapangan")
	fmt.Println("3. Hapus Data Lapangan")
	fmt.Println("4. Tampilkan Semua Data Lapangan")
	fmt.Println("5. Urutkan Data Lapangan (Berdasarkan Harga)")
	fmt.Println("6. Cari Data Lapangan")
	fmt.Println("7. Statistik Pendapatan")
	fmt.Println("8. Keluar")
	fmt.Println("-----------------------------------------------")
}

func menuPenyewa() {
	fmt.Println("\n==================================================")
	fmt.Println("||        MENU DATA PENYEWA FUTSAL              ||")
	fmt.Println("==================================================")
	fmt.Println("1. Tambah Data Penyewa")
	fmt.Println("2. Ubah Data Penyewa")
	fmt.Println("3. Hapus Data Penyewa")
	fmt.Println("4. Tampilkan Semua Data Penyewa")
	fmt.Println("5. Cari Data Penyewa")
	fmt.Println("6. Keluar")
	fmt.Println("--------------------------------------------------")
}

func menuTransaksi() {
	fmt.Println("\n==================================================")
	fmt.Println("||        MENU DATA TRANSAKSI FUTSAL            ||")
	fmt.Println("==================================================")
	fmt.Println("1. Tambah Data Transaksi")
	fmt.Println("2. Ubah Data Transaksi")
	fmt.Println("3. Hapus Data Transaksi")
	fmt.Println("4. Tampilkan Semua Data Transaksi")
	fmt.Println("5. Cari Data Transaksi")
	fmt.Println("6. Keluar")
	fmt.Println("--------------------------------------------------")
}

func main() {
    var pilihUtama, pilihMenuLap, pilihMenuPeny, pilihMenuTransaksi int

    // Flag untuk menjaga program tetap berjalan
    var aktif bool = true

	nLapangan = 10 // Sesuai dengan data dummy yang ada
	nPenyewa = 7 // Sesuai dengan data dummy yang ada
	nTransaksi = 7 // Sesuai dengan data dummy yang ada

    for aktif {
        menuUtama()
        fmt.Print("Pilih menu utama: ")
        fmt.Scan(&pilihUtama)

        // TINGKAT 1: Menu Utama
        if pilihUtama == 1 {
            menuLapangan()
            fmt.Print("Pilih menu lapangan: ")
            fmt.Scan(&pilihMenuLap)
            
            switch pilihMenuLap {
            case 1: 
                tambahLapangan(&lapangan, &nLapangan)
                tampilLapangan(lapangan, nLapangan)
            case 2: 
                ubahLapangan(&lapangan, nLapangan)
				tampilLapangan(lapangan, nLapangan)
			case 3:
                hapusLapangan(&lapangan, &nLapangan)
                tampilLapangan(lapangan, nLapangan)
			case 4:
				tampilLapangan(lapangan, nLapangan)
			case 5:
				urutkanLapanganByPrice(&lapangan, nLapangan)
				tampilLapangan(lapangan, nLapangan)
			case 6:
				cariLapangan(lapangan, nLapangan)
			case 7:
				fmt.Println("Statistik Pendapatan")
				statistik(lapangan, nPenyewa, nLapangan)
            case 8: 
                fmt.Println("Kembali ke menu utama")
            }
            
		// TINGKAT 2: Sub-Menu Penyewa
        } else if pilihUtama == 2 {
            menuPenyewa()
            fmt.Print("Pilih menu penyewa: ")
            fmt.Scan(&pilihMenuPeny)
            
            switch pilihMenuPeny {
            case 1: 
                tambahPenyewa(&daftarPenyewa, &nPenyewa)
                tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
            case 2:
				ubahPenyewa(&daftarPenyewa, nPenyewa)
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 3:
				hapusPenyewa(&daftarPenyewa, &nPenyewa)
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 4:
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 5:
				cariPenyewa(daftarPenyewa, nPenyewa)
			case 6:
				fmt.Println("Kembali ke menu utama")
			}  
		} else if pilihUtama == 3 {
			menuTransaksi()
			fmt.Print("Pilih menu transaksi: ")
            fmt.Scan(&pilihMenuTransaksi)
            
            switch pilihMenuTransaksi {
            case 1: 
                tambahTransaksi(&daftarTransaksi, &nTransaksi)
                tampilDaftarTransaksi(daftarTransaksi, nTransaksi)
            case 2:
				ubahTransaksi(&daftarTransaksi, nTransaksi)
				tampilDaftarTransaksi(daftarTransaksi, nTransaksi)
			case 3:
				hapusTransaksi(&daftarTransaksi, &nTransaksi)
				tampilDaftarTransaksi(daftarTransaksi, nTransaksi)
			case 4:
				tampilDaftarTransaksi(daftarTransaksi, nTransaksi)
			case 5:
				cariTransaksi(daftarTransaksi, nTransaksi)
			case 6:
				fmt.Println("Kembali ke menu utama")
			}  

        } else if pilihUtama == 4 {
            aktif = false // Keluar dari aplikasi
            fmt.Println("Terima kasih!")
        } else {
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func ubahTransaksi(T *tabTransaksi, n int) {
	var namaCari string
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan Nama Penyewa transaksi yang ingin diubah datanya: ")
	fmt.Scan(&namaCari)

	i = 0
	for i < n && !ketemu {
		if T[i].NamaPenyewa == namaCari {
			ketemu = true
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Printf("Data ditemukan: %s\n", T[i].NamaPenyewa)
		fmt.Print("Nama Penyewa: ")
		fmt.Scan(&T[i].NamaPenyewa)
		fmt.Print("Nama Lapangan: ")
		fmt.Scan(&T[i].NamaLapangan)
		fmt.Print("Tanggal Sewa: ")
		fmt.Scan(&T[i].TanggalSewa.Hari, &T[i].TanggalSewa.Bulan, &T[i].TanggalSewa.Tahun)
		fmt.Print("Jam Mulai: ")
		fmt.Scan(&T[i].JamMulai)
		fmt.Print("Total Harga: ")
		fmt.Scan(&T[i].TotalHarga)
		fmt.Print("Status: ")
		fmt.Scan(&T[i].Status)

		fmt.Println("✅ Data berhasil diperbarui!")
	} else {
		fmt.Println("❌ ID transaksi tidak ditemukan.")
	}
}

func ubahLapangan(L *daftarLapangan, n int) {
	var idCari int
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan ID lapangan yang ingin diubah datanya: ")
	fmt.Scan(&idCari)

	i = 0
	for i < n && !ketemu {
		if L[i].ID == idCari {
			ketemu = true
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Printf("Data ditemukan: %s\n", L[i].Nama)
		fmt.Print("Nama Lapangan: ")
		fmt.Scan(&L[i].Nama)
		fmt.Print("Ukuran: ")
		fmt.Scan(&L[i].Ukuran)
		fmt.Print("Jenis: ")
		fmt.Scan(&L[i].Jenis)
		fmt.Print("Harga: ")
		fmt.Scan(&L[i].Harga)
		fmt.Print("Status: ")
		fmt.Scan(&L[i].Status)
		fmt.Print("Durasi: ")
		fmt.Scan(&L[i].Durasi)
		fmt.Print("Jam Buka: ")
		fmt.Scan(&L[i].JamBuka)
		fmt.Print("Jam Tutup: ")
		fmt.Scan(&L[i].JamTutup)
		
		fmt.Println("✅ Data berhasil diperbarui!")
	} else {
		fmt.Println("❌ ID lapangan tidak ditemukan.")
	}
}

func ubahPenyewa(T *tabPenyewa, n int) {
	var idCari int
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan ID penyewa yang ingin diubah datanya: ")
	fmt.Scan(&idCari)

	i = 0
	for i < n && !ketemu {
		if T[i].ID == idCari {
			ketemu = true
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Printf("Data ditemukan: %s\n", T[i].nama)
		fmt.Print("Masukkan Nama baru: ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan No HP baru: ")
		fmt.Scan(&T[i].noHP)
		
		fmt.Println("✅ Data berhasil diperbarui!")
	} else {
		fmt.Println("❌ ID penyewa tidak ditemukan.")
	}
}

func hapusLapangan(L *daftarLapangan, n *int) {
	var idCari int
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan ID lapangan yang ingin dihapus: ")
	fmt.Scan(&idCari)

	i = 0
	for i < *n && !ketemu {
		if L[i].ID == idCari {
			ketemu = true // Loop akan berhenti karena kondisi !ketemu sudah salah
		} else {
			i++
		}
	}
	if ketemu {
		for j := i; j < *n-1; j++ {
			L[j] = L[j+1]
		}
		*n = *n - 1
		fmt.Println("✅ Data berhasil dihapus!")
	} else {
		fmt.Println("❌ ID lapangan tidak ditemukan.")
	}
}

func hapusTransaksi(T *tabTransaksi, n *int) {
	var namaCari string
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan Nama Penyewa transaksi yang ingin dihapus: ")
	fmt.Scan(&namaCari)

	i = 0
	for i < *n && !ketemu {
		if T[i].NamaPenyewa == namaCari {
			ketemu = true // Loop akan berhenti karena kondisi !ketemu sudah salah
		} else {
			i++
		}
	}
	if ketemu {
		for j := i; j < *n-1; j++ {
			T[j] = T[j+1]
		}
		*n = *n - 1
		fmt.Println("\n✅ Data berhasil dihapus!")
	} else {
		fmt.Println("❌ Nama Penyewa tidak ditemukan.")
	}
}

func hapusPenyewa(T *tabPenyewa, n *int) {
	var idCari int
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan ID penyewa yang ingin dihapus: ")
	fmt.Scan(&idCari)

	i = 0
	for i < *n && !ketemu {
		if T[i].ID == idCari {
			ketemu = true // Loop akan berhenti karena kondisi !ketemu sudah salah
		} else {
			i++
		}
	}
	if ketemu {
		for j := i; j < *n-1; j++ {
			T[j] = T[j+1]
		}
		*n = *n - 1
		fmt.Println("\n✅ Data berhasil dihapus!")
	} else {
		fmt.Println("❌ ID penyewa tidak ditemukan.")
	}
}

func urutkanLapanganByPrice(L *daftarLapangan, n int) {
    var i, j, idx int
    var temp Lapangan

    // Selection Sort
    i = 0
    for i < n-1 {
        idx = i
        j = i + 1
        
        for j < n {
            if L[j].Harga < L[idx].Harga {
                idx = j
            }
            j++
        }
        temp = L[idx]
        L[idx] = L[i]
        L[i] = temp
        i++
    }
    fmt.Println("\n✅ Data lapangan berhasil diurutkan berdasarkan Harga.")
}

func cariTransaksi(T [NMAX]Transaksi, n int) {
    var namaLapangan string
    var ditemukan bool = false
    var tanggal string

    fmt.Print("Masukkan Nama Lapangan yang dicari: ")
    fmt.Scan(&namaLapangan)

    fmt.Println("\n=== HASIL PENCARIAN TRANSAKSI ===")
    fmt.Printf("%-15s | %-15s | %-12s | %-8s | %-12s | %-10s\n",
        "Penyewa", "Lapangan", "Tanggal", "Jam", "Total (Rp)", "Status")
    fmt.Println("------------------------------------------------------------------------------------------")

    // Perulangan dari awal sampai akhir
    for i := 0; i < n; i++ {
        if T[i].NamaLapangan == namaLapangan {
            ditemukan = true
            
            // Format tanggal
            tanggal = fmt.Sprintf("%02d/%02d/%d",
                T[i].TanggalSewa.Hari,
                T[i].TanggalSewa.Bulan,
                T[i].TanggalSewa.Tahun)

            // Tampilkan baris yang cocok
            fmt.Printf("%-15s | %-15s | %-12s | %-8s | %-12d | %-10s\n",
                T[i].NamaPenyewa,
                T[i].NamaLapangan,
                tanggal,
                T[i].JamMulai,
                T[i].TotalHarga,
                T[i].Status)
        }
    }

    if !ditemukan {
        fmt.Println("❌ Tidak ada transaksi untuk lapangan tersebut.")
    }
    fmt.Println("==========================================================================================")
}
/*
func cariTransaksi(T tabTransaksi, n int) {
	//menampilkan semua transaksi berdasarkan jenis lapangan yang disewa 
	var i int
	var namaLapangan string

	fmt.Print("Masukkan Nama Lapangan yang dicari: ")
	fmt.Scan(&namaLapangan)

	i = 0
	for i < n  {
		if T[i].NamaLapangan == namaLapangan {
			fmt.Print
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Println("\n✅ Data ditemukan!")
		fmt.Printf("ID    : %d\n", T[i].IDTransaksi)
		fmt.Printf("Nama  : %s\n", T[i].nama)
		fmt.Printf("No HP : %s\n", T[i].noHP)
	} else {
		fmt.Println("❌ ID transaksi tidak ditemukan.")
	}
}
*/
func cariPenyewa(T tabPenyewa, n int) {
	var i int
	var NoPenyewa string
	var ketemu bool = false

	fmt.Print("Masukkan Nomor HP penyewa yang dicari: ")
	fmt.Scan(&NoPenyewa)

	i = 0
	for i < n && !ketemu {
		if T[i].noHP == NoPenyewa {
			ketemu = true
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Println("\n✅ Data ditemukan!")
		fmt.Printf("ID    : %d\n", T[i].ID)
		fmt.Printf("Nama  : %s\n", T[i].nama)
		fmt.Printf("No HP : %s\n", T[i].noHP)
	} else {
		fmt.Println("❌ Nomor HP penyewa tidak ditemukan.")
	}
}

func cariLapangan(L daftarLapangan, n int) {
	var i int
	var ketemu bool = false
	var IDLapangan int

	fmt.Print("Masukkan ID lapangan yang dicari: ")
	fmt.Scan(&IDLapangan)

	i = 0
	for i < n && !ketemu {
		if L[i].ID == IDLapangan {
			ketemu = true
		} else {
			i++
		}
	}

	if ketemu {
		fmt.Println("✅ Data ditemukan!")
		fmt.Printf("Nama  : %s\n", L[i].Nama)
		fmt.Printf("Ukuran: %s\n", L[i].Ukuran)
		fmt.Printf("Jenis : %s\n", L[i].Jenis)
		fmt.Printf("Harga : %d\n", L[i].Harga)
		fmt.Printf("Status: %s\n", L[i].Status)
		fmt.Printf("Durasi: %s\n", L[i].Durasi)
		fmt.Printf("Jam Buka: %s\n", L[i].JamBuka)
		fmt.Printf("Jam Tutup: %s\n", L[i].JamTutup)
	} else {
		fmt.Println("❌ ID lapangan tidak ditemukan.")
	}
}

func statistik(L daftarLapangan, nPenyewa int, nLapangan int) {
	var tersedia, tidakTersedia int
	var totalHarga int
	var i int

	tersedia = 0
	tidakTersedia = 0

	for i := 0; i < nLapangan; i++ {
		if L[i].Status == "Tersedia" {
			tersedia++
		} else {
			tidakTersedia++
		}
	}

	// total pendapatan
	totalHarga = 0
	for i = 0; i < nLapangan; i++ {
		totalHarga += L[i].Harga
	}

	// Tampilan Statistik
	fmt.Println("\n==================================================")
	fmt.Println("||             STATISTIK DATA FUTSAL            ||")
	fmt.Println("==================================================")
	fmt.Printf("Total Penyewa         : %d orang\n", nPenyewa)
	fmt.Printf("Lapangan Tersedia     : %d lapangan\n", tersedia)
	fmt.Printf("Lapangan Dipesan      : %d lapangan\n", tidakTersedia)
	fmt.Printf("Total Nilai Aset/Harga: Rp %d\n", totalHarga)
	fmt.Println("==================================================")
}

/*
NOTES PENTING!
1. Nama lapangannya harus satu kata, contoh: "LapanganA", "LapanganB", dst. Jangan menggunakan spasi.
2. Program belum bisa menentukan apakah lapangan sudah dipesan atau belum, jadi status lapangan 
   masih diinput manual oleh user.
3. Fitur statistik pendapatan masih menghitung total harga dari semua lapangan, 
   bukan berdasarkan transaksi yang sudah terjadi. Jadi ini lebih ke total nilai aset lapangan.
   bukan hanya total keseluruhan. Jadi nanti harus ada input tanggal untuk menghitung pendapatan per minggu.

PROGRESS KESELURUHAN: 75%
*/
