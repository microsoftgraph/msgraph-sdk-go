package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder provides operations to manage the media for the user entity.
type ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderInternal instantiates a new MetadataContentRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) {
    m := &ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}/metadataContent", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder instantiates a new MetadataContentRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the time-aligned metadata of the utterances in the transcript. Read-only.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onlinemeeting-list-transcripts?view=graph-rest-1.0
func (m *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the time-aligned metadata of the utterances in the transcript. Read-only.
func (m *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation the time-aligned metadata of the utterances in the transcript. Read-only.
func (m *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the time-aligned metadata of the utterances in the transcript. Read-only.
func (m *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) {
    return NewItemOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
