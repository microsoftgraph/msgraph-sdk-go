package rejectedsenders

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4e1f588b74fcc251466a08810fb8a3efa51152db86379c9d3ef91dbaf3f02bf9 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/rejectedsenders/ref"
)

// rejectedSendersRequestBuilder builds and executes requests for operations under \groups\{group-id}\rejectedSenders
type RejectedSendersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RejectedSendersRequestBuilderGetOptions options for Get
type RejectedSendersRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RejectedSendersRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// rejectedSendersRequestBuilderGetQueryParameters the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
type RejectedSendersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewRejectedSendersRequestBuilderInternal instantiates a new RejectedSendersRequestBuilder and sets the default values.
func NewRejectedSendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RejectedSendersRequestBuilder) {
    m := &RejectedSendersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/rejectedSenders{?top,skip,filter,count,orderby,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRejectedSendersRequestBuilder instantiates a new RejectedSendersRequestBuilder and sets the default values.
func NewRejectedSendersRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RejectedSendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRejectedSendersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *RejectedSendersRequestBuilder) CreateGetRequestInformation(options *RejectedSendersRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *RejectedSendersRequestBuilder) Get(options *RejectedSendersRequestBuilderGetOptions)(*RejectedSendersResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRejectedSendersResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*RejectedSendersResponse), nil
}
func (m *RejectedSendersRequestBuilder) Ref()(*i4e1f588b74fcc251466a08810fb8a3efa51152db86379c9d3ef91dbaf3f02bf9.RefRequestBuilder) {
    return i4e1f588b74fcc251466a08810fb8a3efa51152db86379c9d3ef91dbaf3f02bf9.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
