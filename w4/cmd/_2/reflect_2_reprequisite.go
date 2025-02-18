package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	// Упаковываем число 1_123_456 (32-битное беззнаковое целое, little-endian)
	if err := binary.Write(&buf, binary.LittleEndian, uint32(1123456)); err != nil {
		panic(err)
	}

	s := "KBTU"

	// Упаковываем длину строки как 32-битное беззнаковое целое
	if err := binary.Write(&buf, binary.LittleEndian, uint32(len(s))); err != nil {
		panic(err)
	}

	// Упаковываем сами байты строки
	if _, err := buf.Write([]byte(s)); err != nil {
		panic(err)
	}

	// Упаковываем число 16 (также 32-битное беззнаковое целое, little-endian)
	if err := binary.Write(&buf, binary.LittleEndian, uint32(16)); err != nil {
		panic(err)
	}

	// Получаем итоговый срез байтов
	b := buf.Bytes()

	fmt.Println(b)
}
