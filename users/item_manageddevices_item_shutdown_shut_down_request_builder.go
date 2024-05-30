package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemShutdownShutDownRequestBuilder provides operations to call the shutDown method.
type ItemManageddevicesItemShutdownShutDownRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemShutdownShutDownRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemShutdownShutDownRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemShutdownShutDownRequestBuilderInternal instantiates a new ItemManageddevicesItemShutdownShutDownRequestBuilder and sets the default values.
func NewItemManageddevicesItemShutdownShutDownRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemShutdownShutDownRequestBuilder) {
    m := &ItemManageddevicesItemShutdownShutDownRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/shutDown", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemShutdownShutDownRequestBuilder instantiates a new ItemManageddevicesItemShutdownShutDownRequestBuilder and sets the default values.
func NewItemManageddevicesItemShutdownShutDownRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemShutdownShutDownRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemShutdownShutDownRequestBuilderInternal(urlParams, requestAdapter)
}
// Post shut down device
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-shutdown?view=graph-rest-1.0
func (m *ItemManageddevicesItemShutdownShutDownRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemShutdownShutDownRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation shut down device
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemShutdownShutDownRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemShutdownShutDownRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemShutdownShutDownRequestBuilder when successful
func (m *ItemManageddevicesItemShutdownShutDownRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemShutdownShutDownRequestBuilder) {
    return NewItemManageddevicesItemShutdownShutDownRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
