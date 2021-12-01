package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1ad3fe5394c35529e6087d62114881d22e7b120baabdc943f14fd614e978f28d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/preview"
    i218a4e8a4c96c784aedfa83171b53c418e3ab42c114b9a8accd67f489d3f5f7f "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/content"
    i62503885920d3da262ecb302fd8a95917237e73be938e66fab1ca94c0bd9741c "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/parentnotebook"
    i8d1830bc452ef248a843f372fa93ef5d4e315d8cccaff07a6954e2bd37801a21 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/onenotepatchcontent"
    i94921cb1adfb83c3646627ccff8b972836f7726b76816795100c23f9e40a5df6 "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/copytosection"
    ia35fe61769385a2a7408c1b84a0bda91edb195d2c63141c5683f667c4a752c0d "github.com/microsoftgraph/msgraph-sdk-go/me/onenote/sectiongroups/item/parentnotebook/sections/item/pages/item/parentsection"
)

// OnenotePageRequestBuilder builds and executes requests for operations under \me\onenote\sectionGroups\{sectionGroup-id}\parentNotebook\sections\{onenoteSection-id}\pages\{onenotePage-id}
type OnenotePageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageRequestBuilderDeleteOptions options for Delete
type OnenotePageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetOptions options for Get
type OnenotePageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type OnenotePageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageRequestBuilderPatchOptions options for Patch
type OnenotePageRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenotePageRequestBuilderInternal instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    m := &OnenotePageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/sectionGroups/{sectionGroup_id}/parentNotebook/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageRequestBuilder instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageRequestBuilder) Content()(*i218a4e8a4c96c784aedfa83171b53c418e3ab42c114b9a8accd67f489d3f5f7f.ContentRequestBuilder) {
    return i218a4e8a4c96c784aedfa83171b53c418e3ab42c114b9a8accd67f489d3f5f7f.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*i94921cb1adfb83c3646627ccff8b972836f7726b76816795100c23f9e40a5df6.CopyToSectionRequestBuilder) {
    return i94921cb1adfb83c3646627ccff8b972836f7726b76816795100c23f9e40a5df6.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateGetRequestInformation(options *OnenotePageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreatePatchRequestInformation(options *OnenotePageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Delete(options *OnenotePageRequestBuilderDeleteOptions)(error) {
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
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Get(options *OnenotePageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnenotePage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePage), nil
}
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*i8d1830bc452ef248a843f372fa93ef5d4e315d8cccaff07a6954e2bd37801a21.OnenotePatchContentRequestBuilder) {
    return i8d1830bc452ef248a843f372fa93ef5d4e315d8cccaff07a6954e2bd37801a21.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i62503885920d3da262ecb302fd8a95917237e73be938e66fab1ca94c0bd9741c.ParentNotebookRequestBuilder) {
    return i62503885920d3da262ecb302fd8a95917237e73be938e66fab1ca94c0bd9741c.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*ia35fe61769385a2a7408c1b84a0bda91edb195d2c63141c5683f667c4a752c0d.ParentSectionRequestBuilder) {
    return ia35fe61769385a2a7408c1b84a0bda91edb195d2c63141c5683f667c4a752c0d.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Patch(options *OnenotePageRequestBuilderPatchOptions)(error) {
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
// Preview builds and executes requests for operations under \me\onenote\sectionGroups\{sectionGroup-id}\parentNotebook\sections\{onenoteSection-id}\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i1ad3fe5394c35529e6087d62114881d22e7b120baabdc943f14fd614e978f28d.PreviewRequestBuilder) {
    return i1ad3fe5394c35529e6087d62114881d22e7b120baabdc943f14fd614e978f28d.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
