package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19 "github.com/microsoftgraph/msgraph-sdk-go/models/callrecords"
)

// CallrecordsItemOrganizer_v2RequestBuilder provides operations to manage the organizer_v2 property of the microsoft.graph.callRecords.callRecord entity.
type CallrecordsItemOrganizer_v2RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsItemOrganizer_v2RequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemOrganizer_v2RequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallrecordsItemOrganizer_v2RequestBuilderGetQueryParameters identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
type CallrecordsItemOrganizer_v2RequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallrecordsItemOrganizer_v2RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemOrganizer_v2RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsItemOrganizer_v2RequestBuilderGetQueryParameters
}
// CallrecordsItemOrganizer_v2RequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsItemOrganizer_v2RequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallrecordsItemOrganizer_v2RequestBuilderInternal instantiates a new CallrecordsItemOrganizer_v2RequestBuilder and sets the default values.
func NewCallrecordsItemOrganizer_v2RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemOrganizer_v2RequestBuilder) {
    m := &CallrecordsItemOrganizer_v2RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/{callRecord%2Did}/organizer_v2{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCallrecordsItemOrganizer_v2RequestBuilder instantiates a new CallrecordsItemOrganizer_v2RequestBuilder and sets the default values.
func NewCallrecordsItemOrganizer_v2RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsItemOrganizer_v2RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsItemOrganizer_v2RequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property organizer_v2 for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemOrganizer_v2RequestBuilder) Delete(ctx context.Context, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderDeleteRequestConfiguration)(error) {
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
// Get identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
// returns a Organizerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemOrganizer_v2RequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderGetRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateOrganizerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable), nil
}
// Patch update the navigation property organizer_v2 in communications
// returns a Organizerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsItemOrganizer_v2RequestBuilder) Patch(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderPatchRequestConfiguration)(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.CreateOrganizerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable), nil
}
// ToDeleteRequestInformation delete navigation property organizer_v2 for communications
// returns a *RequestInformation when successful
func (m *CallrecordsItemOrganizer_v2RequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
// returns a *RequestInformation when successful
func (m *CallrecordsItemOrganizer_v2RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property organizer_v2 in communications
// returns a *RequestInformation when successful
func (m *CallrecordsItemOrganizer_v2RequestBuilder) ToPatchRequestInformation(ctx context.Context, body iaf7085b34cf3df74d75420043707a37fee7e9a355a2db4b4b46244736f7f1d19.Organizerable, requestConfiguration *CallrecordsItemOrganizer_v2RequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CallrecordsItemOrganizer_v2RequestBuilder when successful
func (m *CallrecordsItemOrganizer_v2RequestBuilder) WithUrl(rawUrl string)(*CallrecordsItemOrganizer_v2RequestBuilder) {
    return NewCallrecordsItemOrganizer_v2RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
