package printdocument

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i87aeb4b05dcc01f973ed2d571eb3c7db43911fa6f6ff6d11608a8cb4921d4933 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printdocument/createuploadsession"
)

// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.printDocument
type PrintDocumentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Instantiates a new PrintDocumentRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintDocumentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    m := &PrintDocumentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.printDocument";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrintDocumentRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintDocumentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintDocumentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintDocumentRequestBuilder) CreateUploadSession()(*i87aeb4b05dcc01f973ed2d571eb3c7db43911fa6f6ff6d11608a8cb4921d4933.CreateUploadSessionRequestBuilder) {
    return i87aeb4b05dcc01f973ed2d571eb3c7db43911fa6f6ff6d11608a8cb4921d4933.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
