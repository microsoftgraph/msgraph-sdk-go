package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder provides operations to manage the webinars property of the microsoft.graph.virtualEventsRoot entity.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters read the properties and relationships of a virtualEventWebinar object.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetQueryParameters
}
// VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal instantiates a new VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    m := &VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/virtualEvents/webinars/{virtualEventWebinar%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder instantiates a new VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder and sets the default values.
func NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property webinars for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a virtualEventWebinar object.
// returns a VirtualEventWebinarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualeventwebinar-get?view=graph-rest-1.0
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property webinars in solutions
// returns a VirtualEventWebinarable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Registrations provides operations to manage the registrations property of the microsoft.graph.virtualEventWebinar entity.
// returns a *VirtualeventsWebinarsItemRegistrationsRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Registrations()(*VirtualeventsWebinarsItemRegistrationsRequestBuilder) {
    return NewVirtualeventsWebinarsItemRegistrationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sessions provides operations to manage the sessions property of the microsoft.graph.virtualEvent entity.
// returns a *VirtualeventsWebinarsItemSessionsRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) Sessions()(*VirtualeventsWebinarsItemSessionsRequestBuilder) {
    return NewVirtualeventsWebinarsItemSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property webinars for solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a virtualEventWebinar object.
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property webinars in solutions
// returns a *RequestInformation when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VirtualEventWebinarable, requestConfiguration *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder when successful
func (m *VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) WithUrl(rawUrl string)(*VirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder) {
    return NewVirtualeventsWebinarsVirtualEventWebinarItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
