package managedappregistrations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/managedappregistrations/ref"
    id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc "github.com/microsoftgraph/msgraph-sdk-go/users/item/managedappregistrations/getuseridswithflaggedappregistration"
)

// ManagedAppRegistrationsRequestBuilder builds and executes requests for operations under \users\{user-id}\managedAppRegistrations
type ManagedAppRegistrationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedAppRegistrationsRequestBuilderGetOptions options for Get
type ManagedAppRegistrationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppRegistrationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedAppRegistrationsRequestBuilderGetQueryParameters zero or more managed app registrations that belong to the user.
type ManagedAppRegistrationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewManagedAppRegistrationsRequestBuilderInternal instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    m := &ManagedAppRegistrationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedAppRegistrations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedAppRegistrationsRequestBuilder instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
func NewManagedAppRegistrationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation zero or more managed app registrations that belong to the user.
func (m *ManagedAppRegistrationsRequestBuilder) CreateGetRequestInformation(options *ManagedAppRegistrationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get zero or more managed app registrations that belong to the user.
func (m *ManagedAppRegistrationsRequestBuilder) Get(options *ManagedAppRegistrationsRequestBuilderGetOptions)(*ManagedAppRegistrationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistrationsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedAppRegistrationsResponse), nil
}
// GetUserIdsWithFlaggedAppRegistration builds and executes requests for operations under \users\{user-id}\managedAppRegistrations\microsoft.graph.getUserIdsWithFlaggedAppRegistration()
func (m *ManagedAppRegistrationsRequestBuilder) GetUserIdsWithFlaggedAppRegistration()(*id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc.GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    return id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc.NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationsRequestBuilder) Ref()(*i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9.RefRequestBuilder) {
    return i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
