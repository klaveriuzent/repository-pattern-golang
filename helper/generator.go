package helper

import (
	"errors"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const userIdPrefix = "USR"

func GenerateUserId(numberOfDigits int) (string, error) {
	prefix := userIdPrefix
	uniqueNumber, err := GenerateRandomSecureToken(numberOfDigits)
	if err != nil {
		return "", err
	}
	uniqueYear := getLastTwoDigitsOfYear()
	day := getDayOfMonth()
	timeValue := getCurrentTime()

	var builder strings.Builder
	builder.WriteString(prefix)
	builder.WriteString(strconv.Itoa(uniqueNumber))
	builder.WriteString(uniqueYear)
	builder.WriteString(day)
	builder.WriteString(timeValue)

	return builder.String(), nil
}

func GenerateRandomSecureToken(numberOfDigits int) (int, error) {
	if numberOfDigits < 1 || numberOfDigits > 9 {
		return 0, errors.New("numberOfDigits must be between 1 and 9")
	}

	rand.Seed(time.Now().UnixNano())
	maxLimit := int(math.Pow10(numberOfDigits))
	randomNumber := rand.Intn(maxLimit)
	return randomNumber, nil
}

func getLastTwoDigitsOfYear() string {
	currentTime := time.Now()
	lastTwoDigits := currentTime.Format("06")
	return lastTwoDigits
}

func getDayOfMonth() string {
	currentTime := time.Now()
	dayOfMonth := strconv.Itoa(currentTime.Day())
	return dayOfMonth
}

func getCurrentTime() string {
	currentTime := time.Now()
	timeValue := currentTime.Format("1504")
	return timeValue
}
