package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder provides operations to manage the taskDefinition property of the microsoft.graph.identityGovernance.taskReport entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
// returns a TaskDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskDefinitionable), nil
}
// ToGetRequestInformation the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
