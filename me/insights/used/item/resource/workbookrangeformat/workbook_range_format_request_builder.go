package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i23234a17861c49e8aaaf8d9dd799b4e588196f6f2c0c54c72901f2e70e93eacd "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangeformat/autofitcolumns"
    ie51f6eaf3322f04272ba503f4fdac7aa0e94f9cb0e83f6be7d022757848e95b3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/workbookrangeformat/autofitrows"
)

type WorkbookRangeFormatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*i23234a17861c49e8aaaf8d9dd799b4e588196f6f2c0c54c72901f2e70e93eacd.AutofitColumnsRequestBuilder) {
    return i23234a17861c49e8aaaf8d9dd799b4e588196f6f2c0c54c72901f2e70e93eacd.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*ie51f6eaf3322f04272ba503f4fdac7aa0e94f9cb0e83f6be7d022757848e95b3.AutofitRowsRequestBuilder) {
    return ie51f6eaf3322f04272ba503f4fdac7aa0e94f9cb0e83f6be7d022757848e95b3.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/used/{usedInsight_id}/resource/microsoft.graph.workbookRangeFormat";
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
func NewWorkbookRangeFormatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookRangeFormatRequestBuilderInternal(urlParams, requestAdapter)
}
