package countrymoji

func Alpha2ToFlag(code string) string {
	flg := ""
	for _, c := range code {
		flg += charmap[c]
	}

	return flg
}

func Alpha3ToFlag(code string) string {
	alpha2 := alpha3to2[code]

	return Alpha2ToFlag(alpha2)
}
