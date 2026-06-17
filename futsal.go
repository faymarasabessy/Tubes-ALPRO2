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
	JamTersedia   string
	Status 		  string
}
type Penyewa struct {
	ID 		int
	nama 	string
	noHP 	string
}
type Transaksi struct {
    NamaPenyewa 	string
	NomorPenyewa 	string
	NamaLapangan 	string
	TanggalSewa 	Tanggal
	JamSewa			string
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

var nTransaksi, nPenyewa, nLapangan int

//data dummy
var daftarPenyewa = tabPenyewa{
    {1, "Nolan", "081234567890"},
    {2, "Kagura", "081345678901"},
    {3, "Suyou", "082156789012"},
    {4, "Gusion", "085767890123"},
	{5, "Aamon", "085782749283"},
	{6, "Hirara", "089521159807"},
}

var lapangan daftarLapangan = daftarLapangan{
	//Lapangan A
    {1, "Lapangan_A", "25x15", "Sintetis", 200000, "08.00-10.00", "Tidak_Tersedia"},
    {2, "Lapangan_A", "25x15", "Sintetis", 200000, "14.00-16.00", "Tersedia"},
	//Lapangan B
    {3, "Lapangan_B", "25x15", "Vinyl", 250000, "08.00-10.00", "Tidak_Tersedia"},
    {4, "Lapangan_B", "25x15", "Vinyl", 250000, "12.00-14.00", "Tersedia"},
    {5, "Lapangan_B", "25x15", "Vinyl", 250000, "18.00-20.00", "Tersedia"},
	//Lapangan C
    {6, "Lapangan_C", "30x25", "Semen", 150000, "09.00-11.00", "Tersedia"},
    {7, "Lapangan_C", "30x25", "Semen", 150000, "19.00-21.00", "Tidak_Tersedia"},
	//Lapangan D
    {8, "Lapangan_D", "24x14", "Sintetis", 180000, "10.00-12.00", "Tidak_Tersedia"},
    {9, "Lapangan_D", "24x14", "Sintetis", 180000, "15.00-17.00", "Tersedia"},
	//Lapangan E
    {10, "Lapangan_E", "25x15", "Vinyl", 220000, "13.00-15.00", "Tersedia"},
    {11, "Lapangan_E", "25x15", "Vinyl", 220000, "15.00-17.00", "Tersedia"},
	//Lapangan F
    {12, "Lapangan_F", "26x16", "Semen", 160000, "08.00-10.00", "Tidak_Tersedia"},
}

var daftarTransaksi = tabTransaksi{
    {"Nolan", "081234567890", "Lapangan_A", Tanggal{1, 6, 2026}, "08.00-10.00"},
    {"Kagura", "081345678901", "Lapangan_B", Tanggal{1, 6, 2026}, "10.00-12.00"},
    {"Suyou", "082156789012", "Lapangan_C", Tanggal{2, 6, 2026}, "19.00-21.00"},
    {"Gusion", "085767890123", "Lapangan_D", Tanggal{2, 6, 2026}, "10.00-12.00"},
	{"Aamon", "085782749283", "Lapangan_F", Tanggal{5, 6, 2026}, "08.00-10.00"},
	{"Hirara", "085782749283", "Lapangan_D", Tanggal{5, 6, 2026}, "15.00-17.00"},
}

func tambahLapangan(L *daftarLapangan, n *int) {
	if *n < NMAX {
		var idLap int
		// Input data
		fmt.Print("Masukkan ID: ")
		fmt.Scan(&idLap)
		idExist := false
		for i := 0; i < *n && !idExist; i++ {
			if L[i].ID == idLap {
				idExist = true
			}
		}

		if idExist {
			fmt.Println("❌ ID Sudah terdaftar! Gagal menambahkan data.")
		} else {
			L[*n].ID = idLap
			fmt.Print("Nama Lapangan: ")
			fmt.Scan(&L[*n].Nama)
			fmt.Print("Ukuran: ")
			fmt.Scan(&L[*n].Ukuran)
			fmt.Print("Jenis: ")
			fmt.Scan(&L[*n].Jenis)
			fmt.Print("Harga: ")
			fmt.Scan(&L[*n].Harga)
			fmt.Print("Jam Sewa (ex. 08.00-10.00): ")
			fmt.Scan(&L[*n].JamTersedia)
			fmt.Print("Status: ")
			fmt.Scan(&L[*n].Status)
			*n++

			fmt.Println("✅ Data lapangan berhasil ditambahkan!")
		}
	} else {
		fmt.Println("❌ Kapasitas data penuh!")
	}
}

func tambahTransaksi(T *tabTransaksi, n *int) {
	var ada bool
	if *n < NMAX {
		var tFj Transaksi
		// Input data
		fmt.Print("Masukkan Nama Penyewa: ")
		fmt.Scan(&tFj.NamaPenyewa)
		fmt.Print("Masukkan Nomor HP Penyewa: ")
		fmt.Scan(&tFj.NomorPenyewa)
		fmt.Print("Masukkan Nama Lapangan: ")
		fmt.Scan(&tFj.NamaLapangan)
		fmt.Print("Masukkan Tanggal Sewa (hari bulan tahun): ")
		fmt.Scan(&tFj.TanggalSewa.Hari, &tFj.TanggalSewa.Bulan, &tFj.TanggalSewa.Tahun)
		fmt.Print("Masukkan Jam Sewa (ex. 08.00-10.00): ")
		fmt.Scan(&tFj.JamSewa) 
		
		ada = false
		for i := 0; i < *n && !ada; i++ {
			if T[i].NamaLapangan == tFj.NamaLapangan && 
			T[i].TanggalSewa.Hari == tFj.TanggalSewa.Hari && 
			T[i].TanggalSewa.Bulan == tFj.TanggalSewa.Bulan &&
			T[i].TanggalSewa.Tahun == tFj.TanggalSewa.Tahun &&
			T[i].JamSewa == tFj.JamSewa {
				ada = true
			}
		}
		if ada {
			fmt.Printf("❌ Lapangan %s pada jam tersebut sudah disewa, silahkan pilih jadwal lain! :D\n", tFj.NamaLapangan)
		} else {
			T[*n] = tFj
			*n++
			fmt.Println("✅ Data transaksi berhasil ditambahkan!")
			
			for i := 0; i < nLapangan; i++ {
				if lapangan[i].Nama == tFj.NamaLapangan && lapangan[i].JamTersedia == tFj.JamSewa {
					lapangan[i].Status = "Tidak_Tersedia"
				}
			}

			penyewaExist := false 
			for i := 0; i < nPenyewa && !penyewaExist; i++ {
				if daftarPenyewa[i].noHP == tFj.NomorPenyewa {
					penyewaExist = true
				}
			}

			if !penyewaExist && nPenyewa < NMAX {
				var newID int
				if nPenyewa > 0 {
					newID = daftarPenyewa[nPenyewa-1].ID + 1
				} else {
					newID = 1
				}
				daftarPenyewa[nPenyewa].ID = newID
				daftarPenyewa[nPenyewa].nama = tFj.NamaPenyewa
				daftarPenyewa[nPenyewa].noHP = tFj.NomorPenyewa
				nPenyewa++
			}
		}
	} else {
		fmt.Println("❌ Kapasitas data penuh!")
	}
}

//Print data ketersediaan lapangan
func tampilLapangan(lapangan [NMAX]Lapangan, n int) {
	fmt.Println("\n=====================================================================================================")
	fmt.Println("||                                   DATA KETERSEDIAAN LAPANGAN FUTSAL                             ||")
	fmt.Println("=====================================================================================================")

	// HEADER
	fmt.Printf("| %-3s | %-12s | %-15s | %-10s | %-12s | %-12s | %-15s |\n",
		"ID", "Nama", "Ukuran (Meter)", "Jenis", "Harga (Rp)", "Jam Sewa", "Status")

	fmt.Println("-----------------------------------------------------------------------------------------------------")

	// DATA
	for i := 0; i < n; i++ {
		fmt.Printf("| %-3d | %-12s | %-15s | %-10s | %-12d | %-12s | %-15s |\n",
			lapangan[i].ID,
			lapangan[i].Nama,
			lapangan[i].Ukuran,
			lapangan[i].Jenis,
			lapangan[i].Harga,
			lapangan[i].JamTersedia,
			lapangan[i].Status)
	}
	fmt.Println("=====================================================================================================")
}

func tampilLapanganTersedia(L daftarLapangan, n int) {
	fmt.Println("\n=====================================================================================================")
    fmt.Println("||                          DATA JADWAL LAPANGAN TERSEDIA (Belum Dibooking)                        ||")
    fmt.Println("=====================================================================================================")
    fmt.Printf("| %-3s | %-12s | %-15s | %-10s | %-12s | %-12s | %-15s |\n", 
                "ID", "Nama", "Ukuran", "Jenis", "Harga", "Jam Sewa", "Status")
    fmt.Println("-----------------------------------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		if L[i].Status == "Tersedia" {
		fmt.Printf("| %-3d | %-12s | %-15s | %-10s | %-12d | %-12s | %-15s |\n",
                L[i].ID, L[i].Nama, L[i].Ukuran, L[i].Jenis, L[i].Harga, L[i].JamTersedia, L[i].Status)
        }
    }
    fmt.Println("=====================================================================================================")
}	

