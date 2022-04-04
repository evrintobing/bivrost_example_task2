package models

type Order struct {
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}

type GetOrder struct {
	IDOrder      int `json:"id_order"`
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}
