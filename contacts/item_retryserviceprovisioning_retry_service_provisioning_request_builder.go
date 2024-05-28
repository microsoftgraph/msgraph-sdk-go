package contacts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder provides operations to call the retryServiceProvisioning method.
type ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderInternal instantiates a new ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder and sets the default values.
func NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) {
    m := &ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/contacts/{orgContact%2Did}/retryServiceProvisioning", pathParameters),
    }
    return m
}
// NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder instantiates a new ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder and sets the default values.
func NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action retryServiceProvisioning
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action retryServiceProvisioning
// returns a *RequestInformation when successful
func (m *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder when successful
func (m *ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) WithUrl(rawUrl string)(*ItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder) {
    return NewItemRetryserviceprovisioningRetryServiceProvisioningRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
