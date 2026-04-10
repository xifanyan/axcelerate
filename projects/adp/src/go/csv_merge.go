package adp

import (
	"context"
	"errors"
	"strconv"
)

const (
	csvMergeTaskType        = "CSV Merge"
	csvMergeTaskDescription = "Merges metaData or natives/images by using a csv file."
	csvMergeTaskDisplayName = "Csv merge task"
)

type CSVMergeBuilder struct {
	*builderCommon[CSVMergeBuilder]
	client                        *Client
	csvFile                       *string
	csvIDFieldKey                 *string
	mergeType                     *string
	csvMode                       *string
	engineName                    *string
	engineUser                    *string
	enginePassword                *string
	engineIDFieldKey              *string
	applicationIdentifier         *string
	fieldMappings                 *string
	fieldSeparator                *string
	imageBasePath                 *string
	nativeBasePath                *string
	csvFieldImageLocation         *string
	csvFieldNativeLocation        *string
	multiValueDelimiter           *string
	textIndicator                 *string
	doNotChangeProtectedDocuments *bool
}

func NewCSVMergeBuilder(client *Client) *CSVMergeBuilder {
	b := &CSVMergeBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *CSVMergeBuilder) CSVFile(value string) *CSVMergeBuilder {
	b.csvFile = &value
	return b
}

func (b *CSVMergeBuilder) CSVIDFieldKey(value string) *CSVMergeBuilder {
	b.csvIDFieldKey = &value
	return b
}

func (b *CSVMergeBuilder) MergeType(value string) *CSVMergeBuilder {
	b.mergeType = &value
	return b
}

func (b *CSVMergeBuilder) CSVMode(value string) *CSVMergeBuilder {
	b.csvMode = &value
	return b
}

func (b *CSVMergeBuilder) EngineName(value string) *CSVMergeBuilder {
	b.engineName = &value
	return b
}

func (b *CSVMergeBuilder) EngineUser(value string) *CSVMergeBuilder {
	b.engineUser = &value
	return b
}

func (b *CSVMergeBuilder) EnginePassword(value string) *CSVMergeBuilder {
	b.enginePassword = &value
	return b
}

func (b *CSVMergeBuilder) EngineIDFieldKey(value string) *CSVMergeBuilder {
	b.engineIDFieldKey = &value
	return b
}

func (b *CSVMergeBuilder) ApplicationIdentifier(value string) *CSVMergeBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *CSVMergeBuilder) FieldMappings(value string) *CSVMergeBuilder {
	b.fieldMappings = &value
	return b
}

func (b *CSVMergeBuilder) FieldSeparator(value string) *CSVMergeBuilder {
	b.fieldSeparator = &value
	return b
}

func (b *CSVMergeBuilder) ImageBasePath(value string) *CSVMergeBuilder {
	b.imageBasePath = &value
	return b
}

func (b *CSVMergeBuilder) NativeBasePath(value string) *CSVMergeBuilder {
	b.nativeBasePath = &value
	return b
}

func (b *CSVMergeBuilder) CSVFieldImageLocation(value string) *CSVMergeBuilder {
	b.csvFieldImageLocation = &value
	return b
}

func (b *CSVMergeBuilder) CSVFieldNativeLocation(value string) *CSVMergeBuilder {
	b.csvFieldNativeLocation = &value
	return b
}

func (b *CSVMergeBuilder) MultiValueDelimiter(value string) *CSVMergeBuilder {
	b.multiValueDelimiter = &value
	return b
}

func (b *CSVMergeBuilder) TextIndicator(value string) *CSVMergeBuilder {
	b.textIndicator = &value
	return b
}

func (b *CSVMergeBuilder) DoNotChangeProtectedDocuments(value bool) *CSVMergeBuilder {
	b.doNotChangeProtectedDocuments = &value
	return b
}

func (b *CSVMergeBuilder) buildRequest() (rawTaskRequest, error) {
	if b.csvFile == nil || *b.csvFile == "" {
		return rawTaskRequest{}, errors.New("csvFile is required")
	}

	cfg := map[string]any{
		"adp_csvMerge_csvFile": *b.csvFile,
	}
	if b.csvIDFieldKey != nil {
		cfg["adp_csvMerge_csvIdFieldKey"] = *b.csvIDFieldKey
	}
	if b.mergeType != nil {
		cfg["adp_csvMerge_mergeType"] = *b.mergeType
	}
	if b.csvMode != nil {
		cfg["adp_csvMerge_csvMode"] = *b.csvMode
	}
	if b.engineName != nil {
		cfg["adp_csvMerge_engineName"] = *b.engineName
	}
	if b.engineUser != nil {
		cfg["adp_csvMerge_engineUser"] = *b.engineUser
	}
	if b.enginePassword != nil {
		cfg["adp_csvMerge_enginePassword"] = *b.enginePassword
	}
	if b.engineIDFieldKey != nil {
		cfg["adp_csvMerge_engineIdFieldKey"] = *b.engineIDFieldKey
	}
	if b.applicationIdentifier != nil {
		cfg["adp_csvMerge_applicationIdentifier"] = *b.applicationIdentifier
	}
	if b.fieldMappings != nil {
		cfg["adp_csvMerge_fieldMappings"] = *b.fieldMappings
	}
	if b.fieldSeparator != nil {
		cfg["adp_csvMerge_fieldSeperator"] = *b.fieldSeparator
	}
	if b.imageBasePath != nil {
		cfg["adp_csvMerge_imageBasePath"] = *b.imageBasePath
	}
	if b.nativeBasePath != nil {
		cfg["adp_csvMerge_nativeBasePath"] = *b.nativeBasePath
	}
	if b.csvFieldImageLocation != nil {
		cfg["adp_csvMerge_csvFieldImageLocation"] = *b.csvFieldImageLocation
	}
	if b.csvFieldNativeLocation != nil {
		cfg["adp_csvMerge_csvFieldNativeLocation"] = *b.csvFieldNativeLocation
	}
	if b.multiValueDelimiter != nil {
		cfg["adp_csvMerge_multiValueDelimiter"] = *b.multiValueDelimiter
	}
	if b.textIndicator != nil {
		cfg["adp_csvMerge_textIndicator"] = *b.textIndicator
	}
	if b.doNotChangeProtectedDocuments != nil {
		cfg["adp_csvMerge_doNotChangeProtectedDocuments"] = strconv.FormatBool(*b.doNotChangeProtectedDocuments)
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          csvMergeTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   csvMergeTaskDescription,
		TaskDisplayName:   csvMergeTaskDisplayName,
	}, nil
}

func decodeCSVMerge(meta any) (CSVMergeResult, error) {
	_, err := metaObject(meta)
	if err != nil {
		return CSVMergeResult{}, err
	}

	return CSVMergeResult{}, nil
}

func (b *CSVMergeBuilder) Execute(ctx context.Context) (CSVMergeResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return CSVMergeResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return CSVMergeResult{}, err
	}

	return decodeCSVMerge(resp.ExecutionMetaData)
}
