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

type CoreRuleConfig json.RawMessage

func (r CoreRuleConfig) MarshalJSON() ([]byte, error) {
	return json.RawMessage(r).MarshalJSON()
}

func (r *CoreRuleConfig) UnmarshalJSON(data []byte) error {
	*r = CoreRuleConfig(data)
	return nil
}
