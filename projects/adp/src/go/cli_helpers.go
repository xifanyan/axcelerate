package adp

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

func ParseEngineTaxonomyArg(input string) (EngineTaxonomyArg, error) {
	return parseEngineTaxonomyArg(input)
}

func ParseOutputTaxonomies(input string) ([]OutputTaxonomiesArg, error) {
	return parseOutputTaxonomies(input)
}

func ParseConfigArgs(input string) ([]ConfigArg, error) {
	return parseConfigArgs(input)
}

func ParseBatchScriptParameters(input string) ([]CLIBatchParameter, error) {
	return parseBatchScriptParameters(input)
}

func PrettyJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func MustBaseURL(host string, port int, path string) string {
	host = strings.TrimSpace(host)
	if host == "" {
		panic("host is required")
	}

	path = strings.TrimSpace(path)
	if !strings.HasPrefix(host, "http://") && !strings.HasPrefix(host, "https://") {
		host = "https://" + host
	}

	parsed, err := url.Parse(host)
	if err != nil {
		panic(err)
	}
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}
	if parsed.Host == "" && parsed.Path != "" {
		parsed.Host = parsed.Path
		parsed.Path = ""
	}

	if parsed.Port() == "" {
		parsed.Host = net.JoinHostPort(parsed.Hostname(), strconv.Itoa(port))
	}

	if path != "" {
		if strings.HasPrefix(path, "/") {
			parsed.Path = path
		} else {
			parsed.Path = "/" + path
		}
	}

	return strings.TrimRight(parsed.String(), "/")
}

func parseEngineTaxonomyArg(input string) (EngineTaxonomyArg, error) {
	for _, separator := range []struct {
		value    string
		negation bool
	}{
		{value: "!=", negation: true},
		{value: "=", negation: false},
	} {
		if separator.value == "=" && strings.Contains(input, "!=") {
			continue
		}

		parts := strings.SplitN(input, separator.value, 2)
		if len(parts) != 2 {
			continue
		}

		taxonomy := strings.TrimSpace(parts[0])
		query := strings.TrimSpace(parts[1])
		if taxonomy == "" || query == "" {
			continue
		}

		return EngineTaxonomyArg{
			Taxonomy: taxonomy,
			Negation: separator.negation,
			Query:    query,
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
