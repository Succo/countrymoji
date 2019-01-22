package countrymoji

func Iso2ToFlag(code string) string {
	flg := ""
	for _, c := range code {
		flg += charmap[c]
	}

	return flg
}
