package main

import (
	"PI6/share"
	uuid "github.com/google/uuid"
)

func main() {
	uuidV := share.FloatsAsUUID(1.5342, -19248.390434)

	namespace := uuid.NameSpaceDNS

	// Gera UUID v5 com base no namespace e string
	generatedUUID := uuid.NewMD5(namespace, []byte(uuidV))

	println(generatedUUID.String())
	println(share.FloatsFromUUID([]byte(uuidV)))
}
