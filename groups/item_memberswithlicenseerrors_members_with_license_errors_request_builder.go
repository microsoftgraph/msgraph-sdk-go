package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
type ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetQueryParameters a list of group members with license errors from this group-based license assignment. Read-only.
type ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetQueryParameters struct {
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
// ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
// returns a *ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemMemberswithlicenseerrorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderInternal instantiates a new ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) {
    m := &ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/membersWithLicenseErrors{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder instantiates a new ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder and sets the default values.
func NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemMemberswithlicenseerrorsCountRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) Count()(*ItemMemberswithlicenseerrorsCountRequestBuilder) {
    return NewItemMemberswithlicenseerrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of group members with license errors from this group-based license assignment. Read-only.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *ItemMemberswithlicenseerrorsGraphapplicationGraphApplicationRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphApplication()(*ItemMemberswithlicenseerrorsGraphapplicationGraphApplicationRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemMemberswithlicenseerrorsGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphDevice()(*ItemMemberswithlicenseerrorsGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemMemberswithlicenseerrorsGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphGroup()(*ItemMemberswithlicenseerrorsGraphgroupGraphGroupRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemMemberswithlicenseerrorsGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphOrgContact()(*ItemMemberswithlicenseerrorsGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemMemberswithlicenseerrorsGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphServicePrincipal()(*ItemMemberswithlicenseerrorsGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemMemberswithlicenseerrorsGraphuserGraphUserRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) GraphUser()(*ItemMemberswithlicenseerrorsGraphuserGraphUserRequestBuilder) {
    return NewItemMemberswithlicenseerrorsGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation a list of group members with license errors from this group-based license assignment. Read-only.
// returns a *RequestInformation when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder when successful
func (m *ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) WithUrl(rawUrl string)(*ItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder) {
    return NewItemMemberswithlicenseerrorsMembersWithLicenseErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
