package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

// TargetedManagedAppConfigurationItemRequestBuilder provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type TargetedManagedAppConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions options for Delete
type TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TargetedManagedAppConfigurationItemRequestBuilderGetOptions options for Get
type TargetedManagedAppConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters targeted managed app configurations.
type TargetedManagedAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TargetedManagedAppConfigurationItemRequestBuilderPatchOptions options for Patch
type TargetedManagedAppConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Apps()(*if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.AppsRequestBuilder) {
    return if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.apps.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AppsById(id string)(*i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.ManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.NewManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assign()(*i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.AssignRequestBuilder) {
    return i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Assignments()(*i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.AssignmentsRequestBuilder) {
    return i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.assignments.item collection
func (m *TargetedManagedAppConfigurationItemRequestBuilder) AssignmentsById(id string)(*iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.TargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment_id"] = id
    }
    return iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.NewTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTargetedManagedAppConfigurationItemRequestBuilderInternal instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    m := &TargetedManagedAppConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTargetedManagedAppConfigurationItemRequestBuilder instantiates a new TargetedManagedAppConfigurationItemRequestBuilder and sets the default values.
func NewTargetedManagedAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreateGetRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property targetedManagedAppConfigurations for deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Delete(options *TargetedManagedAppConfigurationItemRequestBuilderDeleteOptions)(error) {
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
func (m *TargetedManagedAppConfigurationItemRequestBuilder) DeploymentSummary()(*i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.DeploymentSummaryRequestBuilder) {
    return i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get targeted managed app configurations.
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Get(options *TargetedManagedAppConfigurationItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTargetedManagedAppConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfigurationable), nil
}
// Patch update the navigation property targetedManagedAppConfigurations in deviceAppManagement
func (m *TargetedManagedAppConfigurationItemRequestBuilder) Patch(options *TargetedManagedAppConfigurationItemRequestBuilderPatchOptions)(error) {
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
func (m *TargetedManagedAppConfigurationItemRequestBuilder) TargetApps()(*ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.TargetAppsRequestBuilder) {
    return ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
