package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0b84a273ded1e6669385f95709d9a7cf84a87df5fd73dc8c7af3710571de2128 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/unsetverifiedpublisher"
    i1a8360833e23698ad94af084d2df35cce4a1972916936379ed6648bfa6ba57ec "github.com/microsoftgraph/msgraph-sdk-go/applications/item/logo"
    i1afcec2f462ce9653f6d9c178d0d0542f5684412e62323ab5a3f979e1b79b5b8 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/tokenissuancepolicies"
    i1cf4d6bdfc37eea0b8dc06a0fbf17ea8c1b58d8dc24649421cd74e454d266909 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/addkey"
    i402e9b9506b4f160b32ff917d1b2fe43f4d4615296648051dc3eb89140278233 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/removekey"
    i486b52d4dad87ceeac08bd843c5504a58e010e9931346242f546e8106ff87250 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/tokenlifetimepolicies"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5a085ba4baea1613f7766106526258f582e52faf4f832b1ce2655815416aa722 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/removepassword"
    i6f17a12e2de6cae814f546d592ddc360f0666722e94dabfd2392eaa13b36740a "github.com/microsoftgraph/msgraph-sdk-go/applications/item/setverifiedpublisher"
    i72720b8b61b527b526dacbdb4f1b2ebc44378e69b3a31f054e52851fc51aee5a "github.com/microsoftgraph/msgraph-sdk-go/applications/item/checkmembergroups"
    i8644212e8e36b861035674e39cad93b1868aca40e9292c7019cd29ae25503583 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/createdonbehalfof"
    i8f079434e6086cc56be6ed0afce670e9aa0f7ff3ac0107e2f5d9ac8d358fe2dc "github.com/microsoftgraph/msgraph-sdk-go/applications/item/getmembergroups"
    i9565e5e51a86644270ab0a3f267a0bbce14c338625e43342715917d0c6d5eec7 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/owners"
    i9dc2689198e150b8c4d03044f025827c3242c48569acd497be6fec72d51f6797 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/homerealmdiscoverypolicies"
    ia607556161094bfdb9a3407887dfa9b53a9d7a0857a0fd091dab27dd31a9459a "github.com/microsoftgraph/msgraph-sdk-go/applications/item/checkmemberobjects"
    iafba803a7908a654425c76ab1e010310bb6222243e602ddaab50bc32848dac3c "github.com/microsoftgraph/msgraph-sdk-go/applications/item/getmemberobjects"
    ib5b5951df8af6837ed6ea3344f5688d5f2483d29ea1c165d2d219ec5edf941da "github.com/microsoftgraph/msgraph-sdk-go/applications/item/restore"
    id3ff5385c48c1f8ddd4e5f53452e89d8e30d638f1b87dc3265bf1f8adb5d3078 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/addpassword"
    ida306ab63547733b6e18148e7f72a169be3442196d415312a52cbd9978ab6961 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/extensionproperties"
    ie9db136445b8a8e3d0b048cc9910b564635f65b63edab67c97c083c411010b77 "github.com/microsoftgraph/msgraph-sdk-go/applications/item/extensionproperties/item"
)

