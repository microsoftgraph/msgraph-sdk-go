package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder provides operations to call the windowsDefenderUpdateSignatures method.
type ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal instantiates a new WindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    m := &ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/windowsDefenderUpdateSignatures", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder instantiates a new WindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-windowsdefenderupdatesignatures?view=graph-rest-1.0
func (m *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation not yet documented
func (m *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    return NewItemManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
