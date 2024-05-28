package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder provides operations to manage the collection of identityGovernance entities.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderInternal instantiates a new EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) {
    m := &EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization%2Did}/externalSponsors/{directoryObject%2Did}/$ref", pathParameters),
    }
    return m
}
// NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder instantiates a new EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder and sets the default values.
func NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a user or a group from the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/connectedorganization-delete-externalsponsors?view=graph-rest-1.0
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove a user or a group from the connected organization's external sponsors. The external sponsors are a set of users who can approve requests on behalf of other users from that connected organization.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder when successful
func (m *EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsItemExternalsponsorsItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
