package devices

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRegisteredOwnersItemRefRequestBuilder provides operations to manage the collection of device entities.
type ItemRegisteredOwnersItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters remove a user as a registered owner of the device. This API is available in the following national cloud deployments.
type ItemRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// ItemRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters
}
// NewItemRegisteredOwnersItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewItemRegisteredOwnersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersItemRefRequestBuilder) {
    m := &ItemRegisteredOwnersItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewItemRegisteredOwnersItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewItemRegisteredOwnersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredOwnersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredOwnersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a user as a registered owner of the device. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/device-delete-registeredowners?view=graph-rest-1.0
func (m *ItemRegisteredOwnersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ToDeleteRequestInformation remove a user as a registered owner of the device. This API is available in the following national cloud deployments.
func (m *ItemRegisteredOwnersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemRegisteredOwnersItemRefRequestBuilder) WithUrl(rawUrl string)(*ItemRegisteredOwnersItemRefRequestBuilder) {
    return NewItemRegisteredOwnersItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
