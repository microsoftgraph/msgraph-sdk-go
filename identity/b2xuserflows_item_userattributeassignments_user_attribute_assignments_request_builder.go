package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetQueryParameters get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
type B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetQueryParameters struct {
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
// B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetQueryParameters
}
// B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIdentityUserFlowAttributeAssignmentId provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
// returns a *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) ByIdentityUserFlowAttributeAssignmentId(identityUserFlowAttributeAssignmentId string)(*B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if identityUserFlowAttributeAssignmentId != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = identityUserFlowAttributeAssignmentId
    }
    return NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderInternal instantiates a new B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) {
    m := &B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder instantiates a new B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *B2xuserflowsItemUserattributeassignmentsCountRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) Count()(*B2xuserflowsItemUserattributeassignmentsCountRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
// returns a IdentityUserFlowAttributeAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-list-userattributeassignments?view=graph-rest-1.0
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentCollectionResponseable), nil
}
// GetOrder provides operations to call the getOrder method.
// returns a *B2xuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) GetOrder()(*B2xuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsGetorderGetOrderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new identityUserFlowAttributeAssignment object in a b2xIdentityUserFlow.
// returns a IdentityUserFlowAttributeAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-post-userattributeassignments?view=graph-rest-1.0
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable), nil
}
// SetOrder provides operations to call the setOrder method.
// returns a *B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) SetOrder()(*B2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsSetorderSetOrderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the identityUserFlowAttributeAssignment resources from the userAttributeAssignments navigation property in a b2xIdentityUserFlow.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new identityUserFlowAttributeAssignment object in a b2xIdentityUserFlow.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
