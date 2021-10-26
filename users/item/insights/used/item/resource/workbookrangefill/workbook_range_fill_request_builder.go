package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i66609d9c4de208a5a0d13b41d20b591149007778059c89feef4a9c8a31756057 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangefill/clear"
)

// Builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeFill
type WorkbookRangeFillRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*i66609d9c4de208a5a0d13b41d20b591149007778059c89feef4a9c8a31756057.ClearRequestBuilder) {
    return i66609d9c4de208a5a0d13b41d20b591149007778059c89feef4a9c8a31756057.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeFill";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWorkbookRangeFillRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFillRequestBuilderInternal(urlParams, requestAdapter)
}
