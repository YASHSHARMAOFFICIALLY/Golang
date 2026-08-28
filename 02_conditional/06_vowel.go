package main

func vowelorConstant(ch byte) string {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return "vowel"
	default:
		return "consonant"
	}
}
