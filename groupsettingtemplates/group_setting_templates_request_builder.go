package groupsettingtemplates

import (
    i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/validateproperties"
    i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getavailableextensionproperties"
    iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/count"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
)

// GroupSettingTemplatesRequestBuilder provides operations to manage the collection of groupSettingTemplate entities.
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
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplateable;
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupSettingTemplatesRequestBuilder instantiates a new GroupSettingTemplatesRequestBuilder and sets the default values.
func NewGroupSettingTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupSettingTemplatesRequestBuilder) Count()(*ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809.CountRequestBuilder) {
    return ie2feb4ec9a72c319a82ed8b82d8ae3403e819512010e7c6dd097f3426e334809.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *GroupSettingTemplatesRequestBuilder) Get(options *GroupSettingTemplatesRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateGroupSettingTemplateCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplateCollectionResponseable), nil
}
func (m *GroupSettingTemplatesRequestBuilder) GetAvailableExtensionProperties()(*i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.GetAvailableExtensionPropertiesRequestBuilder) {
    return i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d.NewGetAvailableExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupSettingTemplatesRequestBuilder) GetByIds()(*iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.GetByIdsRequestBuilder) {
    return iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to groupSettingTemplates
func (m *GroupSettingTemplatesRequestBuilder) Post(options *GroupSettingTemplatesRequestBuilderPostOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateGroupSettingTemplateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplateable), nil
}
func (m *GroupSettingTemplatesRequestBuilder) ValidateProperties()(*i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.ValidatePropertiesRequestBuilder) {
    return i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
