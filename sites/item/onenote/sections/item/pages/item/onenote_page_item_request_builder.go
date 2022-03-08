package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i58ccd0f326a6850c97ccf703b2449085c4aa3dbbaf5475e66ec4200040131469 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/preview"
    i853f326417fded8a86fb5fe2a634b14a610581b4831cb0c96d1c0b3f372b400c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/parentsection"
    i9e3ca00a0159142bca20ec1d002eeeb62ad9e7d073ed953a3ec7d3eb34249500 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/copytosection"
    ia0ee195388b0a393aa3e34a0130313d331e898096011da201f7a71815e6e58b2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/content"
    ida75c6a37de0f4115b7a8ae1274459efade3bbd15668948b890443eec9cb010f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/onenotepatchcontent"
    ie1ab22c49b071ab6bd9fa8522a00e465fea46e9f8fc6a3ee696074a47bee343f "github.com/microsoftgraph/msgraph-sdk-go/sites/item/onenote/sections/item/pages/item/parentnotebook"
)

// OnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type OnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageItemRequestBuilderDeleteOptions options for Delete
type OnenotePageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetOptions options for Get
type OnenotePageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type OnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageItemRequestBuilderPatchOptions options for Patch
type OnenotePageItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    m := &OnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/sections/{onenoteSection_id}/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewOnenotePageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageItemRequestBuilder) Content()(*ia0ee195388b0a393aa3e34a0130313d331e898096011da201f7a71815e6e58b2.ContentRequestBuilder) {
    return ia0ee195388b0a393aa3e34a0130313d331e898096011da201f7a71815e6e58b2.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) CopyToSection()(*i9e3ca00a0159142bca20ec1d002eeeb62ad9e7d073ed953a3ec7d3eb34249500.CopyToSectionRequestBuilder) {
    return i9e3ca00a0159142bca20ec1d002eeeb62ad9e7d073ed953a3ec7d3eb34249500.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for sites
func (m *OnenotePageItemRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *OnenotePageItemRequestBuilder) CreateGetRequestInformation(options *OnenotePageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property pages in sites
func (m *OnenotePageItemRequestBuilder) CreatePatchRequestInformation(options *OnenotePageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property pages for sites
func (m *OnenotePageItemRequestBuilder) Delete(options *OnenotePageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *OnenotePageItemRequestBuilder) Get(options *OnenotePageItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateOnenotePageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnenotePageable), nil
}
func (m *OnenotePageItemRequestBuilder) OnenotePatchContent()(*ida75c6a37de0f4115b7a8ae1274459efade3bbd15668948b890443eec9cb010f.OnenotePatchContentRequestBuilder) {
    return ida75c6a37de0f4115b7a8ae1274459efade3bbd15668948b890443eec9cb010f.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) ParentNotebook()(*ie1ab22c49b071ab6bd9fa8522a00e465fea46e9f8fc6a3ee696074a47bee343f.ParentNotebookRequestBuilder) {
    return ie1ab22c49b071ab6bd9fa8522a00e465fea46e9f8fc6a3ee696074a47bee343f.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageItemRequestBuilder) ParentSection()(*i853f326417fded8a86fb5fe2a634b14a610581b4831cb0c96d1c0b3f372b400c.ParentSectionRequestBuilder) {
    return i853f326417fded8a86fb5fe2a634b14a610581b4831cb0c96d1c0b3f372b400c.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in sites
func (m *OnenotePageItemRequestBuilder) Patch(options *OnenotePageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Preview provides operations to call the preview method.
func (m *OnenotePageItemRequestBuilder) Preview()(*i58ccd0f326a6850c97ccf703b2449085c4aa3dbbaf5475e66ec4200040131469.PreviewRequestBuilder) {
    return i58ccd0f326a6850c97ccf703b2449085c4aa3dbbaf5475e66ec4200040131469.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
