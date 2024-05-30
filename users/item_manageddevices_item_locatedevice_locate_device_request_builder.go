package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder provides operations to call the locateDevice method.
type ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal instantiates a new ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    m := &ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/locateDevice", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder instantiates a new ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post locate a device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-locatedevice?view=graph-rest-1.0
func (m *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation locate a device
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewItemManageddevicesItemLocatedeviceLocateDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
