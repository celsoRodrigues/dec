package mygpg

import (
	"bytes"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func Encrypt(plaintext []byte, password []byte, packetConfig *packet.Config) (ciphertext []byte, err error) {

	encbuf := bytes.NewBuffer(nil)

	w, _ := armor.Encode(encbuf, "PGP MESSAGE", nil)

	pt, _ := openpgp.SymmetricallyEncrypt(w, password, nil, packetConfig)

	_, err = pt.Write(plaintext)
	if err != nil {
		return
	}

	pt.Close()
	w.Close()
	ciphertext = encbuf.Bytes()

	return
}
