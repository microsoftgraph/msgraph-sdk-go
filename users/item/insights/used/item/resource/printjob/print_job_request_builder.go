package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3fdd269e9cc01f5b2ce54c97aea664e17986a40b049d1d5e5acc2304a13ecb03 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printjob/cancel"
    i56e5974478defaffdf1fff8d717b6dea8c54879ce0bd7d38ed98f866e6bd1102 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printjob/start"
    i78b0b951711be90460ac7859eb137bdf1c3a6ea32f5bd2bf5a098ed74ef8aff0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printjob/redirect"
    i8cfea6715264ae5a84ff05853c4358ffc56d0d148e5c824f2d5132d307713934 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/used/item/resource/printjob/abort"
)

type PrintJobRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*i8cfea6715264ae5a84ff05853c4358ffc56d0d148e5c824f2d5132d307713934.AbortRequestBuilder) {
    return i8cfea6715264ae5a84ff05853c4358ffc56d0d148e5c824f2d5132d307713934.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i3fdd269e9cc01f5b2ce54c97aea664e17986a40b049d1d5e5acc2304a13ecb03.CancelRequestBuilder) {
    return i3fdd269e9cc01f5b2ce54c97aea664e17986a40b049d1d5e5acc2304a13ecb03.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/used/{usedInsight_id}/resource/microsoft.graph.printJob";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewPrintJobRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintJobRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintJobRequestBuilder) Redirect()(*i78b0b951711be90460ac7859eb137bdf1c3a6ea32f5bd2bf5a098ed74ef8aff0.RedirectRequestBuilder) {
    return i78b0b951711be90460ac7859eb137bdf1c3a6ea32f5bd2bf5a098ed74ef8aff0.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i56e5974478defaffdf1fff8d717b6dea8c54879ce0bd7d38ed98f866e6bd1102.StartRequestBuilder) {
    return i56e5974478defaffdf1fff8d717b6dea8c54879ce0bd7d38ed98f866e6bd1102.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
