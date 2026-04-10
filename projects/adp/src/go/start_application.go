package adp

import "context"

const (
	startApplicationTaskType        = "Start Application"
	startApplicationTaskDescription = "Starts an application"
	startApplicationTaskDisplayName = "Start application"
)

type StartApplicationBuilder struct {
	*builderCommon[StartApplicationBuilder]
	client                *Client
	applicationIdentifier *string
	useHTTPS              *bool
}

func NewStartApplicationBuilder(client *Client) *StartApplicationBuilder {
	b := &StartApplicationBuilder{client: client}
	common := newBuilderCommon(b)
	b.builderCommon = &common
	return b
}

func (b *StartApplicationBuilder) ApplicationIdentifier(value string) *StartApplicationBuilder {
	b.applicationIdentifier = &value
	return b
}

func (b *StartApplicationBuilder) UseHTTPS(value bool) *StartApplicationBuilder {
	b.useHTTPS = &value
	return b
}

func (b *StartApplicationBuilder) buildRequest() (rawTaskRequest, error) {
	cfg := map[string]any{}
	if b.applicationIdentifier != nil {
		cfg["adp_startApplication_applicationIdentifier"] = *b.applicationIdentifier
	}
	if b.useHTTPS != nil {
		cfg["adp_startApplication_useHttps"] = *b.useHTTPS
	}
	b.apply(cfg)

	return rawTaskRequest{
		TaskType:          startApplicationTaskType,
		TaskConfiguration: cfg,
		TaskDescription:   startApplicationTaskDescription,
		TaskDisplayName:   startApplicationTaskDisplayName,
	}, nil
}

func decodeStartApplication(meta any) (StartApplicationResult, error) {
	obj, err := metaObject(meta)
	if err != nil {
		return StartApplicationResult{}, err
	}

	applicationURL, err := stringField(obj, "adp_started_application_url")
	if err != nil {
		return StartApplicationResult{}, err
	}

	return StartApplicationResult{ApplicationURL: applicationURL}, nil
}

func (b *StartApplicationBuilder) Execute(ctx context.Context) (StartApplicationResult, error) {
	req, err := b.buildRequest()
	if err != nil {
		return StartApplicationResult{}, err
	}

	resp, err := b.client.execute(ctx, "/executeAdpTask", req)
	if err != nil {
		return StartApplicationResult{}, err
	}

	return decodeStartApplication(resp.ExecutionMetaData)
}
