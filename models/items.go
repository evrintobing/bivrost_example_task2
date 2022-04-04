package models

type Items struct {
	ID              int    `json:"id"`
	NamaProduk      string `json:"nama_produk"`
	DeskripsiProduk string `json:"deskripsi_produk"`
	Harga           int    `json:"harga"`
}
