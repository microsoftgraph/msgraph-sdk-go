package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0dc6645d7e0cbe9ed7208bb276e0fe3100e8a57b894161244ad8e3da90b65ca4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printjob/start"
    i7e17a89bd86ce092f7479c22f78f2e8489fe5a32e709d44d1437519155c60dc0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printjob/redirect"
    ib23d731227c773b22cce072b1253ea039721abe06783510499869a2bdd238d90 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printjob/abort"
    ib704f387d9b1af33a522a47750ed92cc8b89f5b8882e76985288aa981f2ea1f6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/printjob/cancel"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*ib23d731227c773b22cce072b1253ea039721abe06783510499869a2bdd238d90.AbortRequestBuilder) {
    return ib23d731227c773b22cce072b1253ea039721abe06783510499869a2bdd238d90.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*ib704f387d9b1af33a522a47750ed92cc8b89f5b8882e76985288aa981f2ea1f6.CancelRequestBuilder) {
    return ib704f387d9b1af33a522a47750ed92cc8b89f5b8882e76985288aa981f2ea1f6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PrintJobRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.printJob";
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
func (m *PrintJobRequestBuilder) Redirect()(*i7e17a89bd86ce092f7479c22f78f2e8489fe5a32e709d44d1437519155c60dc0.RedirectRequestBuilder) {
    return i7e17a89bd86ce092f7479c22f78f2e8489fe5a32e709d44d1437519155c60dc0.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i0dc6645d7e0cbe9ed7208bb276e0fe3100e8a57b894161244ad8e3da90b65ca4.StartRequestBuilder) {
    return i0dc6645d7e0cbe9ed7208bb276e0fe3100e8a57b894161244ad8e3da90b65ca4.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
