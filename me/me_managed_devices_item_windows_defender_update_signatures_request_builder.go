package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder provides operations to call the windowsDefenderUpdateSignatures method.
type MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal instantiates a new WindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    m := &MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/microsoft.graph.windowsDefenderUpdateSignatures";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder instantiates a new WindowsDefenderUpdateSignaturesRequestBuilder and sets the default values.
func NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action windowsDefenderUpdateSignatures
func (m *MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action windowsDefenderUpdateSignatures
func (m *MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilder) Post(ctx context.Context, requestConfiguration *MeManagedDevicesItemWindowsDefenderUpdateSignaturesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
