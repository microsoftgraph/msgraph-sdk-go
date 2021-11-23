package printdocument

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i73251bdb16df2920100ba1dfaa9e83d1784bc236aeb1a4c27e93221aedb09125 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printdocument/createuploadsession"
)

// PrintDocumentRequestBuilder builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.printDocument
type PrintDocumentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewPrintDocumentRequestBuilderInternal instantiates a new PrintDocumentRequestBuilder and sets the default values.
func NewPrintDocumentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    m := &PrintDocumentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.printDocument";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintDocumentRequestBuilder instantiates a new PrintDocumentRequestBuilder and sets the default values.
func NewPrintDocumentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintDocumentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintDocumentRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintDocumentRequestBuilder) CreateUploadSession()(*i73251bdb16df2920100ba1dfaa9e83d1784bc236aeb1a4c27e93221aedb09125.CreateUploadSessionRequestBuilder) {
    return i73251bdb16df2920100ba1dfaa9e83d1784bc236aeb1a4c27e93221aedb09125.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
