package adp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func metaObject(meta any) (map[string]any, error) {
	switch value := meta.(type) {
	case map[string]any:
		return value, nil
	case []any:
		if len(value) == 0 {
			return map[string]any{}, nil
		}
	}

	return nil, fmt.Errorf("executionMetaData must be an object or empty array, got %T", meta)
}

func stringField(meta map[string]any, key string) (string, error) {
	value, ok := meta[key]
	if !ok {
		return "", fmt.Errorf("missing %s", key)
	}

	text, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("%s must be a string", key)
	}

	return text, nil
}

func intStringField(meta map[string]any, key string) (int, error) {
	text, err := stringField(meta, key)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", key, err)
	}

	return value, nil
}

func jsonStringField(meta map[string]any, key string, target any) error {
	text, err := stringField(meta, key)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(text), target); err != nil {
		return fmt.Errorf("parse %s: %w", key, err)
	}

	return nil
}
