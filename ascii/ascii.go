package ascii

import (
	"fmt"
	"strings"
)

const ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// Oppgave 1b
// Implementer en funksjon som eksportere const ascii

// Funksjon tar en "string literal" med kun ASCII tegn og lager en utskrift på
// følgende format:
// [ascii-kode heksadesimalt med store bokstaver A-F][mellomrom]
// [symbol for ascii-kode][mellomrom][ascii-kode binært][linjeskift]
//
// Eksempel (utskriften kommer fra en main.go fil):
//	…
// 3E > 111110
// 3F ? 111111
// 40 @ 1000000
// ...
func IterateOverASCIIStringLiteral(stringLiteral string) {
	// Kode for Oppgave 1a

	stringLiteral = ascii

	for a := range stringLiteral { // itererer gjennom hele stringLiteral

		s := fmt.Sprintf("%c", a) // setter s lik ASCII-symbolet, hvis det finnes

		if len(strings.TrimSpace(s)) != 0 { // strings.TrimSpace(s) fjerner alle istanser av mellomrom(" ") https://stackoverflow.com/questions/18594330/what-is-the-best-way-to-test-for-an-empty-string-in-go
			fmt.Printf("%X %c %b\n", stringLiteral[a], stringLiteral[a], stringLiteral[a]) // %X gir heksades, %c gir symbol, %b gir den binære koden
		} else { // utføres hvis ASCII tegnet ikke har et symbol
			fmt.Printf("%X %b\n", a, a)
		}
	}

	/*for i := 0; i < len(stringLiteral); i++ {
		// using %q here as several of the first characters are unprintable
		// including null, start of header etc
		fmt.Printf("%X %c %b\n", i, i, i)
	}*/

}

// This function returns the constant ascii as a string
func GetASCIIStringLiteral() string {
	return ascii
}

// Unix-like operating systems are known to use it as erase control character,
// i.e. to delete the previous character in the line mode.

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen GreetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
// Gjelder oppgave 1c
func GreetingASCII() string {
	s := []byte{'"', 'H', 'e', 'l', 'l', 'o', ' ', ':', '-', ')', '"'}
	fmt.Printf("\"Hello :-)\" til heksadesimal verdi: % X\n", s)
	return string(s[:])
}
