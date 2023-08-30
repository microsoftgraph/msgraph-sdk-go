package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemValidatePropertiesRequestBuilder provides operations to call the validateProperties method.
type ItemValidatePropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemValidatePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemValidatePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemValidatePropertiesRequestBuilderInternal instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatePropertiesRequestBuilder) {
    m := &ItemValidatePropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/validateProperties", pathParameters),
    }
    return m
}
// NewItemValidatePropertiesRequestBuilder instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemValidatePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use this API to determine whether a display name or mail nickname is valid before trying to update a Microsoft 365 group. To validate the properties before creating a group, use the directoryobject:validateProperties function. The following policy validations are performed for the display name and mail nickname properties: This API only returns the first validation failure that is encountered. If the properties fail multiple validations, only the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy. To learn more about configuring naming policies, see Configure naming policy.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-validateproperties?view=graph-rest-1.0
func (m *ItemValidatePropertiesRequestBuilder) Post(ctx context.Context, body ItemValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatePropertiesRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation validate that a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use this API to determine whether a display name or mail nickname is valid before trying to update a Microsoft 365 group. To validate the properties before creating a group, use the directoryobject:validateProperties function. The following policy validations are performed for the display name and mail nickname properties: This API only returns the first validation failure that is encountered. If the properties fail multiple validations, only the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy. To learn more about configuring naming policies, see Configure naming policy.
func (m *ItemValidatePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemValidatePropertiesRequestBuilder) WithUrl(rawUrl string)(*ItemValidatePropertiesRequestBuilder) {
    return NewItemValidatePropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
