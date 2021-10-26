package managedappregistrations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9 "github.com/microsoftgraph/msgraph-sdk-go/users/item/managedappregistrations/ref"
    id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc "github.com/microsoftgraph/msgraph-sdk-go/users/item/managedappregistrations/getuseridswithflaggedappregistration"
)

// Builds and executes requests for operations under \users\{user-id}\managedAppRegistrations
type ManagedAppRegistrationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Zero or more managed app registrations that belong to the user.
type ManagedAppRegistrationsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
// Instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    m := &ManagedAppRegistrationsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/managedAppRegistrations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppRegistrationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Zero or more managed app registrations that belong to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ManagedAppRegistrationsRequestBuilder) CreateGetRequestInformation(q func (value *ManagedAppRegistrationsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagedAppRegistrationsRequestBuilderGetQueryParameters)
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
// Zero or more managed app registrations that belong to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ManagedAppRegistrationsRequestBuilder) Get(q func (value *ManagedAppRegistrationsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ManagedAppRegistrationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistrationsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedAppRegistrationsResponse), nil
}
// Builds and executes requests for operations under \users\{user-id}\managedAppRegistrations\microsoft.graph.getUserIdsWithFlaggedAppRegistration()
func (m *ManagedAppRegistrationsRequestBuilder) GetUserIdsWithFlaggedAppRegistration()(*id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc.GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    return id0044acc4fa27eb6383b5a73b5f9e0bafc65844f934e448ac3d3cbf1c69e3dcc.NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationsRequestBuilder) Ref()(*i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9.RefRequestBuilder) {
    return i7e975f33d4b01e27c6cd3625ae8f7b30dc51d547286f10ad29474b734ae348d9.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
