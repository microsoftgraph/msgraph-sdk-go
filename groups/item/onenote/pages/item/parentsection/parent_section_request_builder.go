package parentsection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0bc859f0ea50785cf9589f0c7a8b1475d323b3f76f9a15fe64597f0fcc2433b8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentnotebook"
    i108804fb5af9e7152cb89f11dc5b46e2c95e051fc265d6f696f244c7fc077ec2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/parentsectiongroup"
    i3e01a5baca9b96c1dda82911a6ce92595857719cf6623f6213dc2c65adb4ebf5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/copytosectiongroup"
    i7e452f413db9b3d0b3444da9782fa3f62c1d9e265cc3d090ebd96c7ae595a9ba "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/pages"
    ib550ff9dd6a37222c64b0794d65eb94a92442cff279c1131c892ab9f9834d091 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/copytonotebook"
    i17258079b21b69dff9150ce0b4fb7283d332e70e45786661f2b5dd7295a816e2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/onenote/pages/item/parentsection/pages/item"
)

// Builds and executes requests for operations under \groups\{group-id}\onenote\pages\{onenotePage-id}\parentSection
type ParentSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The section that contains the page. Read-only.
type ParentSectionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ParentSectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    m := &ParentSectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/groups/{group_id}/onenote/pages/{onenotePage_id}/parentSection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ParentSectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewParentSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentSectionRequestBuilder) CopyToNotebook()(*ib550ff9dd6a37222c64b0794d65eb94a92442cff279c1131c892ab9f9834d091.CopyToNotebookRequestBuilder) {
    return ib550ff9dd6a37222c64b0794d65eb94a92442cff279c1131c892ab9f9834d091.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) CopyToSectionGroup()(*i3e01a5baca9b96c1dda82911a6ce92595857719cf6623f6213dc2c65adb4ebf5.CopyToSectionGroupRequestBuilder) {
    return i3e01a5baca9b96c1dda82911a6ce92595857719cf6623f6213dc2c65adb4ebf5.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section that contains the page. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ParentSectionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ParentSectionRequestBuilder) CreateGetRequestInformation(q func (value *ParentSectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ParentSectionRequestBuilderGetQueryParameters)
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
// The section that contains the page. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ParentSectionRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentSectionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// The section that contains the page. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentSectionRequestBuilder) Get(q func (value *ParentSectionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, error) {
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
func (m *ParentSectionRequestBuilder) Pages()(*i7e452f413db9b3d0b3444da9782fa3f62c1d9e265cc3d090ebd96c7ae595a9ba.PagesRequestBuilder) {
    return i7e452f413db9b3d0b3444da9782fa3f62c1d9e265cc3d090ebd96c7ae595a9ba.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.onenote.pages.item.parentSection.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ParentSectionRequestBuilder) PagesById(id string)(*i17258079b21b69dff9150ce0b4fb7283d332e70e45786661f2b5dd7295a816e2.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id1"] = id
    }
    return i17258079b21b69dff9150ce0b4fb7283d332e70e45786661f2b5dd7295a816e2.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentNotebook()(*i0bc859f0ea50785cf9589f0c7a8b1475d323b3f76f9a15fe64597f0fcc2433b8.ParentNotebookRequestBuilder) {
    return i0bc859f0ea50785cf9589f0c7a8b1475d323b3f76f9a15fe64597f0fcc2433b8.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) ParentSectionGroup()(*i108804fb5af9e7152cb89f11dc5b46e2c95e051fc265d6f696f244c7fc077ec2.ParentSectionGroupRequestBuilder) {
    return i108804fb5af9e7152cb89f11dc5b46e2c95e051fc265d6f696f244c7fc077ec2.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section that contains the page. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ParentSectionRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
