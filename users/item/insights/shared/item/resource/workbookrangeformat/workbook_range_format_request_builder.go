package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i786c71b1cf6a084fd081649d35f26aa20fbf48652c77f7d3c3c5b720f619defa "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangeformat/autofitrows"
    ic6c84e32d545a9bbb88cf479b2029b2f02d6cbc7546d654c4ed91857b7b72611 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangeformat/autofitcolumns"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRangeFormat
type WorkbookRangeFormatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*ic6c84e32d545a9bbb88cf479b2029b2f02d6cbc7546d654c4ed91857b7b72611.AutofitColumnsRequestBuilder) {
    return ic6c84e32d545a9bbb88cf479b2029b2f02d6cbc7546d654c4ed91857b7b72611.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i786c71b1cf6a084fd081649d35f26aa20fbf48652c77f7d3c3c5b720f619defa.AutofitRowsRequestBuilder) {
    return i786c71b1cf6a084fd081649d35f26aa20fbf48652c77f7d3c3c5b720f619defa.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRangeFormat";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookRangeFormatRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeFormatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFormatRequestBuilderInternal(urlParams, requestAdapter)
}
