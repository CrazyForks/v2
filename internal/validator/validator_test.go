// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package validator // import "miniflux.app/v2/internal/validator"

import (
	"testing"

	"miniflux.app/v2/internal/locale"
)

func TestIsValidURL(t *testing.T) {
	scenarios := map[string]bool{
		"https://www.example.org": true,
		"http://www.example.org/": true,
		"www.example.org":         false,
	}

	for link, expected := range scenarios {
		result := IsValidURL(link)
		if result != expected {
			t.Errorf(`Unexpected result, got %v instead of %v`, result, expected)
		}
	}
}

func TestValidateRange(t *testing.T) {
	if err := ValidateRange(-1, 0); err == nil {
		t.Error(`An invalid offset should generate a error`)
	}

	if err := ValidateRange(0, -1); err == nil {
		t.Error(`An invalid limit should generate a error`)
	}

	if err := ValidateRange(42, 42); err != nil {
		t.Error(`A valid offset and limit should not generate any error`)
	}
}

func TestValidateDirection(t *testing.T) {
	for _, status := range []string{"asc", "desc"} {
		if err := ValidateDirection(status); err != nil {
			t.Error(`A valid direction should not generate any error`)
		}
	}

	if err := ValidateDirection("invalid"); err == nil {
		t.Error(`An invalid direction should generate a error`)
	}
}

func TestIsValidRegex(t *testing.T) {
	scenarios := map[string]bool{
		"(?i)miniflux": true,
		"[":            false,
	}

	for expr, expected := range scenarios {
		result := IsValidRegex(expr)
		if result != expected {
			t.Errorf(`Unexpected result, got %v instead of %v`, result, expected)
		}
	}
}

func TestIsValidDomain(t *testing.T) {
	scenarios := map[string]bool{
		"example.org":          true,
		"example":              false,
		"example.":             false,
		"example..":            false,
		"mail.example.com:443": false,
		"*.example.com":        false,
	}

	for domain, expected := range scenarios {
		result := IsValidDomain(domain)
		if result != expected {
			t.Errorf(`Unexpected result, got %v instead of %v`, result, expected)
		}
	}
}

func TestValidateUsername(t *testing.T) {
	scenarios := map[string]*locale.LocalizedError{
		"jvoisin":          nil,
		"j.voisin":         nil,
		"j@vois.in":        nil,
		"invalid username": locale.NewLocalizedError("error.invalid_username"),
	}

	for username, expected := range scenarios {
		result := validateUsername(username)
		if expected == nil {
			if result != nil {
				t.Errorf(`got an unexpected error for %q instead of nil: %v`, username, result)
			}
		} else {
			if result == nil {
				t.Errorf(`expected an error, got nil.`)
			}
		}
	}
}
