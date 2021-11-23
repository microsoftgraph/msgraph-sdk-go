package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i2951426e975cb9af7743c0f0535cc9456fe72639efd4c8fcd10b8214973a29e6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangeformat/autofitrows"
    ib05114c30a9f1b3b68921e972440009f75c723bce137ae323178e0d38c22e89b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/workbookrangeformat/autofitcolumns"
)

// workbookRangeFormatRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRangeFormat
type WorkbookRangeFormatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*ib05114c30a9f1b3b68921e972440009f75c723bce137ae323178e0d38c22e89b.AutofitColumnsRequestBuilder) {
    return ib05114c30a9f1b3b68921e972440009f75c723bce137ae323178e0d38c22e89b.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i2951426e975cb9af7743c0f0535cc9456fe72639efd4c8fcd10b8214973a29e6.AutofitRowsRequestBuilder) {
    return i2951426e975cb9af7743c0f0535cc9456fe72639efd4c8fcd10b8214973a29e6.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFormatRequestBuilderInternal instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRangeFormat";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookRangeFormatRequestBuilder instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
func NewWorkbookRangeFormatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFormatRequestBuilderInternal(urlParams, requestAdapter)
}
