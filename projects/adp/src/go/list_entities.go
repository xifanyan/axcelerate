package adp

import (
	"context"
	"time"
)

const (
	listEntitiesTaskType        = "List Entities"
	listEntitiesTaskDescription = "Writes a list of entities ot an output variable"
	listEntitiesTaskDisplayName = "List entities"
)

type ListEntitiesBuilder struct {
	*builderCommon[ListEntitiesBuilder]
	client        *Client
	typeValue     *string
	idValue       *string
	relatedEntity *string
	whiteList     *string
	workspace     *string
	status        *string
}

func NewListEntitiesBuilder(client *Client) *ListEntitiesBuilder {
	b := &ListEntitiesBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *ListEntitiesBuilder) Type(value string) *ListEntitiesBuilder {
	b.typeValue = &value
	return b
}

func (b *ListEntitiesBuilder) ID(value string) *ListEntitiesBuilder {
	b.idValue = &value
	return b
}

func (b *ListEntitiesBuilder) RelatedEntity(value string) *ListEntitiesBuilder {
	b.relatedEntity = &value
	return b
}

func (b *ListEntitiesBuilder) WhiteList(value string) *ListEntitiesBuilder {
	b.whiteList = &value
	return b
}

func (b *ListEntitiesBuilder) Workspace(value string) *ListEntitiesBuilder {
	b.workspace = &value
	return b
}

func (b *ListEntitiesBuilder) Status(value string) *ListEntitiesBuilder {
	b.status = &value
	return b
}

func (b *ListEntitiesBuilder) buildRequest() (rawTaskRequest, error) {
	cfg := map[string]any{}
	if b.typeValue != nil {
		cfg["adp_listEntities_type"] = *b.typeValue
	}
	if b.idValue != nil {
		cfg["adp_listEntities_id"] = *b.idValue
	}
	if b.relatedEntity != nil {
		cfg["adp_listEntities_relatedEntity"] = *b.relatedEntity
	}
	if b.whiteList != nil {
		cfg["adp_listEntities_whiteList"] = *b.whiteList
	}
	if b.workspace != nil {
		cfg["adp_listEntities_workspace"] = *b.workspace
	}
	if b.status != nil {
		cfg["adp_listEntities_status"] = *b.status
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          listEntitiesTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   listEntitiesTaskDescription,
		TaskDisplayName:   listEntitiesTaskDisplayName,
	}, nil
}

func decodeListEntities(meta any) (ListEntitiesResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return ListEntitiesResult{}, err
	}

	outputFile, err := stringField(obj, "adp_entities_output_file_name")
	if err != nil {
		return ListEntitiesResult{}, err
	}

	var entities []Entity
	if err := jsonStringField(obj, "adp_entities_json_output", &entities); err != nil {
		return ListEntitiesResult{}, err
	}

	return ListEntitiesResult{
		OutputFile: outputFile,
		Entities:   entities,
	}, nil
}

func (b *ListEntitiesBuilder) Execute(ctx context.Context) (ListEntitiesResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return ListEntitiesResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return ListEntitiesResult{}, err
	}

	return decodeListEntities(resp.ExecutionMetaData)
}

func (b *ListEntitiesBuilder) Start(ctx context.Context) (*TaskResponse, error) {
	req, err := b.buildRequest()
	if err != nil {
		return nil, err
	}

	return b.client.execute(ctx, "/executeAdpTaskAsync", req)
}

func (b *ListEntitiesBuilder) Wait(ctx context.Context, interval time.Duration) (ListEntitiesResult, error) {
	resp, err := b.Start(ctx)
	if err != nil {
		return ListEntitiesResult{}, err
	}

	resp, err = b.client.Wait(ctx, resp.ExecutionID, interval)
	if err != nil {
		return ListEntitiesResult{}, err
	}

	return decodeListEntities(resp.ExecutionMetaData)
}
