package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMembersWithLicenseErrorsRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type ItemMembersWithLicenseErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMembersWithLicenseErrorsRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type ItemMembersWithLicenseErrorsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMembersWithLicenseErrorsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMembersWithLicenseErrorsRequestBuilderInternal instantiates a new MembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsRequestBuilder) {
    m := &ItemMembersWithLicenseErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemMembersWithLicenseErrorsRequestBuilder instantiates a new MembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewItemMembersWithLicenseErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersWithLicenseErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersWithLicenseErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) Count()(*ItemMembersWithLicenseErrorsCountRequestBuilder) {
    return NewItemMembersWithLicenseErrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// GraphApplication casts the previous resource to application.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphApplication()(*ItemMembersWithLicenseErrorsGraphApplicationRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphDevice()(*ItemMembersWithLicenseErrorsGraphDeviceRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphGroup()(*ItemMembersWithLicenseErrorsGraphGroupRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphOrgContact()(*ItemMembersWithLicenseErrorsGraphOrgContactRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphServicePrincipal()(*ItemMembersWithLicenseErrorsGraphServicePrincipalRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) GraphUser()(*ItemMembersWithLicenseErrorsGraphUserRequestBuilder) {
    return NewItemMembersWithLicenseErrorsGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMembersWithLicenseErrorsRequestBuilder) WithUrl(rawUrl string)(*ItemMembersWithLicenseErrorsRequestBuilder) {
    return NewItemMembersWithLicenseErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