func tampilDaftarPenyewa(T [NMAX]Penyewa, n int) {
    fmt.Println("\n=============================================")
    fmt.Println("||           DATA PENYEWA FUTSAL           ||")
    fmt.Println("=============================================")

    fmt.Printf("| %-5s | %-15s | %-15s |\n",
        "ID", "Nama Penyewa", "No HP")

    fmt.Println("---------------------------------------------")

    if n == 0 {
        fmt.Println("Belum ada data penyewa.")
    }else {
		for i := 0; i < n; i++ {
        	fmt.Printf("| %-5d | %-15s | %-15s |\n",
            	T[i].ID,
            	T[i].nama,
            	T[i].noHP)
		}
    }
    fmt.Println("=============================================")
}

func tampilDaftarTransaksi(T [NMAX]Transaksi, n int) {
	fmt.Println("\n=====================================================================================")
	fmt.Println("||                            DATA TRANSAKSI FUTSAL                                ||")
	fmt.Println("=====================================================================================")
	
	// Header Tabel
	fmt.Printf("| %-15s | %-15s | %-15s | %-12s | %-12s | \n",
		"Penyewa", "Nomor Telepon", "Lapangan", "Tanggal", "Jam Sewa")
    fmt.Println("-------------------------------------------------------------------------------------")
	if n == 0 {
		fmt.Println("Belum ada data transaksi.")
	} else {
		for i := 0; i < n; i++ {
			// Format tanggal
			tanggal := fmt.Sprintf("%02d/%02d/%d",
				T[i].TanggalSewa.Hari,
				T[i].TanggalSewa.Bulan,
				T[i].TanggalSewa.Tahun)

			// Menampilkan data
			fmt.Printf("| %-15s | %-15s | %-15s | %-12s | %-12s |\n",
				T[i].NamaPenyewa,
				T[i].NomorPenyewa,
				T[i].NamaLapangan,
				tanggal,
				T[i].JamSewa)
		}
	}
    fmt.Println("=====================================================================================")
}

