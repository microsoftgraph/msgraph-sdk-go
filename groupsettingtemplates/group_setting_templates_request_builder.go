package groupsettingtemplates

import (
    i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/validateproperties"
    i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getavailableextensionproperties"
    iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GroupSettingTemplatesRequestBuilder builds and executes requests for operations under \groupSettingTemplates
type GroupSettingTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupSettingTemplatesRequestBuilderGetOptions options for Get
type GroupSettingTemplatesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupSettingTemplatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupSettingTemplatesRequestBuilderGetQueryParameters get entities from groupSettingTemplates
type GroupSettingTemplatesRequestBuilderGetQueryParameters struct {
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// GroupSettingTemplatesRequestBuilderPostOptions options for Post
type GroupSettingTemplatesRequestBuilderPostOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupSettingTemplatesRequestBuilderInternal instantiates a new GroupSettingTemplatesRequestBuilder and sets the default values.
func NewGroupSettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    m := &GroupSettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groupSettingTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupSettingTemplatesRequestBuilder instantiates a new GroupSettingTemplatesRequestBuilder and sets the default values.
func NewGroupSettingTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get entities from groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) CreateGetRequestInformation(options *GroupSettingTemplatesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePostRequestInformation add new entity to groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) CreatePostRequestInformation(options *GroupSettingTemplatesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get get entities from groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) Get(options *GroupSettingTemplatesRequestBuilderGetOptions)(*GroupSettingTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupSettingTemplatesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GroupSettingTemplatesResponse), nil
}
func (m *GroupSettingTemplatesRequestBuilder) GetAvailableExtensionProperties()(*i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.GetAvailableExtensionPropertiesRequestBuilder) {
    return i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupSettingTemplatesRequestBuilder) GetByIds()(*iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.GetByIdsRequestBuilder) {
    return iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) Post(options *GroupSettingTemplatesRequestBuilderPostOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewGroupSettingTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate), nil
}
func (m *GroupSettingTemplatesRequestBuilder) ValidateProperties()(*i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.ValidatePropertiesRequestBuilder) {
    return i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
