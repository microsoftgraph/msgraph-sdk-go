package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i37bfcfc6976e33c9456e99baffbfc31c84925111972710b991c40da29f1402e2 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printjob/redirect"
    i81e493f352f9b107df88741067eae820042a46b6bfdbcfb07ac0568cef7ecb2e "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printjob/start"
    i867fe7d0c3b9b391fcd78f87acae7f35518ace42398c90ca3e8af6fe435830f6 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printjob/cancel"
    ia012775db1c00bf654b1da3aac10c30abe613b0bdb7d4bf21219b6a1c4c3b957 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/printjob/abort"
)

// PrintJobRequestBuilder builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.printJob
type PrintJobRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*ia012775db1c00bf654b1da3aac10c30abe613b0bdb7d4bf21219b6a1c4c3b957.AbortRequestBuilder) {
    return ia012775db1c00bf654b1da3aac10c30abe613b0bdb7d4bf21219b6a1c4c3b957.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i867fe7d0c3b9b391fcd78f87acae7f35518ace42398c90ca3e8af6fe435830f6.CancelRequestBuilder) {
    return i867fe7d0c3b9b391fcd78f87acae7f35518ace42398c90ca3e8af6fe435830f6.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrintJobRequestBuilderInternal instantiates a new PrintJobRequestBuilder and sets the default values.
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.printJob";
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
func (m *PrintJobRequestBuilder) Redirect()(*i37bfcfc6976e33c9456e99baffbfc31c84925111972710b991c40da29f1402e2.RedirectRequestBuilder) {
    return i37bfcfc6976e33c9456e99baffbfc31c84925111972710b991c40da29f1402e2.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i81e493f352f9b107df88741067eae820042a46b6bfdbcfb07ac0568cef7ecb2e.StartRequestBuilder) {
    return i81e493f352f9b107df88741067eae820042a46b6bfdbcfb07ac0568cef7ecb2e.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
