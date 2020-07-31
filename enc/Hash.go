package enc

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// get unique hash value
func GetHashString() string {
	// current time for seed
	t := time.Now().String()

	// generate the hash
	hashBytes := md5.Sum([]byte(t))

	// return hex encoded string
	return hex.EncodeToString(hashBytes[:])
}
