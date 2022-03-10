package findmeetingtimes

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// FindMeetingTimesRequestBuilder provides operations to call the findMeetingTimes method.
type FindMeetingTimesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// FindMeetingTimesRequestBuilderPostOptions options for Post
type FindMeetingTimesRequestBuilderPostOptions struct {
    // 
    Body FindMeetingTimesRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// FindMeetingTimesResponse union type wrapper for classes meetingTimeSuggestionsResult
type FindMeetingTimesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type meetingTimeSuggestionsResult
    meetingTimeSuggestionsResult i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable;
}
// NewFindMeetingTimesResponse instantiates a new findMeetingTimesResponse and sets the default values.
func NewFindMeetingTimesResponse()(*FindMeetingTimesResponse) {
    m := &FindMeetingTimesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateFindMeetingTimesResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFindMeetingTimesResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FindMeetingTimesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FindMeetingTimesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["meetingTimeSuggestionsResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateMeetingTimeSuggestionsResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingTimeSuggestionsResult(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable))
        }
        return nil
    }
    return res
}
// GetMeetingTimeSuggestionsResult gets the meetingTimeSuggestionsResult property value. Union type representation for type meetingTimeSuggestionsResult
func (m *FindMeetingTimesResponse) GetMeetingTimeSuggestionsResult()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSuggestionsResult
    }
}
func (m *FindMeetingTimesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FindMeetingTimesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("meetingTimeSuggestionsResult", m.GetMeetingTimeSuggestionsResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FindMeetingTimesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMeetingTimeSuggestionsResult sets the meetingTimeSuggestionsResult property value. Union type representation for type meetingTimeSuggestionsResult
func (m *FindMeetingTimesResponse) SetMeetingTimeSuggestionsResult(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable)() {
    if m != nil {
        m.meetingTimeSuggestionsResult = value
    }
}
// FindMeetingTimesResponseable 
type FindMeetingTimesResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetMeetingTimeSuggestionsResult()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable)
    SetMeetingTimeSuggestionsResult(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MeetingTimeSuggestionsResultable)()
}
// NewFindMeetingTimesRequestBuilderInternal instantiates a new FindMeetingTimesRequestBuilder and sets the default values.
func NewFindMeetingTimesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FindMeetingTimesRequestBuilder) {
    m := &FindMeetingTimesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.findMeetingTimes";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFindMeetingTimesRequestBuilder instantiates a new FindMeetingTimesRequestBuilder and sets the default values.
func NewFindMeetingTimesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*FindMeetingTimesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindMeetingTimesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action findMeetingTimes
func (m *FindMeetingTimesRequestBuilder) CreatePostRequestInformation(options *FindMeetingTimesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action findMeetingTimes
func (m *FindMeetingTimesRequestBuilder) Post(options *FindMeetingTimesRequestBuilderPostOptions)(FindMeetingTimesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateFindMeetingTimesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(FindMeetingTimesResponseable), nil
}
