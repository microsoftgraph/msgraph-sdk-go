package workbookrangefill

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i6da293e4b823640bac7b47af9990d3f95f0f6748794c761e1e5d9eb8cbe6aea6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangefill/clear"
)

// workbookRangeFillRequestBuilder builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.workbookRangeFill
type WorkbookRangeFillRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WorkbookRangeFillRequestBuilder) Clear()(*i6da293e4b823640bac7b47af9990d3f95f0f6748794c761e1e5d9eb8cbe6aea6.ClearRequestBuilder) {
    return i6da293e4b823640bac7b47af9990d3f95f0f6748794c761e1e5d9eb8cbe6aea6.NewClearRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWorkbookRangeFillRequestBuilderInternal instantiates a new WorkbookRangeFillRequestBuilder and sets the default values.
func NewWorkbookRangeFillRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFillRequestBuilder) {
    m := &WorkbookRangeFillRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.workbookRangeFill";
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
