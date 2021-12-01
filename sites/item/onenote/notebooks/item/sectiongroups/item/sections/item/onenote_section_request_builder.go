package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i249095d185f63a55ed5e2ec90a11e2096534eedad287802504f5bcdb993d2a5b "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/copytonotebook"
    i562c96484c27305b96089b1a602edaa8c94d317e0106d82b1dd3c3763d184e90 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/parentsectiongroup"
    i76bf563dda9eef61eefe24d01e9bf663f7c48ede372ca407a6fb4dea6592b2d6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/parentnotebook"
    i9514c1b235526a6fc2fa0820d0f8294370d4c13ee58b86cd3b1d1488cf700288 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages"
    iecbd70e7691f6847ef6cf6c2247c2275d23a0a59544d8c2285af39cbda871e83 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/copytosectiongroup"
    id467fce409b3f7e1eb7de9c4053d1987d56cbb170624c42ebbeaf7e1fa9cd7bc "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item"
)

// OnenoteSectionRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote\notebooks\{notebook-id}\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}
type OnenoteSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionRequestBuilderDeleteOptions options for Delete
type OnenoteSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetOptions options for Get
type OnenoteSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type OnenoteSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionRequestBuilderPatchOptions options for Patch
type OnenoteSectionRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionRequestBuilderInternal instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    m := &OnenoteSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/notebooks/{notebook_id}/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionRequestBuilder instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionRequestBuilder) CopyToNotebook()(*i249095d185f63a55ed5e2ec90a11e2096534eedad287802504f5bcdb993d2a5b.CopyToNotebookRequestBuilder) {
    return i249095d185f63a55ed5e2ec90a11e2096534eedad287802504f5bcdb993d2a5b.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) CopyToSectionGroup()(*iecbd70e7691f6847ef6cf6c2247c2275d23a0a59544d8c2285af39cbda871e83.CopyToSectionGroupRequestBuilder) {
    return iecbd70e7691f6847ef6cf6c2247c2275d23a0a59544d8c2285af39cbda871e83.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Delete(options *OnenoteSectionRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Get(options *OnenoteSectionRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenoteSection() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenoteSection), nil
}
func (m *OnenoteSectionRequestBuilder) Pages()(*i9514c1b235526a6fc2fa0820d0f8294370d4c13ee58b86cd3b1d1488cf700288.PagesRequestBuilder) {
    return i9514c1b235526a6fc2fa0820d0f8294370d4c13ee58b86cd3b1d1488cf700288.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.onenote.notebooks.item.sectionGroups.item.sections.item.pages.item collection
func (m *OnenoteSectionRequestBuilder) PagesById(id string)(*id467fce409b3f7e1eb7de9c4053d1987d56cbb170624c42ebbeaf7e1fa9cd7bc.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return id467fce409b3f7e1eb7de9c4053d1987d56cbb170624c42ebbeaf7e1fa9cd7bc.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentNotebook()(*i76bf563dda9eef61eefe24d01e9bf663f7c48ede372ca407a6fb4dea6592b2d6.ParentNotebookRequestBuilder) {
    return i76bf563dda9eef61eefe24d01e9bf663f7c48ede372ca407a6fb4dea6592b2d6.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentSectionGroup()(*i562c96484c27305b96089b1a602edaa8c94d317e0106d82b1dd3c3763d184e90.ParentSectionGroupRequestBuilder) {
    return i562c96484c27305b96089b1a602edaa8c94d317e0106d82b1dd3c3763d184e90.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Patch(options *OnenoteSectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
