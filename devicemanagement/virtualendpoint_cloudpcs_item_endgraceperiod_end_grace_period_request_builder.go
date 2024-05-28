package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder provides operations to call the endGracePeriod method.
type VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderInternal instantiates a new VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) {
    m := &VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/endGracePeriod", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder instantiates a new VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderInternal(urlParams, requestAdapter)
}
// Post end the grace period for a specific cloudPC object. The grace period is triggered when the Cloud PC license is removed or the provisioning policy is unassigned. It allows users to access Cloud PCs for up to seven days before deprovisioning occurs. Ending the grace period immediately deprovisions the Cloud PC without waiting the seven days.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-endgraceperiod?view=graph-rest-1.0
func (m *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) Post(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation end the grace period for a specific cloudPC object. The grace period is triggered when the Cloud PC license is removed or the provisioning policy is unassigned. It allows users to access Cloud PCs for up to seven days before deprovisioning occurs. Ending the grace period immediately deprovisions the Cloud PC without waiting the seven days.
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder when successful
func (m *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) {
    return NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
