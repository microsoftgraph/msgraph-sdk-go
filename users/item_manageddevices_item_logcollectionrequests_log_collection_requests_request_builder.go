package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetQueryParameters list of log collection requests
type ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetQueryParameters struct {
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
// ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceLogCollectionResponseId provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
// returns a *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) ByDeviceLogCollectionResponseId(deviceLogCollectionResponseId string)(*ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceLogCollectionResponseId != "" {
        urlTplParams["deviceLogCollectionResponse%2Did"] = deviceLogCollectionResponseId
    }
    return NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal instantiates a new ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    m := &ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/logCollectionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder instantiates a new ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemManageddevicesItemLogcollectionrequestsCountRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) Count()(*ItemManageddevicesItemLogcollectionrequestsCountRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of log collection requests
// returns a DeviceLogCollectionResponseCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceLogCollectionResponseCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseCollectionResponseable), nil
}
// Post create new navigation property to logCollectionRequests for users
// returns a DeviceLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceLogCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable), nil
}
// ToGetRequestInformation list of log collection requests
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to logCollectionRequests for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsLogCollectionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
