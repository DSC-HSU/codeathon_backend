package utils

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of the given length
func RandomString(length int) string {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(runes)
}

// GenerateRandomEmail creates a random email address
func GenerateRandomEmail() string {
	username := RandomString(rand.Intn(10) + 5) // Random username between 5 and 15 characters
	domain := RandomString(rand.Intn(5) + 3)    // Random domain name between 3 and 8 characters
	tld := RandomString(3)                      // Random TLD of 3 characters
	return fmt.Sprintf("%s@%s.%s", strings.ToLower(username), strings.ToLower(domain), strings.ToLower(tld))
}

func GenerateRandomUuid() string {
	// Generate a random UUID
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return id.String()
}
