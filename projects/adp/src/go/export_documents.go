package adp

import "context"

const (
	exportDocumentsTaskType        = "Export Documents"
	exportDocumentsTaskDescription = "Export documents in CSV format."
	exportDocumentsTaskDisplayName = "Export documents task"
)

type ExportDocumentsBuilder struct {
	*builderCommon[ExportDocumentsBuilder]
	client                *Client
	fieldSeparator        *string
	waitForExport         *bool
	query                 *string
	applicationIdentifier *string
	applicationType       *string
	engineIdentifier      *string
	engineUser            *string
	enginePassword        *string
	exportName            *string
	exportFields          *string
	exportDirectory       *string
	fileEnding            *string
}

func NewExportDocumentsBuilder(client *Client) *ExportDocumentsBuilder {
	b := &ExportDocumentsBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *ExportDocumentsBuilder) FieldSeparator(value string) *ExportDocumentsBuilder {
	b.fieldSeparator = &value
	return b
}

func (b *ExportDocumentsBuilder) WaitForExport(value bool) *ExportDocumentsBuilder {
	b.waitForExport = &value
	return b
}

func (b *ExportDocumentsBuilder) Query(value string) *ExportDocumentsBuilder {
	b.query = &value
	return b
}

func (b *ExportDocumentsBuilder) ApplicationIdentifier(value string) *ExportDocumentsBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *ExportDocumentsBuilder) ApplicationType(value string) *ExportDocumentsBuilder {
	b.applicationType = &value
	return b
}

func (b *ExportDocumentsBuilder) EngineIdentifier(value string) *ExportDocumentsBuilder {
	b.engineIdentifier = &value
	return b
}

func (b *ExportDocumentsBuilder) EngineUser(value string) *ExportDocumentsBuilder {
	b.engineUser = &value
	return b
}

func (b *ExportDocumentsBuilder) EnginePassword(value string) *ExportDocumentsBuilder {
	b.enginePassword = &value
	return b
}

func (b *ExportDocumentsBuilder) ExportName(value string) *ExportDocumentsBuilder {
	b.exportName = &value
	return b
}

func (b *ExportDocumentsBuilder) ExportFields(value string) *ExportDocumentsBuilder {
	b.exportFields = &value
	return b
}

func (b *ExportDocumentsBuilder) ExportDirectory(value string) *ExportDocumentsBuilder {
	b.exportDirectory = &value
	return b
}

func (b *ExportDocumentsBuilder) FileEnding(value string) *ExportDocumentsBuilder {
	b.fileEnding = &value
	return b
}

func (b *ExportDocumentsBuilder) buildRequest() (rawTaskRequest, error) {
	cfg := map[string]any{}
	if b.fieldSeparator != nil {
		cfg["adp_exportDocuments_field_separator"] = *b.fieldSeparator
	}
	if b.waitForExport != nil {
		cfg["adp_exportDocuments_waitForExport"] = *b.waitForExport
	}
	if b.query != nil {
		cfg["adp_exportDocuments_query"] = *b.query
	}
	if b.applicationIdentifier != nil {
		cfg["adp_exportDocuments_applicationIdentifier"] = *b.applicationIdentifier
	}
	if b.applicationType != nil {
		cfg["adp_exportDocuments_applicationType"] = *b.applicationType
	}
	if b.engineIdentifier != nil {
		cfg["adp_exportDocuments_engineIdentifier"] = *b.engineIdentifier
	}
	if b.engineUser != nil {
		cfg["adp_exportDocuments_engineUser"] = *b.engineUser
	}
	if b.enginePassword != nil {
		cfg["adp_exportDocuments_enginePassword"] = *b.enginePassword
	}
	if b.exportName != nil {
		cfg["adp_exportDocuments_exportName"] = *b.exportName
	}
	if b.exportFields != nil {
		cfg["adp_exportDocuments_exportFields"] = *b.exportFields
	}
	if b.exportDirectory != nil {
		cfg["adp_exportDocuments_exportDirectory"] = *b.exportDirectory
	}
	if b.fileEnding != nil {
		cfg["adp_exportDocuments_File_Ending"] = *b.fileEnding
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          exportDocumentsTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   exportDocumentsTaskDescription,
		TaskDisplayName:   exportDocumentsTaskDisplayName,
	}, nil
}

func decodeExportDocuments(meta any) (ExportDocumentsResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	exportFileName, err := stringField(obj, "adp_exportDocuments_exportFileName")
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	exportPath, err := stringField(obj, "adp_exportDocuments_exportPath")
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	searchResultSize, err := intStringField(obj, "adp_exportDocuments_searchResultSize")
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	return ExportDocumentsResult{
		ExportFileName:   exportFileName,
		ExportPath:       exportPath,
		SearchResultSize: searchResultSize,
	}, nil
}

func (b *ExportDocumentsBuilder) Execute(ctx context.Context) (ExportDocumentsResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return ExportDocumentsResult{}, err
	}

	return decodeExportDocuments(resp.ExecutionMetaData)
}
