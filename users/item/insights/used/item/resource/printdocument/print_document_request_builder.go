package printdocument

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ibb2327bfd91b3be331149b58b263c9b743e19a395cda18ed5e45f9490676743a "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printdocument/createuploadsession"
)

// printDocumentRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.printDocument
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.printDocument";
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
func (m *PrintDocumentRequestBuilder) CreateUploadSession()(*ibb2327bfd91b3be331149b58b263c9b743e19a395cda18ed5e45f9490676743a.CreateUploadSessionRequestBuilder) {
    return ibb2327bfd91b3be331149b58b263c9b743e19a395cda18ed5e45f9490676743a.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
