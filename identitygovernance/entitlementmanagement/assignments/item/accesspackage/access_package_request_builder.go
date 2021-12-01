package accesspackage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3d848aae027f2dd182b9eef8787ec841cbfaa8bee6a735d90712b027430ff2f7 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/accesspackage/ref"
    ia56a44ad78235b271fe0cb87818ebd11cc11628da7581fe9e62134d66851ce58 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/assignments/item/accesspackage/getapplicablepolicyrequirements"
)

// AccessPackageRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\assignments\{accessPackageAssignment-id}\accessPackage
type AccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AccessPackageRequestBuilderGetOptions options for Get
type AccessPackageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AccessPackageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AccessPackageRequestBuilderGetQueryParameters read-only. Nullable.
type AccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    m := &AccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment_id}/accessPackage{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewAccessPackageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only. Nullable.
func (m *AccessPackageRequestBuilder) CreateGetRequestInformation(options *AccessPackageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only. Nullable.
func (m *AccessPackageRequestBuilder) Get(options *AccessPackageRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessPackage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAccessPackage() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AccessPackage), nil
}
func (m *AccessPackageRequestBuilder) GetApplicablePolicyRequirements()(*ia56a44ad78235b271fe0cb87818ebd11cc11628da7581fe9e62134d66851ce58.GetApplicablePolicyRequirementsRequestBuilder) {
    return ia56a44ad78235b271fe0cb87818ebd11cc11628da7581fe9e62134d66851ce58.NewGetApplicablePolicyRequirementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessPackageRequestBuilder) Ref()(*i3d848aae027f2dd182b9eef8787ec841cbfaa8bee6a735d90712b027430ff2f7.RefRequestBuilder) {
    return i3d848aae027f2dd182b9eef8787ec841cbfaa8bee6a735d90712b027430ff2f7.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
