package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	exponents  = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
	digitWords = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}

	hundred     = decimal.NewFromInt(100)
	maxBahtPart = decimal.NewFromInt(math.MaxInt64)
)

// numberToThai renders a non-negative integer as Thai number words.
func numberToThai(number int64) string {
	if number == 0 {
		return digitWords[0]
	}

	var thai strings.Builder
	numStr := strconv.FormatInt(number, 10)
	length := len(numStr)
	for i := range length {
		digit := int(numStr[i] - '0')
		pos := length - i - 1

		var exponent string
		if pos >= 6 {
			exponent = exponents[pos%6]
		} else {
			exponent = exponents[pos]
		}

		mod := (length - i) % 6
		switch {
		case mod == 1 && digit == 1 && length > 1:
			thai.WriteString("เอ็ด")
		case mod == 2 && digit == 1:
			thai.WriteString(exponent)
		case mod == 2 && digit == 2:
			thai.WriteString("ยี่")
			thai.WriteString(exponent)
		case digit > 0:
			thai.WriteString(digitWords[digit])
			thai.WriteString(exponent)
		}

		if mod == 1 && (length-i) >= 7 {
			thai.WriteString("ล้าน")
		}
	}

	return thai.String()
}

func decimalToThai(value decimal.Decimal) string {
	numStr := value.String()
	_, after, ok := strings.Cut(numStr, ".")
	if !ok {
		return ""
	}

	var thai strings.Builder
	for _, ch := range after {
		thai.WriteString(digitWords[ch-'0'])
	}

	return thai.String()
}

func convertDecimalToString(decimalValue decimal.Decimal) (string, error) {
	negative := decimalValue.IsNegative()
	amount := decimalValue.Abs()

	wholeBaht := amount.Truncate(0)
	if wholeBaht.GreaterThan(maxBahtPart) {
		return "", fmt.Errorf("baht amount %s exceeds supported range", decimalValue.String())
	}
	bahtPart := wholeBaht.IntPart()

	satang := amount.Sub(wholeBaht).Mul(hundred)
	satangWhole := satang.Truncate(0)
	satangPart := satangWhole.IntPart()
	satangFraction := satang.Sub(satangWhole)

	var thai strings.Builder
	if negative {
		thai.WriteString("ลบ")
	}
	thai.WriteString(numberToThai(bahtPart))
	thai.WriteString("บาท")

	if satang.IsZero() {
		thai.WriteString("ถ้วน")
	} else {
		thai.WriteString(numberToThai(satangPart))
		if !satangFraction.IsZero() {
			thai.WriteString("จุด")
			thai.WriteString(decimalToThai(satangFraction))
		}
		thai.WriteString("สตางค์")
	}

	return thai.String(), nil
}

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(2222222.75),
		decimal.NewFromFloat(1111111111.5692222),
		decimal.NewFromFloat(-1234.56),
		decimal.NewFromFloat(0),
	}

	for _, input := range inputs {
		fmt.Println(input)
		thaiString, err := convertDecimalToString(input)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(thaiString)
	}
}
