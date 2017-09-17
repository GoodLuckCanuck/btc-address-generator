package main

import (
	"os"
        "time"
	"fmt"
	"math/big"
	"strings"
	"crypto/rand"
        "github.com/btcsuite/btcd/btcec"
        "github.com/btcsuite/btcutil"
        "github.com/btcsuite/btcd/chaincfg"
)

func generatePrivateKey() []byte {
	bytes := make([]byte, 32)
	if len(os.Args) < 2 {
		fmt.Println("No startnum passed on command line, generating random starting point")
		rand.Read(bytes)
	} else {
        	// read number from command-line, used for testing or seeding only
		// max 32-byte int is 115792089237316195423570985008687907853269984665640564039457584007913129639935
		arg := os.Args[1]

		// Try treating the supplied number as decimal; if not, try hexadecimal (the 0x prefix can be used to coerce a number into hex)
		startnum, ok := new(big.Int).SetString(arg,10)
		if ok != true { startnum, ok = new(big.Int).SetString(strings.TrimPrefix(arg,"0x"),16) }
		copy(bytes[32-len(startnum.Bytes()):], startnum.Bytes())
	}
	return bytes
}

func printAddressesForever() {
        padded := generatePrivateKey()
        count, one := new(big.Int).SetBytes(padded), big.NewInt(1)
        for {
                // Copy count value's bytes to padded slice
                copy(padded[32-len(count.Bytes()):], count.Bytes())

                // Get public key
                _, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)

                // Get compressed and uncompressed addresses
                caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
                uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

                // Print keys
                fmt.Printf("%s %x Compressed\n",   caddr.EncodeAddress(), padded)
                fmt.Printf("%s %x Uncompressed\n", uaddr.EncodeAddress(), padded)

                // Increment our counter
                count.Add(count, one)
        }
}

func main() {
    for i := 1; i < 180; i++ {
        go printAddressesForever()
        //fmt.Fprintf(os.Stderr, "\n\nNumber of threads: %d\n", i)
        //time.Sleep(1000 * time.Millisecond)
    }
    for {
        time.Sleep(3000 * time.Millisecond)
    }
}
