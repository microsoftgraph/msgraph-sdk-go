package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
type OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetQueryParameters the recordings of an online meeting. Read-only.
type OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetQueryParameters
}
// OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderInternal instantiates a new OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder and sets the default values.
func NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) {
    m := &OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/recordings/{callRecording%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder instantiates a new OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder and sets the default values.
func NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the cloudCommunications entity.
// returns a *OnlinemeetingsItemRecordingsItemContentRequestBuilder when successful
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) Content()(*OnlinemeetingsItemRecordingsItemContentRequestBuilder) {
    return NewOnlinemeetingsItemRecordingsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property recordings for communications
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the recordings of an online meeting. Read-only.
// returns a CallRecordingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallRecordingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable), nil
}
// Patch update the navigation property recordings in communications
// returns a CallRecordingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallRecordingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable), nil
}
// ToDeleteRequestInformation delete navigation property recordings for communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the recordings of an online meeting. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property recordings in communications
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder when successful
func (m *OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) {
    return NewOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
