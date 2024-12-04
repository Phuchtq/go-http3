package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func NormalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func GenerateId() string {
	return fmt.Sprint(rand.Int())
}
