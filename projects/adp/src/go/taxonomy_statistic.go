package adp

import (
	"context"
	"errors"
	"strconv"
	"time"
)

const (
	taxonomyStatisticTaskType        = "Taxonomy Statistic"
	taxonomyStatisticTaskDescription = "Retrieves category counts for a taxonomy"
	taxonomyStatisticTaskDisplayName = "Taxonomy statistic"
)

type TaxonomyStatisticBuilder struct {
	*builderCommon[TaxonomyStatisticBuilder]
	client                 *Client
	engineName             *string
	engineQuery            *string
	computeCounts          *bool
	listCategoryProperties *bool
	engineTaxonomies       *[]EngineTaxonomyArg
	outputTaxonomies       *[]OutputTaxonomiesArg
	applicationIdentifier  *string
}

func NewTaxonomyStatisticBuilder(client *Client) *TaxonomyStatisticBuilder {
	b := &TaxonomyStatisticBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *TaxonomyStatisticBuilder) EngineName(value string) *TaxonomyStatisticBuilder {
	b.engineName = &value
	return b
}

func (b *TaxonomyStatisticBuilder) EngineQuery(value string) *TaxonomyStatisticBuilder {
	b.engineQuery = &value
	return b
}

func (b *TaxonomyStatisticBuilder) ComputeCounts(value bool) *TaxonomyStatisticBuilder {
	b.computeCounts = &value
	return b
}

func (b *TaxonomyStatisticBuilder) ListCategoryProperties(value bool) *TaxonomyStatisticBuilder {
	b.listCategoryProperties = &value
	return b
}

func (b *TaxonomyStatisticBuilder) EngineTaxonomies(value []EngineTaxonomyArg) *TaxonomyStatisticBuilder {
	b.engineTaxonomies = &value
	return b
}

func (b *TaxonomyStatisticBuilder) OutputTaxonomies(value []OutputTaxonomiesArg) *TaxonomyStatisticBuilder {
	b.outputTaxonomies = &value
	return b
}

func (b *TaxonomyStatisticBuilder) ApplicationIdentifier(value string) *TaxonomyStatisticBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *TaxonomyStatisticBuilder) buildRequest() (rawTaskRequest, error) {
	hasEngine := b.engineName != nil && *b.engineName != ""
	hasApplication := b.applicationIdentifier != nil && *b.applicationIdentifier != ""
	if !hasEngine && !hasApplication {
		return rawTaskRequest{}, errors.New("exactly one of engineName or applicationIdentifier is required")
	}
	if hasEngine && hasApplication {
		return rawTaskRequest{}, errors.New("engineName and applicationIdentifier are mutually exclusive")
	}

	cfg := map[string]any{}
	if hasEngine {
		cfg["adp_taxonomyStatistic_engineName"] = *b.engineName
	}
	if hasApplication {
		cfg["adp_taxonomyStatistic_engineName"] = ""
		cfg["adp_taxonomyStatistic_applicationIdentifier"] = *b.applicationIdentifier
	}
	if b.engineQuery != nil {
		cfg["adp_taxonomyStatistic_engineQuery"] = *b.engineQuery
	}
	if b.computeCounts != nil {
		cfg["adp_taxonomyStatistic_computeCounts"] = strconv.FormatBool(*b.computeCounts)
	}
	if b.listCategoryProperties != nil {
		cfg["adp_taxonomyStatistic_listCategoryProperties"] = strconv.FormatBool(*b.listCategoryProperties)
	}
	if b.engineTaxonomies != nil {
		cfg["adp_taxonomyStatistic_engineTaxonomies"] = *b.engineTaxonomies
	}
	if b.outputTaxonomies != nil {
		cfg["adp_taxonomyStatistic_outputTaxonomies"] = *b.outputTaxonomies
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          taxonomyStatisticTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   taxonomyStatisticTaskDescription,
		TaskDisplayName:   taxonomyStatisticTaskDisplayName,
	}, nil
}

func decodeTaxonomyStatistic(meta any) (TaxonomyStatisticResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	outputFile, err := stringField(obj, "adp_taxonomy_statistics_json_file_path")
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	var statistics StatisticsDocument
	if err := jsonStringField(obj, "adp_taxonomy_statistics_json_output", &statistics); err != nil {
		return TaxonomyStatisticResult{}, err
	}

	return TaxonomyStatisticResult{
		OutputFile: outputFile,
		Statistics: statistics,
	}, nil
}

func (b *TaxonomyStatisticBuilder) Execute(ctx context.Context) (TaxonomyStatisticResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	return decodeTaxonomyStatistic(resp.ExecutionMetaData)
}

func (b *TaxonomyStatisticBuilder) Start(ctx context.Context) (*TaskResponse, error) {
	req, err := b.buildRequest()
	if err != nil {
		return nil, err
	}

	return b.client.execute(ctx, "/executeAdpTaskAsync", req)
}

func (b *TaxonomyStatisticBuilder) Wait(ctx context.Context, interval time.Duration) (TaxonomyStatisticResult, error) {
	resp, err := b.Start(ctx)
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	resp, err = b.client.Wait(ctx, resp.ExecutionID, interval)
	if err != nil {
		return TaxonomyStatisticResult{}, err
	}

	return decodeTaxonomyStatistic(resp.ExecutionMetaData)
}
