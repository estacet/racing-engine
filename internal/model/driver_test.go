package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewDriver(t *testing.T) {
	testCases := map[string]struct {
		age          *int
		wantCategory Category
	}{
		"age less than 18": {
			age:          intToIntPointer(15),
			wantCategory: Junior,
		},
		"age more than 18": {
			age:          intToIntPointer(20),
			wantCategory: Amateur,
		},
		"age is 18": {
			age:          intToIntPointer(18),
			wantCategory: Amateur,
		},
		"age undefined": {
			age:          nil,
			wantCategory: Amateur,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			driver := NewDriver(
				uuid.MustParse("f4239e11-8ae7-486a-b12a-0b37dc6cf7cb"),
				"Tester",
				"+380000000000",
				tc.age,
				nil,
			)

			if driver.Category != tc.wantCategory {
				t.Errorf("invalid category: want: %s, got: %s", tc.wantCategory, driver.Category)
			}
		})
	}
}

func intToIntPointer(n int) *int {
	return &n
}
