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
	IllegibleDigit       = errors.New("Illegible digit")
	InvalidChecksum      = errors.New("Invalid checksum in account number")
	InvalidAccountNumber = errors.New("Invalid account number")
	InvalidStrokes       = errors.New("Invalid characters in cell")
	UnknownDigit         = errors.New("The cell doesn't match a known digit")
)

// Cell represents one "digit" in the banks format.
type cell [3]string

// digit represents one number in an account number.
type digit struct {
	Cell  cell
	Error error
}

// accountNum represents a 9 "digit" account number.
type accountNum [9]digit

// digits returns the digits as a slice of strings.
func (a accountNum) digits() []string {
	var text []string
	for _, digit := range a {
		text = append(text, digits[digit.Cell])
	}
	return text
}

// Checksum validates the checksum and returns Inval
func (a accountNum) Checksum() error {
	darr := a.digits()
	var (
		sum = 0
		mul = 1
	)
	for i := len(darr) - 1; i >= 0; i-- {
		d, err := strconv.Atoi(darr[i])
		if err != nil {
			return UnknownDigit
		}
		sum = sum + d*mul
		mul++
	}
	if sum%11 == 0 {
		return nil
	}
	return InvalidChecksum
}

// String returns the account number as a string.
func (a accountNum) String() string {
	return strings.Join(a.digits(), "")
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
			a[offset/3].Cell[cellIndex] = line[offset : offset+3]
			offset = offset + 3
		}
	}
	return a, nil
}
