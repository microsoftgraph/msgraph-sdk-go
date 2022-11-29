package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityConditionalAccessRequestBuilder provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
type IdentityConditionalAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityConditionalAccessRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityConditionalAccessRequestBuilderGetQueryParameters the entry point for the Conditional Access (CA) object model.
type IdentityConditionalAccessRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityConditionalAccessRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityConditionalAccessRequestBuilderGetQueryParameters
}
// IdentityConditionalAccessRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityConditionalAccessRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationContextClassReferences provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) AuthenticationContextClassReferences()(*IdentityConditionalAccessAuthenticationContextClassReferencesRequestBuilder) {
    return NewIdentityConditionalAccessAuthenticationContextClassReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationContextClassReferencesById provides operations to manage the authenticationContextClassReferences property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) AuthenticationContextClassReferencesById(id string)(*IdentityConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationContextClassReference%2Did"] = id
    }
    return NewIdentityConditionalAccessAuthenticationContextClassReferencesAuthenticationContextClassReferenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityConditionalAccessRequestBuilderInternal instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewIdentityConditionalAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessRequestBuilder) {
    m := &IdentityConditionalAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/conditionalAccess{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityConditionalAccessRequestBuilder instantiates a new ConditionalAccessRequestBuilder and sets the default values.
func NewIdentityConditionalAccessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityConditionalAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityConditionalAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property conditionalAccess for identity
func (m *IdentityConditionalAccessRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the entry point for the Conditional Access (CA) object model.
func (m *IdentityConditionalAccessRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityConditionalAccessRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property conditionalAccess in identity
func (m *IdentityConditionalAccessRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, requestConfiguration *IdentityConditionalAccessRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property conditionalAccess for identity
func (m *IdentityConditionalAccessRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityConditionalAccessRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the entry point for the Conditional Access (CA) object model.
func (m *IdentityConditionalAccessRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityConditionalAccessRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable), nil
}
// NamedLocations provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) NamedLocations()(*IdentityConditionalAccessNamedLocationsRequestBuilder) {
    return NewIdentityConditionalAccessNamedLocationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NamedLocationsById provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) NamedLocationsById(id string)(*IdentityConditionalAccessNamedLocationsNamedLocationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["namedLocation%2Did"] = id
    }
    return NewIdentityConditionalAccessNamedLocationsNamedLocationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property conditionalAccess in identity
func (m *IdentityConditionalAccessRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, requestConfiguration *IdentityConditionalAccessRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateConditionalAccessRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ConditionalAccessRootable), nil
}
// Policies provides operations to manage the policies property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) Policies()(*IdentityConditionalAccessPoliciesRequestBuilder) {
    return NewIdentityConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PoliciesById provides operations to manage the policies property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) PoliciesById(id string)(*IdentityConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return NewIdentityConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Templates provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) Templates()(*IdentityConditionalAccessTemplatesRequestBuilder) {
    return NewIdentityConditionalAccessTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.conditionalAccessRoot entity.
func (m *IdentityConditionalAccessRequestBuilder) TemplatesById(id string)(*IdentityConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessTemplate%2Did"] = id
    }
    return NewIdentityConditionalAccessTemplatesConditionalAccessTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
