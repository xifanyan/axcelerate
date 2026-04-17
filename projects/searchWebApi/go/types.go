package searchwebapi

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StatusObject struct {
	Successful    bool   `json:"successful"`
	BackendStatus string `json:"backendStatus"`
	HTTPStatus    int    `json:"httpStatus"`
	ErrorMessage  string `json:"errorMessage"`
}

type LoginResult struct {
	Status StatusObject `json:"status"`
}

type LogoutResult struct {
	Status StatusObject `json:"status"`
}

type Project struct {
	ID string `json:"id"`
}

type ProjectResource struct {
	ID                       string           `json:"id"`
	Description              string           `json:"description"`
	AvailableQueryParameters []QueryParameter `json:"availableQueryParameters,omitempty"`
}

type ProjectsResult struct {
	Status        StatusObject `json:"status"`
	NumberResults int64        `json:"numberResults"`
	Results       []Project    `json:"results"`
}

type ProjectResourcesResult struct {
	Status        StatusObject      `json:"status"`
	NumberResults int64             `json:"numberResults"`
	Results       []ProjectResource `json:"results"`
}

type Collection struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName,omitempty"`
}

type QueryParameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CollectionResource struct {
	ID                       string           `json:"id"`
	Description              string           `json:"description"`
	AvailableQueryParameters []QueryParameter `json:"availableQueryParameters"`
}

type CollectionsResult struct {
	Status        StatusObject `json:"status"`
	NumberResults int64        `json:"numberResults"`
	Results       []Collection `json:"results"`
}

type CollectionResourcesResult struct {
	Status        StatusObject         `json:"status"`
	NumberResults int64                `json:"numberResults"`
	Results       []CollectionResource `json:"results"`
}

type FieldDescription struct {
	ID                           string `json:"id"`
	Type                         string `json:"type"`
	IsSortable                   bool   `json:"isSortable"`
	IsFolderField                bool   `json:"isFolderField"`
	IsMultivalueFolderCollection bool   `json:"isMultivalueFolderCollection"`
	IsPrefixSearchable           bool   `json:"isPrefixSearchable"`
	DisplayName                  string `json:"displayName,omitempty"`
}

type FieldsResult struct {
	Status        StatusObject       `json:"status"`
	NumberResults int64              `json:"numberResults"`
	Results       []FieldDescription `json:"results"`
}

type FolderFieldResource struct {
	ID                       string           `json:"id"`
	Description              string           `json:"description"`
	AvailableQueryParameters []QueryParameter `json:"availableQueryParameters"`
}

type FolderFieldResourcesResult struct {
	Status        StatusObject          `json:"status"`
	NumberResults int64                 `json:"numberResults"`
	Results       []FolderFieldResource `json:"results"`
}

type FolderRecord struct {
	Rank        int64   `json:"rank"`
	Relevance   float64 `json:"relevance"`
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Count       int64   `json:"count"`
}

type FolderValuesResult struct {
	Status        StatusObject   `json:"status"`
	NumberResults int64          `json:"numberResults"`
	Results       []FolderRecord `json:"results"`
}

type Field struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	ValueObject any    `json:"valueObject,omitempty"`
}

type Folder struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Properties  []Field `json:"properties,omitempty"`
}

type FolderSet struct {
	ID    string   `json:"id"`
	Value []Folder `json:"value"`
}

type SponsoredLink struct {
	ExternalLink string  `json:"externalLink,omitempty"`
	RecordID     string  `json:"recordId,omitempty"`
	Title        string  `json:"title,omitempty"`
	Description  string  `json:"description,omitempty"`
	Relevance    float64 `json:"relevance,omitempty"`
}

type SpellingSuggestion struct {
	OriginalWord       string  `json:"originalWord"`
	SuggestedWord      string  `json:"suggestedWord"`
	ImprovementPercent float64 `json:"improvementPercent"`
	IsOriginalFound    bool    `json:"isOriginalFound"`
}

type SpellingSuggestionResult struct {
	TotalImprovementPercent float64              `json:"totalImprovementPercent"`
	SuggestedWords          []SpellingSuggestion `json:"suggestedWords"`
}

