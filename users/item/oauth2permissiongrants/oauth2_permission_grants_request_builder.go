package oauth2permissiongrants

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i73bd1407e8368bdd39a94e2c0e4537b257fd49dfffbb96281fc96a5c59cbbf17 "github.com/microsoftgraph/msgraph-sdk-go/users/item/oauth2permissiongrants/ref"
)

// Oauth2PermissionGrantsRequestBuilder builds and executes requests for operations under \users\{user-id}\oauth2PermissionGrants
type Oauth2PermissionGrantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Oauth2PermissionGrantsRequestBuilderGetOptions options for Get
type Oauth2PermissionGrantsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *Oauth2PermissionGrantsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Oauth2PermissionGrantsRequestBuilderGetQueryParameters get oauth2PermissionGrants from users
type Oauth2PermissionGrantsRequestBuilderGetQueryParameters struct {
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
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewOauth2PermissionGrantsRequestBuilderInternal instantiates a new Oauth2PermissionGrantsRequestBuilder and sets the default values.
func NewOauth2PermissionGrantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*Oauth2PermissionGrantsRequestBuilder) {
    m := &Oauth2PermissionGrantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/oauth2PermissionGrants{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOauth2PermissionGrantsRequestBuilder instantiates a new Oauth2PermissionGrantsRequestBuilder and sets the default values.
func NewOauth2PermissionGrantsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*Oauth2PermissionGrantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOauth2PermissionGrantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get oauth2PermissionGrants from users
func (m *Oauth2PermissionGrantsRequestBuilder) CreateGetRequestInformation(options *Oauth2PermissionGrantsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get oauth2PermissionGrants from users
func (m *Oauth2PermissionGrantsRequestBuilder) Get(options *Oauth2PermissionGrantsRequestBuilderGetOptions)(*Oauth2PermissionGrantsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOauth2PermissionGrantsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*Oauth2PermissionGrantsResponse), nil
}
func (m *Oauth2PermissionGrantsRequestBuilder) Ref()(*i73bd1407e8368bdd39a94e2c0e4537b257fd49dfffbb96281fc96a5c59cbbf17.RefRequestBuilder) {
    return i73bd1407e8368bdd39a94e2c0e4537b257fd49dfffbb96281fc96a5c59cbbf17.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
