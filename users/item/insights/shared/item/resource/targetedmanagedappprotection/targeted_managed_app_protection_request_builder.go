package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i38ea2f6e29ce6076b66b0dc502c0884bfe5806b88791136d36d6bcfecf72ee3b "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/targetedmanagedappprotection/targetapps"
    i613774739b5a2db00c3d21802a28927d9dad6cc5d965dc3e396fe7b7d0af3595 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/resource/targetedmanagedappprotection/assign"
)

// Builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\resource\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i613774739b5a2db00c3d21802a28927d9dad6cc5d965dc3e396fe7b7d0af3595.AssignRequestBuilder) {
    return i613774739b5a2db00c3d21802a28927d9dad6cc5d965dc3e396fe7b7d0af3595.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/insights/shared/{sharedInsight_id}/resource/microsoft.graph.targetedManagedAppProtection";
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
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i38ea2f6e29ce6076b66b0dc502c0884bfe5806b88791136d36d6bcfecf72ee3b.TargetAppsRequestBuilder) {
    return i38ea2f6e29ce6076b66b0dc502c0884bfe5806b88791136d36d6bcfecf72ee3b.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
