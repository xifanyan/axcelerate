package adp

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseEngineTaxonomyArg(input string) (EngineTaxonomyArg, error) {
	for _, separator := range []struct {
		value    string
		negation bool
	}{
		{value: "!=", negation: true},
		{value: "=", negation: false},
	} {
		parts := strings.SplitN(input, separator.value, 2)
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			continue
		}

		return EngineTaxonomyArg{
			Taxonomy: parts[0],
			Negation: separator.negation,
			Query:    parts[1],
		}, nil
	}

	return EngineTaxonomyArg{}, fmt.Errorf("invalid engine taxonomy arg: %q", input)
}

func parseOutputTaxonomies(input string) ([]OutputTaxonomiesArg, error) {
	if input == "" {
		return nil, nil
	}

	if strings.HasPrefix(strings.TrimSpace(input), "[") {
		var result []OutputTaxonomiesArg
		if err := json.Unmarshal([]byte(input), &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	parts := strings.Split(input, ",")
	result := make([]OutputTaxonomiesArg, 0, len(parts))
	for _, part := range parts {
		taxonomy := strings.TrimSpace(part)
		if taxonomy == "" {
			continue
		}
		result = append(result, OutputTaxonomiesArg{
			Taxonomy:                  taxonomy,
			Mode:                      "Category counts",
			MaximumNumberOfCategories: 10,
		})
	}

	return result, nil
}

func parseConfigArgs(input string) ([]ConfigArg, error) {
	if input == "" {
		return nil, nil
	}

	var raw []struct {
		ConfigurationID       string `json:"Configuration ID"`
		DynamicComponentNames string `json:"Dynamic Component Names"`
		FieldList             string `json:"Field list"`
		NameValueList         string `json:"Name value list"`
		ApplicationType       string `json:"Application type"`
		EntityType            string `json:"Entity type"`
	}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return nil, err
	}

	result := make([]ConfigArg, len(raw))
	for i, item := range raw {
		result[i] = ConfigArg{
			ConfigurationID:       item.ConfigurationID,
			DynamicComponentNames: item.DynamicComponentNames,
			FieldList:             item.FieldList,
			NameValueList:         item.NameValueList,
			ApplicationType:       item.ApplicationType,
			EntityType:            item.EntityType,
		}
	}

	return result, nil
}

func parseBatchScriptParameters(input string) ([]CLIBatchParameter, error) {
	if input == "" {
		return nil, nil
	}

	var result []CLIBatchParameter
	if err := json.Unmarshal([]byte(input), &result); err != nil {
		return nil, err
	}

	return result, nil
}
