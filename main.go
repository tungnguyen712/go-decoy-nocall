// Sample app for the ground-truth tests. It pins vulnerable dependency
// versions ON PURPOSE. Never deploy it.
//
// Imports every vulnerable package but never calls the vulnerable function.
// It also defines local functions with the same names as the vulnerable ones:
// matching must be by package, not by bare name.
package main

import (
	"fmt"
	"strings"

	"github.com/Masterminds/goutils"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// Local look-alikes. These are not the library functions.
func ParseAcceptLanguage(s string) string { return strings.TrimSpace(s) }
func Unmarshal(b []byte) string           { return string(b) }
func RandomAlphaNumeric(n int) string     { return strings.Repeat("x", n) }

func main() {
	fmt.Println(language.English, language.Make("fr"))

	out, err := yaml.Marshal(map[string]int{"a": 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	short, err := goutils.Abbreviate("hello world", 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(short)

	fmt.Println(ParseAcceptLanguage(" en "), Unmarshal([]byte("a")), RandomAlphaNumeric(4))
}
