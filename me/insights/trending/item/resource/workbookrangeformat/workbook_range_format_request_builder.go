package workbookrangeformat

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeformat/autofitrows"
    ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/workbookrangeformat/autofitcolumns"
)

type WorkbookRangeFormatRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitColumns()(*ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217.AutofitColumnsRequestBuilder) {
    return ic0cee953dd5c4427d3739f2105b634c597ddb28eba623f6eb1c3ff01e93d0217.NewAutofitColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookRangeFormatRequestBuilder) AutofitRows()(*i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea.AutofitRowsRequestBuilder) {
    return i9993bf75fcfbd53708335879313d82b8bcacfca709dc4decc6bc84a9d63e3fea.NewAutofitRowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookRangeFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookRangeFormatRequestBuilder) {
    m := &WorkbookRangeFormatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.workbookRangeFormat";
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
