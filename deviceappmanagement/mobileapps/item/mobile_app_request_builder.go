package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assignments"
    i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/categories"
    ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assign"
    i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item/assignments/item"
)

// MobileAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}
type MobileAppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MobileAppRequestBuilderDeleteOptions options for Delete
type MobileAppRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppRequestBuilderGetOptions options for Get
type MobileAppRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MobileAppRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MobileAppRequestBuilderGetQueryParameters the mobile apps.
type MobileAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MobileAppRequestBuilderPatchOptions options for Patch
type MobileAppRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MobileApp;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MobileAppRequestBuilder) Assign()(*ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2.AssignRequestBuilder) {
    return ia9cfc52616110c9bb721b0f95f47d861e5afa2655ecff28a0fc66beb7ba78fe2.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Assignments()(*i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975.AssignmentsRequestBuilder) {
    return i41e6a4f5d8af8cbec70d9d8c4f48278ff47465e70b23d4ee2bbdcafa93bae975.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileApps.item.assignments.item collection
func (m *MobileAppRequestBuilder) AssignmentsById(id string)(*i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1.MobileAppAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppAssignment_id"] = id
    }
    return i1de44916490eb612ae701336a2eae5f89d836e0eaae6829aa711a57c559453e1.NewMobileAppAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MobileAppRequestBuilder) Categories()(*i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9.CategoriesRequestBuilder) {
    return i52247cde4874402d4eaf161fa2c9ec2cc4a9257f52df96b9bbbcdf47e1f167f9.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMobileAppRequestBuilderInternal instantiates a new MobileAppRequestBuilder and sets the default values.
func NewMobileAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppRequestBuilder) {
    m := &MobileAppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppRequestBuilder instantiates a new MobileAppRequestBuilder and sets the default values.
func NewMobileAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MobileAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the mobile apps.
func (m *MobileAppRequestBuilder) CreateDeleteRequestInformation(options *MobileAppRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the mobile apps.
func (m *MobileAppRequestBuilder) CreateGetRequestInformation(options *MobileAppRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the mobile apps.
func (m *MobileAppRequestBuilder) CreatePatchRequestInformation(options *MobileAppRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the mobile apps.
func (m *MobileAppRequestBuilder) Delete(options *MobileAppRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the mobile apps.
func (m *MobileAppRequestBuilder) Get(options *MobileAppRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MobileApp, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMobileApp() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MobileApp), nil
}
// Patch the mobile apps.
func (m *MobileAppRequestBuilder) Patch(options *MobileAppRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
