package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

// Builds and executes requests for operations under \deviceAppManagement\targetedManagedAppConfigurations\{targetedManagedAppConfiguration-id}
type TargetedManagedAppConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Targeted managed app configurations.
type TargetedManagedAppConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Apps()(*if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.AppsRequestBuilder) {
    return if5f3e25c7fb7036c0005255cf83f278d6dbdc54efe6806e0f73b94dcc8769f54.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.apps.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TargetedManagedAppConfigurationRequestBuilder) AppsById(id string)(*i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.ManagedMobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i0c58f5c03e1e1e1b14874efa5218a60bb1595c685ea667497c77cd45c45299ab.NewManagedMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assign()(*i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.AssignRequestBuilder) {
    return i1bc0e847e41ebdc918b382210b618f57eecf34594ab8b85ee5961792f938a28b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assignments()(*i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.AssignmentsRequestBuilder) {
    return i2e2aa2658be52cafac5b62880a98fc25ab296fd6408d08702a2e240e8e1ef1af.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *TargetedManagedAppConfigurationRequestBuilder) AssignmentsById(id string)(*iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.TargetedManagedAppPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment_id"] = id
    }
    return iae5cd0b807da9eeca622ea354242f60add20264ab977061a2f592e1e2bb4e2ef.NewTargetedManagedAppPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new TargetedManagedAppConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    m := &TargetedManagedAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TargetedManagedAppConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTargetedManagedAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Targeted managed app configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Targeted managed app configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *TargetedManagedAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TargetedManagedAppConfigurationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Targeted managed app configurations.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *TargetedManagedAppConfigurationRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Targeted managed app configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TargetedManagedAppConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) DeploymentSummary()(*i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.DeploymentSummaryRequestBuilder) {
    return i10df9e1d49ddd8e40cfbeb83f29b10a26d0a9bd66f497b5f5a6f0cac3ad0a238.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Targeted managed app configurations.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TargetedManagedAppConfigurationRequestBuilder) Get(q func (value *TargetedManagedAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTargetedManagedAppConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfiguration), nil
}
// Targeted managed app configurations.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *TargetedManagedAppConfigurationRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TargetedManagedAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) TargetApps()(*ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.TargetAppsRequestBuilder) {
    return ibacc3af011a2e5c778e54915b40576462c31c68b752e932193de352b09cc3d55.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
