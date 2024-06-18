/*
Copyright © 2024 Alex Bedo <alex98hun@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
