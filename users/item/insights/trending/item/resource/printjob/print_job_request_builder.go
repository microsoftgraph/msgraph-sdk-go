package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i407180de38297e0a491480885fb1c5f13886bc604d28b3423dd46714b7d027ef "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printjob/cancel"
    i84a88761230feb1fe8e5ed7fe29166548ffb71bb779d7e3ef658d4a9c087ead0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printjob/redirect"
    id11716614982f65206aa3c64c90303be0b05b22ef8ae83d6395b5dbaf8d19757 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printjob/start"
    ifbb148b173aec976e5b568894d783e79db832b6164b1ba06b15ac0f604892087 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/trending/item/resource/printjob/abort"
)

// printJobRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\trending\{trending-id}\resource\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*ifbb148b173aec976e5b568894d783e79db832b6164b1ba06b15ac0f604892087.AbortRequestBuilder) {
    return ifbb148b173aec976e5b568894d783e79db832b6164b1ba06b15ac0f604892087.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i407180de38297e0a491480885fb1c5f13886bc604d28b3423dd46714b7d027ef.CancelRequestBuilder) {
    return i407180de38297e0a491480885fb1c5f13886bc604d28b3423dd46714b7d027ef.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrintJobRequestBuilderInternal instantiates a new PrintJobRequestBuilder and sets the default values.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/trending/{trending_id}/resource/microsoft.graph.printJob";
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
func (m *PrintJobRequestBuilder) Redirect()(*i84a88761230feb1fe8e5ed7fe29166548ffb71bb779d7e3ef658d4a9c087ead0.RedirectRequestBuilder) {
    return i84a88761230feb1fe8e5ed7fe29166548ffb71bb779d7e3ef658d4a9c087ead0.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*id11716614982f65206aa3c64c90303be0b05b22ef8ae83d6395b5dbaf8d19757.StartRequestBuilder) {
    return id11716614982f65206aa3c64c90303be0b05b22ef8ae83d6395b5dbaf8d19757.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
