package targetedmanagedappprotection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/targetapps"
    ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/assign"
)

// Builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b.AssignRequestBuilder) {
    return ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}/microsoft.graph.targetedManagedAppProtection";
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
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446.TargetAppsRequestBuilder) {
    return id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
