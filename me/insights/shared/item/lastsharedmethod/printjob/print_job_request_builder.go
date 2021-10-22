package printjob

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i20a4f7945d4f27075487b46aa03df219c362f69a453beebd22836809ef512268 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printjob/redirect"
    i65d540e862794ef4e467b02d8f2d4c9bd0daaf7517806969e97b023ddc3b2152 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printjob/cancel"
    i8c1238ff81185e885512c301c4ea19545ac21eef777b326e5da6c9c10f97c437 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printjob/start"
    iea086e9a0e5c814bf8fffa6f32e7e5ee0acad87d9a30672e5faffb72ae926046 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/printjob/abort"
)

type PrintJobRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *PrintJobRequestBuilder) Abort()(*iea086e9a0e5c814bf8fffa6f32e7e5ee0acad87d9a30672e5faffb72ae926046.AbortRequestBuilder) {
    return iea086e9a0e5c814bf8fffa6f32e7e5ee0acad87d9a30672e5faffb72ae926046.NewAbortRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Cancel()(*i65d540e862794ef4e467b02d8f2d4c9bd0daaf7517806969e97b023ddc3b2152.CancelRequestBuilder) {
    return i65d540e862794ef4e467b02d8f2d4c9bd0daaf7517806969e97b023ddc3b2152.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewPrintJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintJobRequestBuilder) {
    m := &PrintJobRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.printJob";
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
func (m *PrintJobRequestBuilder) Redirect()(*i20a4f7945d4f27075487b46aa03df219c362f69a453beebd22836809ef512268.RedirectRequestBuilder) {
    return i20a4f7945d4f27075487b46aa03df219c362f69a453beebd22836809ef512268.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintJobRequestBuilder) Start()(*i8c1238ff81185e885512c301c4ea19545ac21eef777b326e5da6c9c10f97c437.StartRequestBuilder) {
    return i8c1238ff81185e885512c301c4ea19545ac21eef777b326e5da6c9c10f97c437.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
