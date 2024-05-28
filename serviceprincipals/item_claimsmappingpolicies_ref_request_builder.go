package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemClaimsmappingpoliciesRefRequestBuilder provides operations to manage the collection of servicePrincipal entities.
type ItemClaimsmappingpoliciesRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemClaimsmappingpoliciesRefRequestBuilderDeleteQueryParameters remove a claimsMappingPolicy from a servicePrincipal.
type ItemClaimsmappingpoliciesRefRequestBuilderDeleteQueryParameters struct {
    // The delete Uri
    Id *string `uriparametername:"%40id"`
}
// ItemClaimsmappingpoliciesRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemClaimsmappingpoliciesRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemClaimsmappingpoliciesRefRequestBuilderDeleteQueryParameters
}
// ItemClaimsmappingpoliciesRefRequestBuilderGetQueryParameters list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
type ItemClaimsmappingpoliciesRefRequestBuilderGetQueryParameters struct {
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
// ItemClaimsmappingpoliciesRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemClaimsmappingpoliciesRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemClaimsmappingpoliciesRefRequestBuilderGetQueryParameters
}
// ItemClaimsmappingpoliciesRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemClaimsmappingpoliciesRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemClaimsmappingpoliciesRefRequestBuilderInternal instantiates a new ItemClaimsmappingpoliciesRefRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesRefRequestBuilder) {
    m := &ItemClaimsmappingpoliciesRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/$ref?@id={%40id}{&%24count,%24filter,%24orderby,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemClaimsmappingpoliciesRefRequestBuilder instantiates a new ItemClaimsmappingpoliciesRefRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemClaimsmappingpoliciesRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove a claimsMappingPolicy from a servicePrincipal.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-delete-claimsmappingpolicies?view=graph-rest-1.0
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
// returns a StringCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-list-claimsmappingpolicies?view=graph-rest-1.0
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable, error) {
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
// Post assign a claimsMappingPolicy to a servicePrincipal.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-post-claimsmappingpolicies?view=graph-rest-1.0
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderPostRequestConfiguration)(error) {
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
// ToDeleteRequestInformation remove a claimsMappingPolicy from a servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/$ref?@id={%40id}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation list the claimsMappingPolicy objects that are assigned to a servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/$ref{?%24count,%24filter,%24orderby,%24search,%24skip,%24top}", m.BaseRequestBuilder.PathParameters)
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
// ToPostRequestInformation assign a claimsMappingPolicy to a servicePrincipal.
// returns a *RequestInformation when successful
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *ItemClaimsmappingpoliciesRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/$ref", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemClaimsmappingpoliciesRefRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesRefRequestBuilder) WithUrl(rawUrl string)(*ItemClaimsmappingpoliciesRefRequestBuilder) {
    return NewItemClaimsmappingpoliciesRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
