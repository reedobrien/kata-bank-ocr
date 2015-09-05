package account

import (
	"errors"
	"strconv"
	"strings"
)

// Digits maps a cell's "strokes" to corresponding integers.
var digits = map[cell]string{
	[3]string{" _ ", "| |", "|_|"}: "0",
	[3]string{"   ", "  |", "  |"}: "1",
	[3]string{" _ ", " _|", "|_ "}: "2",
	[3]string{" _ ", " _|", " _|"}: "3",
	[3]string{"   ", "|_|", "  |"}: "4",
	[3]string{" _ ", "|_ ", " _|"}: "5",
	[3]string{" _ ", "|_ ", "|_|"}: "6",
	[3]string{" _ ", "  |", "  |"}: "7",
	[3]string{" _ ", "|_|", "|_|"}: "8",
	[3]string{" _ ", "|_|", " _|"}: "9",
}

var (
	InvalidAccountNumber = errors.New("Invalid account number")
	UnknownDigit         = errors.New("The cell doesn't match a known digit")
)

// Cell represents one "digit" in the banks format.
type cell [3]string

// accountNum represents a 9 "digit" account number.
type accountNum [9]cell

// String returns the account number as a string.
func (a accountNum) String() string {
	var text []string
	for _, cell := range a {
		text = append(text, digits[cell])
	}
	return strings.Join(text, "")
}

// Numeric returns the account number as an int32.
func (a accountNum) Numeric() (int32, error) {
	num, err := strconv.ParseInt(a.String(), 10, 32)
	num32 := int32(num)
	if err != nil {
		return num32, InvalidAccountNumber
	}
	return num32, nil
}

// ParseAccountNumber takes the three non blank lines that make up an account
// number from the input file.
func ParseAccountNumber(l []string) (accountNum, error) {
	a := accountNum{}
	for cellIndex, line := range l {
		for offset := 0; offset < 27; {
			a[offset/3][cellIndex] = line[offset : offset+3]
			offset = offset + 3
		}
	}
	return a, nil
}
