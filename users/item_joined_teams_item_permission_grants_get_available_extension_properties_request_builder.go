package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder provides operations to call the getAvailableExtensionProperties method.
type ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderInternal instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) {
    m := &ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/permissionGrants/getAvailableExtensionProperties", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder instantiates a new GetAvailableExtensionPropertiesRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties: This API is supported in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsGetAvailableExtensionPropertiesPostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesResponseable), nil
}
// PostAsGetAvailableExtensionPropertiesPostResponse return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties: This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/directoryobject-getavailableextensionproperties?view=graph-rest-1.0
func (m *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) PostAsGetAvailableExtensionPropertiesPostResponse(ctx context.Context, body ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostResponseable), nil
}
// ToPostRequestInformation return all directory extension definitions that have been registered in a directory, including through multi-tenant apps. The following entities support extension properties: This API is supported in the following national cloud deployments.
func (m *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
func (m *ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder) {
    return NewItemJoinedTeamsItemPermissionGrantsGetAvailableExtensionPropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
