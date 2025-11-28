package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	DecimalValue = 100
	OctalValue   = 075
	HexValue     = 0x1F4
	FloatValue   = 2.71828432
	StringValue  = "go-core-task"
	BoolValue    = true
	ComplexValue = 34 + 4234i
)

const (
	hashSalt = "go-2024"
)

func main() {
	variables := getVariables()

	formattedVariables, err := formatVariablesWithTypes(variables)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range formattedVariables {
		fmt.Println(v)
	}

	combinedString, err := combineVariables(variables)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(combinedString)

	runes := []rune(combinedString)

	hashed, err := hashVariablesWithSalt(runes, hashSalt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hashed)
}

func getVariables() []any {
	return []any{
		DecimalValue,
		OctalValue,
		HexValue,
		FloatValue,
		StringValue,
		BoolValue,
		ComplexValue,
	}
}

// formatVariablesWithTypes форматирует срез переменных, возвращая срез строк с описанием значения и типа.
func formatVariablesWithTypes(variables []any) ([]string, error) {
	if len(variables) == 0 {
		return nil, errors.New("variables slice is empty")
	}

	result := make([]string, len(variables))

	for i, v := range variables {
		result[i] = fmt.Sprintf("Value: %v, Type: %T", v, v)
	}

	return result, nil
}

// combineVariables объединяет все переменные в одну строку через пробел.
func combineVariables(variables []any) (string, error) {
	if len(variables) == 0 {
		return "", errors.New("variables slice is empty")
	}

	var sb strings.Builder
	for _, v := range variables {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}

	return strings.TrimSpace(sb.String()), nil
}

// hashVariablesWithSalt вычисляет SHA-256 хеш от среза рун с добавлением соли посередине.
func hashVariablesWithSalt(runes []rune, salt string) (string, error) {
	if len(runes) == 0 {
		return "", errors.New("runes slice is empty")
	}
	if salt == "" {
		return "", errors.New("hashSalt string is empty")
	}

	mid := len(runes) / 2
	withSalt := make([]rune, 0, len(runes)+len(salt))

	withSalt = append(withSalt, runes[:mid]...)
	withSalt = append(withSalt, []rune(salt)...)
	withSalt = append(withSalt, runes[mid:]...)

	data := []byte(string(withSalt))
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}
