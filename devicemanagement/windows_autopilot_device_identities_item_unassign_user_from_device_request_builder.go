package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder provides operations to call the unassignUserFromDevice method.
type WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderInternal instantiates a new UnassignUserFromDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/unassignUserFromDevice", pathParameters),
    }
    return m
}
// NewWindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder instantiates a new UnassignUserFromDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unassigns the user from an Autopilot device.
func (m *WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder) Post(ctx context.Context, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unassigns the user from an Autopilot device.
func (m *WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemUnassignUserFromDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
