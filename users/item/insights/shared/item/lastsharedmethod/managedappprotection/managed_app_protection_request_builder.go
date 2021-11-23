package managedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i27de29cb6289aef43085ec2740fe5e6e950db4a6cedf48c5811ea70af1190300 "github.com/microsoftgraph/msgraph-sdk-go/users/item/insights/shared/item/lastsharedmethod/managedappprotection/targetapps"
)

// managedAppProtectionRequestBuilder builds and executes requests for operations under \users\{user-id}\insights\shared\{sharedInsight-id}\lastSharedMethod\microsoft.graph.managedAppProtection
type ManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewManagedAppProtectionRequestBuilderInternal instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    m := &ManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/insights/shared/{sharedInsight_id}/lastSharedMethod/microsoft.graph.managedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppProtectionRequestBuilder instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagedAppProtectionRequestBuilder) TargetApps()(*i27de29cb6289aef43085ec2740fe5e6e950db4a6cedf48c5811ea70af1190300.TargetAppsRequestBuilder) {
    return i27de29cb6289aef43085ec2740fe5e6e950db4a6cedf48c5811ea70af1190300.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
