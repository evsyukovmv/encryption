package encryption

const asciiNumberStart = 48

var vowels = [...]rune{'a', 'e', 'i', 'o', 'u'}

var (
	decodeMap = make(map[rune]rune, len(vowels))
	encodeMap = make(map[rune]rune, len(vowels))
)

func init() {
	for i, v := range vowels {
		encodeMap[v] = rune(i + 1 + asciiNumberStart)
		decodeMap[rune(i+1+asciiNumberStart)] = v
	}
}

// Encode replaces vowels to numbers
func Encode(s string) string {
	return applyMapping(s, encodeMap)
}

// Decode replaces numbers to vowels
func Decode(s string) string {
	return applyMapping(s, decodeMap)
}

func applyMapping(s string, m map[rune]rune) string {
	runes := []rune(s)
	for i := range runes {
		if v, ok := m[runes[i]]; ok {
			runes[i] = v
		}
	}
	return string(runes)
}