func menuUtama() {
	fmt.Println("\n==================================================")
	fmt.Println("||      APLIKASI PEMESANAN LAPANGAN FUTSAL      ||")
	fmt.Println("==================================================")
	fmt.Println("⚽ 1. Menu Lapangan")
    fmt.Println("👥 2. Menu Data Penyewa")
    fmt.Println("📝 3. Menu Transaksi")
    fmt.Println("🚪 4. Keluar")
	fmt.Println("--------------------------------------------------")
}

func menuLapangan() {
	fmt.Println("\n===============================================================")
	fmt.Println("||                   MENU DATA LAPANGAN FUTSAL               ||")
	fmt.Println("===============================================================")
	fmt.Println("➕ 1. Tambah Data Lapangan")
    fmt.Println("✏️  2. Ubah Data Lapangan")
    fmt.Println("🗑️  3. Hapus Data Lapangan")
    fmt.Println("📋 4. Tampilkan Semua Data Lapangan")
    fmt.Println("💲 5. Urutkan Data Lapangan (Berdasarkan Harga)")
	fmt.Println("🕒 6. Urutkan Data Lapangan Kosong (Berdasarkan Jam Mulai)")
    fmt.Println("🔍 7. Cari Data Lapangan")
    fmt.Println("📈 8. Statistik Pendapatan")
    fmt.Println("🚪 9. Keluar")
	fmt.Println("--------------------------------------------------------------")
}

