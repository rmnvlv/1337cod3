package main

import "fmt"

/*
---------------------------------------------------------EASY---------------------------------------------------------

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II.
The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII.
Instead, the number four is written as IV. Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX.
There are six instances where subtraction is used:
-----------------------------------------------------------------------
 |   I can be placed before V (5) and X (10) to make 4 and 9.		  |
 |   X can be placed before L (50) and C (100) to make 40 and 90.     |
 |   C can be placed before D (500) and M (1000) to make 400 and 900. |
-----------------------------------------------------------------------
Given a roman numeral, convert it to an integer.


Example 1:
Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:
    1 <= s.length <= 15
    s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
    It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/

// First try   Runtime: 27 ms, Memory Usage: 2.9 MB
func romanToInt_1(s string) int {
	romanInt := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var number int
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i : i+1]
		value := romanInt[char]
		if i != 0 {
			if s[i-1:i] == "I" {
				if char == "V" {
					value = 4
					i--
					if i == 0 {
						number += value
						break
					}
				} else if char == "X" {
					value = 9
					i--
					if i == 0 {
						number += value
						break
					}
				}
			}
			if s[i-1:i] == "X" {
				if char == "L" {
					value = 40
					i--
					if i == 0 {
						number += value
						break
					}
				} else if char == "C" {
					value = 90
					i--
					if i == 0 {
						number += value
						break
					}
				}
			}
			if s[i-1:i] == "C" {
				if char == "D" {
					value = 400
					i--
					if i == 0 {
						number += value
						break
					}
				} else if char == "M" {
					value = 900
					i--
					if i == 0 {
						number += value
						break
					}
				}
			}
		}
		number += value
	}
	return number
}

// Second try Runtime: 13 ms, Memory Usage: 2.9 MB
func romanToInt_2(s string) int {
	romanInt := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	number := 0
	lastValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i : i+1]
		value := romanInt[char]
		if value < lastValue {
			number -= value
		} else if value >= lastValue {
			number += value
		}
		lastValue = value
	}
	return number
}

// Third try Runtime: 8 ms, Memory Usage: 2.8 MB
func romanToInt_3(s string) int {
	romanInt := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	number := 0
	lastValue := 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i : i+1]
		value := romanInt[char]
		if value < lastValue {
			number -= value
			continue
		}
		number += value
		lastValue = value
	}
	return number
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(romanToInt_3(s))
}
