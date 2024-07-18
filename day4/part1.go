package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const secret = "iwrupvqb"

func min5Zeroes(b string) bool {
	return len(b) >= 5 && b[:5] == "00000"
}

func main() {

	index := 0

	for {
		idxStr := strconv.Itoa(index)
		res := secret + idxStr
		hash := md5.Sum([]byte(res))
		hexHash := hex.EncodeToString(hash[:])

		if min5Zeroes(hexHash) {
			fmt.Printf("Solution -> %s, index: %d", hexHash, index)
			break
		}
		if index%100000 == 0 {
			fmt.Printf("Currently at index: %d\n", index)
		}
		index++
	}
}
