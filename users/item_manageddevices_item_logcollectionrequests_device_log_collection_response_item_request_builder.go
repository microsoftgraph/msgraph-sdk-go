package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder provides operations to manage the logCollectionRequests property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetQueryParameters list of log collection requests
type ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderInternal instantiates a new ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) {
    m := &ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/logCollectionRequests/{deviceLogCollectionResponse%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder instantiates a new ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDownloadUrl provides operations to call the createDownloadUrl method.
// returns a *ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) CreateDownloadUrl()(*ItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsItemCreatedownloadurlCreateDownloadUrlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property logCollectionRequests for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of log collection requests
// returns a DeviceLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property logCollectionRequests in users
// returns a DeviceLogCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property logCollectionRequests for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of log collection requests
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property logCollectionRequests in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceLogCollectionResponseable, requestConfiguration *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder when successful
func (m *ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder) {
    return NewItemManageddevicesItemLogcollectionrequestsDeviceLogCollectionResponseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
