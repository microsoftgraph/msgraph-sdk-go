package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages"
    i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/identityproviders"
    i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments"
    ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders"
    i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/identityproviders/item"
    i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userattributeassignments/item"
    i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0 "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/userflowidentityproviders/item"
    iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d "github.com/microsoftgraph/msgraph-sdk-go/identity/b2xuserflows/item/languages/item"
)

// B2xIdentityUserFlowItemRequestBuilder provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
type B2xIdentityUserFlowItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// B2xIdentityUserFlowItemRequestBuilderDeleteOptions options for Delete
type B2xIdentityUserFlowItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// B2xIdentityUserFlowItemRequestBuilderGetOptions options for Get
type B2xIdentityUserFlowItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *B2xIdentityUserFlowItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// B2xIdentityUserFlowItemRequestBuilderGetQueryParameters represents entry point for B2X/self-service sign-up identity userflows.
type B2xIdentityUserFlowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// B2xIdentityUserFlowItemRequestBuilderPatchOptions options for Patch
type B2xIdentityUserFlowItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewB2xIdentityUserFlowItemRequestBuilderInternal instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    m := &B2xIdentityUserFlowItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewB2xIdentityUserFlowItemRequestBuilder instantiates a new B2xIdentityUserFlowItemRequestBuilder and sets the default values.
func NewB2xIdentityUserFlowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xIdentityUserFlowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xIdentityUserFlowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateDeleteRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) CreateGetRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) CreatePatchRequestInformation(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property b2xUserFlows for identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Delete(options *B2xIdentityUserFlowItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get represents entry point for B2X/self-service sign-up identity userflows.
func (m *B2xIdentityUserFlowItemRequestBuilder) Get(options *B2xIdentityUserFlowItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateB2xIdentityUserFlowFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.B2xIdentityUserFlowable), nil
}
// IdentityProviders the identityProviders property
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProviders()(*i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.IdentityProvidersRequestBuilder) {
    return i33898ff3b4aa69a4b3cd382359fade5f14f2042fa16a4ceea1285f765e40c1eb.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.identityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) IdentityProvidersById(id string)(*i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d.IdentityProviderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProvider%2Did"] = id
    }
    return i03e59e65e6bc6e940d593a951d05a91effb188992892de8b3543badf8989532d.NewIdentityProviderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Languages the languages property
func (m *B2xIdentityUserFlowItemRequestBuilder) Languages()(*i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.LanguagesRequestBuilder) {
    return i21e37932240ac829addd6426fbe4f398cd6c1253e77d14a3a2d0ed32b9860f0d.NewLanguagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LanguagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.languages.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) LanguagesById(id string)(*iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.UserFlowLanguageConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userFlowLanguageConfiguration%2Did"] = id
    }
    return iffeb8fd6812998e372f8b86fce623d5b4b055347b27a5e10d253ff43d246f33d.NewUserFlowLanguageConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property b2xUserFlows in identity
func (m *B2xIdentityUserFlowItemRequestBuilder) Patch(options *B2xIdentityUserFlowItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserAttributeAssignments the userAttributeAssignments property
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignments()(*i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.UserAttributeAssignmentsRequestBuilder) {
    return i71e405a8d1fa3f6a758b702c04166c77cbb98f401dadb3fa682f3e2ed2f5589c.NewUserAttributeAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserAttributeAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userAttributeAssignments.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserAttributeAssignmentsById(id string)(*i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.IdentityUserFlowAttributeAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttributeAssignment%2Did"] = id
    }
    return i4e7268fbaade322b200bff897ba7a9ec706172d25d1fac40768df711c1624a0f.NewIdentityUserFlowAttributeAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlowIdentityProviders the userFlowIdentityProviders property
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProviders()(*ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.UserFlowIdentityProvidersRequestBuilder) {
    return ie5d8c8cff8f0ae37cced67067dc8d6c60197107961396166a82aa25c36193142.NewUserFlowIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowIdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identity.b2xUserFlows.item.userFlowIdentityProviders.item collection
func (m *B2xIdentityUserFlowItemRequestBuilder) UserFlowIdentityProvidersById(id string)(*i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0.IdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return i8eca55aa158e5bb2109c9b40a378aea7b85dc3cd79c49336bb5abb4e5653f8d0.NewIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
