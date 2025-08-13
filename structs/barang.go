package structs

// Struct ini digunakan untuk menampilkan data barang sebagai response API
type BarangResponse struct {
	Id        uint    `json:"id"`
	Nama      string  `json:"nama"`
	Harga     float64 `json:"harga"`
	Stok      int     `json:"stok"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
