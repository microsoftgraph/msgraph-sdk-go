package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsLifecycleWorkflowsRequestBuilder provides operations to manage the lifecycleWorkflows property of the microsoft.graph.identityGovernance entity.
type LifecycleworkflowsLifecycleWorkflowsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetQueryParameters get lifecycleWorkflows from identityGovernance
type LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetQueryParameters
}
// LifecycleworkflowsLifecycleWorkflowsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsLifecycleWorkflowsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsLifecycleWorkflowsRequestBuilderInternal instantiates a new LifecycleworkflowsLifecycleWorkflowsRequestBuilder and sets the default values.
func NewLifecycleworkflowsLifecycleWorkflowsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsLifecycleWorkflowsRequestBuilder) {
    m := &LifecycleworkflowsLifecycleWorkflowsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsLifecycleWorkflowsRequestBuilder instantiates a new LifecycleworkflowsLifecycleWorkflowsRequestBuilder and sets the default values.
func NewLifecycleworkflowsLifecycleWorkflowsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsLifecycleWorkflowsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsLifecycleWorkflowsRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomTaskExtensions provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) CustomTaskExtensions()(*LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) {
    return NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property lifecycleWorkflows for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeletedItems provides operations to manage the deletedItems property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) DeletedItems()(*LifecycleworkflowsDeleteditemsDeletedItemsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsDeletedItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get lifecycleWorkflows from identityGovernance
// returns a LifecycleWorkflowsContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateLifecycleWorkflowsContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable), nil
}
// Patch update the navigation property lifecycleWorkflows in identityGovernance
// returns a LifecycleWorkflowsContainerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) Patch(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateLifecycleWorkflowsContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsSettingsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) Settings()(*LifecycleworkflowsSettingsRequestBuilder) {
    return NewLifecycleworkflowsSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskDefinitions provides operations to manage the taskDefinitions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsTaskdefinitionsTaskDefinitionsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) TaskDefinitions()(*LifecycleworkflowsTaskdefinitionsTaskDefinitionsRequestBuilder) {
    return NewLifecycleworkflowsTaskdefinitionsTaskDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property lifecycleWorkflows for identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get lifecycleWorkflows from identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property lifecycleWorkflows in identityGovernance
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.LifecycleWorkflowsContainerable, requestConfiguration *LifecycleworkflowsLifecycleWorkflowsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsLifecycleWorkflowsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsLifecycleWorkflowsRequestBuilder) {
    return NewLifecycleworkflowsLifecycleWorkflowsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// Workflows provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsWorkflowsRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) Workflows()(*LifecycleworkflowsWorkflowsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WorkflowTemplates provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder when successful
func (m *LifecycleworkflowsLifecycleWorkflowsRequestBuilder) WorkflowTemplates()(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
