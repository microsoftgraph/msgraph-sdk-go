package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idc168e9231263e2c311e55fbc5bfc33acf48a5cb2ffb1c6ef2f7aae7f140d850 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/targetedmanagedappprotection/assign"
    if03fbf2512be788b170c83d65cb8357b1ed434919d409e0ec39fe1fbc9b2a915 "github.com/microsoftgraph/msgraph-sdk-go/me/insights/used/item/resource/targetedmanagedappprotection/targetapps"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \me\insights\used\{usedInsight-id}\resource\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*idc168e9231263e2c311e55fbc5bfc33acf48a5cb2ffb1c6ef2f7aae7f140d850.AssignRequestBuilder) {
    return idc168e9231263e2c311e55fbc5bfc33acf48a5cb2ffb1c6ef2f7aae7f140d850.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/insights/used/{usedInsight_id}/resource/microsoft.graph.targetedManagedAppProtection";
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
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*if03fbf2512be788b170c83d65cb8357b1ed434919d409e0ec39fe1fbc9b2a915.TargetAppsRequestBuilder) {
    return if03fbf2512be788b170c83d65cb8357b1ed434919d409e0ec39fe1fbc9b2a915.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
