package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder provides operations to call the cleanWindowsDevice method.
type ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal instantiates a new ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    m := &ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/cleanWindowsDevice", pathParameters),
    }
    return m
}
// NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder instantiates a new ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder and sets the default values.
func NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clean Windows device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-cleanwindowsdevice?view=graph-rest-1.0
func (m *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) Post(ctx context.Context, body ManageddevicesItemCleanwindowsdeviceCleanWindowsDevicePostRequestBodyable, requestConfiguration *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation clean Windows device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManageddevicesItemCleanwindowsdeviceCleanWindowsDevicePostRequestBodyable, requestConfiguration *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder when successful
func (m *ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder) {
    return NewManageddevicesItemCleanwindowsdeviceCleanWindowsDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
