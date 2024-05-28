package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder provides operations to call the locateDevice method.
type ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal instantiates a new ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    m := &ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/locateDevice", pathParameters),
    }
    return m
}
// NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilder instantiates a new ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post locate a device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-locatedevice?view=graph-rest-1.0
func (m *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) Post(ctx context.Context, requestConfiguration *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder when successful
func (m *ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemLocatedeviceLocateDeviceRequestBuilder) {
    return NewManageddevicesItemLocatedeviceLocateDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
