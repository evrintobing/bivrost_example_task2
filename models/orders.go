package models

type Order struct {
	IDOrder      int `json:"id_order"`
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}
