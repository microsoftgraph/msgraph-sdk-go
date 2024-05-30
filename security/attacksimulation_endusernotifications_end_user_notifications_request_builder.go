package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder provides operations to manage the endUserNotifications property of the microsoft.graph.attackSimulationRoot entity.
type AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetQueryParameters get a list of endUserNotification objects and their properties.
type AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetQueryParameters struct {
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
// AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetQueryParameters
}
// AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEndUserNotificationId provides operations to manage the endUserNotifications property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) ByEndUserNotificationId(endUserNotificationId string)(*AttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if endUserNotificationId != "" {
        urlTplParams["endUserNotification%2Did"] = endUserNotificationId
    }
    return NewAttacksimulationEndusernotificationsEndUserNotificationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderInternal instantiates a new AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder and sets the default values.
func NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) {
    m := &AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/endUserNotifications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder instantiates a new AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder and sets the default values.
func NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AttacksimulationEndusernotificationsCountRequestBuilder when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) Count()(*AttacksimulationEndusernotificationsCountRequestBuilder) {
    return NewAttacksimulationEndusernotificationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of endUserNotification objects and their properties.
// returns a EndUserNotificationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attacksimulationroot-list-endusernotifications?view=graph-rest-1.0
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEndUserNotificationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationCollectionResponseable), nil
}
// Post create new navigation property to endUserNotifications for security
// returns a EndUserNotificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationable, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEndUserNotificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationable), nil
}
// ToGetRequestInformation get a list of endUserNotification objects and their properties.
// returns a *RequestInformation when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to endUserNotifications for security
// returns a *RequestInformation when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EndUserNotificationable, requestConfiguration *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder when successful
func (m *AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder) {
    return NewAttacksimulationEndusernotificationsEndUserNotificationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
