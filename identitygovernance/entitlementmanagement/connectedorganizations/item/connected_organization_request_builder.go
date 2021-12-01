package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors"
    i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8 "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors"
    i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/externalsponsors/item"
    i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance/entitlementmanagement/connectedorganizations/item/internalsponsors/item"
)

// ConnectedOrganizationRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\connectedOrganizations\{connectedOrganization-id}
type ConnectedOrganizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ConnectedOrganizationRequestBuilderDeleteOptions options for Delete
type ConnectedOrganizationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectedOrganizationRequestBuilderGetOptions options for Get
type ConnectedOrganizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConnectedOrganizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectedOrganizationRequestBuilderGetQueryParameters get connectedOrganizations from identityGovernance
type ConnectedOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ConnectedOrganizationRequestBuilderPatchOptions options for Patch
type ConnectedOrganizationRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConnectedOrganization;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewConnectedOrganizationRequestBuilderInternal instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewConnectedOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectedOrganizationRequestBuilder) {
    m := &ConnectedOrganizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/connectedOrganizations/{connectedOrganization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectedOrganizationRequestBuilder instantiates a new ConnectedOrganizationRequestBuilder and sets the default values.
func NewConnectedOrganizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectedOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectedOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationRequestBuilder) CreateDeleteRequestInformation(options *ConnectedOrganizationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get connectedOrganizations from identityGovernance
func (m *ConnectedOrganizationRequestBuilder) CreateGetRequestInformation(options *ConnectedOrganizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationRequestBuilder) CreatePatchRequestInformation(options *ConnectedOrganizationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property connectedOrganizations for identityGovernance
func (m *ConnectedOrganizationRequestBuilder) Delete(options *ConnectedOrganizationRequestBuilderDeleteOptions)(error) {
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
func (m *ConnectedOrganizationRequestBuilder) ExternalSponsors()(*i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6.ExternalSponsorsRequestBuilder) {
    return i46b288d9d8f774d7fa38d373a828c5bc7cf81c90735be3014be7f4d14860c0a6.NewExternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.externalSponsors.item collection
func (m *ConnectedOrganizationRequestBuilder) ExternalSponsorsById(id string)(*i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a.DirectoryObjectRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i38c48f98eb917d486190d72c8f1c4ae0a96a04a29df568065e9cdfbc7a1bcc7a.NewDirectoryObjectRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get connectedOrganizations from identityGovernance
func (m *ConnectedOrganizationRequestBuilder) Get(options *ConnectedOrganizationRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConnectedOrganization, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewConnectedOrganization() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ConnectedOrganization), nil
}
func (m *ConnectedOrganizationRequestBuilder) InternalSponsors()(*i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8.InternalSponsorsRequestBuilder) {
    return i6f5b9f8ecd653f7004e3a87b6fd9f1edd541123ee01576fc601583fd7a8ee2f8.NewInternalSponsorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InternalSponsorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.identityGovernance.entitlementManagement.connectedOrganizations.item.internalSponsors.item collection
func (m *ConnectedOrganizationRequestBuilder) InternalSponsorsById(id string)(*i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa.DirectoryObjectRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i6a8a0317ba69a6948d7f14baa6206d53d8ce7fc2cbce939e0ac4ece8b4a838aa.NewDirectoryObjectRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property connectedOrganizations in identityGovernance
func (m *ConnectedOrganizationRequestBuilder) Patch(options *ConnectedOrganizationRequestBuilderPatchOptions)(error) {
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
