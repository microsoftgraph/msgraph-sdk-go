package communications

import (
    i3c47b88ab2c0939d08c5ce7f7d03611ddbd6513767887f2f85b50574b229da2e "github.com/microsoftgraph/msgraph-sdk-go/communications/getpresencesbyuserid"
    i7094ab8c1a06bafeeedd3a2746844e0ec2ac0e2e7d2937ef9abdad9734836b3b "github.com/microsoftgraph/msgraph-sdk-go/communications/presences"
    i8ffdea05c8e955294fca9f00741f2c593d8a678f56d226eb2403adf492ba07b0 "github.com/microsoftgraph/msgraph-sdk-go/communications/onlinemeetings"
    i97bbd24556e25a02196df7f369bb3b5f98f59c02b80c2e99a6f1a240704122ec "github.com/microsoftgraph/msgraph-sdk-go/communications/calls"
    ia16cdc2d9b3bc568ec8cc434362cdc775f06d33edc59461753307ae81aa82be2 "github.com/microsoftgraph/msgraph-sdk-go/communications/callrecords"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i18c4da38d0089983229c741b411e29ea20d5089f44a41039b28b9300ac6b2bbc "github.com/microsoftgraph/msgraph-sdk-go/communications/presences/item"
    i477b1ae6bd5bbfab5a3df641da78a8d4e8d13bcb144b2cc255a5232c225c8e02 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i86a5a2d6d1df2d6f16bc44eca70678c6e13a6e96de856db7943e4d10d0e565f9 "github.com/microsoftgraph/msgraph-sdk-go/communications/callrecords/item"
    ia12cb51356bed2ae549e4383f0c613f371754648330c864ba0a94eed03acad34 "github.com/microsoftgraph/msgraph-sdk-go/communications/onlinemeetings/item"
)

// CommunicationsRequestBuilder provides operations to manage the cloudCommunications singleton.
type CommunicationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CommunicationsRequestBuilderGetOptions options for Get
type CommunicationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CommunicationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CommunicationsRequestBuilderGetQueryParameters get communications
type CommunicationsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CommunicationsRequestBuilderPatchOptions options for Patch
type CommunicationsRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CloudCommunicationsable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CommunicationsRequestBuilder) CallRecords()(*ia16cdc2d9b3bc568ec8cc434362cdc775f06d33edc59461753307ae81aa82be2.CallRecordsRequestBuilder) {
    return ia16cdc2d9b3bc568ec8cc434362cdc775f06d33edc59461753307ae81aa82be2.NewCallRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallRecordsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.callRecords.item collection
func (m *CommunicationsRequestBuilder) CallRecordsById(id string)(*i86a5a2d6d1df2d6f16bc44eca70678c6e13a6e96de856db7943e4d10d0e565f9.CallRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["callRecord_id"] = id
    }
    return i86a5a2d6d1df2d6f16bc44eca70678c6e13a6e96de856db7943e4d10d0e565f9.NewCallRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CommunicationsRequestBuilder) Calls()(*i97bbd24556e25a02196df7f369bb3b5f98f59c02b80c2e99a6f1a240704122ec.CallsRequestBuilder) {
    return i97bbd24556e25a02196df7f369bb3b5f98f59c02b80c2e99a6f1a240704122ec.NewCallsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.calls.item collection
func (m *CommunicationsRequestBuilder) CallsById(id string)(*i477b1ae6bd5bbfab5a3df641da78a8d4e8d13bcb144b2cc255a5232c225c8e02.CallItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["call_id"] = id
    }
    return i477b1ae6bd5bbfab5a3df641da78a8d4e8d13bcb144b2cc255a5232c225c8e02.NewCallItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCommunicationsRequestBuilderInternal instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CommunicationsRequestBuilder) {
    m := &CommunicationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCommunicationsRequestBuilder instantiates a new CommunicationsRequestBuilder and sets the default values.
func NewCommunicationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CommunicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get communications
func (m *CommunicationsRequestBuilder) CreateGetRequestInformation(options *CommunicationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update communications
func (m *CommunicationsRequestBuilder) CreatePatchRequestInformation(options *CommunicationsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get communications
func (m *CommunicationsRequestBuilder) Get(options *CommunicationsRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CloudCommunicationsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCloudCommunicationsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CloudCommunicationsable), nil
}
func (m *CommunicationsRequestBuilder) GetPresencesByUserId()(*i3c47b88ab2c0939d08c5ce7f7d03611ddbd6513767887f2f85b50574b229da2e.GetPresencesByUserIdRequestBuilder) {
    return i3c47b88ab2c0939d08c5ce7f7d03611ddbd6513767887f2f85b50574b229da2e.NewGetPresencesByUserIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CommunicationsRequestBuilder) OnlineMeetings()(*i8ffdea05c8e955294fca9f00741f2c593d8a678f56d226eb2403adf492ba07b0.OnlineMeetingsRequestBuilder) {
    return i8ffdea05c8e955294fca9f00741f2c593d8a678f56d226eb2403adf492ba07b0.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.onlineMeetings.item collection
func (m *CommunicationsRequestBuilder) OnlineMeetingsById(id string)(*ia12cb51356bed2ae549e4383f0c613f371754648330c864ba0a94eed03acad34.OnlineMeetingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return ia12cb51356bed2ae549e4383f0c613f371754648330c864ba0a94eed03acad34.NewOnlineMeetingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update communications
func (m *CommunicationsRequestBuilder) Patch(options *CommunicationsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *CommunicationsRequestBuilder) Presences()(*i7094ab8c1a06bafeeedd3a2746844e0ec2ac0e2e7d2937ef9abdad9734836b3b.PresencesRequestBuilder) {
    return i7094ab8c1a06bafeeedd3a2746844e0ec2ac0e2e7d2937ef9abdad9734836b3b.NewPresencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresencesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.presences.item collection
func (m *CommunicationsRequestBuilder) PresencesById(id string)(*i18c4da38d0089983229c741b411e29ea20d5089f44a41039b28b9300ac6b2bbc.PresenceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["presence_id"] = id
    }
    return i18c4da38d0089983229c741b411e29ea20d5089f44a41039b28b9300ac6b2bbc.NewPresenceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
