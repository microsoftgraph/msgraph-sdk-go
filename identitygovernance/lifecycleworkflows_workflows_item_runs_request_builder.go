package identitygovernance

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemRunsRequestBuilder provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsWorkflowsItemRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetQueryParameters get a list of the run objects and their properties for a lifecycle workflow.
type LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetQueryParameters
}
// ByRunId provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemRunsRunItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) ByRunId(runId string)(*LifecycleworkflowsWorkflowsItemRunsRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if runId != "" {
        urlTplParams["run%2Did"] = runId
    }
    return NewLifecycleworkflowsWorkflowsItemRunsRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsWorkflowsItemRunsRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemRunsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemRunsRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemRunsRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemRunsRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsWorkflowsItemRunsCountRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) Count()(*LifecycleworkflowsWorkflowsItemRunsCountRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the run objects and their properties for a lifecycle workflow.
// returns a RunCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-list-runs?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.RunCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateRunCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.RunCollectionResponseable), nil
}
// MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTime provides operations to call the summary method.
// returns a *LifecycleworkflowsWorkflowsItemRunsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleworkflowsWorkflowsItemRunsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// ToGetRequestInformation get a list of the run objects and their properties for a lifecycle workflow.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemRunsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemRunsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
