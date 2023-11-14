package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/raafly/inventory-management/pkg/helper"
)

func randomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

func TestRandomString(t *testing.T) {
	id := randomString(10)
	fmt.Println(id)
}

func TestId(t *testing.T) {
	id := helper.GenerateUUID()
	fmt.Println(id)
}