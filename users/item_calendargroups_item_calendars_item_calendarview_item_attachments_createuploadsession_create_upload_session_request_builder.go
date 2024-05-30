package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/attachments/createUploadSession", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an upload session that allows an app to iteratively upload ranges of a file, so as to attach the file to the specified Outlook item. The item can be a message or event. Use this approach to attach a file if the file size is between 3 MB and 150 MB. To attach a file that's smaller than 3 MB, do a POST operation on the attachments navigation property of the Outlook item; see how to do this for a message or for an event. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. Request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. The following are the steps to attach a file to an Outlook item using an upload session: See attach large files to Outlook messages or events for an example.
// returns a UploadSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attachment-createuploadsession?view=graph-rest-1.0
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable), nil
}
// ToPostRequestInformation create an upload session that allows an app to iteratively upload ranges of a file, so as to attach the file to the specified Outlook item. The item can be a message or event. Use this approach to attach a file if the file size is between 3 MB and 150 MB. To attach a file that's smaller than 3 MB, do a POST operation on the attachments navigation property of the Outlook item; see how to do this for a message or for an event. As part of the response, this action returns an upload URL that you can use in subsequent sequential PUT queries. Request headers for each PUT operation let you specify the exact range of bytes to be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. The following are the steps to attach a file to an Outlook item using an upload session: See attach large files to Outlook messages or events for an example.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemCalendarviewItemAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
