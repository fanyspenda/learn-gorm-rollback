# DB-ROLLBACK (GORM.IO)
Rollback artinya menarik ulang proses yang sudah berjalan seolah tidak pernah terjadi.

Misal kita akan melakukan 2 proses untuk sebuah pembelian. Pertama kurangin saldo, terus buat invoice.

Jika tanpa rollback, ketika kita berhasil kurangin saldo pembeli, tapi gagal saat membuat invoice, maka saldo akan tetap berkurang tapi barang tidak terbeli.

Jika menggunakan rollback, maka ketika gagal membuat invoice, kita bisa memanggil fungsi rollback untuk mengembalikan jumlah saldo ke semula.

Rollback hanya bisa berjalan ketika kita memulai proses `transaction` di DB kita.

## Rollback di Gorm.io
Gorm.io secara default menggunakan fitur `transaction`. Sehingga, kita langsung bisa menggunakan fitur `rollback`. Kita bisa mematikan setting default ini ketika init DB.