package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AdministrativeunitsItemMembersRequestBuilder provides operations to manage the members property of the microsoft.graph.administrativeUnit entity.
type AdministrativeunitsItemMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdministrativeunitsItemMembersRequestBuilderGetQueryParameters use this API to get the members list (users, groups, or devices) in an administrative unit.
type AdministrativeunitsItemMembersRequestBuilderGetQueryParameters struct {
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
// AdministrativeunitsItemMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdministrativeunitsItemMembersRequestBuilderGetQueryParameters
}
// AdministrativeunitsItemMembersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdministrativeunitsItemMembersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDirectoryObjectId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.directory.administrativeUnits.item.members.item collection
// returns a *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdministrativeunitsItemMembersRequestBuilderInternal instantiates a new AdministrativeunitsItemMembersRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersRequestBuilder) {
    m := &AdministrativeunitsItemMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAdministrativeunitsItemMembersRequestBuilder instantiates a new AdministrativeunitsItemMembersRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeunitsItemMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AdministrativeunitsItemMembersCountRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) Count()(*AdministrativeunitsItemMembersCountRequestBuilder) {
    return NewAdministrativeunitsItemMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get use this API to get the members list (users, groups, or devices) in an administrative unit.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/administrativeunit-list-members?view=graph-rest-1.0
func (m *AdministrativeunitsItemMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
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
// returns a *AdministrativeunitsItemMembersGraphapplicationGraphApplicationRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphApplication()(*AdministrativeunitsItemMembersGraphapplicationGraphApplicationRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *AdministrativeunitsItemMembersGraphdeviceGraphDeviceRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphDevice()(*AdministrativeunitsItemMembersGraphdeviceGraphDeviceRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *AdministrativeunitsItemMembersGraphgroupGraphGroupRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphGroup()(*AdministrativeunitsItemMembersGraphgroupGraphGroupRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *AdministrativeunitsItemMembersGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphOrgContact()(*AdministrativeunitsItemMembersGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphServicePrincipal()(*AdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *AdministrativeunitsItemMembersGraphuserGraphUserRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) GraphUser()(*AdministrativeunitsItemMembersGraphuserGraphUserRequestBuilder) {
    return NewAdministrativeunitsItemMembersGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post use this API to add a member (user, group, or device) to an administrative unit. Currently it's only possible to add one member at a time to an administrative unit.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/administrativeunit-post-members?view=graph-rest-1.0
func (m *AdministrativeunitsItemMembersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *AdministrativeunitsItemMembersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// Ref provides operations to manage the collection of directory entities.
// returns a *AdministrativeunitsItemMembersRefRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) Ref()(*AdministrativeunitsItemMembersRefRequestBuilder) {
    return NewAdministrativeunitsItemMembersRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation use this API to get the members list (users, groups, or devices) in an administrative unit.
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdministrativeunitsItemMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation use this API to add a member (user, group, or device) to an administrative unit. Currently it's only possible to add one member at a time to an administrative unit.
// returns a *RequestInformation when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, requestConfiguration *AdministrativeunitsItemMembersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AdministrativeunitsItemMembersRequestBuilder when successful
func (m *AdministrativeunitsItemMembersRequestBuilder) WithUrl(rawUrl string)(*AdministrativeunitsItemMembersRequestBuilder) {
    return NewAdministrativeunitsItemMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
