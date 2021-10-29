package authentication

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/fido2methods"
    i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods"
    iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods"
    i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/methods/item"
    i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/windowshelloforbusinessmethods/item"
    id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/fido2methods/item"
    ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4 "github.com/microsoftgraph/msgraph-sdk-go/me/authentication/microsoftauthenticatormethods/item"
)

// Builds and executes requests for operations under \me\authentication
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type AuthenticationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type AuthenticationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AuthenticationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get authentication from me
type AuthenticationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type AuthenticationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new AuthenticationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/authentication{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new AuthenticationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property authentication for me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(options *AuthenticationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get authentication from me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(options *AuthenticationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update the navigation property authentication in me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(options *AuthenticationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete navigation property authentication for me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) Delete(options *AuthenticationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) Fido2Methods()(*i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521.Fido2MethodsRequestBuilder) {
    return i085509635ea64d5fbb3a67c6d642d98b7f3c4dde9aa9e524741fac6bc34ea521.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.authentication.fido2Methods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100.Fido2AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod_id"] = id
    }
    return id52d22a8f89e6b3eaca8d0a4bf66b74846ed7c246fc3b1528ed62713ddb3c100.NewFido2AuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get authentication from me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) Get(options *AuthenticationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAuthentication() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication), nil
}
func (m *AuthenticationRequestBuilder) Methods()(*i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f.MethodsRequestBuilder) {
    return i5124ca9d3dbcd8647924646fbf2e408a497f546a270610b963a0c4645202098f.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.authentication.methods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5.AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod_id"] = id
    }
    return i19c5c40e0312d049b21b8875eb90bff2ae1d634fbec124fef57e20ddd2da47f5.NewAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return i2ce1e5bbe67f22b89b9a62b1ec82ed9532485b3c32368fa2f160cb67999943d3.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.authentication.microsoftAuthenticatorMethods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4.MicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return ieef8280a3bcc4019ab7c2f7f81ff70f8682ed1a5643972b20c7af236b531eeb4.NewMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property authentication in me
// Parameters:
//  - options : Options for the request
func (m *AuthenticationRequestBuilder) Patch(options *AuthenticationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c.WindowsHelloForBusinessMethodsRequestBuilder) {
    return iddc094cb0fe07ad41fd27a0a8204433d0ccb63cc6098ca73c2f608b1b8caef8c.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.me.authentication.windowsHelloForBusinessMethods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462.WindowsHelloForBusinessAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod_id"] = id
    }
    return i27d2ecc6d9c0761197a1e664d1ac58196849d406d8afd8a246158ca273dcb462.NewWindowsHelloForBusinessAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