type Record struct {
	Rank        int64       `json:"rank"`
	Relevance   float64     `json:"relevance"`
	ID          string      `json:"id"`
	UniqueField string      `json:"uniqueField"`
	Fields      []Field     `json:"fields"`
	FolderSets  []FolderSet `json:"folderSets"`
	Body        string      `json:"body"`
	Page        int         `json:"page"`
	PageCount   int         `json:"pageCount"`
}

type SearchResult struct {
	Status              StatusObject              `json:"status"`
	NumberResults       int64                     `json:"numberResults"`
	Results             []Record                  `json:"results"`
	SponsoredLinks      []SponsoredLink           `json:"sponsoredLinks,omitempty"`
	SpellingSuggestions *SpellingSuggestionResult `json:"spellingSuggestions,omitempty"`
}

type SearchResultHighlight struct {
	Field             string `json:"field"`
	RegularExpression string `json:"regularExpression"`
}

type SearchResultHighlightingResult struct {
	Results []SearchResultHighlight `json:"results"`
}

type HighlightResultEntity struct {
	TermsToHighlight                    []string  `json:"termsToHighlight,omitempty"`
	RegularExpression                   string    `json:"regularExpression,omitempty"`
	NumberHits                          int       `json:"numberHits,omitempty"`
	NumberHitsByPage                    []int     `json:"numberHitsByPage,omitempty"`
	HitLocationsRel                     []float64 `json:"hitLocationsRel,omitempty"`
	TermToHighlightIsWordBoundaryBefore []bool    `json:"termToHighlightIsWordBoundaryBefore,omitempty"`
	TermToHighlightIsWordBoundaryAfter  []bool    `json:"termToHighlightIsWordBoundaryAfter,omitempty"`
}

type StoredSearchHighlightResult struct {
	Field  string                `json:"field,omitempty"`
	Folder string                `json:"folder,omitempty"`
	Terms  HighlightResultEntity `json:"terms"`
}

type HighlightedWordResult struct {
	SearchTerms                        *HighlightResultEntity        `json:"searchTerms,omitempty"`
	ConceptTerms                       *HighlightResultEntity        `json:"conceptTerms,omitempty"`
	TrainingTerms                      *HighlightResultEntity        `json:"trainingTerms,omitempty"`
	UserTerms                          *HighlightResultEntity        `json:"userTerms,omitempty"`
	StoredSearchHighlightingByCategory []StoredSearchHighlightResult `json:"storedSearchHighlightingByCategory,omitempty"`
}

type SearchResultToken struct {
	Token string `json:"token"`
}

type SearchResultTokenResponse struct {
	Status        StatusObject `json:"status"`
	Token         string       `json:"token"`
	NumberResults int64        `json:"numberResults"`
	EOL           string       `json:"eol,omitempty"`
}

type RecordResource struct {
	ID                       string           `json:"id"`
	Description              string           `json:"description"`
	AvailableQueryParameters []QueryParameter `json:"availableQueryParameters"`
}

type RecordResourcesResult struct {
	Status        StatusObject     `json:"status"`
	NumberResults int64            `json:"numberResults"`
	Results       []RecordResource `json:"results"`
}

type ChangeRequest struct {
	Field     string   `json:"field"`
	Type      string   `json:"type"`
	FolderIDs []string `json:"folderIds,omitempty"`
	Text      string   `json:"text,omitempty"`
	Texts     []string `json:"texts,omitempty"`
}

type ChangeResult struct {
	Status StatusObject `json:"status"`
}

type BinaryResponse struct {
	Header      http.Header
	ContentType string
	Body        io.ReadCloser
}

type MeasureDimensionMember struct {
	Identifier string  `json:"identifier"`
	Fields     []Field `json:"fields,omitempty"`
}

type MeasureDimensionResult struct {
	Size                  int64                    `json:"size"`
	FieldName             string                   `json:"fieldName,omitempty"`
	DocumentsWithAnyValue int64                    `json:"documentsWithAnyValue,omitempty"`
	DocumentsWithNoValue  int64                    `json:"documentsWithNoValue,omitempty"`
	Members               []MeasureDimensionMember `json:"members,omitempty"`
}

