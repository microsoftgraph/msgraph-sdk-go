package authentication

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/fido2methods"
    i1ab105702cd983545027eaccaef83b254c7f05b9cd0320b0430186c1b094a8e1 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/operations"
    i2b453d28c632a5df5e8f1feb8a4b79237d455166afc5b2e6594a52fb03eb040f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/softwareoathmethods"
    i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods"
    i330cc629c4c64eea37d63603b23fd6b7dcdf631ccb9226c77b586cc7eecdd592 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/passwordmethods"
    i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods"
    i578762c28c58199bca46c878eb92de065ccb247f40835d229a0d27724421dedf "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/emailmethods"
    i94e0b23423489cfe0cdbb5bec1cba5f3d30c4d9553bb4a5215cda18153ebb34f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/temporaryaccesspassmethods"
    iab87de02a0d283f2ebc8cb67555b1fffded0f7a536efdd631c7e9771bb529d80 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/phonemethods"
    iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods"
    i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods/item"
    i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item"
    i55878e36b82d9cb48f3a125722dddd8f7ae172dca480280800e318f503e34c61 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/phonemethods/item"
    i5a9396125be9068dc078820cc6d4814598cf64b1fc2bd94d11163a1f0a5e77bd "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/emailmethods/item"
    i6624e87a0ce0621594236e05e9a999aff0e84a7bfa74104ca407022a6083c85f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/temporaryaccesspassmethods/item"
    i7bbaf3ee5cd83c376758a2fc0e8c2a721bda7f7913e9f3ef81038cc65ddfd5c1 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/operations/item"
    i97a1e8335ed4a95f235bcf60a57963d4f7017e4e3143435514f0fb92bca30352 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/passwordmethods/item"
    ic0966f3e95c57647ba075867a3d39ee0dbf2dfbb4f2db8141766804e838fc11a "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/softwareoathmethods/item"
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
// AuthenticationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationRequestBuilderGetQueryParameters the authentication methods that are supported for the user.
type AuthenticationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationRequestBuilderGetQueryParameters
}
// AuthenticationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property authentication for me
func (m *AuthenticationRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *AuthenticationRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EmailMethods the emailMethods property
func (m *AuthenticationRequestBuilder) EmailMethods()(*i578762c28c58199bca46c878eb92de065ccb247f40835d229a0d27724421dedf.EmailMethodsRequestBuilder) {
    return i578762c28c58199bca46c878eb92de065ccb247f40835d229a0d27724421dedf.NewEmailMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmailMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.emailMethods.item collection
func (m *AuthenticationRequestBuilder) EmailMethodsById(id string)(*i5a9396125be9068dc078820cc6d4814598cf64b1fc2bd94d11163a1f0a5e77bd.EmailAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["emailAuthenticationMethod%2Did"] = id
    }
    return i5a9396125be9068dc078820cc6d4814598cf64b1fc2bd94d11163a1f0a5e77bd.NewEmailAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *AuthenticationRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the authentication methods that are supported for the user.
func (m *AuthenticationRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AuthenticationRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationFromDiscriminatorValue, responseHandler, errorMapping)
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
// Operations the operations property
func (m *AuthenticationRequestBuilder) Operations()(*i1ab105702cd983545027eaccaef83b254c7f05b9cd0320b0430186c1b094a8e1.OperationsRequestBuilder) {
    return i1ab105702cd983545027eaccaef83b254c7f05b9cd0320b0430186c1b094a8e1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.operations.item collection
func (m *AuthenticationRequestBuilder) OperationsById(id string)(*i7bbaf3ee5cd83c376758a2fc0e8c2a721bda7f7913e9f3ef81038cc65ddfd5c1.LongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return i7bbaf3ee5cd83c376758a2fc0e8c2a721bda7f7913e9f3ef81038cc65ddfd5c1.NewLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PasswordMethods the passwordMethods property
func (m *AuthenticationRequestBuilder) PasswordMethods()(*i330cc629c4c64eea37d63603b23fd6b7dcdf631ccb9226c77b586cc7eecdd592.PasswordMethodsRequestBuilder) {
    return i330cc629c4c64eea37d63603b23fd6b7dcdf631ccb9226c77b586cc7eecdd592.NewPasswordMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PasswordMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.passwordMethods.item collection
func (m *AuthenticationRequestBuilder) PasswordMethodsById(id string)(*i97a1e8335ed4a95f235bcf60a57963d4f7017e4e3143435514f0fb92bca30352.PasswordAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passwordAuthenticationMethod%2Did"] = id
    }
    return i97a1e8335ed4a95f235bcf60a57963d4f7017e4e3143435514f0fb92bca30352.NewPasswordAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property authentication in me
func (m *AuthenticationRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Authenticationable, requestConfiguration *AuthenticationRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PhoneMethods the phoneMethods property
func (m *AuthenticationRequestBuilder) PhoneMethods()(*iab87de02a0d283f2ebc8cb67555b1fffded0f7a536efdd631c7e9771bb529d80.PhoneMethodsRequestBuilder) {
    return iab87de02a0d283f2ebc8cb67555b1fffded0f7a536efdd631c7e9771bb529d80.NewPhoneMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhoneMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.phoneMethods.item collection
func (m *AuthenticationRequestBuilder) PhoneMethodsById(id string)(*i55878e36b82d9cb48f3a125722dddd8f7ae172dca480280800e318f503e34c61.PhoneAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["phoneAuthenticationMethod%2Did"] = id
    }
    return i55878e36b82d9cb48f3a125722dddd8f7ae172dca480280800e318f503e34c61.NewPhoneAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareOathMethods the softwareOathMethods property
func (m *AuthenticationRequestBuilder) SoftwareOathMethods()(*i2b453d28c632a5df5e8f1feb8a4b79237d455166afc5b2e6594a52fb03eb040f.SoftwareOathMethodsRequestBuilder) {
    return i2b453d28c632a5df5e8f1feb8a4b79237d455166afc5b2e6594a52fb03eb040f.NewSoftwareOathMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SoftwareOathMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.softwareOathMethods.item collection
func (m *AuthenticationRequestBuilder) SoftwareOathMethodsById(id string)(*ic0966f3e95c57647ba075867a3d39ee0dbf2dfbb4f2db8141766804e838fc11a.SoftwareOathAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["softwareOathAuthenticationMethod%2Did"] = id
    }
    return ic0966f3e95c57647ba075867a3d39ee0dbf2dfbb4f2db8141766804e838fc11a.NewSoftwareOathAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemporaryAccessPassMethods the temporaryAccessPassMethods property
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethods()(*i94e0b23423489cfe0cdbb5bec1cba5f3d30c4d9553bb4a5215cda18153ebb34f.TemporaryAccessPassMethodsRequestBuilder) {
    return i94e0b23423489cfe0cdbb5bec1cba5f3d30c4d9553bb4a5215cda18153ebb34f.NewTemporaryAccessPassMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemporaryAccessPassMethodsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.authentication.temporaryAccessPassMethods.item collection
func (m *AuthenticationRequestBuilder) TemporaryAccessPassMethodsById(id string)(*i6624e87a0ce0621594236e05e9a999aff0e84a7bfa74104ca407022a6083c85f.TemporaryAccessPassAuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["temporaryAccessPassAuthenticationMethod%2Did"] = id
    }
    return i6624e87a0ce0621594236e05e9a999aff0e84a7bfa74104ca407022a6083c85f.NewTemporaryAccessPassAuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
