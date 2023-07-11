package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder provides operations to call the wipeManagedAppRegistrationsByDeviceTag method.
type ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    m := &ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/wipeManagedAppRegistrationsByDeviceTag", pathParameters),
    }
    return m
}
// NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder instantiates a new WipeManagedAppRegistrationsByDeviceTagRequestBuilder and sets the default values.
func NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Post issues a wipe operation on an app registration with specified device tag.
func (m *ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Post(ctx context.Context, body ItemWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation issues a wipe operation on an app registration with specified device tag.
func (m *ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemWipeManagedAppRegistrationsByDeviceTagPostRequestBodyable, requestConfiguration *ItemWipeManagedAppRegistrationsByDeviceTagRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
