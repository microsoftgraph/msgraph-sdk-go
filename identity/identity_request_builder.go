package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityRequestBuilder provides operations to manage the identityContainer singleton.
type IdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityRequestBuilderGetQueryParameters
}
// IdentityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApiConnectors provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectors()(*IdentityApiConnectorsRequestBuilder) {
    return NewIdentityApiConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiConnectorsById provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectorsById(id string)(*IdentityApiConnectorsIdentityApiConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityApiConnector%2Did"] = id
    }
    return NewIdentityApiConnectorsIdentityApiConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2xUserFlows provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlows()(*IdentityB2xUserFlowsRequestBuilder) {
    return NewIdentityB2xUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2xUserFlowsById provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlowsById(id string)(*IdentityB2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2xIdentityUserFlow%2Did"] = id
    }
    return NewIdentityB2xUserFlowsB2xIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccess provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ConditionalAccess()(*IdentityConditionalAccessRequestBuilder) {
    return NewIdentityConditionalAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identity
func (m *IdentityRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update identity
func (m *IdentityRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get identity
func (m *IdentityRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProviders()(*IdentityIdentityProvidersRequestBuilder) {
    return NewIdentityIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProvidersById(id string)(*IdentityIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return NewIdentityIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update identity
func (m *IdentityRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IdentityContainerable), nil
}
// UserFlowAttributes provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributes()(*IdentityUserFlowAttributesRequestBuilder) {
    return NewIdentityUserFlowAttributesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowAttributesById provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributesById(id string)(*IdentityUserFlowAttributesIdentityUserFlowAttributeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttribute%2Did"] = id
    }
    return NewIdentityUserFlowAttributesIdentityUserFlowAttributeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
