package riskyusers

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/dismiss"
    ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/confirmcompromised"
    icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779 "github.com/microsoftgraph/msgraph-sdk-go/identityprotection/riskyusers/count"
)

// RiskyUsersRequestBuilder provides operations to manage the riskyUsers property of the microsoft.graph.identityProtectionRoot entity.
type RiskyUsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RiskyUsersRequestBuilderGetOptions options for Get
type RiskyUsersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RiskyUsersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RiskyUsersRequestBuilderGetQueryParameters users that are flagged as at-risk by Azure AD Identity Protection.
type RiskyUsersRequestBuilderGetQueryParameters struct {
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
// RiskyUsersRequestBuilderPostOptions options for Post
type RiskyUsersRequestBuilderPostOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUserable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RiskyUsersRequestBuilder) ConfirmCompromised()(*ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc.ConfirmCompromisedRequestBuilder) {
    return ic0a458e5f0f3b529bf41d03f94eacd03ae5394d2801818addefd2ef0ad4c34bc.NewConfirmCompromisedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRiskyUsersRequestBuilderInternal instantiates a new RiskyUsersRequestBuilder and sets the default values.
func NewRiskyUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyUsersRequestBuilder) {
    m := &RiskyUsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyUsers{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyUsersRequestBuilder instantiates a new RiskyUsersRequestBuilder and sets the default values.
func NewRiskyUsersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RiskyUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyUsersRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RiskyUsersRequestBuilder) Count()(*icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779.CountRequestBuilder) {
    return icf8431a0bea7943ea4c27e26967395f2983c85c27c9bf1afbd36da7444ddb779.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) CreateGetRequestInformation(options *RiskyUsersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) CreatePostRequestInformation(options *RiskyUsersRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *RiskyUsersRequestBuilder) Dismiss()(*i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e.DismissRequestBuilder) {
    return i5a8d5869e42d3c9c54ea08a341c5297a0b0865928dbe0569fabaee2c8e65008e.NewDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get users that are flagged as at-risk by Azure AD Identity Protection.
func (m *RiskyUsersRequestBuilder) Get(options *RiskyUsersRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUserCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateRiskyUserCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUserCollectionResponseable), nil
}
// Post create new navigation property to riskyUsers for identityProtection
func (m *RiskyUsersRequestBuilder) Post(options *RiskyUsersRequestBuilderPostOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUserable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateRiskyUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RiskyUserable), nil
}
