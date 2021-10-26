package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i446c381dd71cf346c15c96de02dfad0a0edaa1558a94c1cf1b0adbd589d416f6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sections"
    i7ace1b771041569115a2f2d49bd50d00de0110690be34469e578dbb78b5e5894 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/copynotebook"
    ic382ce896bb9041641fdbf1bdd428e7714ecda2914594af39e54bc10ee41a3fa "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups"
    i87e362b314454f0d2f73973c92c38f4592c8089d0ec93e5642c4a65f4d264fc8 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item"
    i9265110ae8278f358dc20c310e9bb579d39674c985988118e40c1a3621919fa7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sections/item"
)

// Builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}\parentSection\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The notebook that contains the section.  Read-only.
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
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}/onenote/pages/{onenotePage_id}/parentSection/parentNotebook{?select,expand}";
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
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*i7ace1b771041569115a2f2d49bd50d00de0110690be34469e578dbb78b5e5894.CopyNotebookRequestBuilder) {
    return i7ace1b771041569115a2f2d49bd50d00de0110690be34469e578dbb78b5e5894.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The notebook that contains the section.  Read-only.
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
// The notebook that contains the section.  Read-only.
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
// The notebook that contains the section.  Read-only.
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
// The notebook that contains the section.  Read-only.
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
// The notebook that contains the section.  Read-only.
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
// The notebook that contains the section.  Read-only.
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*ic382ce896bb9041641fdbf1bdd428e7714ecda2914594af39e54bc10ee41a3fa.SectionGroupsRequestBuilder) {
    return ic382ce896bb9041641fdbf1bdd428e7714ecda2914594af39e54bc10ee41a3fa.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.pages.item.parentSection.parentNotebook.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i87e362b314454f0d2f73973c92c38f4592c8089d0ec93e5642c4a65f4d264fc8.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i87e362b314454f0d2f73973c92c38f4592c8089d0ec93e5642c4a65f4d264fc8.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*i446c381dd71cf346c15c96de02dfad0a0edaa1558a94c1cf1b0adbd589d416f6.SectionsRequestBuilder) {
    return i446c381dd71cf346c15c96de02dfad0a0edaa1558a94c1cf1b0adbd589d416f6.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.pages.item.parentSection.parentNotebook.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*i9265110ae8278f358dc20c310e9bb579d39674c985988118e40c1a3621919fa7.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i9265110ae8278f358dc20c310e9bb579d39674c985988118e40c1a3621919fa7.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
