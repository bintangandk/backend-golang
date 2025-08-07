package helpers

import (
	"santrikoding/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Ambil secret key dari environment variable
var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func GenerateToken(username string) string {

	// Buat klaim token
	expirationTime := time.Now().Add(60 * time.Minute) // Token berlaku selama 60 menit

	claims := &jwt.RegisteredClaims{
		Subject:   username,                           // Simpan username sebagai sub
		ExpiresAt: jwt.NewNumericDate(expirationTime), // Waktu kedaluwarsa token
	}

	// Buat token dengan klaim dan tanda tangan menggunakan jwtKey
	// Gunakan metode HS256 untuk tanda tanga
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	return token
}
