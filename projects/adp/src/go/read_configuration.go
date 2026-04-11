package adp

import "context"

const (
	readConfigurationTaskType        = "Read Configuration"
	readConfigurationTaskDescription = "A Task to read configurations into JSON or XML."
	readConfigurationTaskDisplayName = "Read Configuration"
)

type ReadConfigurationBuilder struct {
	*builderCommon[ReadConfigurationBuilder]
	client         *Client
	entityIDToRead *string
	configsToRead  *[]ConfigArg
	fileFormat     *string
}

func NewReadConfigurationBuilder(client *Client) *ReadConfigurationBuilder {
	b := &ReadConfigurationBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *ReadConfigurationBuilder) EntityIDToRead(value string) *ReadConfigurationBuilder {
	b.entityIDToRead = &value
	return b
}

func (b *ReadConfigurationBuilder) ConfigsToRead(value []ConfigArg) *ReadConfigurationBuilder {
	b.configsToRead = &value
	return b
}

func (b *ReadConfigurationBuilder) FileFormat(value string) *ReadConfigurationBuilder {
	b.fileFormat = &value
	return b
}

func readConfigRaw(items []ConfigArg) []map[string]any {
	raw := make([]map[string]any, 0, len(items))
	for _, item := range items {
		entry := map[string]any{}
		if item.ConfigurationID != "" {
			entry["Configuration ID"] = item.ConfigurationID
		}
		if item.DynamicComponentNames != "" {
			entry["Dynamic Component Names"] = item.DynamicComponentNames
		}
		if item.FieldList != "" {
			entry["Field list"] = item.FieldList
		}
		if item.NameValueList != "" {
			entry["Name value list"] = item.NameValueList
		}
		if item.ApplicationType != "" {
			entry["Application type"] = item.ApplicationType
		}
		if item.EntityType != "" {
			entry["Entity type"] = item.EntityType
		}
		raw = append(raw, entry)
	}
	return raw
}

func (b *ReadConfigurationBuilder) buildRequest() (rawTaskRequest, error) {
	cfg := map[string]any{}
	if b.entityIDToRead != nil {
		cfg["adp_readConfiguration_entityIdToRead"] = *b.entityIDToRead
	}
	if b.configsToRead != nil && len(*b.configsToRead) > 0 {
		cfg["adp_readConfiguration_configsToRead"] = readConfigRaw(*b.configsToRead)
	}
	if b.fileFormat != nil {
		cfg["adp_readConfiguration_fileFormat"] = *b.fileFormat
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          readConfigurationTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   readConfigurationTaskDescription,
		TaskDisplayName:   readConfigurationTaskDisplayName,
	}, nil
}

func decodeReadConfiguration(meta any) (ReadConfigurationResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return ReadConfigurationResult{}, err
	}

	outputFile, err := stringField(obj, "adp_entities_output_file_name")
	if err != nil {
		return ReadConfigurationResult{}, err
	}

	configuration := map[string]ConfigurationInfo{}
	if err := jsonStringField(obj, "adp_entities_json_output", &configuration); err != nil {
		return ReadConfigurationResult{}, err
	}

	return ReadConfigurationResult{
		OutputFile:    outputFile,
		Configuration: configuration,
	}, nil
}

func (b *ReadConfigurationBuilder) Execute(ctx context.Context) (ReadConfigurationResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return ReadConfigurationResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return ReadConfigurationResult{}, err
	}

	return decodeReadConfiguration(resp.ExecutionMetaData)
}
