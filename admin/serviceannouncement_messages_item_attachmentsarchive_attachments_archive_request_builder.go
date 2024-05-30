package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder provides operations to manage the media for the admin entity.
type ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderInternal instantiates a new ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) {
    m := &ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}/attachmentsArchive", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder instantiates a new ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder and sets the default values.
func NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete the zip file that contains all attachments for a message.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the list of attachments associated with a service message.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-list-attachments?view=graph-rest-1.0
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the zip file that contains all attachments for a message.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderPutRequestConfiguration)([]byte, error) {
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
// ToDeleteRequestInformation the zip file that contains all attachments for a message.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the list of attachments associated with a service message.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the zip file that contains all attachments for a message.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) {
    return NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
