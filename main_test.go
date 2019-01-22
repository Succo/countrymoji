package countrymoji

import "testing"

func TestAlpha2(t *testing.T) {
	if "🇫🇷" != Alpha2ToFlag("fr") {
		t.Error("alpha 2 invalid result for fr")
	}
	if "🇫🇷" != Alpha2ToFlag("Fr") {
		t.Error("alpha 2 invalid result for Fr")
	}
}

func TestAlpha3(t *testing.T) {
	if "🇫🇷" != Alpha3ToFlag("FRA") {
		t.Error("alpha 3 invalid result for FRA")
	}
}
