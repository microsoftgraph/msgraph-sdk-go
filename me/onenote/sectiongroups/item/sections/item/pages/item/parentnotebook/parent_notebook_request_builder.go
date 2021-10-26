package parentnotebook

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5a148cf780b9696d2da89b2611271ef227665919a495d7b89062d385ed2f01b6 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sections"
    i757f7a7a797ea74e0655b0c5342544d98f6f18ad814630ccb59c88e73177bbd1 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sectiongroups"
    if0a369d7e8a42d36552ea1a4d01b4220fc7dc83015647fa7345459768f0b8a47 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/copynotebook"
    i193e3e470fba058449334c1f2c2f28da77363587c5eda69b89e9df0f4b7a11e6 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sections/item"
    i8d2044057c9dda25059ac72ea6f39bff415ee84f37369af897807272c7334a51 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/sections/item/pages/item/parentnotebook/sectiongroups/item"
)

// Builds and executes requests for operations under \me\onenote\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}\parentNotebook
type ParentNotebookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The notebook that contains the page.  Read-only.
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
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}/parentNotebook{?select,expand}";
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
func (m *ParentNotebookRequestBuilder) CopyNotebook()(*if0a369d7e8a42d36552ea1a4d01b4220fc7dc83015647fa7345459768f0b8a47.CopyNotebookRequestBuilder) {
    return if0a369d7e8a42d36552ea1a4d01b4220fc7dc83015647fa7345459768f0b8a47.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The notebook that contains the page.  Read-only.
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
// The notebook that contains the page.  Read-only.
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
// The notebook that contains the page.  Read-only.
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
// The notebook that contains the page.  Read-only.
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
// The notebook that contains the page.  Read-only.
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
// The notebook that contains the page.  Read-only.
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
func (m *ParentNotebookRequestBuilder) SectionGroups()(*i757f7a7a797ea74e0655b0c5342544d98f6f18ad814630ccb59c88e73177bbd1.SectionGroupsRequestBuilder) {
    return i757f7a7a797ea74e0655b0c5342544d98f6f18ad814630ccb59c88e73177bbd1.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sectionGroups.item.sections.item.pages.item.parentNotebook.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionGroupsById(id string)(*i8d2044057c9dda25059ac72ea6f39bff415ee84f37369af897807272c7334a51.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return i8d2044057c9dda25059ac72ea6f39bff415ee84f37369af897807272c7334a51.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentNotebookRequestBuilder) Sections()(*i5a148cf780b9696d2da89b2611271ef227665919a495d7b89062d385ed2f01b6.SectionsRequestBuilder) {
    return i5a148cf780b9696d2da89b2611271ef227665919a495d7b89062d385ed2f01b6.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.onenote.sectionGroups.item.sections.item.pages.item.parentNotebook.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentNotebookRequestBuilder) SectionsById(id string)(*i193e3e470fba058449334c1f2c2f28da77363587c5eda69b89e9df0f4b7a11e6.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id1"] = id
    }
    return i193e3e470fba058449334c1f2c2f28da77363587c5eda69b89e9df0f4b7a11e6.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
