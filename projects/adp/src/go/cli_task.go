package adp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	cliTaskType        = "CLI"
	cliTaskDescription = "Runs a native task in its own process"
	cliTaskDisplayName = "Command Line Task"
)

type CLITaskBuilder struct {
	*builderCommon[CLITaskBuilder]
	client                            *Client
	batchScriptPath                   *string
	batchScriptParameters             *[]CLIBatchParameter
	workingDirectory                  *string
	batchScriptJSONLogOutput          *string
	batchScriptRedirectLogging        *bool
	batchScriptPositiveExecutionCodes *string
	batchScriptFilterPasswords        *bool
	batchScriptLoggingDirectory       *string
	batchScriptResultCode             *string
	batchScriptResultLogPath          *string
	batchScriptErrorLogPath           *string
}

func NewCLITaskBuilder(client *Client) *CLITaskBuilder {
	b := &CLITaskBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *CLITaskBuilder) BatchScriptPath(value string) *CLITaskBuilder {
	b.batchScriptPath = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptParameters(value []CLIBatchParameter) *CLITaskBuilder {
	b.batchScriptParameters = &value
	return b
}

func (b *CLITaskBuilder) WorkingDirectory(value string) *CLITaskBuilder {
	b.workingDirectory = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptJSONLogOutput(value string) *CLITaskBuilder {
	b.batchScriptJSONLogOutput = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptRedirectLogging(value bool) *CLITaskBuilder {
	b.batchScriptRedirectLogging = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptPositiveExecutionCodes(value string) *CLITaskBuilder {
	b.batchScriptPositiveExecutionCodes = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptFilterPasswords(value bool) *CLITaskBuilder {
	b.batchScriptFilterPasswords = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptLoggingDirectory(value string) *CLITaskBuilder {
	b.batchScriptLoggingDirectory = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptResultCode(value string) *CLITaskBuilder {
	b.batchScriptResultCode = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptResultLogPath(value string) *CLITaskBuilder {
	b.batchScriptResultLogPath = &value
	return b
}

func (b *CLITaskBuilder) BatchScriptErrorLogPath(value string) *CLITaskBuilder {
	b.batchScriptErrorLogPath = &value
	return b
}

func (b *CLITaskBuilder) buildRequest() (rawTaskRequest, error) {
	if b.batchScriptPath == nil || strings.TrimSpace(*b.batchScriptPath) == "" {
		return rawTaskRequest{}, errors.New("batchScriptPath is required")
	}

	cfg := map[string]any{
		"adp_batchScriptPath": *b.batchScriptPath,
	}
	if b.batchScriptParameters != nil && len(*b.batchScriptParameters) > 0 {
		cfg["adp_batchScriptParameters"] = *b.batchScriptParameters
	}
	if b.workingDirectory != nil {
		cfg["adp_workingDirectory"] = *b.workingDirectory
	}
	if b.batchScriptJSONLogOutput != nil {
		cfg["adp_batchScriptJsonLogOutput"] = *b.batchScriptJSONLogOutput
	}
	if b.batchScriptRedirectLogging != nil {
		cfg["adp_batchScriptRedirectLogging"] = strconv.FormatBool(*b.batchScriptRedirectLogging)
	}
	if b.batchScriptPositiveExecutionCodes != nil {
		cfg["adp_batchScriptPositiveExecutionCodes"] = *b.batchScriptPositiveExecutionCodes
	}
	if b.batchScriptFilterPasswords != nil {
		cfg["adp_batchScriptFilterPasswords"] = *b.batchScriptFilterPasswords
	}
	if b.batchScriptLoggingDirectory != nil {
		cfg["adp_batchScriptLoggingDirectory"] = *b.batchScriptLoggingDirectory
	}
	if b.batchScriptResultCode != nil {
		cfg["adp_batchScriptResultCode"] = *b.batchScriptResultCode
	}
	if b.batchScriptResultLogPath != nil {
		cfg["adp_batchScriptResultLogPath"] = *b.batchScriptResultLogPath
	}
	if b.batchScriptErrorLogPath != nil {
		cfg["adp_batchScriptErrorLogPath"] = *b.batchScriptErrorLogPath
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          cliTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   cliTaskDescription,
		TaskDisplayName:   cliTaskDisplayName,
	}, nil
}

func decodeCLITask(meta any) (CLIResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return CLIResult{}, err
	}

	resultValue, ok := obj["cli_result"]
	if !ok {
		return CLIResult{}, errors.New("missing cli_result")
	}

	var result int
	switch value := resultValue.(type) {
	case float64:
		if value != math.Trunc(value) {
			return CLIResult{}, errors.New("cli_result must be an integer")
		}
		if value < math.MinInt || value > math.MaxInt {
			return CLIResult{}, errors.New("cli_result out of range")
		}
		result = int(value)
	case int:
		result = value
	case json.Number:
		if strings.ContainsAny(value.String(), ".eE") {
			return CLIResult{}, errors.New("cli_result must be an integer")
		}
		parsed, ok := new(big.Int).SetString(value.String(), 10)
		if !ok {
			return CLIResult{}, fmt.Errorf("cli_result must be numeric, got %T", resultValue)
		}
		if !parsed.IsInt64() {
			return CLIResult{}, errors.New("cli_result out of range")
		}
		parsed64 := parsed.Int64()
		if strconv.IntSize == 32 && (parsed64 < math.MinInt32 || parsed64 > math.MaxInt32) {
			return CLIResult{}, errors.New("cli_result out of range")
		}
		result = int(parsed64)
	default:
		return CLIResult{}, fmt.Errorf("cli_result must be numeric, got %T", resultValue)
	}

	errorPath, err := stringField(obj, "cli_error_path")
	if err != nil {
		return CLIResult{}, err
	}

	resultPath, err := stringField(obj, "cli_result_path")
	if err != nil {
		return CLIResult{}, err
	}

	var jsonOutput map[string]any
	if value, ok := obj["json_output"]; ok {
		text, ok := value.(string)
		if !ok {
			return CLIResult{}, errors.New("json_output must be a string")
		}
		if err := json.Unmarshal([]byte(text), &jsonOutput); err != nil {
			return CLIResult{}, fmt.Errorf("parse json_output: %w", err)
		}
	}

	return CLIResult{
		Result:     result,
		JSONOutput: jsonOutput,
		ErrorPath:  errorPath,
		ResultPath: resultPath,
	}, nil
}

func (b *CLITaskBuilder) Execute(ctx context.Context) (CLIResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return CLIResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return CLIResult{}, err
	}

	return decodeCLITask(resp.ExecutionMetaData)
}
