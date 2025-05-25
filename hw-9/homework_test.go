package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func Serialize(person Person) string {
	b := &strings.Builder{}
	obj := reflect.TypeOf(person)

	for i := range obj.NumField() {
		fName := obj.Field(i).Tag.Get("properties")
		fTagList := strings.Split(fName, ",")
		fVal := reflect.ValueOf(person).Field(i)

		if fVal.IsZero() && slices.Contains(fTagList, "omitempty") || fName == "" {
			continue
		}

		_, _ = b.WriteString(fmt.Sprintf("%s=%v\n", fTagList[0], fVal))
	}

	return strings.TrimSpace(b.String())
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
