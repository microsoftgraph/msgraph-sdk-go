package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
type ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters read the properties and relationships of a conditionalAccessTemplate object.
type ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetQueryParameters
}
// NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal instantiates a new ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder and sets the default values.
func NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    m := &ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/templates/{conditionalAccessTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder instantiates a new ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder and sets the default values.
func NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read the properties and relationships of a conditionalAccessTemplate object.
// returns a ConditionalAccessTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccesstemplate-get?view=graph-rest-1.0
func (m *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessTemplateable), nil
}
// ToGetRequestInformation read the properties and relationships of a conditionalAccessTemplate object.
// returns a *RequestInformation when successful
func (m *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder when successful
func (m *ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    return NewConditionalaccessTemplatesConditionalAccessTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
