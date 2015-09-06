# Bank OCR Kata


This is an excercise in programming. See the [coding dojo page](http://codingdojo.org/cgi-bin/index.pl?KataBankOCR) for more details regarding this specific kata. FOr more information on coding katas in general see the [wikipedia page](https://en.wikipedia.org/wiki/Kata_(programming)) or [Dave Thonmas' kata site](http://codekata.com/).

## Overview

### User Story 1

You work for a bank, which has recently purchased an ingenious machine to assist in reading letters and faxes sent in by branch offices. The machine scans the paper documents, and produces a file with a number of entries which each look like this:

```
    _  _     _  _  _  _  _
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _| 
 
```
                           
Each entry is 4 lines long, and each line has 27 characters. The first 3 lines of each entry contain an account number written using pipes and underscores, and the fourth line is blank. Each account number should have 9 digits, all of which should be in the range 0-9. A normal file contains around 500 entries.

Your first task is to write a program that can take this file and parse it into actual account numbers.

### User Story 2

Having done that, you quickly realize that the ingenious machine is not in fact infallible. Sometimes it goes wrong in its scanning. The next step therefore is to validate that the numbers you read are in fact valid account numbers. A valid account number has a valid checksum. This can be calculated as follows:

account number:  `3  4  5  8  8  2  8  6  5`
position names:  `d9 d8 d7 d6 d5 d4 d3 d2 d1`

checksum calculation:
`(d1+2*d2+3*d3 +..+9*d9) mod 11 = 0`

So now you should also write some code that calculates the checksum for a given number, and identifies if it is a valid account number.
