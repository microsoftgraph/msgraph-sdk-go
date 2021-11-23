package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/targetedmanagedappprotection/assign"
    ibd4538acea21b5ba8dd7f77ee1abdee1c11946ee780f11d87bda5666fe398f10 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/trending/item/resource/targetedmanagedappprotection/targetapps"
)

// targetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \me\insights\trending\{trending-id}\resource\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111.AssignRequestBuilder) {
    return i3272a439ec24aa01bb0c5c5eab90b0412bea2231615fa2033f9ed7c454b9a111.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/trending/{trending_id}/resource/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppProtectionRequestBuilder instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*ibd4538acea21b5ba8dd7f77ee1abdee1c11946ee780f11d87bda5666fe398f10.TargetAppsRequestBuilder) {
    return ibd4538acea21b5ba8dd7f77ee1abdee1c11946ee780f11d87bda5666fe398f10.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
