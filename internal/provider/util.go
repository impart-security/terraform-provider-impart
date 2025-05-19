package provider

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"reflect"
)

func calculateSha256(input string) (string, error) {
	h := sha256.New()

	_, err := h.Write([]byte(input))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// equalJSON compares two JSON strings for structural equality
func equalJSON(a, b string) bool {
	var ja, jb interface{}
	if err := json.Unmarshal([]byte(a), &ja); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &jb); err != nil {
		return false
	}
	return reflect.DeepEqual(ja, jb)
}

func subsetEqualJSON(a, b string) bool {
	var ja, jb interface{}

	if err := json.Unmarshal([]byte(a), &ja); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &jb); err != nil {
		return false
	}

	return deepSubsetEqual(ja, jb)
}

func deepSubsetEqual(a, b interface{}) bool {
	switch aVal := a.(type) {
	case map[string]interface{}:
		bVal, ok := b.(map[string]interface{})
		if !ok {
			return false
		}
		for key, aSub := range aVal {
			bSub, exists := bVal[key]
			if !exists || !deepSubsetEqual(aSub, bSub) {
				return false
			}
		}
		return true

	case []interface{}:
		bVal, ok := b.([]interface{})
		if !ok || len(aVal) > len(bVal) {
			return false
		}
		// Ensure all elements of `a` exist in the same positions in `b`
		for i := range aVal {
			if !deepSubsetEqual(aVal[i], bVal[i]) {
				return false
			}
		}
		return true

	default:
		return reflect.DeepEqual(a, b)
	}
}
