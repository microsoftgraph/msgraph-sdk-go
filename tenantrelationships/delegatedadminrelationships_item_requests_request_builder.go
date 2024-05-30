package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DelegatedadminrelationshipsItemRequestsRequestBuilder provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
type DelegatedadminrelationshipsItemRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadminrelationshipsItemRequestsRequestBuilderGetQueryParameters get a list of the delegatedAdminRelationshipRequest objects and their properties.
type DelegatedadminrelationshipsItemRequestsRequestBuilderGetQueryParameters struct {
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
// DelegatedadminrelationshipsItemRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadminrelationshipsItemRequestsRequestBuilderGetQueryParameters
}
// DelegatedadminrelationshipsItemRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDelegatedAdminRelationshipRequestId provides operations to manage the requests property of the microsoft.graph.delegatedAdminRelationship entity.
// returns a *DelegatedadminrelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilder when successful
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) ByDelegatedAdminRelationshipRequestId(delegatedAdminRelationshipRequestId string)(*DelegatedadminrelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if delegatedAdminRelationshipRequestId != "" {
        urlTplParams["delegatedAdminRelationshipRequest%2Did"] = delegatedAdminRelationshipRequestId
    }
    return NewDelegatedadminrelationshipsItemRequestsDelegatedAdminRelationshipRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDelegatedadminrelationshipsItemRequestsRequestBuilderInternal instantiates a new DelegatedadminrelationshipsItemRequestsRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemRequestsRequestBuilder) {
    m := &DelegatedadminrelationshipsItemRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}/requests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDelegatedadminrelationshipsItemRequestsRequestBuilder instantiates a new DelegatedadminrelationshipsItemRequestsRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadminrelationshipsItemRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DelegatedadminrelationshipsItemRequestsCountRequestBuilder when successful
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) Count()(*DelegatedadminrelationshipsItemRequestsCountRequestBuilder) {
    return NewDelegatedadminrelationshipsItemRequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the delegatedAdminRelationshipRequest objects and their properties.
// returns a DelegatedAdminRelationshipRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-list-requests?view=graph-rest-1.0
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestCollectionResponseable), nil
}
// Post create a new delegatedAdminRelationshipRequest object.
// returns a DelegatedAdminRelationshipRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationship-post-requests?view=graph-rest-1.0
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestable, requestConfiguration *DelegatedadminrelationshipsItemRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminRelationshipRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestable), nil
}
// ToGetRequestInformation get a list of the delegatedAdminRelationshipRequest objects and their properties.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new delegatedAdminRelationshipRequest object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminRelationshipRequestable, requestConfiguration *DelegatedadminrelationshipsItemRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadminrelationshipsItemRequestsRequestBuilder when successful
func (m *DelegatedadminrelationshipsItemRequestsRequestBuilder) WithUrl(rawUrl string)(*DelegatedadminrelationshipsItemRequestsRequestBuilder) {
    return NewDelegatedadminrelationshipsItemRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
