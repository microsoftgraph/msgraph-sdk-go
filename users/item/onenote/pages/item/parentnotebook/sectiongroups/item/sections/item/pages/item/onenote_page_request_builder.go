package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1cc799fd65ade3ddb5a3804bfb7ae7d969fc7006f993d0ca0f320a2c2f372f6e "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages/item/content"
    i6b11497ba6b7cbcda4b65d5159fb5575af4dfec1e3905a448dfa0db1ad32be78 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages/item/preview"
    i6fd61fcbb92d706e422bae72b03c9078c294059b4dba9843fca272bd99b1689e "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages/item/onenotepatchcontent"
    id10f84471dc0ff41f57535f1ea4cc67e423517f2a38abf3eed86dddf7af238f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/onenote/pages/item/parentnotebook/sectiongroups/item/sections/item/pages/item/copytosection"
)

// OnenotePageRequestBuilder builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}\parentNotebook\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\pages\{onenotePage-id1}
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/onenote/pages/{onenotePage_id}/parentNotebook/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/pages/{onenotePage_id1}{?select,expand}";
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
func (m *OnenotePageRequestBuilder) Content()(*i1cc799fd65ade3ddb5a3804bfb7ae7d969fc7006f993d0ca0f320a2c2f372f6e.ContentRequestBuilder) {
    return i1cc799fd65ade3ddb5a3804bfb7ae7d969fc7006f993d0ca0f320a2c2f372f6e.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*id10f84471dc0ff41f57535f1ea4cc67e423517f2a38abf3eed86dddf7af238f5.CopyToSectionRequestBuilder) {
    return id10f84471dc0ff41f57535f1ea4cc67e423517f2a38abf3eed86dddf7af238f5.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*i6fd61fcbb92d706e422bae72b03c9078c294059b4dba9843fca272bd99b1689e.OnenotePatchContentRequestBuilder) {
    return i6fd61fcbb92d706e422bae72b03c9078c294059b4dba9843fca272bd99b1689e.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Preview builds and executes requests for operations under \users\{user-id}\onenote\pages\{onenotePage-id}\parentNotebook\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\pages\{onenotePage-id1}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i6b11497ba6b7cbcda4b65d5159fb5575af4dfec1e3905a448dfa0db1ad32be78.PreviewRequestBuilder) {
    return i6b11497ba6b7cbcda4b65d5159fb5575af4dfec1e3905a448dfa0db1ad32be78.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
