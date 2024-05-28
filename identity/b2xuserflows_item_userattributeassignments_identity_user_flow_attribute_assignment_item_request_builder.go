package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
type B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters read the properties and relationships of an identityUserFlowAttributeAssignment object.
type B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetQueryParameters
}
// B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal instantiates a new B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    m := &B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userAttributeAssignments/{identityUserFlowAttributeAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder instantiates a new B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an identityUserFlowAttributeAssignment object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-delete?view=graph-rest-1.0
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an identityUserFlowAttributeAssignment object.
// returns a IdentityUserFlowAttributeAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-get?view=graph-rest-1.0
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of a identityUserFlowAttributeAssignment object.
// returns a IdentityUserFlowAttributeAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identityuserflowattributeassignment-update?view=graph-rest-1.0
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete an identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a identityUserFlowAttributeAssignment object.
// returns a *RequestInformation when successful
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityUserFlowAttributeAssignmentable, requestConfiguration *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserAttribute provides operations to manage the userAttribute property of the microsoft.graph.identityUserFlowAttributeAssignment entity.
// returns a *B2xuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) UserAttribute()(*B2xuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsItemUserattributeUserAttributeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder when successful
func (m *B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsIdentityUserFlowAttributeAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
