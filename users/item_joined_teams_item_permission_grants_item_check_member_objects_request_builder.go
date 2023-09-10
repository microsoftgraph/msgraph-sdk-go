package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder provides operations to call the checkMemberObjects method.
type ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    m := &ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/permissionGrants/{resourceSpecificPermissionGrant%2Did}/checkMemberObjects", pathParameters),
    }
    return m
}
// NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder instantiates a new CheckMemberObjectsRequestBuilder and sets the default values.
func NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action checkMemberObjects
func (m *ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) Post(ctx context.Context, body ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration)(ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsResponseable), nil
}
// ToPostRequestInformation invoke action checkMemberObjects
func (m *ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsPostRequestBodyable, requestConfiguration *ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder) {
    return NewItemJoinedTeamsItemPermissionGrantsItemCheckMemberObjectsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
