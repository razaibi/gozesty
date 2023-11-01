package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func normalizeJSON(jsonStr string) (string, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return "", err
	}

	normalizedJSON, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(normalizedJSON), nil
}

func CompareJSON(obj1, obj2 string) bool {

	// Normalize JSON strings
	normalizedStr1, err := normalizeJSON(obj1)
	if err != nil {
		fmt.Println("Error normalizing JSON string 1:", err)
		return false
	}

	normalizedStr2, err := normalizeJSON(obj2)
	if err != nil {
		fmt.Println("Error normalizing JSON string 2:", err)
		return false
	}
	// Compare the normalized strings
	if strings.Compare(normalizedStr1, normalizedStr2) == 0 {
		return true
	} else {
		fmt.Println("The response payloads is not as expected.")
		return false
	}
}
