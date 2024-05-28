package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderInternal instantiates a new EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) {
    m := &EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/internalSponsors/{directoryObject%2Did}/$ref", pathParameters),
    }
    return m
}
// NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder instantiates a new EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a user or a group from the connected organization's internal sponsors. The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-delete-internalsponsors?view=graph-rest-1.0
func (m *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove a user or a group from the connected organization's internal sponsors. The internal sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsItemInternalsponsorsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