type MeasureCube struct {
	Dimensions   []MeasureDimensionResult `json:"dimensions,omitempty"`
	Values       [][]int64                `json:"values,omitempty"`
	ValuesDouble [][]float64              `json:"valuesDouble,omitempty"`
}

type CachedSearchDescription struct {
	CreationTraceID string `json:"creationTraceId"`
}

type WaitForPendingChangesResult struct {
	Success bool `json:"success"`
}

type SearchRequest struct {
	Query           string `json:"query,omitempty"`
	Language        string `json:"language,omitempty"`
	JoinRestriction string `json:"joinRestriction,omitempty"`
}

type FieldData struct {
	FieldName string   `json:"fieldName,omitempty"`
	Value     string   `json:"value,omitempty"`
	ValueList []string `json:"valueList,omitempty"`
}

type RecordData struct {
	UniqueID  string      `json:"uniqueId,omitempty"`
	FieldData []FieldData `json:"fieldData,omitempty"`
}

type InsertRemoveRequest struct {
	NewRecords             []RecordData   `json:"newRecords,omitempty"`
	RecordsToDelete        []string       `json:"recordsToDelete,omitempty"`
	RecordsToDeleteByQuery *SearchRequest `json:"recordsToDeleteByQuery,omitempty"`
	DeletionMode           string         `json:"deletionMode,omitempty"`
}

type InsertRemoveResult struct {
	Status StatusObject `json:"status"`
}

type StartTransactionRequest struct {
	DataSourceID     string `json:"dataSourceId,omitempty"`
	IndexingBufferID string `json:"indexingBufferId,omitempty"`
}

type StartTransactionResult struct {
	IndexingBufferID string `json:"indexingBufferId"`
}

type FinishTransactionRequest struct {
	NumberDocuments int64 `json:"numberDocuments,omitempty"`
	AnalyzedSizeGb  int64 `json:"analyzedSizeGb,omitempty"`
}

type FinishTransactionResponse struct {
	JobID string `json:"jobId"`
}

type JobStatusResponse struct {
	JobStatus string `json:"jobStatus"`
}

type DimensionRequest struct {
	Field                  string   `json:"field,omitempty"`
	RestrictFoldersByQuery string   `json:"restrictFoldersByQuery,omitempty"`
	RestrictValuesByList   []string `json:"restrictValuesByList,omitempty"`
	ReturnEmptyMembers     bool     `json:"returnEmptyMembers,omitempty"`
	FolderOrder            string   `json:"folderOrder,omitempty"`
	ReturnedFields         string   `json:"returnedFields,omitempty"`
	Page                   int      `json:"page,omitempty"`
	PageSize               int      `json:"pageSize,omitempty"`
	Offset                 int      `json:"offset,omitempty"`
}

type RecordStream struct {
	Meta   SearchResult
	reader io.Reader
	body   io.Closer
}

func (s *RecordStream) Next() (*Record, error) {
	if s.reader == nil {
		return nil, io.EOF
	}
	line, err := bufioReadLine(s.reader)
	if err != nil {
		return nil, err
	}
	var record Record
	if err := jsonUnmarshal(line, &record); err != nil {
		return nil, err
	}
	return &record, nil
}

func (s *RecordStream) Close() error {
	if s.body == nil {
		return nil
	}
	return s.body.Close()
}

type SearchRecordsOptions struct {
	Query                      string
	Language                   string
	JoinRestriction            string
	Order                      string
	Fields                     string
	FolderFields               string
	FolderFieldsWithProperties string
	Body                       *bool
	Highlight                  *bool
	Page                       *int
	Limit                      *int
	SponsoredLinks             *bool
	SpellingSuggestions        *bool
	SearchCacheControl         string
}

