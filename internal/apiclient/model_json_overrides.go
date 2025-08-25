package apiclient

import "encoding/json"

// this file contains overrides for the generated sdk code.
// since terraform uses json.RawMessage we don't include corresponding models of the sdk.
// instead we use the json.RawMessage type to represent the data.

type RuleRecipeComponents json.RawMessage

func (r RuleRecipeComponents) MarshalJSON() ([]byte, error) {
	return json.RawMessage(r).MarshalJSON()
}

func (r *RuleRecipeComponents) UnmarshalJSON(data []byte) error {
	*r = RuleRecipeComponents(data)
	return nil
}

type CoreRuleConfig struct {
	// Slug is readable in Go, but not emitted (we emit Raw below).
	Slug string `json:"-"`
	// Raw holds the entire original config JSON.
	Raw json.RawMessage `json:"-"`
}

func NewCoreRuleConfigFromString(s string) (CoreRuleConfig, error) {
	var c CoreRuleConfig
	if err := c.UnmarshalJSON([]byte(s)); err != nil {
		return CoreRuleConfig{}, err
	}
	return c, nil
}

func (c *CoreRuleConfig) UnmarshalJSON(data []byte) error {
	// Keep a copy of the entire config as-is for lossless round-tripping.
	c.Raw = append([]byte(nil), data...)

	// Best-effort extract of slug.
	var aux struct {
		Slug string `json:"app_slug"`
	}
	if err := json.Unmarshal(data, &aux); err == nil {
		c.Slug = aux.Slug
	} else {
		c.Slug = "" // not present or not an object
	}
	return nil
}

func (c CoreRuleConfig) MarshalJSON() ([]byte, error) {
	// Always send the exact original bytes if we have them.
	if c.Raw != nil {
		return c.Raw, nil
	}
	// Fallback if built programmatically without Raw.
	aux := struct {
		Slug string `json:"app_slug"`
	}{
		Slug: c.Slug,
	}
	return json.Marshal(aux)
}
