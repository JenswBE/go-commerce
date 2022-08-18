package entities

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AmountInCents struct {
	amount int
}

// FormatPrice formats a price in cents into a decimal string,
// without using float conversions.
func (a AmountInCents) String() string {
	// Fast path
	if a.amount == 0 {
		return "0.00"
	}

	// Determine sign
	sign := ""
	if a.amount < 0 {
		sign = "-"
	}

	// Format price
	absAmount := int(math.Abs(float64(a.amount)))
	return fmt.Sprintf("%s%d.%02d", sign, absAmount/100, absAmount%100)
}

func (a AmountInCents) Int() int {
	return a.amount
}

func NewAmountInCents(amountInCents int) AmountInCents {
	return AmountInCents{amount: amountInCents}
}

// NewAmountInCentsFromString converts a decimal string into an AmountInCents,
// without using float conversions
func NewAmountInCentsFromString(amount string) (AmountInCents, error) {
	// Return zero amount if string is empty
	if amount == "" {
		return NewAmountInCents(0), nil
	}

	// Determine sign of amount
	sign := 1
	if strings.HasPrefix(amount, "-") {
		sign = -1
		amount = strings.TrimPrefix(amount, "-")
	}

	// Extract amount parts
	amountParts := strings.Split(amount, ".")
	if len(amountParts) == 1 {
		amountParts = append(amountParts, "00")
	}
	if len(amountParts) != 2 {
		return AmountInCents{}, fmt.Errorf("expected decimal string with exactly a single dot as separator, received: %s", amount)
	}

	// Parse integer parts
	var intPart int
	var err error
	if amountParts[0] != "" && amountParts[0] != "0" {
		intPart, err = strconv.Atoi(amountParts[0])
		if err != nil {
			return AmountInCents{}, fmt.Errorf("failed to parse integer part of decimal string: %w", err)
		}
	}

	// Parse decimal part
	var decPart int
	switch len(amountParts[1]) {
	case 1:
		amountParts[1] += "0"
	case 2:
		// No action required
	default:
		return AmountInCents{}, errors.New("decimal part of decimal string must not have more than 2 numbers precision")
	}
	amountParts[1] = strings.TrimLeft(amountParts[1], "0")
	if amountParts[1] != "" {
		decPart, err = strconv.Atoi(amountParts[1])
		if err != nil {
			return AmountInCents{}, fmt.Errorf("failed to parse decimal part of decimal string: %w", err)
		}
	}

	// Build result
	return NewAmountInCents(sign * (intPart*100 + decPart)), nil
}