func menuPenyewa() {
	fmt.Println("\n==================================================")
	fmt.Println("||        MENU DATA PENYEWA FUTSAL              ||")
	fmt.Println("==================================================")
	fmt.Println("✍️  1. Ubah Data Penyewa")
    fmt.Println("🗑️  2. Hapus Data Penyewa")
    fmt.Println("👥 3. Tampilkan Semua Data Penyewa")
    fmt.Println("🔍 4. Cari Data Penyewa (Berdasarkan No.HP)")
	fmt.Println("🔍 5. Cari Data Penyewa (Berdasarkan Nama)")
    fmt.Println("🚪 6. Keluar")
	fmt.Println("--------------------------------------------------")
}

func menuTransaksi() {
	fmt.Println("\n==================================================")
	fmt.Println("||        MENU DATA TRANSAKSI FUTSAL            ||")
	fmt.Println("==================================================")
	fmt.Println("➕ 1. Tambah Data Transaksi")
    fmt.Println("✏️  2. Ubah Data Transaksi")
    fmt.Println("🗑️  3. Hapus Data Transaksi")
    fmt.Println("📄 4. Tampilkan Semua Data Transaksi")
    fmt.Println("🔍 5. Cari Data Transaksi")
    fmt.Println("🚪 6. Keluar")
	fmt.Println("--------------------------------------------------")
}

