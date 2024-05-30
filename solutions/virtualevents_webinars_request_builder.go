package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsRequestBuilder provides operations to manage the webinars property of the microsoft.graph.virtualEventsRoot entity.
type VirtualeventsWebinarsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsRequestBuilderGetQueryParameters get the list of all virtualEventWebinar objects created in the tenant.
type VirtualeventsWebinarsRequestBuilderGetQueryParameters struct {
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
// VirtualeventsWebinarsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsRequestBuilderGetQueryParameters
}
// VirtualeventsWebinarsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByVirtualEventWebinarId provides operations to manage the webinars property of the microsoft.graph.virtualEventsRoot entity.
// returns a *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder when successful
func (m *VirtualeventsWebinarsRequestBuilder) ByVirtualEventWebinarId(virtualEventWebinarId string)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if virtualEventWebinarId != "" {
        urlTplParams["virtualEventWebinar%2Did"] = virtualEventWebinarId
    }
    return NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualeventsWebinarsRequestBuilderInternal instantiates a new VirtualeventsWebinarsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsRequestBuilder) {
    m := &VirtualeventsWebinarsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsRequestBuilder instantiates a new VirtualeventsWebinarsRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *VirtualeventsWebinarsCountRequestBuilder when successful
func (m *VirtualeventsWebinarsRequestBuilder) Count()(*VirtualeventsWebinarsCountRequestBuilder) {
    return NewVirtualeventsWebinarsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of all virtualEventWebinar objects created in the tenant.
// returns a VirtualEventWebinarCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventsroot-list-webinars?view=graph-rest-1.0
func (m *VirtualeventsWebinarsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVirtualEventWebinarCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarCollectionResponseable), nil
}
// GetByUserIdAndRoleWithUserIdWithRole provides operations to call the getByUserIdAndRole method.
// returns a *VirtualeventsWebinarsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder when successful
func (m *VirtualeventsWebinarsRequestBuilder) GetByUserIdAndRoleWithUserIdWithRole(role *string, userId *string)(*VirtualeventsWebinarsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilder) {
    return NewVirtualeventsWebinarsGetbyuseridandrolewithuseridwithroleGetByUserIdAndRoleWithUserIdWithRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, role, userId)
}
// GetByUserRoleWithRole provides operations to call the getByUserRole method.
// returns a *VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder when successful
func (m *VirtualeventsWebinarsRequestBuilder) GetByUserRoleWithRole(role *string)(*VirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilder) {
    return NewVirtualeventsWebinarsGetbyuserrolewithroleGetByUserRoleWithRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, role)
}
// Post create new navigation property to webinars for solutions
// returns a VirtualEventWebinarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVirtualEventWebinarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable), nil
}
// ToGetRequestInformation get the list of all virtualEventWebinar objects created in the tenant.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to webinars for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsRequestBuilder when successful
func (m *VirtualeventsWebinarsRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsRequestBuilder) {
    return NewVirtualeventsWebinarsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
