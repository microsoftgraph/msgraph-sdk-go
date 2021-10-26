package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sectiongroups"
    ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections"
    iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/parentnotebook"
    ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/parentsectiongroup"
    i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item"
    i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentnotebook/sectiongroups/item/sectiongroups/item"
)

// Builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentNotebook\sectionGroups\{sectionGroup-id}
type SectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The section groups in the notebook. Read-only. Nullable.
type SectionGroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    m := &SectionGroupRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/onenote/pages/{onenotePage_id}/parentNotebook/sectionGroups/{sectionGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *SectionGroupRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *SectionGroupRequestBuilder) CreateGetRequestInformation(q func (value *SectionGroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SectionGroupRequestBuilderGetQueryParameters)
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *SectionGroupRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *SectionGroupRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *SectionGroupRequestBuilder) Get(q func (value *SectionGroupRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSectionGroup() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup), nil
}
func (m *SectionGroupRequestBuilder) ParentNotebook()(*iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114.ParentNotebookRequestBuilder) {
    return iabd06dfa023f86494ccd0adf4612084319849eebeafc8626f13f6e8e3a901114.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a.ParentSectionGroupRequestBuilder) {
    return ie54d3ae495ee37efadd6e2191a33ee0ce5b6fac69255f0ba77b835f821831d1a.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *SectionGroupRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SectionGroup, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SectionGroupRequestBuilder) SectionGroups()(*i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e.SectionGroupsRequestBuilder) {
    return i2d4405d61073097820d21feaeb8cfb4b00b8b97ca196c8d61adaaa139a98553e.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentNotebook.sectionGroups.item.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i29f4f1b45400549b046daaacd804ee1174916b917848e3c7a941a72a9026979d.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615.SectionsRequestBuilder) {
    return ia63f4929dba5b38acb07057cefc6c982492d9f0f8bd23fbf594b9af33e358615.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentNotebook.sectionGroups.item.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i1d1d7ffa0230017a6a201b1784298e9991cc81e8f030794f7db0f9755ec106ef.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
