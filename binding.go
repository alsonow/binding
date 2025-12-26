// Package binding
// Copyright 2025 alsonow. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package binding

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSON binds the JSON payload from the request body to provided struct
// dst destination
func JSON(r *http.Request, dst any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %v", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, dst); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}

// Query binds the query string parameters to provided struct
// application/x-www-form-urlencoded
func Query(r *http.Request, dst any) error {
	return nil
}

// Form binds the form data to provided struct
func Form(r *http.Request, dst any) error {
	return nil
}

// MultipartForm binds multipart form data(including files) to provided struct
func MultipartForm(r *http.Request, dst any) error {
	return nil
}
