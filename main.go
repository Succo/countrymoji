package countrymoji

// Alpha2ToFlag returns the flag emoji for a given ISO 3166-1 alpha-2 country code if possible
// it returns an empty string if nothing matches
func Alpha2ToFlag(code string) string {
	flg := ""
	for _, c := range code {
		flg += charmap[c]
	}

	return flg
}

// Alpha3ToFlag returns the flag emoji for a given ISO 3166-1 alpha-3 country code if possible
// it returns an empty string if nothing matches
func Alpha3ToFlag(code string) string {
	alpha2 := alpha3to2[code]

	return Alpha2ToFlag(alpha2)
}
