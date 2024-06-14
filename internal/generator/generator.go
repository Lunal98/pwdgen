package generator

import (
	"crypto/rand"
	"math/big"
)

func Generate(length int, characterset string) string {
	//set := CreateCharacterSet(true, true, true)
	if len(characterset) <= 1 {
		characterset = string(CreateCharacterSet(true, true, true))
	}
	pwd := make([]byte, length)
	for i := range pwd {
		rnd, _ := getrandom(len(characterset))
		pwd[i] = characterset[rnd]
	}
	return string(pwd)
}
func getrandom(max int) (int, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return -1, err
	}
	return int(num.Int64()), nil

}
func CreateCharacterSet(uppercase bool, specials bool, numbers bool) []byte {
	const lowercaseString string = "abcdefghijklmnopqrstuvwxyz"
	const uppercaseString string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const SpecialsString string = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	const numbersString string = "0123456789"
	characterset := lowercaseString
	if uppercase {
		characterset += uppercaseString
	}
	if specials {
		characterset += SpecialsString
	}
	if numbers {
		characterset += numbersString
	}
	set := []byte(characterset)
	return set
}
