package strange_advertising_test

import (
	. "hackerRankGo/strange_advertising"
	"testing"
)

func Test_ViralAdvertising(t *testing.T) {
	advertising := ViralAdvertising(5)
	if advertising != 24 {
		t.Error("Failed ", advertising)
	}
}
