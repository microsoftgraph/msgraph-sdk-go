package targetedmanagedappprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/targetapps"
    ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/targetedmanagedappprotection/assign"
)

// TargetedManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}\microsoft.graph.targetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Assign the assign property
func (m *TargetedManagedAppProtectionRequestBuilder) Assign()(*ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b.AssignRequestBuilder) {
    return ied40a1921a418126b1497588d6e60eb5122aea6a506e3315cae486375479574b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTargetedManagedAppProtectionRequestBuilderInternal instantiates a new TargetedManagedAppProtectionRequestBuilder and sets the default values.
func NewTargetedManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedManagedAppProtectionRequestBuilder) {
    m := &TargetedManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy%2Did}/microsoft.graph.targetedManagedAppProtection";
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
func (m *TargetedManagedAppProtectionRequestBuilder) TargetApps()(*id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446.TargetAppsRequestBuilder) {
    return id1b02e4dee553990e9f1801893643dbc07a99861f9f8aee255af66881d72f446.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
