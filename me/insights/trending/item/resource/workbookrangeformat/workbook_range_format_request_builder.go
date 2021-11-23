package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeformat/autofitrows"
    ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeformat/autofitcolumns"
)

// WorkbookRangeFormatRequestBuilder builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRangeFormat
type WorkbookRangeFormatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217.AutofitColumnsRequestBuilder) {
    return ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea.AutofitRowsRequestBuilder) {
    return i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFormatRequestBuilderInternal instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.workbookRangeFormat";
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