func main() {
    var pilihUtama, pilihMenuLap, pilihMenuPeny, pilihMenuTransaksi int

    // Flag untuk menjaga program tetap berjalan
    var aktif bool = true

	nLapangan = 12 // Sesuai dengan data dummy yang ada
	nPenyewa = 6 // Sesuai dengan data dummy yang ada
	nTransaksi = 6 // Sesuai dengan data dummy yang ada

	fmt.Println()
	fmt.Println("⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽")
    fmt.Println("=============================================================")
    fmt.Println("|| 🏟️     SELAMAT DATANG DI APLIKASI PEMESANAN          🏟️   ||")
    fmt.Println("|| 👟                 LAPANGAN FUTSAL                  👟  ||")
    fmt.Println("=============================================================")
    fmt.Println("⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽ 🥅 ⚽")
	tampilLapangan(lapangan, nLapangan)

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
				urutkanJadwalByHour(&lapangan, nLapangan)
				tampilLapanganTersedia(lapangan, nLapangan)
			case 7:
				cariLapangan(lapangan, nLapangan)
			case 8:
				fmt.Println("Statistik Pendapatan")
				statistikPerBulan(daftarTransaksi, nTransaksi, lapangan, nLapangan)
            case 9: 
                fmt.Println("Kembali ke menu utama")
            }
            
		// TINGKAT 2: Sub-Menu Penyewa
        } else if pilihUtama == 2 {
            menuPenyewa()
            fmt.Print("Pilih menu penyewa: ")
            fmt.Scan(&pilihMenuPeny)
            
            switch pilihMenuPeny {
            case 1:
				ubahPenyewa(&daftarPenyewa, nPenyewa)
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 2:
				hapusPenyewa(&daftarPenyewa, &nPenyewa)
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 3:
				tampilDaftarPenyewa(daftarPenyewa, nPenyewa)
			case 4:
				cariPenyewaByNoHP(daftarPenyewa, nPenyewa)
			case 5:
				sortingNamaPenyewa(&daftarPenyewa, nPenyewa)
				cariPenyewaByNama(daftarPenyewa, nPenyewa)
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

		oldLap := T[i].NamaLapangan
		oldJam := T[i].JamSewa

		fmt.Print("Nama Penyewa: ")
		fmt.Scan(&T[i].NamaPenyewa)
		fmt.Print("Nama Lapangan: ")
		fmt.Scan(&T[i].NamaLapangan)
		fmt.Print("Tanggal Sewa: ")
		fmt.Scan(&T[i].TanggalSewa.Hari, &T[i].TanggalSewa.Bulan, &T[i].TanggalSewa.Tahun)
		fmt.Print("Jam Sewa (ex. 08.00-10.00): ")
		fmt.Scan(&T[i].JamSewa)

		for k := 0; k < nLapangan; k++ {
			if lapangan[k].Nama == oldLap && lapangan[k].JamTersedia == oldJam {
				lapangan[k].Status = "Tersedia"
			}
		}

		penyewaExist := false 
		for j := 0; j < nPenyewa && !penyewaExist; j++ {
			if daftarPenyewa[j].noHP == T[i].NomorPenyewa {
				penyewaExist = true
			}
		}

		if !penyewaExist && nPenyewa < NMAX {
			var newID int
			if nPenyewa > 0 {
				newID = daftarPenyewa[nPenyewa-1].ID
			} else {
				newID = 1
			}
			daftarPenyewa[nPenyewa].ID = newID
				daftarPenyewa[nPenyewa].nama = T[i].NamaPenyewa
				daftarPenyewa[nPenyewa].noHP = T[i].NomorPenyewa
				nPenyewa++
		}
		fmt.Println("✅ Data berhasil diperbarui!")
	} else {
		fmt.Println("❌ ID transaksi tidak ditemukan.")
	}
}

