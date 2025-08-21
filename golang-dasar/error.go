package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return nilai / pembagi, nil
}

func main() {
	result, err := pembagian(100, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Hasil:", result)
	}

	error := saveData("")
	if error != nil {
		fmt.Println("Error:", error.Error())
	} else {
		fmt.Println("Data berhasil disimpan")
	}

}

func saveData(id string) error {
	if id == "" {
		return &validationError{Message: "ID tidak boleh kosong"}
	}
	// Simulasi penyimpanan data
	fmt.Println("Data disimpan:", id)
	return nil
}

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}
