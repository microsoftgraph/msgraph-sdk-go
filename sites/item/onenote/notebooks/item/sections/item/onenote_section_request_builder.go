package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0788ee8cc0f88534473ad867f6033f79c94d3ced1d865def09e887a4af6505f0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/copytosectiongroup"
    i090eb63c39e2c392547c70c8250c04ea9dc1ea28fd27551825475fe51dedb728 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/pages"
    i44235afe4d281e29e2cd6b9abe12f703ec1959f43cd3be6ea227b53da7492dd7 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/copytonotebook"
    i7fda827a8abe7ff3582c46ee5993acec0f2c7f7fe245e0e35e98367e8db423c5 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/parentnotebook"
    i84bdfed9ec30dc209a4c530b508d815c14183bd652a5999a2f6f72613165ac5b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/parentsectiongroup"
    icd8f355cf1b63abcb90f8942bfacbc3a61d12900f6fa5545795d228102f51471 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sections/item/pages/item"
)

// Builds and executes requests for operations under \sites\{site-id}\onenote\notebooks\{notebook-id}\sections\{onenoteSection-id}
type OnenoteSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The sections in the notebook. Read-only. Nullable.
type OnenoteSectionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new OnenoteSectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    m := &OnenoteSectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/sites/{site_id}/onenote/notebooks/{notebook_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenoteSectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionRequestBuilder) CopyToNotebook()(*i44235afe4d281e29e2cd6b9abe12f703ec1959f43cd3be6ea227b53da7492dd7.CopyToNotebookRequestBuilder) {
    return i44235afe4d281e29e2cd6b9abe12f703ec1959f43cd3be6ea227b53da7492dd7.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) CopyToSectionGroup()(*i0788ee8cc0f88534473ad867f6033f79c94d3ced1d865def09e887a4af6505f0.CopyToSectionGroupRequestBuilder) {
    return i0788ee8cc0f88534473ad867f6033f79c94d3ced1d865def09e887a4af6505f0.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *OnenoteSectionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *OnenoteSectionRequestBuilder) CreateGetRequestInformation(q func (value *OnenoteSectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(OnenoteSectionRequestBuilderGetQueryParameters)
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
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *OnenoteSectionRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteSectionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteSectionRequestBuilder) Get(q func (value *OnenoteSectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenoteSection() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection), nil
}
func (m *OnenoteSectionRequestBuilder) Pages()(*i090eb63c39e2c392547c70c8250c04ea9dc1ea28fd27551825475fe51dedb728.PagesRequestBuilder) {
    return i090eb63c39e2c392547c70c8250c04ea9dc1ea28fd27551825475fe51dedb728.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.notebooks.item.sections.item.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteSectionRequestBuilder) PagesById(id string)(*icd8f355cf1b63abcb90f8942bfacbc3a61d12900f6fa5545795d228102f51471.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return icd8f355cf1b63abcb90f8942bfacbc3a61d12900f6fa5545795d228102f51471.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentNotebook()(*i7fda827a8abe7ff3582c46ee5993acec0f2c7f7fe245e0e35e98367e8db423c5.ParentNotebookRequestBuilder) {
    return i7fda827a8abe7ff3582c46ee5993acec0f2c7f7fe245e0e35e98367e8db423c5.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentSectionGroup()(*i84bdfed9ec30dc209a4c530b508d815c14183bd652a5999a2f6f72613165ac5b.ParentSectionGroupRequestBuilder) {
    return i84bdfed9ec30dc209a4c530b508d815c14183bd652a5999a2f6f72613165ac5b.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The sections in the notebook. Read-only. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *OnenoteSectionRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
