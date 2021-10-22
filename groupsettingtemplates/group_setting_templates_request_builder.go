package groupsettingtemplates

import (
    i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/validateproperties"
    i876aea553bfa1c412214415b5c2d08dc7ae79cdcecf4bf4e6f912c57c0de078d "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getavailableextensionproperties"
    iccb6d65dafe6f1cf697180a62c7210b026edfb6f165ab42626bb6c7cf8a7836a "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/getbyids"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type GroupSettingTemplatesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type GroupSettingTemplatesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewGroupSettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    m := &GroupSettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groupSettingTemplates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewGroupSettingTemplatesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupSettingTemplatesRequestBuilder) CreateGetRequestInformation(q func (value *GroupSettingTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GroupSettingTemplatesRequestBuilderGetQueryParameters)
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
func (m *GroupSettingTemplatesRequestBuilder) CreatePostRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupSettingTemplatesRequestBuilder) Get(q func (value *GroupSettingTemplatesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*GroupSettingTemplatesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupSettingTemplatesResponse() }, responseHandler)
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
func (m *GroupSettingTemplatesRequestBuilder) Post(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, error) {
    requestInfo, err := m.CreatePostRequestInformation(body, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewGroupSettingTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate), nil
}
func (m *GroupSettingTemplatesRequestBuilder) ValidateProperties()(*i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.ValidatePropertiesRequestBuilder) {
    return i5c2514cac71bd42e8d3ff0b5f2392805dbb5175ce8500f2f4066069fd3f8a025.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
