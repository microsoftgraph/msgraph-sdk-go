package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/checkmemberobjects"
    ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/getmemberobjects"
    ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/checkmembergroups"
    id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/getmembergroups"
    if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0 "github.com/microsoftgraph/msgraph-sdk-go/groupsettingtemplates/item/restore"
)

// Builds and executes requests for operations under \groupSettingTemplates\{groupSettingTemplate-id}
type GroupSettingTemplateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get entity from groupSettingTemplates by key
type GroupSettingTemplateRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *GroupSettingTemplateRequestBuilder) CheckMemberGroups()(*ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b.CheckMemberGroupsRequestBuilder) {
    return ic9284ba20a018bdf9cc611bd4cc87612370fd87e308417912c51e24a09a84b7b.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupSettingTemplateRequestBuilder) CheckMemberObjects()(*i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c.CheckMemberObjectsRequestBuilder) {
    return i535340b77c27094ee236bb42cd6d6ecb0741e1f8ed70e47b43e92472c1f9535c.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new GroupSettingTemplateRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupSettingTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplateRequestBuilder) {
    m := &GroupSettingTemplateRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groupSettingTemplates/{groupSettingTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GroupSettingTemplateRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupSettingTemplateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupSettingTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupSettingTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from groupSettingTemplates
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *GroupSettingTemplateRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from groupSettingTemplates by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *GroupSettingTemplateRequestBuilder) CreateGetRequestInformation(q func (value *GroupSettingTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(GroupSettingTemplateRequestBuilderGetQueryParameters)
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
// Update entity in groupSettingTemplates
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *GroupSettingTemplateRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from groupSettingTemplates
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *GroupSettingTemplateRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// Get entity from groupSettingTemplates by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *GroupSettingTemplateRequestBuilder) Get(q func (value *GroupSettingTemplateRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewGroupSettingTemplate() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate), nil
}
func (m *GroupSettingTemplateRequestBuilder) GetMemberGroups()(*id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c.GetMemberGroupsRequestBuilder) {
    return id343ec0d598ad136300b8599ebf6da41179c7224734d67f8c9ca6914ab993a9c.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupSettingTemplateRequestBuilder) GetMemberObjects()(*ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb.GetMemberObjectsRequestBuilder) {
    return ib24705b5e21850aefb572da669580b25cff74a7513ba53674efeff14fe51e5cb.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in groupSettingTemplates
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *GroupSettingTemplateRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.GroupSettingTemplate, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *GroupSettingTemplateRequestBuilder) Restore()(*if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0.RestoreRequestBuilder) {
    return if0de78395fe1dccb1dbdb7b7fe16ca5dd70edd86fe4f481d0b82bca471e231f0.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
