package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(ud.password))

	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}
