package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder provides operations to manage the appManagementPolicies property of the microsoft.graph.application entity.
type ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetQueryParameters the appManagementPolicy applied to this application.
type ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetQueryParameters struct {
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
// ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetQueryParameters
}
// ByAppManagementPolicyId gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.appManagementPolicies.item collection
// returns a *ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder when successful
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) ByAppManagementPolicyId(appManagementPolicyId string)(*ItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appManagementPolicyId != "" {
        urlTplParams["appManagementPolicy%2Did"] = appManagementPolicyId
    }
    return NewItemAppmanagementpoliciesAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderInternal instantiates a new ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder and sets the default values.
func NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) {
    m := &ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/appManagementPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder instantiates a new ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder and sets the default values.
func NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAppmanagementpoliciesCountRequestBuilder when successful
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) Count()(*ItemAppmanagementpoliciesCountRequestBuilder) {
    return NewItemAppmanagementpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the appManagementPolicy applied to this application.
// returns a AppManagementPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppManagementPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppManagementPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppManagementPolicyCollectionResponseable), nil
}
// Ref provides operations to manage the collection of application entities.
// returns a *ItemAppmanagementpoliciesRefRequestBuilder when successful
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) Ref()(*ItemAppmanagementpoliciesRefRequestBuilder) {
    return NewItemAppmanagementpoliciesRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the appManagementPolicy applied to this application.
// returns a *RequestInformation when successful
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder when successful
func (m *ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) WithUrl(rawUrl string)(*ItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder) {
    return NewItemAppmanagementpoliciesAppManagementPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
