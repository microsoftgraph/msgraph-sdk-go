package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \identityGovernance\lifecycleWorkflows\customTaskExtensions\{customTaskExtension-id}\createdBy\serviceProvisioningErrors
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetQueryParameters get serviceProvisioningErrors property value
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
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
// LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderInternal instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) {
    m := &LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/customTaskExtensions/{customTaskExtension%2Did}/createdBy/serviceProvisioningErrors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) Count()(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsCountRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get serviceProvisioningErrors property value
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceProvisioningErrorCollectionResponseable, error) {
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
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder) {
    return NewLifecycleWorkflowsCustomTaskExtensionsItemCreatedByServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
