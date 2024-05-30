package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder provides operations to call the getByUserRole method.
type VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetQueryParameters get a virtualEventWebinar collection where the signed-in user is either the organizer or a coorganizer.
type VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetQueryParameters struct {
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
// VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetQueryParameters
}
// NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderInternal instantiates a new VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, role *string)(*VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) {
    m := &VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/getByUserRole(role='{role}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if role != nil {
        m.BaseRequestBuilder.PathParameters["role"] = *role
    }
    return m
}
// NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder instantiates a new VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a virtualEventWebinar collection where the signed-in user is either the organizer or a coorganizer.
// Deprecated: This method is obsolete. Use GetAsGetByUserRoleWithRoleGetResponse instead.
// returns a VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuserrole?view=graph-rest-1.0
func (m *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleResponseable), nil
}
// GetAsGetByUserRoleWithRoleGetResponse get a virtualEventWebinar collection where the signed-in user is either the organizer or a coorganizer.
// returns a VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-getbyuserrole?view=graph-rest-1.0
func (m *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) GetAsGetByUserRoleWithRoleGetResponse(ctx context.Context, requestConfiguration *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleGetResponseable), nil
}
// ToGetRequestInformation get a virtualEventWebinar collection where the signed-in user is either the organizer or a coorganizer.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder when successful
func (m *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) {
    return NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