func ubahLapangan(L *daftarLapangan, n int) {
	var idCari int
	var i, idx int
	var ketemu bool = false
	
	fmt.Print("Masukkan ID lapangan yang ingin diubah datanya: ")
	fmt.Scan(&idCari)

	i = 0
	for i < n && !ketemu {
		if L[i].ID == idCari {
			ketemu = true
			idx = i
		} else {
			i++
		}
	}

	if ketemu {
		var namaBaru, jamBaru string

		fmt.Printf("Data ditemukan: %s\n", L[idx].Nama)

		NamaLama := L[idx].Nama
		JamSewaLama := L[idx].JamTersedia

		fmt.Print("Update Nama Lapangan: ")
		fmt.Scan(&namaBaru)
		fmt.Print("Update Jam Tersedia (ex. 08.00-10.00): ")
		fmt.Scan(&jamBaru)
		
		L[idx].Nama = namaBaru
		L[idx].JamTersedia = jamBaru

		for j := 0; j < nTransaksi; j++ {
			if daftarTransaksi[j].NamaLapangan == NamaLama && daftarTransaksi[j].JamSewa == JamSewaLama {
				daftarTransaksi[j].NamaLapangan = namaBaru
				daftarTransaksi[j].JamSewa = jamBaru
			}
		}
		fmt.Println("✅ Nama lapangan dan jam sesi berhasil diubah!")
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

		oldNoHP := T[i].noHP

		var newNama, newNoHP string
		fmt.Print("Masukkan Nama baru: ")
		fmt.Scan(&newNama)
		fmt.Print("Masukkan No HP baru: ")
		fmt.Scan(&newNoHP)

		T[i].nama = newNama
		T[i].noHP = newNoHP

		for j := 0; j < nTransaksi; j++ {
			if daftarTransaksi[j].NomorPenyewa == oldNoHP {
				daftarTransaksi[j].NamaPenyewa = newNama
				daftarTransaksi[j].NomorPenyewa = newNoHP
			}

		}
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

		hapusLap := T[i].NamaLapangan
		hapusJam := T[i].JamSewa

		for j := i; j < *n-1; j++ {
			T[j] = T[j+1]
		}
		*n = *n - 1

		for k := 0; k < nLapangan; k++ {
			if lapangan[k].Nama == hapusLap && lapangan[k].JamTersedia == hapusJam {
				lapangan[k].Status = "Tersedia"
			}
		}
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

 // Selection Sort untuk mengurutkan lapangan berdasarkan harga
func urutkanLapanganByPrice(L *daftarLapangan, n int) {
    var i, j, idx int
    var temp Lapangan

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

// Ambil jam mulai sebagai integer dari format "HH-HH"
func jamMulaiInt(jam string) int {
    var jamMulai int
	fmt.Sscanf(jam, "%d", &jamMulai)
	return jamMulai
}

// Urutkan ascending berdasarkan jam mulai
func urutkanJadwalByHour(L *daftarLapangan, n int) {
    var i, j int
    var temp Lapangan

	for i = 1; i < n; i++ {
		if L[i].Status == "Tersedia" {
			temp = L[i]
			j = i - 1
		}
	}
    for i = 1; i < n; i++ {
        temp = L[i]
        j = i - 1

        for j >= 0 &&
            (jamMulaiInt(L[j].JamTersedia) > jamMulaiInt(temp.JamTersedia)) {
            L[j+1] = L[j]
            j--
        }
        L[j+1] = temp
    }
    fmt.Println("\n✅ Data jadwal berhasil diurutkan berdasarkan Jam Mulai.")
}

func cariTransaksi(T [NMAX]Transaksi, n int) {
    var namaLapangan string
    var ditemukan bool = false
    var tanggal string

    fmt.Print("Masukkan Nama Lapangan yang dicari: ")
    fmt.Scan(&namaLapangan)

    fmt.Println("\n=== HASIL PENCARIAN TRANSAKSI ===")
    fmt.Printf("%-15s | %-15s | %-12s | %-8s | %-12s | %-10s\n",
        "Penyewa", "Lapangan", "Tanggal", "Jam Awal", "Jam Selesai", "Status")
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
            fmt.Printf("%-15s | %-15s | %-12s | %-8s \n",
                T[i].NamaPenyewa,
                T[i].NamaLapangan,
                tanggal,
                T[i].JamSewa)
        }
    }

    if !ditemukan {
        fmt.Println("❌ Tidak ada transaksi untuk lapangan tersebut.")
    }
    fmt.Println("==========================================================================================")
}

//Mencari data penyewa melalui nomor telepon dengan Sequential Search
func cariPenyewaByNoHP(T tabPenyewa, n int) {
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

//Mengurutkan data nama penyewa sebelum digunakan untuk Binary Search
func sortingNamaPenyewa(T *tabPenyewa, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if T[j].nama > T[j+1].nama {
				temp := T[j]
				T[j] = T[j+1]
				T[j+1] = temp
			}
		}
	}
}

//Mencari data penyewa melalui nama dengan Binary Search
func cariPenyewaByNama(T tabPenyewa, n int) {
	var cariNama string
	fmt.Print("Masukkan nama penyewa yang dicari: ")
	fmt.Scan(&cariNama)

	left := 0
	right := n - 1
	idx := -1

	for left <= right {
		mid := (right + left) / 2
		if T[mid].nama == cariNama {
			idx = mid
			left = right + 1
		} else if T[mid].nama < cariNama {
			left = mid + 1
		} else {
			right = mid -1
		}
	}

	if idx != -1 {
		fmt.Println("\n✅ Data ditemukan!")
		fmt.Printf("ID    : %d\n", T[idx].ID)
		fmt.Printf("Nama  : %s\n", T[idx].nama)
		fmt.Printf("No HP : %s\n", T[idx].noHP)
	} else {
		fmt.Println("❌ Nama tidak ditemukan.")
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
	} else {
		fmt.Println("❌ ID lapangan tidak ditemukan.")
	}
}

