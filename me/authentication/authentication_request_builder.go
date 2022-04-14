package authentication

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/fido2methods"
    i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods"
    i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods"
    iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods"
    i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods/item"
    i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item"
    id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/fido2methods/item"
    ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item"
)

// AuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationRequestBuilderDeleteOptions options for Delete
type AuthenticationRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AuthenticationRequestBuilderGetOptions options for Get
type AuthenticationRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationRequestBuilderPatchOptions options for Patch
type AuthenticationRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(options *AuthenticationRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(options *AuthenticationRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) Delete(options *AuthenticationRequestBuilderDeleteOptions)(error) {
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
// Fido2Methods the fido2Methods property
func (m *AuthenticationRequestBuilder) Fido2Methods()(*i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521.Fido2MethodsRequestBuilder) {
    return i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.fido2Methods.item collection
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100.Fido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100.NewFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) Get(options *AuthenticationRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable), nil
}
// Methods the methods property
func (m *AuthenticationRequestBuilder) Methods()(*i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f.MethodsRequestBuilder) {
    return i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.methods.item collection
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5.AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5.NewAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftAuthenticatorMethods the microsoftAuthenticatorMethods property
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.microsoftAuthenticatorMethods.item collection
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4.MicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4.NewMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) Patch(options *AuthenticationRequestBuilderPatchOptions)(error) {
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
// WindowsHelloForBusinessMethods the windowsHelloForBusinessMethods property
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c.WindowsHelloForBusinessMethodsRequestBuilder) {
    return iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.windowsHelloForBusinessMethods.item collection
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462.WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462.NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
