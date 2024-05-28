package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
type ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a serviceUpdateMessage object. This operation retrieves a specified service update message for the tenant. The operation returns an error if the message does not exist for the tenant.
type ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetQueryParameters
}
// ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.serviceUpdateMessage entity.
// returns a *ServiceannouncementMessagesItemAttachmentsRequestBuilder when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) Attachments()(*ServiceannouncementMessagesItemAttachmentsRequestBuilder) {
    return NewServiceannouncementMessagesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttachmentsArchive provides operations to manage the media for the admin entity.
// returns a *ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) AttachmentsArchive()(*ServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilder) {
    return NewServiceannouncementMessagesItemAttachmentsarchiveAttachmentsArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderInternal instantiates a new ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder and sets the default values.
func NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) {
    m := &ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages/{serviceUpdateMessage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder instantiates a new ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder and sets the default values.
func NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property messages for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a serviceUpdateMessage object. This operation retrieves a specified service update message for the tenant. The operation returns an error if the message does not exist for the tenant.
// returns a ServiceUpdateMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceupdatemessage-get?view=graph-rest-1.0
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceUpdateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable), nil
}
// Patch update the navigation property messages in admin
// returns a ServiceUpdateMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceUpdateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable), nil
}
// ToDeleteRequestInformation delete navigation property messages for admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a serviceUpdateMessage object. This operation retrieves a specified service update message for the tenant. The operation returns an error if the message does not exist for the tenant.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property messages in admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder when successful
func (m *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) {
    return NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
