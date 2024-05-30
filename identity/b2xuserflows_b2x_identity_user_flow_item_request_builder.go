package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// B2xuserflowsB2xIdentityUserFlowItemRequestBuilder provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
type B2xuserflowsB2xIdentityUserFlowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// B2xuserflowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters retrieve the properties and relationships of a b2xIdentityUserFlow object.
type B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetQueryParameters
}
// B2xuserflowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type B2xuserflowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApiConnectorConfiguration the apiConnectorConfiguration property
// returns a *B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) ApiConnectorConfiguration()(*B2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilder) {
    return NewB2xuserflowsItemApiconnectorconfigurationApiConnectorConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilderInternal instantiates a new B2xuserflowsB2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) {
    m := &B2xuserflowsB2xIdentityUserFlowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilder instantiates a new B2xuserflowsB2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a b2xIdentityUserFlow object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-delete?view=graph-rest-1.0
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of a b2xIdentityUserFlow object.
// returns a B2xIdentityUserFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/b2xidentityuserflow-get?view=graph-rest-1.0
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateB2xIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
// returns a *B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) IdentityProviders()(*B2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilder) {
    return NewB2xuserflowsItemIdentityprovidersIdentityProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Languages provides operations to manage the languages property of the microsoft.graph.b2xIdentityUserFlow entity.
// returns a *B2xuserflowsItemLanguagesRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) Languages()(*B2xuserflowsItemLanguagesRequestBuilder) {
    return NewB2xuserflowsItemLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property b2xUserFlows in identity
// returns a B2xIdentityUserFlowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateB2xIdentityUserFlowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable), nil
}
// ToDeleteRequestInformation delete a b2xIdentityUserFlow object.
// returns a *RequestInformation when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of a b2xIdentityUserFlow object.
// returns a *RequestInformation when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property b2xUserFlows in identity
// returns a *RequestInformation when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, requestConfiguration *B2xuserflowsB2xIdentityUserFlowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserAttributeAssignments provides operations to manage the userAttributeAssignments property of the microsoft.graph.b2xIdentityUserFlow entity.
// returns a *B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*B2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilder) {
    return NewB2xuserflowsItemUserattributeassignmentsUserAttributeAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserFlowIdentityProviders provides operations to manage the userFlowIdentityProviders property of the microsoft.graph.b2xIdentityUserFlow entity.
// returns a *B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*B2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilder) {
    return NewB2xuserflowsItemUserflowidentityprovidersUserFlowIdentityProvidersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder when successful
func (m *B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) WithUrl(rawUrl string)(*B2xuserflowsB2xIdentityUserFlowItemRequestBuilder) {
    return NewB2xuserflowsB2xIdentityUserFlowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
