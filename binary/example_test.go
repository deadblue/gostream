package binary

import (
	"bytes"
	"encoding/hex"
	"log"
)

func ExampleReader_ReadUint32() {
	buf := bytes.NewReader([]byte{
		0x12, 0x34, 0x56, 0x78,
	})
	br := NewReader(buf, LittleEndian)
	if u, err := br.ReadUint32(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Uint value: 0x%x", u)
	}
	// Output:
	// Uint value: 0x78563412
}

func ExampleWriter_WriteUint32() {
	buf := &bytes.Buffer{}
	bw := NewWriter(buf, BigEndian)
	if err := bw.WriteUint32(0x12345678); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Binary data: %s", hex.EncodeToString(buf.Bytes()))
	}
	// Output:
	// Binary data: 12345678
}
