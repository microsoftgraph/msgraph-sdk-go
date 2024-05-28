package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderInternal instantiates a new OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder and sets the default values.
func NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) {
    m := &OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}/metadataContent", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder instantiates a new OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder and sets the default values.
func NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the time-aligned metadata of the utterances in the transcript. Read-only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the time-aligned metadata of the utterances in the transcript. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation the time-aligned metadata of the utterances in the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the time-aligned metadata of the utterances in the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the time-aligned metadata of the utterances in the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder when successful
func (m *OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder) {
    return NewOnlinemeetingsItemTranscriptsItemMetadatacontentMetadataContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
