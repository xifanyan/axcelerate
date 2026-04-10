package adp

import "fmt"

type Entity struct {
	ID                          string `json:"id"`
	DisplayName                 string `json:"displayName,omitempty"`
	ProcessStatus               string `json:"processStatus,omitempty"`
	HostID                      string `json:"hostId,omitempty"`
	HostName                    string `json:"hostName,omitempty"`
	SourceForCreateFromExisting bool   `json:"sourceForCreateFromExisting,omitempty"`
}

type ListEntitiesResult struct {
	OutputFile string   `json:"outputFile"`
	Entities   []Entity `json:"entities"`
}

type QueryEngineResult struct {
	DocumentsCount  int    `json:"documentsCount"`
	AggregatedValue string `json:"aggregatedValue"`
}

type SearchParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Category struct {
	ID          string              `json:"id"`
	DisplayName string              `json:"displayName"`
	Count       int                 `json:"count"`
	Properties  map[string][]string `json:"properties,omitempty"`
}

type TaxonomyEntry struct {
	ID       string     `json:"id"`
	Category []Category `json:"category"`
}

type Statistics struct {
	Taxonomy []TaxonomyEntry `json:"taxonomy"`
}

type StatisticsDocument struct {
	Date            string            `json:"date"`
	SearchParameter []SearchParameter `json:"searchParameter"`
	Statistics      Statistics        `json:"statistics"`
}

type TaxonomyStatisticResult struct {
	OutputFile string             `json:"outputFile"`
	Statistics StatisticsDocument `json:"statistics"`
}

type StartApplicationResult struct {
	ApplicationURL string `json:"applicationUrl"`
}

type CSVMergeResult struct{}

type ExportDocumentsResult struct {
	ExportFileName   string `json:"exportFileName"`
	ExportPath       string `json:"exportPath"`
	SearchResultSize int    `json:"searchResultSize"`
}

type Cell struct {
	Value any    `json:"Value"`
	Name  string `json:"Name"`
}

type ConfigurationParameter struct {
	Cells [][]Cell `json:"Cells"`
	Name  string   `json:"Name"`
	Value any      `json:"Value"`
}

type ConfigurationStatic struct {
	Parameters []ConfigurationParameter `json:"Parameters"`
}

type ConfigurationGlobal struct {
	Static ConfigurationStatic `json:"Static"`
}

type ConfigurationInfo struct {
	DynamicComponents map[string]any      `json:"DynamicComponents"`
	Global            ConfigurationGlobal `json:"Global"`
}

type ReadConfigurationResult struct {
	OutputFile    string                       `json:"outputFile"`
	Configuration map[string]ConfigurationInfo `json:"configuration"`
}

type CreateOcrJobResult struct{}

type CLIResult struct {
	Result     int            `json:"result"`
	JSONOutput map[string]any `json:"jsonOutput,omitempty"`
	ErrorPath  string         `json:"errorPath"`
	ResultPath string         `json:"resultPath"`
}

type EngineTaxonomyArg struct {
	Taxonomy string `json:"taxonomy"`
	Negation bool   `json:"negation"`
	Query    string `json:"query"`
}

type OutputTaxonomiesArg struct {
	Taxonomy                  string `json:"taxonomy"`
	Mode                      string `json:"mode"`
	MaximumNumberOfCategories int    `json:"maximumNumberOfCategories"`
}

type ConfigArg struct {
	ConfigurationID       string `json:"configurationId"`
	DynamicComponentNames string `json:"dynamicComponentNames"`
	FieldList             string `json:"fieldList"`
	NameValueList         string `json:"nameValueList"`
	ApplicationType       string `json:"applicationType"`
	EntityType            string `json:"entityType"`
}

type CLIBatchParameter struct {
	Parameter string `json:"parameter"`
}

type builderCommon[B any] struct {
	owner               *B
	taskActive          *bool
	taskTimeout         *int
	executionPersistent *bool
	abortWfOnFailure    *bool
	loggingEnabled      *bool
	cleanUpHistory      *bool
}

func newBuilderCommon[B any](owner *B) builderCommon[B] {
	return builderCommon[B]{owner: owner}
}

func (b *builderCommon[B]) TaskActive(value bool) *B {
	b.taskActive = &value
	return b.owner
}

func (b *builderCommon[B]) TaskTimeout(value int) *B {
	b.taskTimeout = &value
	return b.owner
}

func (b *builderCommon[B]) ExecutionPersistent(value bool) *B {
	b.executionPersistent = &value
	return b.owner
}

func (b *builderCommon[B]) AbortWfOnFailure(value bool) *B {
	b.abortWfOnFailure = &value
	return b.owner
}

func (b *builderCommon[B]) LoggingEnabled(value bool) *B {
	b.loggingEnabled = &value
	return b.owner
}

func (b *builderCommon[B]) CleanUpHistory(value bool) *B {
	b.cleanUpHistory = &value
	return b.owner
}

func (b builderCommon[B]) apply(dst map[string]any) {
	if b.taskActive != nil {
		dst["adp_taskActive"] = *b.taskActive
	}
	if b.taskTimeout != nil {
		dst["adp_taskTimeout"] = *b.taskTimeout
	}
	if b.executionPersistent != nil {
		dst["adp_executionPersistent"] = *b.executionPersistent
	}
	if b.abortWfOnFailure != nil {
		dst["adp_abortWfOnFailure"] = *b.abortWfOnFailure
	}
	if b.loggingEnabled != nil {
		dst["adp_loggingEnabled"] = *b.loggingEnabled
	}
	if b.cleanUpHistory != nil {
		dst["adp_cleanUpHistory"] = *b.cleanUpHistory
	}
}

type TaskResponse struct {
	ExecutionID         string  `json:"executionId"`
	TaskType            string  `json:"taskType"`
	LoggingEnabled      string  `json:"loggingEnabled"`
	ProgressMax         int     `json:"progressMax"`
	ExecutionStatus     string  `json:"executionStatus"`
	ExecutionRootDir    string  `json:"executionRootDir"`
	ContextID           string  `json:"contextId"`
	ExecutionPersistent string  `json:"executionPersistent"`
	ProgressCurrent     int     `json:"progressCurrent"`
	ProgressPercentage  float64 `json:"progressPercentage"`
	TaskDisplayName     string  `json:"taskDisplayName"`
	ExecutionMetaData   any     `json:"executionMetaData"`
	ErrorMessage        string  `json:"errorMessage,omitempty"`
}

type TaskExecutionError struct {
	ExecutionID       string
	TaskType          string
	ExecutionStatus   string
	ErrorMessage      string
	ExecutionMetaData any
}

func (e *TaskExecutionError) Error() string {
	return fmt.Sprintf("task %s failed: %s (executionId=%s)", e.TaskType, e.ErrorMessage, e.ExecutionID)
}
