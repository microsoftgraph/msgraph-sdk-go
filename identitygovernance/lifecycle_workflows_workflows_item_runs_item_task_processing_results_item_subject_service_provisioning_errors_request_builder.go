package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \identityGovernance\lifecycleWorkflows\workflows\{workflow-id}\runs\{run-id}\taskProcessingResults\{taskProcessingResult-id}\subject\serviceProvisioningErrors
type LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetQueryParameters get serviceProvisioningErrors property value
type LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
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
// LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderInternal instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs/{run%2Did}/taskProcessingResults/{taskProcessingResult%2Did}/subject/serviceProvisioningErrors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) Count()(*LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsCountRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get serviceProvisioningErrors property value
func (m *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceProvisioningErrorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable), nil
}
// ToGetRequestInformation get serviceProvisioningErrors property value
func (m *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemRunsItemTaskProcessingResultsItemSubjectServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
