package windowsinformationprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ia3c09fae49593008a78467a5ab1287391d45f11f4b32c1969eb55b2ec24a77f3 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/windowsinformationprotection/assign"
)

// Builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.windowsInformationProtection
type WindowsInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *WindowsInformationProtectionRequestBuilder) Assign()(*ia3c09fae49593008a78467a5ab1287391d45f11f4b32c1969eb55b2ec24a77f3.AssignRequestBuilder) {
    return ia3c09fae49593008a78467a5ab1287391d45f11f4b32c1969eb55b2ec24a77f3.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWindowsInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    m := &WindowsInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/insights/trending/{trending_id}/resource/microsoft.graph.windowsInformationProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewWindowsInformationProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
