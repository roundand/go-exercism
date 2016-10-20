// Package diffiehellman provide hand rolled functions for public key encryption.
package diffiehellman

import (
	"fmt"
	"math/big"
	"math/rand"
)

// using closures to provide cross-call state without need for vulnerable global variables.
func uniqueRandomGen() func(*big.Int) *big.Int {
	one := big.NewInt(1)
	r := rand.New(rand.NewSource(9))
	return func(max *big.Int) *big.Int {
		var x = big.NewInt(0)
		for n := 0; n < 1024; n++ {
			x = x.Rand(r, max)
			// need a result greater than one
			if x.Cmp(one) == 1 {
				return x
			}
		}
		panic(fmt.Sprintf("ran out of random numbers in range [2, %s)\n", max.String()))
	}
}

// urg holds a seeded random number generator
var urg = uniqueRandomGen()

// PrivateKey returns a private key greater than 1 and smaller than p.
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
	return private, public
}

// SecretKey generates a secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	// Alice calculates secret key s, where s = B**a mod p
	return big.NewInt(0).Exp(public2, private1, p)
}
