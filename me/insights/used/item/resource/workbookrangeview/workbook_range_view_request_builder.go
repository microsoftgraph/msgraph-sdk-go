package workbookrangeview

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ic05fecab786dbace1a31b4bc589e370751f9ee06abda243e9194fa7908f5b130 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangeview/range_escaped"
)

// workbookRangeViewRequestBuilder builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeView
type WorkbookRangeViewRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewWorkbookRangeViewRequestBuilderInternal instantiates a new WorkbookRangeViewRequestBuilder and sets the default values.
func NewWorkbookRangeViewRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeViewRequestBuilder) {
    m := &WorkbookRangeViewRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeView";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWorkbookRangeViewRequestBuilder instantiates a new WorkbookRangeViewRequestBuilder and sets the default values.
func NewWorkbookRangeViewRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeViewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeViewRequestBuilderInternal(urlParams, requestAdapter)
}
// Range_escaped builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeView\microsoft.graph.range()
func (m *WorkbookRangeViewRequestBuilder) Range_escaped()(*ic05fecab786dbace1a31b4bc589e370751f9ee06abda243e9194fa7908f5b130.RangeRequestBuilder) {
    return ic05fecab786dbace1a31b4bc589e370751f9ee06abda243e9194fa7908f5b130.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
