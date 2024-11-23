
# Assignment_Day_26

## Deskripsi Proyek
Sistem backend untuk **manajemen inventaris** yang mencakup pengelolaan produk, inventaris, dan pesanan. Proyek ini dibangun menggunakan **Golang** dengan framework **Gin**, serta mendukung operasi **CRUD** melalui **RESTful API**.

## Langkah-langkah Pengaturan

### 1. Persiapan Lingkungan
Pastikan Anda memiliki:
- **Go** 1.19 atau versi terbaru.
- **MySQL** (atau database relasional lain yang kompatibel).
- Alat pengujian API seperti **Postman** atau **Curl**.

### 2. Menyiapkan Database

1. Buat database baru di MySQL:
   ```sql
   CREATE DATABASE assignment_day_26;
   ```

2. Import file `ddl.sql` untuk membuat tabel dan memasukkan data awal:
   ```bash
   mysql -u [username] -p assignment_day_26 < ddl.sql
   ```

### 3. Mengonfigurasi Proyek
1. Pastikan koneksi ke database disesuaikan. Tambahkan informasi koneksi di file konfigurasi (misalnya `config/db.go`) atau langsung di `main.go`:
   ```go
   dsn := "username:password@tcp(localhost:3306)/assignment_day_26?parseTime=true"
   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
   if err != nil {
       log.Fatalf("failed to connect database: %v", err)
   }
   ```

2. Sesuaikan file konfigurasi jika Anda menggunakan lingkungan yang berbeda (contoh: `.env`).

### 4. Menjalankan Proyek
1. **Instal dependensi**:
   ```bash
   go mod tidy
   ```

2. **Jalankan server**:
   ```bash
   go run main.go
   ```

3. **Akses server**:
   - API akan berjalan di `http://localhost:8080`.
   - Gunakan **Postman** atau **Curl** untuk mengakses endpoint yang tersedia.


### 5. Endpoint

| **Method** | **Endpoint**                      | **Deskripsi**                             |
|------------|-----------------------------------|-------------------------------------------|
| **Produk** |                                   |                                           |
| POST       | `/product`                        | Membuat produk baru                       |
| GET        | `/product`                        | Membaca daftar semua produk               |
| GET        | `/product/:id`                    | Membaca detail produk berdasarkan ID      |
| PUT        | `/product/:id`                    | Memperbarui produk berdasarkan ID         |
| DELETE     | `/product/:id`                    | Menghapus produk berdasarkan ID           |
| POST       | `/upload-product-image`           | Mengunggah gambar produk                  |
| POST       | `/upload-product-image/:id`       | Mengunggah gambar produk berdasarkan ID   |
| GET        | `/download-product-image/:id`     | Mengunduh gambar produk berdasarkan ID    |
| **Inventaris** |                                |                                           |
| POST       | `/inventory`                      | Membuat inventaris baru                   |
| GET        | `/inventory/:id`                  | Membaca detail inventaris berdasarkan ID  |
| PUT        | `/inventory/:id`                  | Memperbarui inventaris berdasarkan ID     |
| **Pesanan** |                                  |                                           |
| POST       | `/order`                          | Membuat pesanan baru                      |
| GET        | `/order/:id`                      | Membaca detail pesanan berdasarkan ID     |



