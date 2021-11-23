package workbookrangeview

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i82988629ba479cd4e3cb571ce52ef062a21e46a4a21002d26c10c3cef7e01673 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/workbookrangeview/range_escaped"
)

// WorkbookRangeViewRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRangeView
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.workbookRangeView";
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
// Range_escaped builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.workbookRangeView\microsoft.graph.range()
func (m *WorkbookRangeViewRequestBuilder) Range_escaped()(*i82988629ba479cd4e3cb571ce52ef062a21e46a4a21002d26c10c3cef7e01673.RangeRequestBuilder) {
    return i82988629ba479cd4e3cb571ce52ef062a21e46a4a21002d26c10c3cef7e01673.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
