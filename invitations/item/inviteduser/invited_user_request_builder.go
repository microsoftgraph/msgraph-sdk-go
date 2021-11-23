package inviteduser

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i89266147459be47604a3bc71c8b8a481a9f0e411a5aa2e53b24d4a9ffdea1e5d "github.com/microsoftgraph/msgraph-sdk-go/invitations/item/inviteduser/ref"
)

// InvitedUserRequestBuilder builds and executes requests for operations under \invitations\{invitation-id}\invitedUser
type InvitedUserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InvitedUserRequestBuilderGetOptions options for Get
type InvitedUserRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InvitedUserRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InvitedUserRequestBuilderGetQueryParameters the user created as part of the invitation creation. Read-Only
type InvitedUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// NewInvitedUserRequestBuilderInternal instantiates a new InvitedUserRequestBuilder and sets the default values.
func NewInvitedUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InvitedUserRequestBuilder) {
    m := &InvitedUserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/invitations/{invitation_id}/invitedUser{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInvitedUserRequestBuilder instantiates a new InvitedUserRequestBuilder and sets the default values.
func NewInvitedUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InvitedUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInvitedUserRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user created as part of the invitation creation. Read-Only
func (m *InvitedUserRequestBuilder) CreateGetRequestInformation(options *InvitedUserRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the user created as part of the invitation creation. Read-Only
func (m *InvitedUserRequestBuilder) Get(options *InvitedUserRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewUser() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.User), nil
}
func (m *InvitedUserRequestBuilder) Ref()(*i89266147459be47604a3bc71c8b8a481a9f0e411a5aa2e53b24d4a9ffdea1e5d.RefRequestBuilder) {
    return i89266147459be47604a3bc71c8b8a481a9f0e411a5aa2e53b24d4a9ffdea1e5d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
