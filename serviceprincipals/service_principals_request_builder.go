package serviceprincipals

import (
    i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/delta"
    i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/validateproperties"
    i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/getavailableextensionproperties"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5 "github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals/getbyids"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \servicePrincipals
type ServicePrincipalsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get entities from servicePrincipals
type ServicePrincipalsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    m := &ServicePrincipalsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/servicePrincipals{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ServicePrincipalsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewServicePrincipalsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ServicePrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get entities from servicePrincipals
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ServicePrincipalsRequestBuilder) CreateGetRequestInformation(q func (value *ServicePrincipalsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ServicePrincipalsRequestBuilderGetQueryParameters)
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
// Add new entity to servicePrincipals
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ServicePrincipalsRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Builds and executes requests for operations under \servicePrincipals\microsoft.graph.delta()
func (m *ServicePrincipalsRequestBuilder) Delta()(*i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080.DeltaRequestBuilder) {
    return i3f07418133571e1113338a9bf728861821f52ca4198d88567184067eda8c8080.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entities from servicePrincipals
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ServicePrincipalsRequestBuilder) Get(q func (value *ServicePrincipalsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ServicePrincipalsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServicePrincipalsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ServicePrincipalsResponse), nil
}
func (m *ServicePrincipalsRequestBuilder) GetAvailableExtensionProperties()(*i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc.GetAvailableExtensionPropertiesRequestBuilder) {
    return i8febfd579535614c02c9f7bb12392954732efb7d66b1a1f24cb4f5cb774ad2dc.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ServicePrincipalsRequestBuilder) GetByIds()(*idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5.GetByIdsRequestBuilder) {
    return idebb258791ee268c51b14dfc22d19aef5e0100389ec9771b23c53e2010b284f5.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Add new entity to servicePrincipals
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ServicePrincipalsRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewServicePrincipal() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ServicePrincipal), nil
}
func (m *ServicePrincipalsRequestBuilder) ValidateProperties()(*i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd.ValidatePropertiesRequestBuilder) {
    return i78b819e66260c21f46560645db29002f1711583967d1a6a2a6cd9e6788fc77fd.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
