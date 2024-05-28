package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemTranscriptsItemContentRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type OnlinemeetingsItemTranscriptsItemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemTranscriptsItemContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemTranscriptsItemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemTranscriptsItemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemTranscriptsItemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemTranscriptsItemContentRequestBuilderInternal instantiates a new OnlinemeetingsItemTranscriptsItemContentRequestBuilder and sets the default values.
func NewOnlinemeetingsItemTranscriptsItemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemTranscriptsItemContentRequestBuilder) {
    m := &OnlinemeetingsItemTranscriptsItemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}/content", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemTranscriptsItemContentRequestBuilder instantiates a new OnlinemeetingsItemTranscriptsItemContentRequestBuilder and sets the default values.
func NewOnlinemeetingsItemTranscriptsItemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemTranscriptsItemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemTranscriptsItemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the content of the transcript. Read-only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the content of the transcript. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the content of the transcript. Read-only.
// returns a CallTranscriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallTranscriptable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallTranscriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CallTranscriptable), nil
}
// ToDeleteRequestInformation the content of the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the content of the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content of the transcript. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemTranscriptsItemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemTranscriptsItemContentRequestBuilder when successful
func (m *OnlinemeetingsItemTranscriptsItemContentRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemTranscriptsItemContentRequestBuilder) {
    return NewOnlinemeetingsItemTranscriptsItemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
