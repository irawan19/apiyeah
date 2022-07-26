package helper

import (
	"apiyeah/models"
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	var generateNoRegistrasi = string(b)
	ceknoregistrasi, err := models.CekNoRegistrasi(generateNoRegistrasi)
	if err != nil {
		fmt.Println("data: ", err)
	} else {
		if ceknoregistrasi != 0 {
			generateNoRegistrasi = RandStringBytes(6)
		} else {
			generateNoRegistrasi = string(b)
		}
	}

	return generateNoRegistrasi
}
