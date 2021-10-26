package authentication

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i17beca6845c8590a7993cf25977350ea5aff5795257953969bade065fb0a8c8e "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/fido2methods"
    i36d51fdd7beea498cf790dc7e6370ee500c1ae341f5d8a4c36d081226f91ee96 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/methods"
    i3d41e7ad8d690559b414262fba543314187fe0d02cf0420f778314c36a9270af "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods"
    i96e8d458309e5631e00b3cbc4d031fbc1d45f1137d2f6149c3e103603bd3dcf2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods"
    i33902446e72a35eb92a8d90a86021f50ebf6a7badd666e5dd29d39829615e81f "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/microsoftauthenticatormethods/item"
    i4f5c6f02fac27458fe06277174a65fe094c80b096c7dceec66f3bc6ab61be13b "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/fido2methods/item"
    i9204deeb9308a81051ecfbd988907f606e6330dbfc4b953a713d2803d2d5735a "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/windowshelloforbusinessmethods/item"
    ia6b9ba2f6ac4b661e1925c06219918528114879cf86ac1234330ab29aa5c8f58 "github.com/microsoftgraph/msgraph-sdk-go/users/item/authentication/methods/item"
)

// Builds and executes requests for operations under \users\{user-id}\authentication
type AuthenticationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get authentication from users
type AuthenticationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new AuthenticationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewAuthenticationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AuthenticationRequestBuilder) {
    m := &AuthenticationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/authentication{?select,expand}";
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
// Delete navigation property authentication for users
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *AuthenticationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get authentication from users
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *AuthenticationRequestBuilder) CreateGetRequestInformation(q func (value *AuthenticationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AuthenticationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update the navigation property authentication in users
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *AuthenticationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete navigation property authentication for users
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *AuthenticationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) Fido2Methods()(*i17beca6845c8590a7993cf25977350ea5aff5795257953969bade065fb0a8c8e.Fido2MethodsRequestBuilder) {
    return i17beca6845c8590a7993cf25977350ea5aff5795257953969bade065fb0a8c8e.NewFido2MethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.fido2Methods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) Fido2MethodsById(id string)(*i4f5c6f02fac27458fe06277174a65fe094c80b096c7dceec66f3bc6ab61be13b.Fido2AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fido2AuthenticationMethod_id"] = id
    }
    return i4f5c6f02fac27458fe06277174a65fe094c80b096c7dceec66f3bc6ab61be13b.NewFido2AuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get authentication from users
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *AuthenticationRequestBuilder) Get(q func (value *AuthenticationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAuthentication() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication), nil
}
func (m *AuthenticationRequestBuilder) Methods()(*i36d51fdd7beea498cf790dc7e6370ee500c1ae341f5d8a4c36d081226f91ee96.MethodsRequestBuilder) {
    return i36d51fdd7beea498cf790dc7e6370ee500c1ae341f5d8a4c36d081226f91ee96.NewMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.methods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) MethodsById(id string)(*ia6b9ba2f6ac4b661e1925c06219918528114879cf86ac1234330ab29aa5c8f58.AuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationMethod_id"] = id
    }
    return ia6b9ba2f6ac4b661e1925c06219918528114879cf86ac1234330ab29aa5c8f58.NewAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethods()(*i96e8d458309e5631e00b3cbc4d031fbc1d45f1137d2f6149c3e103603bd3dcf2.MicrosoftAuthenticatorMethodsRequestBuilder) {
    return i96e8d458309e5631e00b3cbc4d031fbc1d45f1137d2f6149c3e103603bd3dcf2.NewMicrosoftAuthenticatorMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.microsoftAuthenticatorMethods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) MicrosoftAuthenticatorMethodsById(id string)(*i33902446e72a35eb92a8d90a86021f50ebf6a7badd666e5dd29d39829615e81f.MicrosoftAuthenticatorAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftAuthenticatorAuthenticationMethod_id"] = id
    }
    return i33902446e72a35eb92a8d90a86021f50ebf6a7badd666e5dd29d39829615e81f.NewMicrosoftAuthenticatorAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update the navigation property authentication in users
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *AuthenticationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Authentication, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethods()(*i3d41e7ad8d690559b414262fba543314187fe0d02cf0420f778314c36a9270af.WindowsHelloForBusinessMethodsRequestBuilder) {
    return i3d41e7ad8d690559b414262fba543314187fe0d02cf0420f778314c36a9270af.NewWindowsHelloForBusinessMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.authentication.windowsHelloForBusinessMethods.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *AuthenticationRequestBuilder) WindowsHelloForBusinessMethodsById(id string)(*i9204deeb9308a81051ecfbd988907f606e6330dbfc4b953a713d2803d2d5735a.WindowsHelloForBusinessAuthenticationMethodRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsHelloForBusinessAuthenticationMethod_id"] = id
    }
    return i9204deeb9308a81051ecfbd988907f606e6330dbfc4b953a713d2803d2d5735a.NewWindowsHelloForBusinessAuthenticationMethodRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
