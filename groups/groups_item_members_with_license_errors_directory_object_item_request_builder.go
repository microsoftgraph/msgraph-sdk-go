package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Application()(*GroupsItemMembersWithLicenseErrorsItemApplicationRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    m := &GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device casts the previous resource to device.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Device()(*GroupsItemMembersWithLicenseErrorsItemDeviceRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) Group()(*GroupsItemMembersWithLicenseErrorsItemGroupRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) OrgContact()(*GroupsItemMembersWithLicenseErrorsItemOrgContactRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) ServicePrincipal()(*GroupsItemMembersWithLicenseErrorsItemServicePrincipalRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *GroupsItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) User()(*GroupsItemMembersWithLicenseErrorsItemUserRequestBuilder) {
    return NewGroupsItemMembersWithLicenseErrorsItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
