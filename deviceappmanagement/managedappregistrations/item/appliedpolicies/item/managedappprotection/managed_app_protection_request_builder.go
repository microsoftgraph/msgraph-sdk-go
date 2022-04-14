package managedappprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ic9ec1fdab3ac558677ef9b5eacb630506937d9a4f57797c67241d25c130c1add "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item/managedappprotection/targetapps"
)

// ManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}\appliedPolicies\{managedAppPolicy-id}\microsoft.graph.managedAppProtection
type ManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewManagedAppProtectionRequestBuilderInternal instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    m := &ManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/appliedPolicies/{managedAppPolicy%2Did}/microsoft.graph.managedAppProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppProtectionRequestBuilder instantiates a new ManagedAppProtectionRequestBuilder and sets the default values.
func NewManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// TargetApps the targetApps property
func (m *ManagedAppProtectionRequestBuilder) TargetApps()(*ic9ec1fdab3ac558677ef9b5eacb630506937d9a4f57797c67241d25c130c1add.TargetAppsRequestBuilder) {
    return ic9ec1fdab3ac558677ef9b5eacb630506937d9a4f57797c67241d25c130c1add.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
