# convert-decimal-to-thai-baht-string

A Go program that converts a decimal number into its Thai baht text representation (การอ่านจำนวนเงินเป็นภาษาไทย), similar to how amounts are written out on Thai checks and invoices.

For example:
- `1234` → `หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน`
- `2222222.75` → `สองล้านสองแสนสองหมื่นสองพันสองร้อยยี่สิบสองบาทเจ็ดสิบห้าสตางค์`
- `-1234.56` → `ลบหนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์`

## How it works

- `numberToThai` converts a non-negative integer into Thai number words, handling special cases like `เอ็ด` (one, at the end of a number) and `ยี่` (two, in the tens place).
- `decimalToThai` renders any leftover fractional digits (beyond satang) as spoken-out digits after `จุด`.
- `convertDecimalToString` combines the baht (whole number) and satang (fractional/cents) parts into the full Thai string, handling negative amounts (`ลบ`) and exact amounts (`ถ้วน`).

Amounts are parsed with [`github.com/shopspring/decimal`](https://github.com/shopspring/decimal) to avoid floating-point rounding issues.

## Requirements

- Go 1.26.5 or later

## Running the code

Clone the repository and run the program directly:

```sh
go run convert-decimal.go
```

This runs the sample inputs defined in `main()` and prints each decimal value alongside its Thai baht string.

Alternatively, build a binary:

```sh
go build -o convert-decimal .
./convert-decimal
```

## Dependencies

Fetch dependencies (if not already present) with:

```sh
go mod download
```
