package main

func isIsomorphic(s string, t string) bool {
	isoS := make(map[byte]byte)
	isoS[s[0]] = t[0]

	isoT := make(map[byte]byte)
	isoT[t[0]] = s[0]

	for i := 1; i < len(s); i++ {
		if (isoS[s[i]] != t[i] && isoS[s[i]] != 0) || (isoT[t[i]] != s[i] && isoT[t[i]] != 0) {
			return false
		}

		isoS[s[i]] = t[i]
		isoT[t[i]] = s[i]
	}

	return true
}

func main() {
	isIsomorphic("egg", "add")
}
