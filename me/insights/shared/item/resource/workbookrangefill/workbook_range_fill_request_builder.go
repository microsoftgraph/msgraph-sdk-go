package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i82c6bf85b2492322ba6fb9c982c87ddc6ebcb18961a81c582df903cee20da37b "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/workbookrangefill/clear"
)

// workbookRangeFillRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRangeFill
type WorkbookRangeFillRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*i82c6bf85b2492322ba6fb9c982c87ddc6ebcb18961a81c582df903cee20da37b.ClearRequestBuilder) {
    return i82c6bf85b2492322ba6fb9c982c87ddc6ebcb18961a81c582df903cee20da37b.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFillRequestBuilderInternal instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRangeFill";
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
