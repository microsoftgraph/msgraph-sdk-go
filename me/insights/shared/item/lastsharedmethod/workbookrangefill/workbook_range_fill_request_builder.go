package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01f6159efd2bc7d105eb4296aa0bf8586fabcbb0ca8e507764b98348b29370a1 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/workbookrangefill/clear"
)

// workbookRangeFillRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.workbookRangeFill
type WorkbookRangeFillRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*i01f6159efd2bc7d105eb4296aa0bf8586fabcbb0ca8e507764b98348b29370a1.ClearRequestBuilder) {
    return i01f6159efd2bc7d105eb4296aa0bf8586fabcbb0ca8e507764b98348b29370a1.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFillRequestBuilderInternal instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.workbookRangeFill";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookRangeFillRequestBuilder instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
func NewWorkbookRangeFillRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFillRequestBuilderInternal(urlParams, requestAdapter)
}
