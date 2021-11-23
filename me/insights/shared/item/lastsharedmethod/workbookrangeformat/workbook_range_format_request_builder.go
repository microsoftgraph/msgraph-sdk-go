package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0afa10a52583757a95ea9ed84975e6eddba16db8b0960784b5fd95200f64309e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangeformat/autofitcolumns"
    ic5de5a458465cad03098c751beb3ef14e3924a3b515ded8cad8bae811707e3a4 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangeformat/autofitrows"
)

// WorkbookRangeFormatRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRangeFormat
type WorkbookRangeFormatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*i0afa10a52583757a95ea9ed84975e6eddba16db8b0960784b5fd95200f64309e.AutofitColumnsRequestBuilder) {
    return i0afa10a52583757a95ea9ed84975e6eddba16db8b0960784b5fd95200f64309e.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*ic5de5a458465cad03098c751beb3ef14e3924a3b515ded8cad8bae811707e3a4.AutofitRowsRequestBuilder) {
    return ic5de5a458465cad03098c751beb3ef14e3924a3b515ded8cad8bae811707e3a4.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFormatRequestBuilderInternal instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRangeFormat";
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
