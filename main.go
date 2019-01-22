package countrymoji

var charmap = map[rune]string{
	'a': "🇦",
	'b': "🇧",
	'c': "🇨",
	'd': "🇩",
	'e': "🇪",
	'f': "🇫",
	'g': "🇬",
	'h': "🇭",
	'i': "🇮",
	'j': "🇯",
	'k': "🇰",
	'l': "🇱",
	'm': "🇲",
	'n': "🇳",
	'o': "🇴",
	'p': "🇵",
	'q': "🇶",
	'r': "🇷",
	's': "🇸",
	't': "🇹",
	'u': "🇺",
	'v': "🇻",
	'w': "🇼",
	'x': "🇽",
	'y': "🇾",
	'z': "🇿",
}

func ToFlag(code string) string {
	flg := ""
	for _, c := range code {
		flg += charmap[c]
	}

	return flg
}
