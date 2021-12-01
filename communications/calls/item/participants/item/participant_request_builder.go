package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i34d02e15d3dbbd84be27b43a7676cbd08e07b5748e464a86fa76072e0b53219d "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/participants/item/startholdmusic"
    i3971077a7a850e86ed619e297dd893137fa9fe6257be92bbea7e0542da02be7b "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/participants/item/mute"
    i61980b8fddd9659b7077791f193bc75e71a71475ff9efa632dd2c50bc47e5cc7 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/participants/item/stopholdmusic"
)

// ParticipantRequestBuilder builds and executes requests for operations under \communications\calls\{call-id}\participants\{participant-id}
type ParticipantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParticipantRequestBuilderDeleteOptions options for Delete
type ParticipantRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParticipantRequestBuilderGetOptions options for Get
type ParticipantRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParticipantRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParticipantRequestBuilderGetQueryParameters read-only. Nullable.
type ParticipantRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ParticipantRequestBuilderPatchOptions options for Patch
type ParticipantRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Participant;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewParticipantRequestBuilderInternal instantiates a new ParticipantRequestBuilder and sets the default values.
func NewParticipantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParticipantRequestBuilder) {
    m := &ParticipantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}/participants/{participant_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParticipantRequestBuilder instantiates a new ParticipantRequestBuilder and sets the default values.
func NewParticipantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParticipantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParticipantRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreateDeleteRequestInformation(options *ParticipantRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreateGetRequestInformation(options *ParticipantRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable.
func (m *ParticipantRequestBuilder) CreatePatchRequestInformation(options *ParticipantRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable.
func (m *ParticipantRequestBuilder) Delete(options *ParticipantRequestBuilderDeleteOptions)(error) {
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
// Get read-only. Nullable.
func (m *ParticipantRequestBuilder) Get(options *ParticipantRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Participant, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewParticipant() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Participant), nil
}
func (m *ParticipantRequestBuilder) Mute()(*i3971077a7a850e86ed619e297dd893137fa9fe6257be92bbea7e0542da02be7b.MuteRequestBuilder) {
    return i3971077a7a850e86ed619e297dd893137fa9fe6257be92bbea7e0542da02be7b.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch read-only. Nullable.
func (m *ParticipantRequestBuilder) Patch(options *ParticipantRequestBuilderPatchOptions)(error) {
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
func (m *ParticipantRequestBuilder) StartHoldMusic()(*i34d02e15d3dbbd84be27b43a7676cbd08e07b5748e464a86fa76072e0b53219d.StartHoldMusicRequestBuilder) {
    return i34d02e15d3dbbd84be27b43a7676cbd08e07b5748e464a86fa76072e0b53219d.NewStartHoldMusicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParticipantRequestBuilder) StopHoldMusic()(*i61980b8fddd9659b7077791f193bc75e71a71475ff9efa632dd2c50bc47e5cc7.StopHoldMusicRequestBuilder) {
    return i61980b8fddd9659b7077791f193bc75e71a71475ff9efa632dd2c50bc47e5cc7.NewStopHoldMusicRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