func (o SearchRecordsOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	addQuery(values, "order", o.Order)
	addQuery(values, "fields", o.Fields)
	addQuery(values, "folderFields", o.FolderFields)
	addQuery(values, "folderFieldsWithProperties", o.FolderFieldsWithProperties)
	if v := boolPtrString(o.Body); v != "" {
		values.Set("body", v)
	}
	if v := boolPtrString(o.Highlight); v != "" {
		values.Set("highlight", v)
	}
	if v := intPtrString(o.Page); v != "" {
		values.Set("page", v)
	}
	if v := intPtrString(o.Limit); v != "" {
		values.Set("limit", v)
	}
	if v := boolPtrString(o.SponsoredLinks); v != "" {
		values.Set("sponsoredLinks", v)
	}
	if v := boolPtrString(o.SpellingSuggestions); v != "" {
		values.Set("spellingSuggestions", v)
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type FetchRecordContentOptions struct {
	Fields                             string
	FolderFields                       string
	FolderFieldsWithProperties         string
	Body                               *bool
	Page                               *int
	HighlightSearchTermQuery           string
	HighlightSearchTermLanguage        string
	HighlightSearchTermJoinRestriction string
	FieldsHighlighted                  string
	HighlightHitNavigation             string
	HighlightUserTerms                 string
	HighlightFolderFieldList           string
	Summarize                          *bool
	SearchCacheControl                 string
}

func (o FetchRecordContentOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "fields", o.Fields)
	addQuery(values, "folderFields", o.FolderFields)
	addQuery(values, "folderFieldsWithProperties", o.FolderFieldsWithProperties)
	if v := boolPtrString(o.Body); v != "" {
		values.Set("body", v)
	}
	if v := intPtrString(o.Page); v != "" {
		values.Set("page", v)
	}
	addQuery(values, "highlightSearchTermQuery", o.HighlightSearchTermQuery)
	addQuery(values, "highlightSearchTermLanguage", o.HighlightSearchTermLanguage)
	addQuery(values, "highlightSearchTermJoinRestriction", o.HighlightSearchTermJoinRestriction)
	addQuery(values, "fieldsHighlighted", o.FieldsHighlighted)
	addQuery(values, "highlightHitNavigation", o.HighlightHitNavigation)
	addQuery(values, "highlightUserTerms", o.HighlightUserTerms)
	addQuery(values, "highlightFolderFieldList", o.HighlightFolderFieldList)
	if v := boolPtrString(o.Summarize); v != "" {
		values.Set("summarize", v)
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type FolderValuesOptions struct {
	Query                  string
	Language               string
	JoinRestriction        string
	Prefix                 string
	RestrictFoldersByQuery string
	ReturnEmptyFolders     *bool
	Order                  string
	Offset                 *int
	Limit                  *int
	SearchCacheControl     string
}

func (o FolderValuesOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	addQuery(values, "prefix", o.Prefix)
	addQuery(values, "restrictFoldersByQuery", o.RestrictFoldersByQuery)
	if v := boolPtrString(o.ReturnEmptyFolders); v != "" {
		values.Set("returnEmptyFolders", v)
	}
	addQuery(values, "order", o.Order)
	if v := intPtrString(o.Offset); v != "" {
		values.Set("offset", v)
	}
	if v := intPtrString(o.Limit); v != "" {
		values.Set("limit", v)
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type BinarySearchOptions struct {
	Query              string
	Language           string
	JoinRestriction    string
	Order              string
	Field              string
	SelectedIndex      *int
	SearchCacheControl string
}

func (o BinarySearchOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	addQuery(values, "order", o.Order)
	addQuery(values, "field", o.Field)
	if v := intPtrString(o.SelectedIndex); v != "" {
		values.Set("selectedIndex", v)
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type CachedSearchesDeleteOptions struct {
	CreationTraceIDs string
}

func (o CachedSearchesDeleteOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "creationTraceIds", o.CreationTraceIDs)
	return values
}

type WaitForPendingChangesOptions struct {
	TimeoutMillis           *int64
	OnlyHighPriorityChanges *bool
}

func (o WaitForPendingChangesOptions) values() url.Values {
	values := url.Values{}
	if v := int64PtrString(o.TimeoutMillis); v != "" {
		values.Set("timeoutMillis", v)
	}
	if v := boolPtrString(o.OnlyHighPriorityChanges); v != "" {
		values.Set("onlyHighPriorityChanges", v)
	}
	return values
}

type InDocumentSearchOptions struct {
	HighlightSearchTermQuery            string
	HighlightSearchTermLanguage         string
	HighlightSearchTermJoinRestriction  string
	HighlightUserTerms                  string
	HighlightFolderFieldList            string
	HighlightFolderFieldsAggregation    string
	ContentFieldNames                   string
	PageTag                             string
	OmitHitsPerPage                     *bool
	RequestHitLocationsPageRelative     *bool
	RequestHitLocationsDocumentRelative *bool
	SearchCacheControl                  string
}

func (o InDocumentSearchOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "highlightSearchTermQuery", o.HighlightSearchTermQuery)
	addQuery(values, "highlightSearchTermLanguage", o.HighlightSearchTermLanguage)
	addQuery(values, "highlightSearchTermJoinRestriction", o.HighlightSearchTermJoinRestriction)
	addQuery(values, "highlightUserTerms", o.HighlightUserTerms)
	addQuery(values, "highlightFolderFieldList", o.HighlightFolderFieldList)
	addQuery(values, "highlightFolderFieldsAggregation", o.HighlightFolderFieldsAggregation)
	addQuery(values, "contentFieldNames", o.ContentFieldNames)
	addQuery(values, "pageTag", o.PageTag)
	if v := boolPtrString(o.OmitHitsPerPage); v != "" {
		values.Set("omitHitsPerPage", v)
	}
	if v := boolPtrString(o.RequestHitLocationsPageRelative); v != "" {
		values.Set("requestHitLocationsPageRelative", v)
	}
	if v := boolPtrString(o.RequestHitLocationsDocumentRelative); v != "" {
		values.Set("requestHitLocationsDocumentRelative", v)
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type GetHighlightExpressionsOptions struct {
	Query              string
	Language           string
	JoinRestriction    string
	SearchCacheControl string
}

func (o GetHighlightExpressionsOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

type SearchTokenOptions struct {
	Query           string
	Language        string
	JoinRestriction string
	Order           string
}

func (o SearchTokenOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	addQuery(values, "order", o.Order)
	return values
}

type SortOrderSnapshotOptions struct {
	TopN *int64
}

func (o SortOrderSnapshotOptions) values() url.Values {
	values := url.Values{}
	if v := int64PtrString(o.TopN); v != "" {
		values.Set("topN", v)
	}
	return values
}

type ChangeOptions struct {
	BlockUntilComplete *bool
}

func (o ChangeOptions) values() url.Values {
	values := url.Values{}
	if v := boolPtrString(o.BlockUntilComplete); v != "" {
		values.Set("blockUntilComplete", v)
	}
	return values
}

type MeasureType struct {
	TypeName            string   `json:"typeName,omitempty"`
	FieldName           string   `json:"fieldName,omitempty"`
	Universe            string   `json:"universe,omitempty"`
	SelectedFieldValues []string `json:"selectedFieldValues,omitempty"`
}

type MeasureOptions struct {
	Query              string
	Language           string
	JoinRestriction    string
	MeasureType        *MeasureType
	SearchCacheControl string
}

func (o MeasureOptions) values() url.Values {
	values := url.Values{}
	addQuery(values, "query", o.Query)
	addQuery(values, "language", o.Language)
	addQuery(values, "joinRestriction", o.JoinRestriction)
	if o.MeasureType != nil {
		values.Set("measureType", marshalCompact(o.MeasureType))
	}
	addQuery(values, "SWA-searchCacheControl", o.SearchCacheControl)
	return values
}

func marshalCompact(v any) string {
	data, err := jsonMarshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func encodeJSONQuery(v any) string {
	data, err := jsonMarshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

func bufioReadLine(r io.Reader) ([]byte, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	line, err := br.ReadBytes('\n')
	if err == io.EOF && len(strings.TrimSpace(string(line))) > 0 {
		return []byte(strings.TrimSpace(string(line))), nil
	}
	if err != nil {
		return nil, err
	}
	return []byte(strings.TrimSpace(string(line))), nil
}

var jsonMarshal = json.Marshal
var jsonUnmarshal = json.Unmarshal
