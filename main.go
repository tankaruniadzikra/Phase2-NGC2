package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// 'Character' untuk merepresentasikan karakter (Heroes atau Villains).
type Character struct {
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImageURL string `json:"image_url"`
}

func main() {
	// Membuat koneksi ke database MySQL.
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phase2-ngc2")
	if err != nil {
		log.Fatal(err) // Jika terjadi kesalahan dalam menghubungkan ke database, hentikan program.
	}
	defer db.Close() // Pastikan untuk menutup koneksi ke database setelah selesai digunakan.

	mux := http.NewServeMux() // Membuat multiplexer untuk menangani rute HTTP.

	// Handler untuk '/heroes'.
	mux.HandleFunc("/heroes", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM Heroes") // Eksekusi query SQL untuk mendapatkan data Heroes.
		if err != nil {
			log.Fatal(err) // Jika terjadi kesalahan dalam eksekusi query, hentikan program.
		}
		defer rows.Close() // Menutup hasil query setelah selesai digunakan.

		var heroes []Character // Membuat slice untuk menyimpan karakter-karakter Heroes.

		// Loop melalui hasil query dan memindai data karakter Heroes.
		for rows.Next() {
			var id int
			var name, universe, skill, imageURL string
			err := rows.Scan(&id, &name, &universe, &skill, &imageURL) // Memindai data dari hasil query.
			if err != nil {
				log.Fatal(err) // Jika terjadi kesalahan dalam pemindaian data, hentikan program.
			}
			heroes = append(heroes, Character{Name: name, Universe: universe, Skill: skill, ImageURL: imageURL}) // Tambahkan karakter pahlawan ke slice.
		}

		w.Header().Set("Content-Type", "application/json") // Atur header respons untuk tipe konten JSON.
		json.NewEncoder(w).Encode(heroes)                  // Kodekan slice Heroes ke JSON dan kirim sebagai respons.
	})

	// Handler untuk '/villains'.
	mux.HandleFunc("/villains", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM Villains") // Eksekusi query SQL untuk mendapatkan data villains.
		if err != nil {
			log.Fatal(err) // Jika terjadi kesalahan dalam eksekusi query, hentikan program.
		}
		defer rows.Close() // Menutup hasil query setelah selesai digunakan.

		var villains []Character // Membuat slice untuk menyimpan karakter-karakter villains.

		// Loop melalui hasil query dan memindai data karakter villains.
		for rows.Next() {
			var id int
			var name, universe, imageURL string
			err := rows.Scan(&id, &name, &universe, &imageURL) // Memindai data dari hasil query.
			if err != nil {
				log.Fatal(err) // Jika terjadi kesalahan dalam pemindaian data, hentikan program.
			}
			villains = append(villains, Character{Name: name, Universe: universe, ImageURL: imageURL}) // Tambahkan karakter villains ke slice.
		}

		w.Header().Set("Content-Type", "application/json") // Atur header respons untuk tipe konten JSON.
		json.NewEncoder(w).Encode(villains)                // Kodekan slice villains ke JSON dan kirim sebagai respons.
	})

	// Membuat server HTTP dengan konfigurasi.
	app := http.Server{
		Addr:    "localhost:8001", // Server akan berjalan di localhost pada port 8001.
		Handler: mux,              // Menggunakan multiplexer yang telah kita buat untuk menangani rute HTTP.
	}

	err = app.ListenAndServe() // Menjalankan server HTTP.
	if err != nil {
		log.Fatalf("failed to run app: %s\n", err.Error()) // Jika terjadi kesalahan dalam menjalankan server, hentikan program.
	}
}
