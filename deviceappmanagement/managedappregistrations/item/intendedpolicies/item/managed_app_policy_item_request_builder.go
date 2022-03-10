package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i243d8966cf38814c6b048ef9b3f5af1da82f45a88be962698bd507cef3f8f1d7 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetedmanagedappprotection"
    i3216bc8e50c50cd971b05bb4b9bbf372c09864803a7b408274e09d06cc235ba3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/windowsinformationprotection"
    i61b3e44e6f75f54ccdfc163d68a0fec47a9154658fb6ce84fb7da4896a9d8636 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/managedappprotection"
    ie41ce94cd59afd0132802c8e484fc0404ee07b13babb164aee36865069d56a6d "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item/targetapps"
)

// ManagedAppPolicyItemRequestBuilder provides operations to manage the intendedPolicies property of the microsoft.graph.managedAppRegistration entity.
type ManagedAppPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppPolicyItemRequestBuilderDeleteOptions options for Delete
type ManagedAppPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyItemRequestBuilderGetOptions options for Get
type ManagedAppPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppPolicyItemRequestBuilderGetQueryParameters zero or more policies admin intended for the app as of now.
type ManagedAppPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedAppPolicyItemRequestBuilderPatchOptions options for Patch
type ManagedAppPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewManagedAppPolicyItemRequestBuilderInternal instantiates a new ManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedAppPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyItemRequestBuilder) {
    m := &ManagedAppPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}/intendedPolicies/{managedAppPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppPolicyItemRequestBuilder instantiates a new ManagedAppPolicyItemRequestBuilder and sets the default values.
func NewManagedAppPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property intendedPolicies for deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyItemRequestBuilder) CreateGetRequestInformation(options *ManagedAppPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property intendedPolicies in deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) CreatePatchRequestInformation(options *ManagedAppPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property intendedPolicies for deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) Delete(options *ManagedAppPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get zero or more policies admin intended for the app as of now.
func (m *ManagedAppPolicyItemRequestBuilder) Get(options *ManagedAppPolicyItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateManagedAppPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppPolicyable), nil
}
func (m *ManagedAppPolicyItemRequestBuilder) ManagedAppProtection()(*i61b3e44e6f75f54ccdfc163d68a0fec47a9154658fb6ce84fb7da4896a9d8636.ManagedAppProtectionRequestBuilder) {
    return i61b3e44e6f75f54ccdfc163d68a0fec47a9154658fb6ce84fb7da4896a9d8636.NewManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property intendedPolicies in deviceAppManagement
func (m *ManagedAppPolicyItemRequestBuilder) Patch(options *ManagedAppPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ManagedAppPolicyItemRequestBuilder) TargetApps()(*ie41ce94cd59afd0132802c8e484fc0404ee07b13babb164aee36865069d56a6d.TargetAppsRequestBuilder) {
    return ie41ce94cd59afd0132802c8e484fc0404ee07b13babb164aee36865069d56a6d.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyItemRequestBuilder) TargetedManagedAppProtection()(*i243d8966cf38814c6b048ef9b3f5af1da82f45a88be962698bd507cef3f8f1d7.TargetedManagedAppProtectionRequestBuilder) {
    return i243d8966cf38814c6b048ef9b3f5af1da82f45a88be962698bd507cef3f8f1d7.NewTargetedManagedAppProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppPolicyItemRequestBuilder) WindowsInformationProtection()(*i3216bc8e50c50cd971b05bb4b9bbf372c09864803a7b408274e09d06cc235ba3.WindowsInformationProtectionRequestBuilder) {
    return i3216bc8e50c50cd971b05bb4b9bbf372c09864803a7b408274e09d06cc235ba3.NewWindowsInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
