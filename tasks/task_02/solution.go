package main

func rotateRunes(s string, shift int) string {
	r := []rune(s)
	n := len(r)
	if n == 0 {
		return s
	}
	k := ((shift % n) + n) % n
	out := []rune{}
	out = append(out, r[k:]...)
	out = append(out, r[:k]...)
	return string(out)
}
