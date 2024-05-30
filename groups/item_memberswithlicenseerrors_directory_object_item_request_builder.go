package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderInternal instantiates a new ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) {
    m := &ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder instantiates a new ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
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
// GraphApplication casts the previous resource to application.
// returns a *ItemMemberswithlicenseerrorsItemGraphapplicationGraphApplicationRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphApplication()(*ItemMemberswithlicenseerrorsItemGraphapplicationGraphApplicationRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemMemberswithlicenseerrorsItemGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphDevice()(*ItemMemberswithlicenseerrorsItemGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemMemberswithlicenseerrorsItemGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphOrgContact()(*ItemMemberswithlicenseerrorsItemGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemMemberswithlicenseerrorsItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemMemberswithlicenseerrorsItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemMemberswithlicenseerrorsItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) GraphUser()(*ItemMemberswithlicenseerrorsItemGraphuserGraphUserRequestBuilder) {
    return NewItemMemberswithlicenseerrorsItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
// returns a *RequestInformation when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) {
    return NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
