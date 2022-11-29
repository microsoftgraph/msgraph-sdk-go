package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationRequestBuilder provides operations to manage the authentication property of the microsoft.graph.user entity.
type UsersItemAuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemAuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type UsersItemAuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemAuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemAuthenticationRequestBuilderGetQueryParameters
}
// UsersItemAuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAuthenticationRequestBuilderInternal instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewUsersItemAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationRequestBuilder) {
    m := &UsersItemAuthenticationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationRequestBuilder instantiates a new AuthenticationRequestBuilder and sets the default values.
func NewUsersItemAuthenticationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property authentication for users
func (m *UsersItemAuthenticationRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the authentication methods that are supported for the user.
func (m *UsersItemAuthenticationRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property authentication in users
func (m *UsersItemAuthenticationRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *UsersItemAuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property authentication for users
func (m *UsersItemAuthenticationRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemAuthenticationRequestBuilderDeleteRequestConfiguration)(error) {
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
// EmailMethods provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) EmailMethods()(*UsersItemAuthenticationEmailMethodsRequestBuilder) {
    return NewUsersItemAuthenticationEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById provides operations to manage the emailMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) EmailMethodsById(id string)(*UsersItemAuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationEmailMethodsEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Fido2Methods provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) Fido2Methods()(*UsersItemAuthenticationFido2MethodsRequestBuilder) {
    return NewUsersItemAuthenticationFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fido2MethodsById provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) Fido2MethodsById(id string)(*UsersItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationFido2MethodsFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the authentication methods that are supported for the user.
func (m *UsersItemAuthenticationRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAuthenticationRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable), nil
}
// Methods provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) Methods()(*UsersItemAuthenticationMethodsRequestBuilder) {
    return NewUsersItemAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MethodsById provides operations to manage the methods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) MethodsById(id string)(*UsersItemAuthenticationMethodsAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationMethodsAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftAuthenticatorMethods provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsRequestBuilder) {
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftAuthenticatorMethodsById provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) Operations()(*UsersItemAuthenticationOperationsRequestBuilder) {
    return NewUsersItemAuthenticationOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) OperationsById(id string)(*UsersItemAuthenticationOperationsLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return NewUsersItemAuthenticationOperationsLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordMethods provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) PasswordMethods()(*UsersItemAuthenticationPasswordMethodsRequestBuilder) {
    return NewUsersItemAuthenticationPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById provides operations to manage the passwordMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) PasswordMethodsById(id string)(*UsersItemAuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationPasswordMethodsPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in users
func (m *UsersItemAuthenticationRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *UsersItemAuthenticationRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable), nil
}
// PhoneMethods provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) PhoneMethods()(*UsersItemAuthenticationPhoneMethodsRequestBuilder) {
    return NewUsersItemAuthenticationPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById provides operations to manage the phoneMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) PhoneMethodsById(id string)(*UsersItemAuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationPhoneMethodsPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareOathMethods provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) SoftwareOathMethods()(*UsersItemAuthenticationSoftwareOathMethodsRequestBuilder) {
    return NewUsersItemAuthenticationSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById provides operations to manage the softwareOathMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*UsersItemAuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationSoftwareOathMethodsSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemporaryAccessPassMethods provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) TemporaryAccessPassMethods()(*UsersItemAuthenticationTemporaryAccessPassMethodsRequestBuilder) {
    return NewUsersItemAuthenticationTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById provides operations to manage the temporaryAccessPassMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*UsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationTemporaryAccessPassMethodsTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsHelloForBusinessMethods provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*UsersItemAuthenticationWindowsHelloForBusinessMethodsRequestBuilder) {
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsHelloForBusinessMethodsById provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
func (m *UsersItemAuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*UsersItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod%2Did"] = id
    }
    return NewUsersItemAuthenticationWindowsHelloForBusinessMethodsWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
