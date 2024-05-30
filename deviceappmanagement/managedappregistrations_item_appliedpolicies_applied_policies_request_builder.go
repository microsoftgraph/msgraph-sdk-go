package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
type ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetQueryParameters zero or more policys already applied on the registered app when it last synchronized with managment service.
type ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetQueryParameters struct {
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
// ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetQueryParameters
}
// ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedAppPolicyId provides operations to manage the appliedPolicies property of the microsoft.graph.managedAppRegistration entity.
// returns a *ManagedappregistrationsItemAppliedpoliciesManagedAppPolicyItemRequestBuilder when successful
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) ByManagedAppPolicyId(managedAppPolicyId string)(*ManagedappregistrationsItemAppliedpoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedAppPolicyId != "" {
        urlTplParams["managedAppPolicy%2Did"] = managedAppPolicyId
    }
    return NewManagedappregistrationsItemAppliedpoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderInternal instantiates a new ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder and sets the default values.
func NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) {
    m := &ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/appliedPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder instantiates a new ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder and sets the default values.
func NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedappregistrationsItemAppliedpoliciesCountRequestBuilder when successful
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) Count()(*ManagedappregistrationsItemAppliedpoliciesCountRequestBuilder) {
    return NewManagedappregistrationsItemAppliedpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more policys already applied on the registered app when it last synchronized with managment service.
// returns a ManagedAppPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedAppPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyCollectionResponseable), nil
}
// Post create new navigation property to appliedPolicies for deviceAppManagement
// returns a ManagedAppPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation zero or more policys already applied on the registered app when it last synchronized with managment service.
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to appliedPolicies for deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder when successful
func (m *ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) WithUrl(rawUrl string)(*ManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder) {
    return NewManagedappregistrationsItemAppliedpoliciesAppliedPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
