package windowsinformationprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4945824f71056d5d89251fddd5ef4058d7d2319e5b2dc4b6ba9fd3c902bde3cf "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item/windowsinformationprotection/assign"
)

// WindowsInformationProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppPolicies\{managedAppPolicy-id}\microsoft.graph.windowsInformationProtection
type WindowsInformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Assign the assign property
func (m *WindowsInformationProtectionRequestBuilder) Assign()(*i4945824f71056d5d89251fddd5ef4058d7d2319e5b2dc4b6ba9fd3c902bde3cf.AssignRequestBuilder) {
    return i4945824f71056d5d89251fddd5ef4058d7d2319e5b2dc4b6ba9fd3c902bde3cf.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsInformationProtectionRequestBuilderInternal instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
func NewWindowsInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    m := &WindowsInformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppPolicies/{managedAppPolicy_id}/microsoft.graph.windowsInformationProtection";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsInformationProtectionRequestBuilder instantiates a new WindowsInformationProtectionRequestBuilder and sets the default values.
func NewWindowsInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsInformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
