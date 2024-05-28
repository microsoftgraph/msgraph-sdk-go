package identitygovernance

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetQueryParameters the associated individual user execution.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetQueryParameters
}
// ByUserProcessingResultId provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) ByUserProcessingResultId(userProcessingResultId string)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userProcessingResultId != "" {
        urlTplParams["userProcessingResult%2Did"] = userProcessingResultId
    }
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsCountRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) Count()(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsCountRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the associated individual user execution.
// returns a UserProcessingResultCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateUserProcessingResultCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultCollectionResponseable), nil
}
// MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTime provides operations to call the summary method.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) MicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsMicrosoftgraphidentitygovernancesummarywithstartdatetimewithenddatetimeMicrosoftGraphIdentityGovernanceSummaryWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, startDateTime)
}
// ToGetRequestInformation the associated individual user execution.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
