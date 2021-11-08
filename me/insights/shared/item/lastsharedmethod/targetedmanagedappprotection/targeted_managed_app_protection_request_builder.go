package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    if565fe0a010186cffb7c0ed5a0a1eaebda62b57e0ada6af129b7fdf33c9e1aff "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/targetedmanagedappprotection/targetapps"
    ifd00ab38ca0b190c2bd53d8267259834d5a5821cbe6dcf7f9010c4fc5b0f3e04 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/shared/item/lastsharedmethod/targetedmanagedappprotection/assign"
)

// Builds and executes requests for operations under \me\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*ifd00ab38ca0b190c2bd53d8267259834d5a5821cbe6dcf7f9010c4fc5b0f3e04.AssignRequestBuilder) {
    return ifd00ab38ca0b190c2bd53d8267259834d5a5821cbe6dcf7f9010c4fc5b0f3e04.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*if565fe0a010186cffb7c0ed5a0a1eaebda62b57e0ada6af129b7fdf33c9e1aff.TargetAppsRequestBuilder) {
    return if565fe0a010186cffb7c0ed5a0a1eaebda62b57e0ada6af129b7fdf33c9e1aff.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
