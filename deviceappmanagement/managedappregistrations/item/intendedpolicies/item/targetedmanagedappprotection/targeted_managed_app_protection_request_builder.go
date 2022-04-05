package targetedmanagedappprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i1b1d2009b45f35e9541da2124133967092be5fce43613a9944c97592b141204f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection/assign"
    i3a69285403c1034683cbb302413a5e479732d4301b178d825d75f71b24aab025 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection/targetapps"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\intendedPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Assign the assign property
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i1b1d2009b45f35e9541da2124133967092be5fce43613a9944c97592b141204f.AssignRequestBuilder) {
    return i1b1d2009b45f35e9541da2124133967092be5fce43613a9944c97592b141204f.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/intendedPolicies/{managedAppPolicy_id}/microsoft.graph.targetedManagedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppProtectionRequestBuilder instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// TargetApps the targetApps property
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i3a69285403c1034683cbb302413a5e479732d4301b178d825d75f71b24aab025.TargetAppsRequestBuilder) {
    return i3a69285403c1034683cbb302413a5e479732d4301b178d825d75f71b24aab025.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
