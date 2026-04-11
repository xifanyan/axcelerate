package adp

import (
	"context"
	"strconv"
	"time"
)

const (
	createOcrJobTaskType        = "Create OCR Job"
	createOcrJobTaskDescription = "Changes metaData by using regEx replacement."
	createOcrJobTaskDisplayName = "Create OCR Job"
)

type CreateOcrJobBuilder struct {
	*builderCommon[CreateOcrJobBuilder]
	client                *Client
	engineName            *string
	query                 *string
	engineUserName        *string
	engineUserPassword    *string
	jobName               *string
	jobDescription        *string
	jobPriority           *int
	applicationIdentifier *string
	applicationType       *string
	wait                  *bool
	engineType            *string
	listOfJobProperties   *string
	globalSearchJSON      *string
	globalSearchID        *string
	restrictions          *[]EngineTaxonomyArg
	advancedRestrictions  *[]EngineTaxonomyArg
	mainQueryType         *string
}

func NewCreateOcrJobBuilder(client *Client) *CreateOcrJobBuilder {
	b := &CreateOcrJobBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *CreateOcrJobBuilder) EngineName(value string) *CreateOcrJobBuilder {
	b.engineName = &value
	return b
}

func (b *CreateOcrJobBuilder) Query(value string) *CreateOcrJobBuilder {
	b.query = &value
	return b
}

func (b *CreateOcrJobBuilder) EngineUserName(value string) *CreateOcrJobBuilder {
	b.engineUserName = &value
	return b
}

func (b *CreateOcrJobBuilder) EngineUserPassword(value string) *CreateOcrJobBuilder {
	b.engineUserPassword = &value
	return b
}

func (b *CreateOcrJobBuilder) JobName(value string) *CreateOcrJobBuilder {
	b.jobName = &value
	return b
}

func (b *CreateOcrJobBuilder) JobDescription(value string) *CreateOcrJobBuilder {
	b.jobDescription = &value
	return b
}

func (b *CreateOcrJobBuilder) JobPriority(value int) *CreateOcrJobBuilder {
	b.jobPriority = &value
	return b
}

func (b *CreateOcrJobBuilder) ApplicationIdentifier(value string) *CreateOcrJobBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *CreateOcrJobBuilder) ApplicationType(value string) *CreateOcrJobBuilder {
	b.applicationType = &value
	return b
}

func (b *CreateOcrJobBuilder) WaitFlag(value bool) *CreateOcrJobBuilder {
	b.wait = &value
	return b
}

func (b *CreateOcrJobBuilder) EngineType(value string) *CreateOcrJobBuilder {
	b.engineType = &value
	return b
}

func (b *CreateOcrJobBuilder) ListOfJobProperties(value string) *CreateOcrJobBuilder {
	b.listOfJobProperties = &value
	return b
}

func (b *CreateOcrJobBuilder) GlobalSearchJSON(value string) *CreateOcrJobBuilder {
	b.globalSearchJSON = &value
	return b
}

func (b *CreateOcrJobBuilder) GlobalSearchID(value string) *CreateOcrJobBuilder {
	b.globalSearchID = &value
	return b
}

func (b *CreateOcrJobBuilder) Restrictions(value []EngineTaxonomyArg) *CreateOcrJobBuilder {
	b.restrictions = &value
	return b
}

func (b *CreateOcrJobBuilder) AdvancedRestrictions(value []EngineTaxonomyArg) *CreateOcrJobBuilder {
	b.advancedRestrictions = &value
	return b
}

func (b *CreateOcrJobBuilder) MainQueryType(value string) *CreateOcrJobBuilder {
	b.mainQueryType = &value
	return b
}

func (b *CreateOcrJobBuilder) buildRequest() (rawTaskRequest, error) {
	cfg := map[string]any{}
	if b.engineName != nil {
		cfg["adp_createOcrJob_engineName"] = *b.engineName
	}
	if b.query != nil {
		cfg["adp_createOcrJob_query"] = *b.query
	}
	if b.engineUserName != nil {
		cfg["adp_createOcrJob_engineUserName"] = *b.engineUserName
	}
	if b.engineUserPassword != nil {
		cfg["adp_createOcrJob_engineUserPassword"] = *b.engineUserPassword
	}
	if b.jobName != nil {
		cfg["adp_createOcrJob_jobName"] = *b.jobName
	}
	if b.jobDescription != nil {
		cfg["adp_createOcrJob_jobDescription"] = *b.jobDescription
	}
	if b.jobPriority != nil {
		cfg["adp_createOcrJob_jobPriority"] = strconv.Itoa(*b.jobPriority)
	}
	if b.applicationIdentifier != nil {
		cfg["adp_createOcrJob_applicationIdentifier"] = *b.applicationIdentifier
	}
	if b.applicationType != nil {
		cfg["adp_createOcrJob_applicationType"] = *b.applicationType
	}
	if b.wait != nil {
		cfg["adp_createOcrJob_wait"] = strconv.FormatBool(*b.wait)
	}
	if b.engineType != nil {
		cfg["adp_createOcrJob_engineType"] = *b.engineType
	}
	if b.listOfJobProperties != nil {
		cfg["adp_createOcrJob_listOfJobProperties"] = *b.listOfJobProperties
	}
	if b.globalSearchJSON != nil {
		cfg["adp_createOcrJob_globalSearchJson"] = *b.globalSearchJSON
	}
	if b.globalSearchID != nil {
		cfg["adp_createOcrJob_globalSearchId"] = *b.globalSearchID
	}
	if b.restrictions != nil && len(*b.restrictions) > 0 {
		cfg["adp_createOcrJob_restrictions"] = *b.restrictions
	}
	if b.advancedRestrictions != nil && len(*b.advancedRestrictions) > 0 {
		cfg["adp_createOcrJob_AdvancedRestrictions"] = *b.advancedRestrictions
	}
	if b.mainQueryType != nil {
		cfg["adp_createOcrJob_mainQueryType"] = *b.mainQueryType
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          createOcrJobTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   createOcrJobTaskDescription,
		TaskDisplayName:   createOcrJobTaskDisplayName,
	}, nil
}

func decodeCreateOcrJob(meta any) (CreateOcrJobResult, error) {
	_, err := metaObject(meta)
	if err != nil {
		return CreateOcrJobResult{}, err
	}

	return CreateOcrJobResult{}, nil
}

func (b *CreateOcrJobBuilder) Start(ctx context.Context) (*TaskResponse, error) {
	req, err := b.buildRequest()
	if err != nil {
		return nil, err
	}

	return b.client.execute(ctx, "/executeAdpTaskAsync", req)
}

func (b *CreateOcrJobBuilder) Wait(ctx context.Context, interval time.Duration) (CreateOcrJobResult, error) {
	resp, err := b.Start(ctx)
	if err != nil {
		return CreateOcrJobResult{}, err
	}

	resp, err = b.client.Wait(ctx, resp.ExecutionID, interval)
	if err != nil {
		return CreateOcrJobResult{}, err
	}

	return decodeCreateOcrJob(resp.ExecutionMetaData)
}
