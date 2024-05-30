package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
type ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetQueryParameters zero or more policies admin intended for the app as of now.
type ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetQueryParameters struct {
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
// ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetQueryParameters
}
// ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedAppPolicyId provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
// returns a *ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder when successful
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) ByManagedAppPolicyId(managedAppPolicyId string)(*ManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedAppPolicyId != "" {
        urlTplParams["managedAppPolicy%2Did"] = managedAppPolicyId
    }
    return NewManagedappregistrationsItemIntendedpoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderInternal instantiates a new ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder and sets the default values.
func NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) {
    m := &ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/intendedPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder instantiates a new ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder and sets the default values.
func NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedappregistrationsItemIntendedpoliciesCountRequestBuilder when successful
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) Count()(*ManagedappregistrationsItemIntendedpoliciesCountRequestBuilder) {
    return NewManagedappregistrationsItemIntendedpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more policies admin intended for the app as of now.
// returns a ManagedAppPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyCollectionResponseable, error) {
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
// Post create new navigation property to intendedPolicies for deviceAppManagement
// returns a ManagedAppPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, error) {
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
// ToGetRequestInformation zero or more policies admin intended for the app as of now.
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to intendedPolicies for deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedAppPolicyable, requestConfiguration *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder when successful
func (m *ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) WithUrl(rawUrl string)(*ManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder) {
    return NewManagedappregistrationsItemIntendedpoliciesIntendedPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
