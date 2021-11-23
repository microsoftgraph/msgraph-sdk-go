package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i487b40a03cdb07a384dea2221eada41f575f7de0cbd127bce368b9c933286cb8 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printjob/cancel"
    i6351dd0ddea973fa8b7f26f412707580c55b50f47df7ebc45c228e789517c577 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printjob/start"
    i8aac8ef8c35d15b986b93d01aaa723fff3e43b89f4bb6d7ed7d871e195dd0b60 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printjob/redirect"
    i90d257538f345b5c95ed98dc8e652b3831bcebeaaed03456503caa48d844e50d "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/resource/printjob/abort"
)

// PrintJobRequestBuilder builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\resource\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*i90d257538f345b5c95ed98dc8e652b3831bcebeaaed03456503caa48d844e50d.AbortRequestBuilder) {
    return i90d257538f345b5c95ed98dc8e652b3831bcebeaaed03456503caa48d844e50d.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i487b40a03cdb07a384dea2221eada41f575f7de0cbd127bce368b9c933286cb8.CancelRequestBuilder) {
    return i487b40a03cdb07a384dea2221eada41f575f7de0cbd127bce368b9c933286cb8.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrintJobRequestBuilderInternal instantiates a new PrintJobRequestBuilder and sets the default values.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/resource/microsoft.graph.printJob";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrintJobRequestBuilder instantiates a new PrintJobRequestBuilder and sets the default values.
func NewPrintJobRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintJobRequestBuilder) Redirect()(*i8aac8ef8c35d15b986b93d01aaa723fff3e43b89f4bb6d7ed7d871e195dd0b60.RedirectRequestBuilder) {
    return i8aac8ef8c35d15b986b93d01aaa723fff3e43b89f4bb6d7ed7d871e195dd0b60.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i6351dd0ddea973fa8b7f26f412707580c55b50f47df7ebc45c228e789517c577.StartRequestBuilder) {
    return i6351dd0ddea973fa8b7f26f412707580c55b50f47df7ebc45c228e789517c577.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
