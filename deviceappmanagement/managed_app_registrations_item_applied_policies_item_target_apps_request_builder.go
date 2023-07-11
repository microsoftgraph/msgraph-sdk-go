package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder provides operations to call the targetApps method.
type ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderInternal instantiates a new TargetAppsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder) {
    m := &ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/appliedPolicies/{managedAppPolicy%2Did}/targetApps", pathParameters),
    }
    return m
}
// NewManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder instantiates a new TargetAppsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action targetApps
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder) Post(ctx context.Context, body ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBodyable, requestConfiguration *ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action targetApps
func (m *ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsPostRequestBodyable, requestConfiguration *ManagedAppRegistrationsItemAppliedPoliciesItemTargetAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
