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
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	func TestGetrandom(t *testing.T) {
		got, err := getrandom(1)
		require.NoError(t, err, "Error generating random number")
		require.Condition()
	}
*/
func TestGenerate(t *testing.T) {
	length := 64
	got := Generate(64, "germanic") //"eariotnslcudpmhg"
	require.Condition(t, func() bool {
		return (len(got) == length)
	})

}
func TestGetrandom(t *testing.T) {
	assert.Panics(t, func() { getrandom(0) })
	assert.NotPanics(t, func() { getrandom(1) })
}

func TestCreateCharacterSet(t *testing.T) {
	type args struct {
		uppercase bool
		specials  bool
		numbers   bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"lower case test", args{false, false, false}, []byte("abcdefghijklmnopqrstuvwxyz")},
		{"upper case test", args{true, false, false}, []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")},
		{"special test", args{false, true, false}, []byte("abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~")},
		{"alphanumeric test", args{true, false, true}, []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")},
		{"full house test", args{true, true, true}, []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~0123456789")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateCharacterSet(tt.args.uppercase, tt.args.specials, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCharacterSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
