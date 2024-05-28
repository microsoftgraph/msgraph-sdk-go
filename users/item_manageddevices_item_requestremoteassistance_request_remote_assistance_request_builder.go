package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder provides operations to call the requestRemoteAssistance method.
type ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal instantiates a new ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder and sets the default values.
func NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    m := &ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/requestRemoteAssistance", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder instantiates a new ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder and sets the default values.
func NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post request remote assistance
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-devices-manageddevice-requestremoteassistance?view=graph-rest-1.0
func (m *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation request remote assistance
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder when successful
func (m *ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder) {
    return NewItemManageddevicesItemRequestremoteassistanceRequestRemoteAssistanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
