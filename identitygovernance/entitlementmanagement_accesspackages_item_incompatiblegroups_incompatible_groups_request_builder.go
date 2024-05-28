package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetQueryParameters retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetQueryParameters
}
// ByGroupId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.accessPackages.item.incompatibleGroups.item collection
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsGroupItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) ByGroupId(groupId string)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupId != "" {
        urlTplParams["group%2Did"] = groupId
    }
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsCountRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) Count()(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsCountRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
// returns a GroupCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-list-incompatiblegroups?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.GroupCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) Ref()(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
