package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemTroubleshootRequestBuilder provides operations to call the troubleshoot method.
type VirtualendpointCloudpcsItemTroubleshootRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemTroubleshootRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemTroubleshootRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemTroubleshootRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemTroubleshootRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemTroubleshootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemTroubleshootRequestBuilder) {
    m := &VirtualendpointCloudpcsItemTroubleshootRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/troubleshoot", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemTroubleshootRequestBuilder instantiates a new VirtualendpointCloudpcsItemTroubleshootRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemTroubleshootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemTroubleshootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemTroubleshootRequestBuilderInternal(urlParams, requestAdapter)
}
// Post troubleshoot a specific cloudPC object. Use this API to check the health status of the Cloud PC and the session host.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-troubleshoot?view=graph-rest-1.0
func (m *VirtualendpointCloudpcsItemTroubleshootRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemTroubleshootRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation troubleshoot a specific cloudPC object. Use this API to check the health status of the Cloud PC and the session host.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemTroubleshootRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemTroubleshootRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemTroubleshootRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemTroubleshootRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemTroubleshootRequestBuilder) {
    return NewVirtualendpointCloudpcsItemTroubleshootRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
