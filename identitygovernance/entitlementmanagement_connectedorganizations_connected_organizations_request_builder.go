package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetQueryParameters retrieve a list of connectedOrganization objects.
type EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetQueryParameters
}
// EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConnectedOrganizationId provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementConnectedorganizationsConnectedOrganizationItemRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) ByConnectedOrganizationId(connectedOrganizationId string)(*EntitlementmanagementConnectedorganizationsConnectedOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if connectedOrganizationId != "" {
        urlTplParams["connectedOrganization%2Did"] = connectedOrganizationId
    }
    return NewEntitlementmanagementConnectedorganizationsConnectedOrganizationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderInternal instantiates a new EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) {
    m := &EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder instantiates a new EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementConnectedorganizationsCountRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) Count()(*EntitlementmanagementConnectedorganizationsCountRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of connectedOrganization objects.
// returns a ConnectedOrganizationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/entitlementmanagement-list-connectedorganizations?view=graph-rest-1.0
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConnectedOrganizationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationCollectionResponseable), nil
}
// Post create a new connectedOrganization object.
// returns a ConnectedOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/entitlementmanagement-post-connectedorganizations?view=graph-rest-1.0
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, requestConfiguration *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConnectedOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable), nil
}
// ToGetRequestInformation retrieve a list of connectedOrganization objects.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new connectedOrganization object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConnectedOrganizationable, requestConfiguration *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
