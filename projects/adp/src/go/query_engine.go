package adp

import (
	"context"
	"errors"
	"time"
)

const (
	queryEngineTaskType        = "Query Engine"
	queryEngineTaskDescription = "Queries an engine"
	queryEngineTaskDisplayName = "Query engine"
)

type QueryEngineBuilder struct {
	*builderCommon[QueryEngineBuilder]
	client                *Client
	engineName            *string
	engineQuery           *string
	engineUserName        *string
	engineUserPassword    *string
	engineTaxonomies      *[]EngineTaxonomyArg
	applicationIdentifier *string
}

func NewQueryEngineBuilder(client *Client) *QueryEngineBuilder {
	b := &QueryEngineBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *QueryEngineBuilder) EngineName(value string) *QueryEngineBuilder {
	b.engineName = &value
	return b
}

func (b *QueryEngineBuilder) EngineQuery(value string) *QueryEngineBuilder {
	b.engineQuery = &value
	return b
}

func (b *QueryEngineBuilder) EngineUserName(value string) *QueryEngineBuilder {
	b.engineUserName = &value
	return b
}

func (b *QueryEngineBuilder) EngineUserPassword(value string) *QueryEngineBuilder {
	b.engineUserPassword = &value
	return b
}

func (b *QueryEngineBuilder) EngineTaxonomies(value []EngineTaxonomyArg) *QueryEngineBuilder {
	b.engineTaxonomies = &value
	return b
}

func (b *QueryEngineBuilder) ApplicationIdentifier(value string) *QueryEngineBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *QueryEngineBuilder) buildRequest() (rawTaskRequest, error) {
	if b.engineName == nil {
		return rawTaskRequest{}, errors.New("engineName is required")
	}

	cfg := map[string]any{
		"adp_queryEngine_engineName": *b.engineName,
	}
	if b.engineQuery != nil {
		cfg["adp_queryEngine_engineQuery"] = *b.engineQuery
	}
	if b.engineUserName != nil {
		cfg["adp_queryEngine_engineUserName"] = *b.engineUserName
	}
	if b.engineUserPassword != nil {
		cfg["adp_queryEngine_engineUserPassword"] = *b.engineUserPassword
	}
	if b.engineTaxonomies != nil {
		cfg["adp_queryEngine_engineTaxonomies"] = *b.engineTaxonomies
	}
	if b.applicationIdentifier != nil {
		cfg["adp_queryEngine_applicationIdentifier"] = *b.applicationIdentifier
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          queryEngineTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   queryEngineTaskDescription,
		TaskDisplayName:   queryEngineTaskDisplayName,
	}, nil
}

func decodeQueryEngine(meta any) (QueryEngineResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return QueryEngineResult{}, err
	}

	documentsCount, err := intStringField(obj, "adp_query_engine_documents_count")
	if err != nil {
		return QueryEngineResult{}, err
	}

	aggregatedValue, err := stringField(obj, "adp_query_engine_aggregated_value")
	if err != nil {
		return QueryEngineResult{}, err
	}

	return QueryEngineResult{
		DocumentsCount:  documentsCount,
		AggregatedValue: aggregatedValue,
	}, nil
}

func (b *QueryEngineBuilder) Execute(ctx context.Context) (QueryEngineResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return QueryEngineResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return QueryEngineResult{}, err
	}

	return decodeQueryEngine(resp.ExecutionMetaData)
}

func (b *QueryEngineBuilder) Start(ctx context.Context) (*TaskResponse, error) {
	req, err := b.buildRequest()
	if err != nil {
		return nil, err
	}

	return b.client.execute(ctx, "/executeAdpTaskAsync", req)
}

func (b *QueryEngineBuilder) Wait(ctx context.Context, interval time.Duration) (QueryEngineResult, error) {
	resp, err := b.Start(ctx)
	if err != nil {
		return QueryEngineResult{}, err
	}

	resp, err = b.client.Wait(ctx, resp.ExecutionID, interval)
	if err != nil {
		return QueryEngineResult{}, err
	}

	return decodeQueryEngine(resp.ExecutionMetaData)
}
