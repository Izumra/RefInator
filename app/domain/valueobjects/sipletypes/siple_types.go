package simpletypes

import (
	"fmt"
	mathrand "math/rand/v2"
	"strings"
	"unicode"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	Int8 = iota
	UInt8
	Int16
	UInt16
	Int32
	UInt32
	Int64
	UInt64
	Int
	UInt
	Float
	Double
	Bool
	String
)

type Type struct {
	Value string
	Type  string
}

// Generates new perem one of presented types.
// If passed type param, create perem with passed type
/*
 * Int8
 * UInt8
 * Int16
 * UInt16
 * Int32
 * UInt32
 * Int64
 * UInt64
 * Int
 * UInt
 * Float
 * Double
 * Bool
 * String
 */
func New(t int) *Type {
	choosenType := mathrand.IntN(14)
	perem := &Type{}

	switch choosenType {
	case Int8:
		randValue := mathrand.IntN(127)

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%d", randValue)

	case UInt8:
		randValue := mathrand.IntN(254)
		perem.Value = fmt.Sprintf("%d", randValue)

	case Int16:
		randValue := gofakeit.Int16()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%d", randValue)

	case UInt16:
		randValue := mathrand.IntN(65534)

		perem.Value = fmt.Sprintf("%d", randValue)

	case Int32:
		randValue := gofakeit.Int32()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%d", randValue)

	case UInt32:
		randValue := mathrand.Uint32()

		perem.Value = fmt.Sprintf("%d", randValue)
	case Int64:
		randValue := mathrand.Int64()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%d", randValue)

	case UInt64:
		randValue := mathrand.Uint64()

		perem.Value = fmt.Sprintf("%d", randValue)

	case Int:
		randValue := mathrand.Int()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%d", randValue)

	case UInt:
		randValue := mathrand.Uint64()

		perem.Value = fmt.Sprintf("%d", randValue)

	case Float:
		randValue := gofakeit.Float32()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%4.f", randValue)
	case Double:
		randValue := gofakeit.Float64()

		typeSign := mathrand.IntN(2)
		if typeSign == 0 && randValue != 0 {
			randValue = 0 - randValue
		}

		perem.Value = fmt.Sprintf("%4.f", randValue)
	case Bool:
		randValue := false

		typeBool := mathrand.IntN(2)
		if typeBool == 0 {
			randValue = true
		}

		perem.Value = fmt.Sprintf("%v", randValue)

	case String:
		randValue := ""

		typeString := mathrand.IntN(5)
		switch typeString {
		case 0:
			randValue = strings.ReplaceAll(fmt.Sprintf(
				"%s%s%s",
				capitalizeFirstLetter(gofakeit.VerbAction()),
				gofakeit.City(),
				capitalizeFirstLetter(gofakeit.Animal()),
			), " ", "")
		case 1:
			randValue = strings.ReplaceAll(fmt.Sprintf(
				"%s%s%s",
				gofakeit.Adjective(),
				gofakeit.BookTitle(),
				gofakeit.Username(),
			), " ", "")
		case 2:
			randValue = strings.ReplaceAll(fmt.Sprintf(
				"%s%s%s",
				capitalizeFirstLetter(gofakeit.Animal()),
				gofakeit.Dog(),
				gofakeit.Word(),
			), " ", "")
		case 3:
			randValue = strings.ReplaceAll(fmt.Sprintf(
				"%s%s%s",
				gofakeit.Word(),
				gofakeit.BeerAlcohol(),
				capitalizeFirstLetter(gofakeit.Cat()),
			), " ", "")
		case 4:
			randValue = strings.ReplaceAll(fmt.Sprintf(
				"%s%s%s",
				gofakeit.Color(),
				gofakeit.VerbAction(),
				capitalizeFirstLetter(gofakeit.CelebrityActor()),
			), " ", "")
		}

		perem.Value = fmt.Sprintf("\"%s\"", randValue)
	}

	perem.Type = DetermineType(choosenType)

	//log.Printf("Сгенерированная переменная значение: %v, тип: %v", perem.Value, perem.Type)

	return perem
}

func DetermineType(typeId int) string {
	switch typeId {
	case Int8:
		return "Int8"
	case UInt8:
		return "UInt8"
	case Int16:
		return "Int16"
	case UInt16:
		return "UInt16"
	case Int32:
		return "Int32"
	case UInt32:
		return "UInt32"
	case Int64:
		return "Int64"
	case UInt64:
		return "UInt64"
	case Int:
		return "Int"
	case UInt:
		return "UInt"
	case Float:
		return "Float"
	case Double:
		return "Double"
	case Bool:
		return "Bool"
	case String:
		return "String"
	default:
		panic("Unknown type")
	}
}

func (t *Type) UpdateValueByRandomMathOp() string {
	return ""
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	for i := range runes {
		if i == 0 || !unicode.IsLetter(runes[i-1]) {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return string(runes)
}
