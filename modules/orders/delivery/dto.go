package delivery

type AddOrder struct {
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}
