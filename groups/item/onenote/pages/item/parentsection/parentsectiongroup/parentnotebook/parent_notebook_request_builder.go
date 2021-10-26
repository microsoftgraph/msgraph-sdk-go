package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0ede73f9505946f38f5d2589c381d3ecd310707c5057a5e320034bdff1d08511 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook/copynotebook"
    i89193adf7acd9c990c4d077963b291d3d7ae9d0f6b9a2497e3631443069bd5f0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook/sectiongroups"
    ib867be4e0674053633d60023944d2a0ad98a32e410f34d5f5fb6e20fe5bb5b89 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook/sections"
    i32e3ca3b61578826250877f4a308c755341b80991247bbdebcec1373175ee1d9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook/sections/item"
    i875f4114b3d5b7254aecfed4e6e3ab54e0075af82bb5fab74e4745bc77a2d069 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup/parentnotebook/sectiongroups/item"
)

// Builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentSection\parentSectionGroup\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The notebook that contains the section group. Read-only.
type ParentNotebookRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    m := &ParentNotebookRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/onenote/pages/{onenotePage_id}/parentSection/parentSectionGroup/parentNotebook{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentNotebookRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentNotebookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentNotebookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentNotebookRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*i0ede73f9505946f38f5d2589c381d3ecd310707c5057a5e320034bdff1d08511.CopyNotebookRequestBuilder) {
    return i0ede73f9505946f38f5d2589c381d3ecd310707c5057a5e320034bdff1d08511.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The notebook that contains the section group. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ParentNotebookRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ParentNotebookRequestBuilder) CreateGetRequestInformation(q func (value *ParentNotebookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ParentNotebookRequestBuilderGetQueryParameters)
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ParentNotebookRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentNotebookRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// The notebook that contains the section group. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentNotebookRequestBuilder) Get(q func (value *ParentNotebookRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewNotebook() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook), nil
}
// The notebook that contains the section group. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentNotebookRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Notebook, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*i89193adf7acd9c990c4d077963b291d3d7ae9d0f6b9a2497e3631443069bd5f0.SectionGroupsRequestBuilder) {
    return i89193adf7acd9c990c4d077963b291d3d7ae9d0f6b9a2497e3631443069bd5f0.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentSection.parentSectionGroup.parentNotebook.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i875f4114b3d5b7254aecfed4e6e3ab54e0075af82bb5fab74e4745bc77a2d069.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i875f4114b3d5b7254aecfed4e6e3ab54e0075af82bb5fab74e4745bc77a2d069.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*ib867be4e0674053633d60023944d2a0ad98a32e410f34d5f5fb6e20fe5bb5b89.SectionsRequestBuilder) {
    return ib867be4e0674053633d60023944d2a0ad98a32e410f34d5f5fb6e20fe5bb5b89.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentSection.parentSectionGroup.parentNotebook.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*i32e3ca3b61578826250877f4a308c755341b80991247bbdebcec1373175ee1d9.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i32e3ca3b61578826250877f4a308c755341b80991247bbdebcec1373175ee1d9.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
