package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder provides operations to call the checkGrantedPermissionsForApp method.
type ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderInternal instantiates a new ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) {
    m := &ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/checkGrantedPermissionsForApp", pathParameters),
    }
    return m
}
// NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder instantiates a new ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder and sets the default values.
func NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkGrantedPermissionsForApp
// Deprecated: This method is obsolete. Use PostAsCheckGrantedPermissionsForAppPostResponse instead.
// returns a ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppResponseable), nil
}
// PostAsCheckGrantedPermissionsForAppPostResponse invoke action checkGrantedPermissionsForApp
// returns a ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) PostAsCheckGrantedPermissionsForAppPostResponse(ctx context.Context, requestConfiguration *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppPostResponseable), nil
}
// ToPostRequestInformation invoke action checkGrantedPermissionsForApp
// returns a *RequestInformation when successful
func (m *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder when successful
func (m *ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) WithUrl(rawUrl string)(*ItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder) {
    return NewItemCheckgrantedpermissionsforappCheckGrantedPermissionsForAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
