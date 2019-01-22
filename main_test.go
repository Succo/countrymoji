package countrymoji

import "testing"

func TestIso2(t *testing.T) {
	if "🇫🇷" != Iso2ToFlag("fr") {
		t.Error("Invalid result for fr")
	}
	if "🇫🇷" != Iso2ToFlag("Fr") {
		t.Error("Invalid result for Fr")
	}
}
