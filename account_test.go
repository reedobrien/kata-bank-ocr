package account_test

import (
	"strings"
	"testing"

	. "github.com/reedobrien/kata-bank-ocr"
)

func TestParseAccountNumbers(t *testing.T) {
	for _, tv := range case1Tests {
		lines := strings.Split(tv.in, "\n")
		anum, err := ParseAccountNumber(lines[:3])
		if err != nil {
			t.Error("Unexpected error parsing account input.")
		}
		got, err := anum.Numeric()
		if err != nil {
			if err != tv.err {
				t.Errorf("Got error %s for %s when trying to get numeric value", err, tv.in)
			}
		}

		if got != tv.want {
			t.Errorf("Got: %v but wanted %d", got, tv.want)
		}
	}
}

var case1Tests = []struct {
	in   string
	want int32
	err  error
}{
	{` _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|

`, 000000000, nil},
	{`                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |

`, 111111111, nil},
	{` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 

`, 222222222, nil},
	{` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|
`, 333333333, nil},
	{`                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |
`, 444444444, nil},
	{` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|
`, 555555555, nil},
	{` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|
`, 666666666, nil},
	{` _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
`, 777777777, nil},
	{` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|
`, 888888888, nil},
	{` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|
`, 999999999, nil},
	{`    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
`, 123456789, nil},
	{`This isn't valid, but it is
 a really long multiline   
 string that should faili.:)
 `, 0, InvalidAccountNumber},
}
