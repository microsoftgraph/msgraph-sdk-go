package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
type ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetQueryParameters zero or more policies admin intended for the app as of now.
type ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetQueryParameters
}
// ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderInternal instantiates a new ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) {
    m := &ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/intendedPolicies/{managedAppPolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder instantiates a new ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property intendedPolicies for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get zero or more policies admin intended for the app as of now.
// returns a ManagedAppPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable), nil
}
// Patch update the navigation property intendedPolicies in deviceAppManagement
// returns a ManagedAppPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable), nil
}
// TargetApps provides operations to call the targetApps method.
// returns a *ManagedappregistrationsItemIntendedpoliciesItemTargetappsTargetAppsRequestBuilder when successful
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) TargetApps()(*ManagedappregistrationsItemIntendedpoliciesItemTargetappsTargetAppsRequestBuilder) {
    return NewManagedappregistrationsItemIntendedpoliciesItemTargetappsTargetAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property intendedPolicies for deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation zero or more policies admin intended for the app as of now.
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property intendedPolicies in deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder when successful
func (m *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) WithUrl(rawUrl string)(*ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) {
    return NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
