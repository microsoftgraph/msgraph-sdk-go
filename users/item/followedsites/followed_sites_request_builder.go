package followedsites

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6679b1724d9da04ca2cfb8a7899bbc65aae875e9d06f5baa6e12c779636563db "github.com/microsoftgraph/msgraph-sdk-go/users/item/followedsites/ref"
)

// FollowedSitesRequestBuilder builds and executes requests for operations under \users\{user-id}\followedSites
type FollowedSitesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FollowedSitesRequestBuilderGetOptions options for Get
type FollowedSitesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *FollowedSitesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FollowedSitesRequestBuilderGetQueryParameters get followedSites from users
type FollowedSitesRequestBuilderGetQueryParameters struct {
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
// NewFollowedSitesRequestBuilderInternal instantiates a new FollowedSitesRequestBuilder and sets the default values.
func NewFollowedSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FollowedSitesRequestBuilder) {
    m := &FollowedSitesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/followedSites{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFollowedSitesRequestBuilder instantiates a new FollowedSitesRequestBuilder and sets the default values.
func NewFollowedSitesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FollowedSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFollowedSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get followedSites from users
func (m *FollowedSitesRequestBuilder) CreateGetRequestInformation(options *FollowedSitesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get followedSites from users
func (m *FollowedSitesRequestBuilder) Get(options *FollowedSitesRequestBuilderGetOptions)(*FollowedSitesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFollowedSitesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*FollowedSitesResponse), nil
}
func (m *FollowedSitesRequestBuilder) Ref()(*i6679b1724d9da04ca2cfb8a7899bbc65aae875e9d06f5baa6e12c779636563db.RefRequestBuilder) {
    return i6679b1724d9da04ca2cfb8a7899bbc65aae875e9d06f5baa6e12c779636563db.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
