package beautiful_days_at_the_movies_test

import (
	. "hackerRankGo/beautiful_days_at_the_movies"
	"testing"
)

func Test_BeautifulDays(t *testing.T) {
	days := BeautifulDays(20, 23, 6)
	if days != 2 {
		t.Error("Test case failed expected result 2 but ", days)
	}
}
