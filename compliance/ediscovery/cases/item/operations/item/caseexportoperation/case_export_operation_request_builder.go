package caseexportoperation

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ib30b530beedf01f2e3a045161ede6fef6e7772f37e2234e3c94de0856b8f5275 "github.com/microsoftgraph/msgraph-sdk-go/compliance/ediscovery/cases/item/operations/item/caseexportoperation/getdownloadurl"
)

// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\operations\{caseOperation-id}\microsoft.graph.ediscovery.caseExportOperation
type CaseExportOperationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Instantiates a new CaseExportOperationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseExportOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    m := &CaseExportOperationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/operations/{caseOperation_id}/microsoft.graph.ediscovery.caseExportOperation";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CaseExportOperationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCaseExportOperationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CaseExportOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCaseExportOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\operations\{caseOperation-id}\microsoft.graph.ediscovery.caseExportOperation\microsoft.graph.ediscovery.getDownloadUrl()
func (m *CaseExportOperationRequestBuilder) GetDownloadUrl()(*ib30b530beedf01f2e3a045161ede6fef6e7772f37e2234e3c94de0856b8f5275.GetDownloadUrlRequestBuilder) {
    return ib30b530beedf01f2e3a045161ede6fef6e7772f37e2234e3c94de0856b8f5275.NewGetDownloadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
