package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder provides operations to call the updateWindowsDeviceAccount method.
type ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal instantiates a new ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder and sets the default values.
func NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    m := &ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/updateWindowsDeviceAccount", pathParameters),
    }
    return m
}
// NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder instantiates a new ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder and sets the default values.
func NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-updatewindowsdeviceaccount?view=graph-rest-1.0
func (m *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) Post(ctx context.Context, body ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBodyable, requestConfiguration *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation not yet documented
// returns a *RequestInformation when successful
func (m *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountPostRequestBodyable, requestConfiguration *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder when successful
func (m *ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder) {
    return NewManageddevicesItemUpdatewindowsdeviceaccountUpdateWindowsDeviceAccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
