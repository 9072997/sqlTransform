package storage

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

func regExp(patternI, subjectI interface{}) (bool, error) {
	pattern := fmt.Sprintf("%v", patternI)
	subject := fmt.Sprintf("%v", subjectI)

	return regexp.MatchString(pattern, subject)
}

func regReplace(patternI, replacementI, subjectI interface{}) string {
	pattern := fmt.Sprintf("%v", patternI)
	replacement := fmt.Sprintf("%v", replacementI)
	subject := fmt.Sprintf("%v", subjectI)

	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return err.Error()
	}
	return matcher.ReplaceAllString(subject, replacement)
}

func regFind(patternI, subjectI interface{}) string {
	pattern := fmt.Sprintf("%v", patternI)
	subject := fmt.Sprintf("%v", subjectI)

	matcher, err := regexp.Compile(pattern)
	if err != nil {
		return err.Error()
	}
	return matcher.FindString(subject)
}

func splitPart(subjectI, deliminiterI interface{}, fieldNumber uint) string {
	subject := fmt.Sprintf("%v", subjectI)
	deliminiter := fmt.Sprintf("%v", deliminiterI)

	fields := strings.Split(subject, deliminiter)
	maxIndex := len(fields) - 1
	if fieldNumber == 0 {
		return "Can't get field 0 (fields start from 1)"
	}
	fieldIndex := int(fieldNumber) - 1
	if fieldIndex > maxIndex {
		return fmt.Sprintf(`Requested field %d but only %d fields in "%s"`,
			fieldNumber, len(fields), subject)
	}
	return fields[fieldIndex]
}

func parseDate(dateStringI, exampleI interface{}) string {
	dateString := fmt.Sprintf("%v", dateStringI)
	example := fmt.Sprintf("%v", exampleI)

	format, err := dateparse.ParseFormat(example)
	if err != nil {
		return fmt.Sprintf("Error parsing example: %s", err)
	}
	timezone, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Sprintf("Error getting timezone: %s", err)
	}
	date, err := time.ParseInLocation(format, dateString, timezone)
	if err != nil {
		return fmt.Sprintf("Error parseing date: %s", err)
	}
	return date.Format(time.RFC3339)
}