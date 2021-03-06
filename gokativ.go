package gokativ

import "strings"

func Vokativ(name string) string {
	lname := strings.ToLower(name)
	if IsFem(lname) {
		_, LastFirst := getMatchingSuffix(lname, wfvls)
		if LastFirst == "l" {
			return name
		} else {
			return vokativWomanFirstName(name)
		}
	}
	return vokativMan(name)
}

func vokativMan(n string) string {
	sfx, vsfx := getMatchingSuffix(n, ms)
	sfxl := len(sfx)
	nl := len(n)
	if nl != 0 {
		n = n[0 : nl-sfxl]
	}
	return n + vsfx
}

func vokativWomanFirstName(n string) string {
	nl := len(n)
	if n[nl-1] == 'a' {
		return n[0:nl-1] + "o"
	}
	return n

}

func IsFem(name string) bool {
	lname := strings.ToLower(name)
	_, sx := getMatchingSuffix(lname, mvws)
	return sx == "w"
}

func getMatchingSuffix(name string, suffices map[string]string) (string, string) {
	ln := len(name)
	for i := 0; i < ln; i++ {
		sfx := name[i:ln]
		v, ok := suffices[sfx]
		if ok {
			return sfx, v
		}
	}
	return "", suffices[""]
}
