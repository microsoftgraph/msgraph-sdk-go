package workbookrangeview

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i8e4ee9cdd99c04cfbfb4358aa46e55d8023843ad4daf185890ce257b73ce7907 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/workbookrangeview/range_escaped"
)

// workbookRangeViewRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeView
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeView";
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
// Range_escaped builds and executes requests for operations under \users\{user-id}\insights\used\{usedInsight-id}\resource\microsoft.graph.workbookRangeView\microsoft.graph.range()
func (m *WorkbookRangeViewRequestBuilder) Range_escaped()(*i8e4ee9cdd99c04cfbfb4358aa46e55d8023843ad4daf185890ce257b73ce7907.RangeRequestBuilder) {
    return i8e4ee9cdd99c04cfbfb4358aa46e55d8023843ad4daf185890ce257b73ce7907.NewRangeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
