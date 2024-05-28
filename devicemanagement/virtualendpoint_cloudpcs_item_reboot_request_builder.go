package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemRebootRequestBuilder provides operations to call the reboot method.
type VirtualendpointCloudpcsItemRebootRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemRebootRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemRebootRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemRebootRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemRebootRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRebootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRebootRequestBuilder) {
    m := &VirtualendpointCloudpcsItemRebootRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/reboot", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemRebootRequestBuilder instantiates a new VirtualendpointCloudpcsItemRebootRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemRebootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemRebootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemRebootRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reboot a specific cloudPC object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-reboot?view=graph-rest-1.0
func (m *VirtualendpointCloudpcsItemRebootRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemRebootRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation reboot a specific cloudPC object.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemRebootRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemRebootRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemRebootRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemRebootRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemRebootRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRebootRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
