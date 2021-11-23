package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies"
    i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/operations"
    ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies"
    i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/operations/item"
    i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/intendedpolicies/item"
    i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item/appliedpolicies/item"
)

// ManagedAppRegistrationRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedAppRegistrations\{managedAppRegistration-id}
type ManagedAppRegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppRegistrationRequestBuilderDeleteOptions options for Delete
type ManagedAppRegistrationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppRegistrationRequestBuilderGetOptions options for Get
type ManagedAppRegistrationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppRegistrationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppRegistrationRequestBuilderGetQueryParameters the managed app registrations.
type ManagedAppRegistrationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// ManagedAppRegistrationRequestBuilderPatchOptions options for Patch
type ManagedAppRegistrationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppRegistration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedAppRegistrationRequestBuilder) AppliedPolicies()(*ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304.AppliedPoliciesRequestBuilder) {
    return ifa8f167e8cac58c1a5037bc99e8f6f25e58a3455715484a9419cebb030ee4304.NewAppliedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppliedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.appliedPolicies.item collection
func (m *ManagedAppRegistrationRequestBuilder) AppliedPoliciesById(id string)(*i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i8c014196246f79aab4e3ce9c3fa6da286aeeb8d4af36d5f9b7fe584d08030fe6.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedAppRegistrationRequestBuilderInternal instantiates a new ManagedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    m := &ManagedAppRegistrationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationRequestBuilder instantiates a new ManagedAppRegistrationRequestBuilder and sets the default values.
func NewManagedAppRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreateDeleteRequestInformation(options *ManagedAppRegistrationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreateGetRequestInformation(options *ManagedAppRegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) CreatePatchRequestInformation(options *ManagedAppRegistrationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Delete(options *ManagedAppRegistrationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Get(options *ManagedAppRegistrationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppRegistration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedAppRegistration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedAppRegistration), nil
}
func (m *ManagedAppRegistrationRequestBuilder) IntendedPolicies()(*i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d.IntendedPoliciesRequestBuilder) {
    return i3eed464a4a4f7a056ae78f7d5084484b8e4c7071ced5cd16d8978c75a211418d.NewIntendedPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntendedPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.intendedPolicies.item collection
func (m *ManagedAppRegistrationRequestBuilder) IntendedPoliciesById(id string)(*i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i0d4a89b01c459a241cf1b150996622b188ac6eefcbb7901958e20402e3fbda2f.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ManagedAppRegistrationRequestBuilder) Operations()(*i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0.OperationsRequestBuilder) {
    return i798793d33f3b0c4349f4a5beaee290369942f7f1d78a4207726bf30670bbf0d0.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item.operations.item collection
func (m *ManagedAppRegistrationRequestBuilder) OperationsById(id string)(*i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f.ManagedAppOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppOperation_id"] = id
    }
    return i0a22e7eeb322ca1994c33c38e16e5fe50b47a7810496c52fd2e2c94195f6c08f.NewManagedAppOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the managed app registrations.
func (m *ManagedAppRegistrationRequestBuilder) Patch(options *ManagedAppRegistrationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
