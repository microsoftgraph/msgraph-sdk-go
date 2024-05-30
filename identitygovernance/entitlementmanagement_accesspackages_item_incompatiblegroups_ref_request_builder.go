package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteQueryParameters remove a group from the list of groups that have been marked as incompatible on an accessPackage.  
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteQueryParameters struct {
    // The delete Uri
    Id *string `uriparametername:"%40id"`
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteQueryParameters
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetQueryParameters retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/$ref?@id={%40id}{&%24count,%24filter,%24orderby,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder instantiates a new EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a group from the list of groups that have been marked as incompatible on an accessPackage.  
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-delete-incompatiblegroup?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
// returns a StringCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-list-incompatiblegroups?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable), nil
}
// Post add a group to the list of groups that have been marked as incompatible on an accessPackage.  
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-post-incompatiblegroup?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation remove a group from the list of groups that have been marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/$ref?@id={%40id}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation retrieve a list of the group objects that have been marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/$ref{?%24count,%24filter,%24orderby,%24search,%24skip,%24top}", m.BaseRequestBuilder.PathParameters)
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
// ToPostRequestInformation add a group to the list of groups that have been marked as incompatible on an accessPackage.  
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/$ref", m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
