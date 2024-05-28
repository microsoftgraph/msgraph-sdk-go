package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ServiceannouncementMessagesRequestBuilder provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
type ServiceannouncementMessagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceannouncementMessagesRequestBuilderGetQueryParameters retrieve the serviceUpdateMessage resources from the messages navigation property. This operation retrieves all service update messages that exist for the tenant.
type ServiceannouncementMessagesRequestBuilderGetQueryParameters struct {
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
// ServiceannouncementMessagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceannouncementMessagesRequestBuilderGetQueryParameters
}
// ServiceannouncementMessagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceannouncementMessagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Archive provides operations to call the archive method.
// returns a *ServiceannouncementMessagesArchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) Archive()(*ServiceannouncementMessagesArchiveRequestBuilder) {
    return NewServiceannouncementMessagesArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByServiceUpdateMessageId provides operations to manage the messages property of the microsoft.graph.serviceAnnouncement entity.
// returns a *ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) ByServiceUpdateMessageId(serviceUpdateMessageId string)(*ServiceannouncementMessagesServiceUpdateMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if serviceUpdateMessageId != "" {
        urlTplParams["serviceUpdateMessage%2Did"] = serviceUpdateMessageId
    }
    return NewServiceannouncementMessagesServiceUpdateMessageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceannouncementMessagesRequestBuilderInternal instantiates a new ServiceannouncementMessagesRequestBuilder and sets the default values.
func NewServiceannouncementMessagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesRequestBuilder) {
    m := &ServiceannouncementMessagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/serviceAnnouncement/messages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewServiceannouncementMessagesRequestBuilder instantiates a new ServiceannouncementMessagesRequestBuilder and sets the default values.
func NewServiceannouncementMessagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceannouncementMessagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceannouncementMessagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ServiceannouncementMessagesCountRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) Count()(*ServiceannouncementMessagesCountRequestBuilder) {
    return NewServiceannouncementMessagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Favorite provides operations to call the favorite method.
// returns a *ServiceannouncementMessagesFavoriteRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) Favorite()(*ServiceannouncementMessagesFavoriteRequestBuilder) {
    return NewServiceannouncementMessagesFavoriteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the serviceUpdateMessage resources from the messages navigation property. This operation retrieves all service update messages that exist for the tenant.
// returns a ServiceUpdateMessageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceannouncement-list-messages?view=graph-rest-1.0
func (m *ServiceannouncementMessagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceannouncementMessagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateServiceUpdateMessageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageCollectionResponseable), nil
}
// MarkRead provides operations to call the markRead method.
// returns a *ServiceannouncementMessagesMarkreadMarkReadRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) MarkRead()(*ServiceannouncementMessagesMarkreadMarkReadRequestBuilder) {
    return NewServiceannouncementMessagesMarkreadMarkReadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MarkUnread provides operations to call the markUnread method.
// returns a *ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) MarkUnread()(*ServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilder) {
    return NewServiceannouncementMessagesMarkunreadMarkUnreadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to messages for admin
// returns a ServiceUpdateMessageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ServiceannouncementMessagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceannouncementMessagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve the serviceUpdateMessage resources from the messages navigation property. This operation retrieves all service update messages that exist for the tenant.
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceannouncementMessagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to messages for admin
// returns a *RequestInformation when successful
func (m *ServiceannouncementMessagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ServiceUpdateMessageable, requestConfiguration *ServiceannouncementMessagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unarchive provides operations to call the unarchive method.
// returns a *ServiceannouncementMessagesUnarchiveRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) Unarchive()(*ServiceannouncementMessagesUnarchiveRequestBuilder) {
    return NewServiceannouncementMessagesUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Unfavorite provides operations to call the unfavorite method.
// returns a *ServiceannouncementMessagesUnfavoriteRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) Unfavorite()(*ServiceannouncementMessagesUnfavoriteRequestBuilder) {
    return NewServiceannouncementMessagesUnfavoriteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServiceannouncementMessagesRequestBuilder when successful
func (m *ServiceannouncementMessagesRequestBuilder) WithUrl(rawUrl string)(*ServiceannouncementMessagesRequestBuilder) {
    return NewServiceannouncementMessagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
