package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3c8a6f7c8215e9486f914e55271481bfc45e8064b6eb3509d2d9a2e91268b73f "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printjob/abort"
    i4e14e6dbb0605c0a103791f08c3032e72c79bffc964e1b58e1747dfe4a27e8f9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printjob/start"
    id3e2e383b9030e72364223dac4deb6060a517e2c07fa4ce205969bda75d3ea9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printjob/cancel"
    if8636381f75fc79aa28d0be5a77c579b52432e4b8470d09141bf7ea48ebb3996 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/printjob/redirect"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*i3c8a6f7c8215e9486f914e55271481bfc45e8064b6eb3509d2d9a2e91268b73f.AbortRequestBuilder) {
    return i3c8a6f7c8215e9486f914e55271481bfc45e8064b6eb3509d2d9a2e91268b73f.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*id3e2e383b9030e72364223dac4deb6060a517e2c07fa4ce205969bda75d3ea9d.CancelRequestBuilder) {
    return id3e2e383b9030e72364223dac4deb6060a517e2c07fa4ce205969bda75d3ea9d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrintJobRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.printJob";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrintJobRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintJobRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintJobRequestBuilder) Redirect()(*if8636381f75fc79aa28d0be5a77c579b52432e4b8470d09141bf7ea48ebb3996.RedirectRequestBuilder) {
    return if8636381f75fc79aa28d0be5a77c579b52432e4b8470d09141bf7ea48ebb3996.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i4e14e6dbb0605c0a103791f08c3032e72c79bffc964e1b58e1747dfe4a27e8f9.StartRequestBuilder) {
    return i4e14e6dbb0605c0a103791f08c3032e72c79bffc964e1b58e1747dfe4a27e8f9.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
