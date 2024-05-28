package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    m := &GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/getAvailableExtensionProperties", pathParameters),
    }
    return m
}
// NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder instantiates a new GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// Deprecated: This method is obsolete. Use PostAsGetAvailableExtensionPropertiesPostResponse instead.
// returns a GetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(GetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetavailableextensionpropertiesGetAvailableExtensionPropertiesResponseable), nil
}
// PostAsGetAvailableExtensionPropertiesPostResponse return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// returns a GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) PostAsGetAvailableExtensionPropertiesPostResponse(ctx context.Context, body GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostResponseable), nil
}
// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties:
// returns a *RequestInformation when successful
func (m *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body GetavailableextensionpropertiesGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder when successful
func (m *GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) WithUrl(rawUrl string)(*GetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder) {
    return NewGetavailableextensionpropertiesGetAvailableExtensionPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
