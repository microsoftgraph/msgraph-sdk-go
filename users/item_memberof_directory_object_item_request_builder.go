package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMemberofDirectoryObjectItemRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.user entity.
type ItemMemberofDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberofDirectoryObjectItemRequestBuilderGetQueryParameters the groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
type ItemMemberofDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMemberofDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberofDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberofDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemMemberofDirectoryObjectItemRequestBuilderInternal instantiates a new ItemMemberofDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMemberofDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberofDirectoryObjectItemRequestBuilder) {
    m := &ItemMemberofDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/memberOf/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMemberofDirectoryObjectItemRequestBuilder instantiates a new ItemMemberofDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMemberofDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberofDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberofDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberofDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// GraphAdministrativeUnit casts the previous resource to administrativeUnit.
// returns a *ItemMemberofItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilder when successful
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) GraphAdministrativeUnit()(*ItemMemberofItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilder) {
    return NewItemMemberofItemGraphadministrativeunitGraphAdministrativeUnitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDirectoryRole casts the previous resource to directoryRole.
// returns a *ItemMemberofItemGraphdirectoryroleGraphDirectoryRoleRequestBuilder when successful
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) GraphDirectoryRole()(*ItemMemberofItemGraphdirectoryroleGraphDirectoryRoleRequestBuilder) {
    return NewItemMemberofItemGraphdirectoryroleGraphDirectoryRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemMemberofItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemMemberofItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemMemberofItemGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the groups and directory roles that the user is a member of. Read-only. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberofDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMemberofDirectoryObjectItemRequestBuilder when successful
func (m *ItemMemberofDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*ItemMemberofDirectoryObjectItemRequestBuilder) {
    return NewItemMemberofDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
