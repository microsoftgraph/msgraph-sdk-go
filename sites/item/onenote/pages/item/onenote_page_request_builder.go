package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i00c1ac7a0b1d419c6e8068d7a1ae18ae47333e7d6aceec9db8a7dc0cb4be948e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/content"
    i1f835ad8becc1fd09ce10176dd39f1d478afabf28fe9c53e39bfabe255bd0f7e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentnotebook"
    i85b6526773790bbcbaeb5a084d3b25327a8fc96c6a3cdcc30c7b7ec8bc19b9d4 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/copytosection"
    i9714ef5f205ccd0f2a637602b59694b2db42ea9f89263bd1ce6baf1bf303ab9c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/parentsection"
    i9dddb33025d0233edfdde534781ea991ab776b435455942bcde0a1097108f3b2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/preview"
    icab4f6eb0e867ecdd726162d5966bd8a8ecc3eef1bb6a959a36156dd8b09847d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/pages/item/onenotepatchcontent"
)

// OnenotePageRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}
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
// OnenotePageRequestBuilderGetQueryParameters the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/pages/{onenotePage_id}{?select,expand}";
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
func (m *OnenotePageRequestBuilder) Content()(*i00c1ac7a0b1d419c6e8068d7a1ae18ae47333e7d6aceec9db8a7dc0cb4be948e.ContentRequestBuilder) {
    return i00c1ac7a0b1d419c6e8068d7a1ae18ae47333e7d6aceec9db8a7dc0cb4be948e.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*i85b6526773790bbcbaeb5a084d3b25327a8fc96c6a3cdcc30c7b7ec8bc19b9d4.CopyToSectionRequestBuilder) {
    return i85b6526773790bbcbaeb5a084d3b25327a8fc96c6a3cdcc30c7b7ec8bc19b9d4.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
// CreateGetRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
// CreatePatchRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
// Delete the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
// Get the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*icab4f6eb0e867ecdd726162d5966bd8a8ecc3eef1bb6a959a36156dd8b09847d.OnenotePatchContentRequestBuilder) {
    return icab4f6eb0e867ecdd726162d5966bd8a8ecc3eef1bb6a959a36156dd8b09847d.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i1f835ad8becc1fd09ce10176dd39f1d478afabf28fe9c53e39bfabe255bd0f7e.ParentNotebookRequestBuilder) {
    return i1f835ad8becc1fd09ce10176dd39f1d478afabf28fe9c53e39bfabe255bd0f7e.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*i9714ef5f205ccd0f2a637602b59694b2db42ea9f89263bd1ce6baf1bf303ab9c.ParentSectionRequestBuilder) {
    return i9714ef5f205ccd0f2a637602b59694b2db42ea9f89263bd1ce6baf1bf303ab9c.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
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
// Preview builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i9dddb33025d0233edfdde534781ea991ab776b435455942bcde0a1097108f3b2.PreviewRequestBuilder) {
    return i9dddb33025d0233edfdde534781ea991ab776b435455942bcde0a1097108f3b2.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
