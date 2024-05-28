package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    m := &DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/deletedItems/getAvailableExtensionProperties", pathParameters),
    }
    return m
}
// NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder instantiates a new DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// Deprecated: This method is obsolete. Use PostAsGetAvailableExtensionPropertiesPostResponse instead.
// returns a DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable), nil
}
// PostAsGetAvailableExtensionPropertiesPostResponse return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// returns a DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) PostAsGetAvailableExtensionPropertiesPostResponse(ctx context.Context, body DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable), nil
}
// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// returns a *RequestInformation when successful
func (m *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder when successful
func (m *DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) WithUrl(rawUrl string)(*DeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    return NewDeleteditemsGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
