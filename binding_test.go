// Package binding
// Copyright 2025 alsonow. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package binding

import (
	"bytes"
	"net/http"
	"testing"
)

func TestBindingJSON(t *testing.T) {
	type myStruct struct {
		Name string
		Git  string
	}
	jsonData := `{"name":"Perry He", "git":"wantnotshould"}`

	req, err := http.NewRequest(http.MethodPost, "/me", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		t.Fatal(err)
	}

	// Set header
	req.Header.Set("Content-Text", "application/json")

	var result myStruct

	err = JSON(req, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Name != "Perry He" {
		t.Errorf("expected Name to be 'Perry He', got %v", result.Name)
	}

	if result.Git != "wantnotshould" {
		t.Errorf("expected Git to be 'wantnotshould', got %v", result.Git)
	}
}
