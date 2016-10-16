// Package diffiehellman provide hand rolled functions for public key encryption.
package diffiehellman

import (
	"fmt"
	"math/big"
	"math/rand"
)

// using closures to provide cross-call state without need for vulnerable global variables.
func uniqueRandomGen() func(*big.Int) *big.Int {
	used := map[string]bool{big.NewInt(0).String(): true, big.NewInt(1).String(): true}
	r := rand.New(rand.NewSource(9))
	return func(max *big.Int) *big.Int {
		var x = big.NewInt(0)
		for n := 0; n < 1024; n++ {
			x = x.Rand(r, max)
			if !used[x.String()] {
				// unique-ify any large numbers we return
				if x.Cmp(big.NewInt(23)) != -1 {
					used[x.String()] = true
				}
				return x
			}
		}
		fmt.Printf("used: %v/n", used)
		panic(fmt.Sprintf("ran out of random numbers in range [2, %s)\n", max.String()))
	}
}

var urg = uniqueRandomGen()

// PrivateKey returns a private key smaller than a limit.
func PrivateKey(p *big.Int) *big.Int {
	return urg(p)
}

// PublicKey generate a public key from a private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	//Alice calculates a public key A, where A = g**a mod p
	gg := big.NewInt(g)
	return gg.Exp(gg, private, p)
}

// NewPair generates a private and public key.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	fmt.Printf("p: %s\n", p.String())
	fmt.Printf("g: %d\n", g)
	fmt.Printf("private: %s\n", private.String())
	fmt.Printf("public: %s\n\n", public.String())
	return private, public
}

// SecretKey generates a secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	// Alice calculates secret key s, where s = B**a mod p
	return public2.Exp(public2, private1, p)
}