// Automatisasi bulan
func NamaBulan(bulan int) string {
	switch bulan {
	case 1: return "JANUARI"
	case 2: return "FEBRUARI"
	case 3: return "MARET"
    case 4: return "APRIL"
    case 5: return "MEI"
    case 6: return "JUNI"
    case 7: return "JULI"
    case 8: return "AGUSTUS"
    case 9: return "SEPTEMBER"
    case 10: return "OKTOBER"
    case 11: return "NOVEMBER"
    case 12: return "DESEMBER"
    default: return "TIDAK VALID"
	}
}

// Menghitung statistik data futsal perbulan
func statistikPerBulan(T tabTransaksi, nTransaksi int, L daftarLapangan, nLapangan int) {
    var bulanInput, j int
    var totalPemasukan int = 0
    var ditemukan bool = false
    
    // Array untuk menyimpan hitungan frekuensi tiap lapangan
    type hitung struct {
        nama string
        jumlah int
    }

    var count [NMAX]hitung
    var nHitung int = 0

    fmt.Print("Masukkan Data Statistik Bulan Ke-(1-12): ")
    fmt.Scan(&bulanInput)

    for i := 0; i < nTransaksi; i++ {
        if T[i].TanggalSewa.Bulan == bulanInput {
            ditemukan = true
            
            // Tambahkan harga (cari harga lapangan berdasarkan nama)
			var findHarga bool = false
            for j < nLapangan && !findHarga {
                if L[j].Nama == T[i].NamaLapangan {
                    totalPemasukan += L[j].Harga
					findHarga = true
                }
				j++
            }

            // Hitung frekuensi lapangan
            idx := -1
            for j := 0; j < nHitung; j++ {
                if count[j].nama == T[i].NamaLapangan {
                    idx = j
                }
            }
            if idx == -1 {
                count[nHitung].nama = T[i].NamaLapangan
                count[nHitung].jumlah = 1
                nHitung++
            } else {
                count[idx].jumlah++
            }
        }
    }

    if !ditemukan {
        fmt.Println("❌ Tidak ada transaksi pada bulan tersebut.")
    } else {
        namaBulan := NamaBulan(bulanInput)
		judul := "STATISTIK PENDAPAT DI BULAN " + namaBulan
        
        // Cari yang paling sering dipesan
        max := 0
        idxMax := 0
        for i := 0; i < nHitung; i++ {
            if count[i].jumlah > max {
                max = count[i].jumlah
                idxMax = i
            }
        }
        // Tampilan Tabel
		cetakTabelStatistik(judul, totalPemasukan, count[idxMax].nama, max)
    }
}

func cetakTabelStatistik(judul string, totalPemasukan int, lapanganFavorit string, jumlahFavorit int) {
	fmt.Println("\n========================================================")
    fmt.Printf("|| %-50s ||\n", judul)
    fmt.Println("========================================================")
    fmt.Printf("|| Total Pemasukan    : Rp %-26d ||\n", totalPemasukan)
    fmt.Printf("|| Lapangan Favorit   : %-29s ||\n", lapanganFavorit)
    fmt.Printf("||                      (dipesan sebanyak %-d kali)     ||\n", jumlahFavorit)
    fmt.Println("========================================================")
}

/*
NOTES PENTING!
1. Nama lapangannya harus satu kata, contoh: "LapanganA", "LapanganB", dst. Jangan menggunakan spasi.

FITUR YANG BELUM ADA: (Updated 16/06/2026)
1. Urut data jadwal kosong berdasarkan jam mulai (insertion sort)			[DONE]
2. Urut data lapangan berdasarkan harga (selection sort) 					[DONE]
3. Statistik perbulan & lapangan serta sesi/jam yang paling sering dipesan	[DONE]
4. Cari data penyewa berdasarkan nama dengan Binary Search					[DONE]
PROGRESS KESELURUHAN: 100%
*/
