package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder provides operations to manage the media for the cloudCommunications entity.
type OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderInternal instantiates a new OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder and sets the default values.
func NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) {
    m := &OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings/{onlineMeeting%2Did}/attendeeReport", pathParameters),
    }
    return m
}
// NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder instantiates a new OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder and sets the default values.
func NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderPutRequestConfiguration)([]byte, error) {
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
// ToDeleteRequestInformation the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content stream of the attendee report of a Microsoft Teams live event. Read-only.
// returns a *RequestInformation when successful
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder when successful
func (m *OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder) {
    return NewOnlinemeetingsItemAttendeereportAttendeeReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
