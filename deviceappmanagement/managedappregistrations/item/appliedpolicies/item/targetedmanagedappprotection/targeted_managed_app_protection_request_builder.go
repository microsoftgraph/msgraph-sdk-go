package targetedmanagedappprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5ecc9e7c112c0c20d1ba56137b27962b7efb5f39d04ba4e9628e4432a0c6186a "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/targetedmanagedappprotection/targetapps"
    i849a70182732d094bef1e9f964a0263760b815bef1cf1a3b5ce26277c7c88fda "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/targetedmanagedappprotection/assign"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\appliedPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Assign the assign property
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*i849a70182732d094bef1e9f964a0263760b815bef1cf1a3b5ce26277c7c88fda.AssignRequestBuilder) {
    return i849a70182732d094bef1e9f964a0263760b815bef1cf1a3b5ce26277c7c88fda.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/appliedPolicies/{managedAppPolicy_id}/microsoft.graph.targetedManagedAppProtection";
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
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*i5ecc9e7c112c0c20d1ba56137b27962b7efb5f39d04ba4e9628e4432a0c6186a.TargetAppsRequestBuilder) {
    return i5ecc9e7c112c0c20d1ba56137b27962b7efb5f39d04ba4e9628e4432a0c6186a.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
