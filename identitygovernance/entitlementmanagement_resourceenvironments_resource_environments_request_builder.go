package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder provides operations to manage the resourceEnvironments property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetQueryParameters retrieve a list of accessPackageResourceEnvironment objects and their properties.
type EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetQueryParameters
}
// EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceEnvironmentId provides operations to manage the resourceEnvironments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementResourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilder when successful
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) ByAccessPackageResourceEnvironmentId(accessPackageResourceEnvironmentId string)(*EntitlementmanagementResourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceEnvironmentId != "" {
        urlTplParams["accessPackageResourceEnvironment%2Did"] = accessPackageResourceEnvironmentId
    }
    return NewEntitlementmanagementResourceenvironmentsAccessPackageResourceEnvironmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderInternal instantiates a new EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) {
    m := &EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resourceEnvironments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder instantiates a new EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder and sets the default values.
func NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementResourceenvironmentsCountRequestBuilder when successful
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) Count()(*EntitlementmanagementResourceenvironmentsCountRequestBuilder) {
    return NewEntitlementmanagementResourceenvironmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of accessPackageResourceEnvironment objects and their properties.
// returns a AccessPackageResourceEnvironmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/entitlementmanagement-list-resourceenvironments?view=graph-rest-1.0
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceEnvironmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentCollectionResponseable), nil
}
// Post create new navigation property to resourceEnvironments for identityGovernance
// returns a AccessPackageResourceEnvironmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, requestConfiguration *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceEnvironmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable), nil
}
// ToGetRequestInformation retrieve a list of accessPackageResourceEnvironment objects and their properties.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to resourceEnvironments for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceEnvironmentable, requestConfiguration *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder when successful
func (m *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) {
    return NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
