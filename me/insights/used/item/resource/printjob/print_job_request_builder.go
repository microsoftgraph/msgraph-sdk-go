package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i87287b9a7b2bb536a07331e857f663d21c0c06ad932f2bb4a7d0c8e77c6d17f7 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printjob/redirect"
    i8dfaaa307ab7d0409b310e5017b2003ab4baa7ff6a2eb8165e2302e62b6cddd3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printjob/abort"
    i9c05938ee4a36ea474a3b1585b9dc783c85f3a6b6d8d326aea450369688c5af9 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printjob/cancel"
    if3310ca6b788a2306a23917469f9361a0218fae09344d900fdd19f9e1bd8ab74 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/printjob/start"
)

// Builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*i8dfaaa307ab7d0409b310e5017b2003ab4baa7ff6a2eb8165e2302e62b6cddd3.AbortRequestBuilder) {
    return i8dfaaa307ab7d0409b310e5017b2003ab4baa7ff6a2eb8165e2302e62b6cddd3.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i9c05938ee4a36ea474a3b1585b9dc783c85f3a6b6d8d326aea450369688c5af9.CancelRequestBuilder) {
    return i9c05938ee4a36ea474a3b1585b9dc783c85f3a6b6d8d326aea450369688c5af9.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrintJobRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource/microsoft.graph.printJob";
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
func (m *PrintJobRequestBuilder) Redirect()(*i87287b9a7b2bb536a07331e857f663d21c0c06ad932f2bb4a7d0c8e77c6d17f7.RedirectRequestBuilder) {
    return i87287b9a7b2bb536a07331e857f663d21c0c06ad932f2bb4a7d0c8e77c6d17f7.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*if3310ca6b788a2306a23917469f9361a0218fae09344d900fdd19f9e1bd8ab74.StartRequestBuilder) {
    return if3310ca6b788a2306a23917469f9361a0218fae09344d900fdd19f9e1bd8ab74.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
