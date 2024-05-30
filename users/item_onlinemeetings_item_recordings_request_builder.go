package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemRecordingsRequestBuilder provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
type ItemOnlinemeetingsItemRecordingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemRecordingsRequestBuilderGetQueryParameters get a callRecording object associated with a scheduled onlineMeeting. This API doesn't support getting call recordings from channel meetings. For a recording, this API returns the metadata of the single recording associated with the online meeting. For the content of a recording, this API returns the stream of bytes associated with the recording.
type ItemOnlinemeetingsItemRecordingsRequestBuilderGetQueryParameters struct {
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
// ItemOnlinemeetingsItemRecordingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemRecordingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlinemeetingsItemRecordingsRequestBuilderGetQueryParameters
}
// ItemOnlinemeetingsItemRecordingsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemRecordingsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCallRecordingId provides operations to manage the recordings property of the microsoft.graph.onlineMeeting entity.
// returns a *ItemOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder when successful
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) ByCallRecordingId(callRecordingId string)(*ItemOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if callRecordingId != "" {
        urlTplParams["callRecording%2Did"] = callRecordingId
    }
    return NewItemOnlinemeetingsItemRecordingsCallRecordingItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemOnlinemeetingsItemRecordingsRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemRecordingsRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemRecordingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemRecordingsRequestBuilder) {
    m := &ItemOnlinemeetingsItemRecordingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/recordings{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemRecordingsRequestBuilder instantiates a new ItemOnlinemeetingsItemRecordingsRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemRecordingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemRecordingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemRecordingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemOnlinemeetingsItemRecordingsCountRequestBuilder when successful
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) Count()(*ItemOnlinemeetingsItemRecordingsCountRequestBuilder) {
    return NewItemOnlinemeetingsItemRecordingsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a callRecording object associated with a scheduled onlineMeeting. This API doesn't support getting call recordings from channel meetings. For a recording, this API returns the metadata of the single recording associated with the online meeting. For the content of a recording, this API returns the stream of bytes associated with the recording.
// returns a CallRecordingCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRecordingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallRecordingCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingCollectionResponseable), nil
}
// Post create new navigation property to recordings for users
// returns a CallRecordingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *ItemOnlinemeetingsItemRecordingsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation get a callRecording object associated with a scheduled onlineMeeting. This API doesn't support getting call recordings from channel meetings. For a recording, this API returns the metadata of the single recording associated with the online meeting. For the content of a recording, this API returns the stream of bytes associated with the recording.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlinemeetingsItemRecordingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to recordings for users
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallRecordingable, requestConfiguration *ItemOnlinemeetingsItemRecordingsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlinemeetingsItemRecordingsRequestBuilder when successful
func (m *ItemOnlinemeetingsItemRecordingsRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemRecordingsRequestBuilder) {
    return NewItemOnlinemeetingsItemRecordingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
