package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder provides operations to manage the media for the admin entity.
type AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderInternal instantiates a new AttachmentsArchiveRequestBuilder and sets the default values.
func NewAdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) {
    m := &AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}/attachmentsArchive";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder instantiates a new AttachmentsArchiveRequestBuilder and sets the default values.
func NewAdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the zip file that contains all attachments for a message.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePutRequestInformation the zip file that contains all attachments for a message.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the zip file that contains all attachments for a message.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) Get(ctx context.Context, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put the zip file that contains all attachments for a message.
func (m *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *AdminServiceAnnouncementMessagesItemAttachmentsArchiveRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
