# Go E-commerce API

API e-commerce sederhana yang dibangun menggunakan Go dan Fiber framework.

## Teknologi yang Digunakan

- Go (Golang)
- Fiber (Web Framework)
- GORM (ORM)
- MySQL (Database)
- JWT (Autentikasi)

## Endpoint API

### Autentikasi

| Endpoint             | Method | Deskripsi              | Akses  |
| -------------------- | ------ | ---------------------- | ------ |
| `/api/auth/register` | POST   | Mendaftarkan user baru | Public |
| `/api/auth/login`    | POST   | Login user             | Public |
| `/api/auth/logout`   | POST   | Logout user            | Public |

### User

| Endpoint         | Method | Deskripsi                            | Akses     |
| ---------------- | ------ | ------------------------------------ | --------- |
| `/api/users`     | GET    | Mendapatkan semua data user          | Public    |
| `/api/users/:id` | GET    | Mendapatkan data user berdasarkan ID | Public    |
| `/api/welcome`   | GET    | Halaman welcome (test autentikasi)   | Protected |

### Toko

| Endpoint        | Method | Deskripsi                            | Akses  |
| --------------- | ------ | ------------------------------------ | ------ |
| `/api/toko`     | GET    | Mendapatkan semua data toko          | Public |
| `/api/toko/:id` | GET    | Mendapatkan data toko berdasarkan ID | Public |

### Alamat

| Endpoint          | Method | Deskripsi                         | Akses     |
| ----------------- | ------ | --------------------------------- | --------- |
| `/api/alamat`     | POST   | Menambahkan alamat baru           | Protected |
| `/api/alamat`     | GET    | Mendapatkan semua alamat user     | Protected |
| `/api/alamat/:id` | GET    | Mendapatkan alamat berdasarkan ID | Protected |
| `/api/alamat/:id` | PUT    | Mengupdate alamat                 | Protected |
| `/api/alamat/:id` | DELETE | Menghapus alamat                  | Protected |

### Kategori

| Endpoint          | Method | Deskripsi                  | Akses      |
| ----------------- | ------ | -------------------------- | ---------- |
| `/api/categories` | GET    | Mendapatkan semua kategori | Public     |
| `/api/categories` | POST   | Menambahkan kategori baru  | Admin Only |

### Produk

| Endpoint                   | Method | Deskripsi                | Akses     |
| -------------------------- | ------ | ------------------------ | --------- |
| `/api/products`            | GET    | Mendapatkan semua produk | Public    |
| `/api/products`            | POST   | Menambahkan produk baru  | Protected |
| `/api/products/:id/upload` | POST   | Upload foto produk       | Protected |

### Transaksi

| Endpoint            | Method | Deskripsi              | Akses     |
| ------------------- | ------ | ---------------------- | --------- |
| `/api/transactions` | POST   | Membuat transaksi baru | Protected |