// Builds and executes requests for operations under \applications\{application-id}
type ApplicationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get entity from applications by key
type ApplicationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *ApplicationRequestBuilder) AddKey()(*i1cf4d6bdfc37eea0b8dc06a0fbf17ea8c1b58d8dc24649421cd74e454d266909.AddKeyRequestBuilder) {
    return i1cf4d6bdfc37eea0b8dc06a0fbf17ea8c1b58d8dc24649421cd74e454d266909.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) AddPassword()(*id3ff5385c48c1f8ddd4e5f53452e89d8e30d638f1b87dc3265bf1f8adb5d3078.AddPasswordRequestBuilder) {
    return id3ff5385c48c1f8ddd4e5f53452e89d8e30d638f1b87dc3265bf1f8adb5d3078.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) CheckMemberGroups()(*i72720b8b61b527b526dacbdb4f1b2ebc44378e69b3a31f054e52851fc51aee5a.CheckMemberGroupsRequestBuilder) {
    return i72720b8b61b527b526dacbdb4f1b2ebc44378e69b3a31f054e52851fc51aee5a.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) CheckMemberObjects()(*ia607556161094bfdb9a3407887dfa9b53a9d7a0857a0fd091dab27dd31a9459a.CheckMemberObjectsRequestBuilder) {
    return ia607556161094bfdb9a3407887dfa9b53a9d7a0857a0fd091dab27dd31a9459a.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ApplicationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationRequestBuilder) {
    m := &ApplicationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/applications/{application_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ApplicationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewApplicationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from applications
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ApplicationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ApplicationRequestBuilder) CreatedOnBehalfOf()(*i8644212e8e36b861035674e39cad93b1868aca40e9292c7019cd29ae25503583.CreatedOnBehalfOfRequestBuilder) {
    return i8644212e8e36b861035674e39cad93b1868aca40e9292c7019cd29ae25503583.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entity from applications by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ApplicationRequestBuilder) CreateGetRequestInformation(q func (value *ApplicationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ApplicationRequestBuilderGetQueryParameters)
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
// Update entity in applications
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ApplicationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Application, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from applications
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ApplicationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ApplicationRequestBuilder) ExtensionProperties()(*ida306ab63547733b6e18148e7f72a169be3442196d415312a52cbd9978ab6961.ExtensionPropertiesRequestBuilder) {
    return ida306ab63547733b6e18148e7f72a169be3442196d415312a52cbd9978ab6961.NewExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.applications.item.extensionProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ApplicationRequestBuilder) ExtensionPropertiesById(id string)(*ie9db136445b8a8e3d0b048cc9910b564635f65b63edab67c97c083c411010b77.ExtensionPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extensionProperty_id"] = id
    }
    return ie9db136445b8a8e3d0b048cc9910b564635f65b63edab67c97c083c411010b77.NewExtensionPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity from applications by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ApplicationRequestBuilder) Get(q func (value *ApplicationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Application, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewApplication() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Application), nil
}
func (m *ApplicationRequestBuilder) GetMemberGroups()(*i8f079434e6086cc56be6ed0afce670e9aa0f7ff3ac0107e2f5d9ac8d358fe2dc.GetMemberGroupsRequestBuilder) {
    return i8f079434e6086cc56be6ed0afce670e9aa0f7ff3ac0107e2f5d9ac8d358fe2dc.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) GetMemberObjects()(*iafba803a7908a654425c76ab1e010310bb6222243e602ddaab50bc32848dac3c.GetMemberObjectsRequestBuilder) {
    return iafba803a7908a654425c76ab1e010310bb6222243e602ddaab50bc32848dac3c.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) HomeRealmDiscoveryPolicies()(*i9dc2689198e150b8c4d03044f025827c3242c48569acd497be6fec72d51f6797.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return i9dc2689198e150b8c4d03044f025827c3242c48569acd497be6fec72d51f6797.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Logo()(*i1a8360833e23698ad94af084d2df35cce4a1972916936379ed6648bfa6ba57ec.LogoRequestBuilder) {
    return i1a8360833e23698ad94af084d2df35cce4a1972916936379ed6648bfa6ba57ec.NewLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Owners()(*i9565e5e51a86644270ab0a3f267a0bbce14c338625e43342715917d0c6d5eec7.OwnersRequestBuilder) {
    return i9565e5e51a86644270ab0a3f267a0bbce14c338625e43342715917d0c6d5eec7.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in applications
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ApplicationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Application, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ApplicationRequestBuilder) RemoveKey()(*i402e9b9506b4f160b32ff917d1b2fe43f4d4615296648051dc3eb89140278233.RemoveKeyRequestBuilder) {
    return i402e9b9506b4f160b32ff917d1b2fe43f4d4615296648051dc3eb89140278233.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) RemovePassword()(*i5a085ba4baea1613f7766106526258f582e52faf4f832b1ce2655815416aa722.RemovePasswordRequestBuilder) {
    return i5a085ba4baea1613f7766106526258f582e52faf4f832b1ce2655815416aa722.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) Restore()(*ib5b5951df8af6837ed6ea3344f5688d5f2483d29ea1c165d2d219ec5edf941da.RestoreRequestBuilder) {
    return ib5b5951df8af6837ed6ea3344f5688d5f2483d29ea1c165d2d219ec5edf941da.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) SetVerifiedPublisher()(*i6f17a12e2de6cae814f546d592ddc360f0666722e94dabfd2392eaa13b36740a.SetVerifiedPublisherRequestBuilder) {
    return i6f17a12e2de6cae814f546d592ddc360f0666722e94dabfd2392eaa13b36740a.NewSetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) TokenIssuancePolicies()(*i1afcec2f462ce9653f6d9c178d0d0542f5684412e62323ab5a3f979e1b79b5b8.TokenIssuancePoliciesRequestBuilder) {
    return i1afcec2f462ce9653f6d9c178d0d0542f5684412e62323ab5a3f979e1b79b5b8.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) TokenLifetimePolicies()(*i486b52d4dad87ceeac08bd843c5504a58e010e9931346242f546e8106ff87250.TokenLifetimePoliciesRequestBuilder) {
    return i486b52d4dad87ceeac08bd843c5504a58e010e9931346242f546e8106ff87250.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ApplicationRequestBuilder) UnsetVerifiedPublisher()(*i0b84a273ded1e6669385f95709d9a7cf84a87df5fd73dc8c7af3710571de2128.UnsetVerifiedPublisherRequestBuilder) {
    return i0b84a273ded1e6669385f95709d9a7cf84a87df5fd73dc8c7af3710571de2128.NewUnsetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
