package findmeetingtimes

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// FindMeetingTimesRequestBody provides operations to call the findMeetingTimes method.
type FindMeetingTimesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    attendees []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AttendeeBaseable;
    // 
    isOrganizerOptional *bool;
    // 
    locationConstraint i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LocationConstraintable;
    // 
    maxCandidates *int32;
    // 
    meetingDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // 
    minimumAttendeePercentage *float64;
    // 
    returnSuggestionReasons *bool;
    // 
    timeConstraint i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeConstraintable;
}
// NewFindMeetingTimesRequestBody instantiates a new findMeetingTimesRequestBody and sets the default values.
func NewFindMeetingTimesRequestBody()(*FindMeetingTimesRequestBody) {
    m := &FindMeetingTimesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFindMeetingTimesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFindMeetingTimesRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFindMeetingTimesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FindMeetingTimesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttendees gets the attendees property value. 
func (m *FindMeetingTimesRequestBody) GetAttendees()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AttendeeBaseable) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FindMeetingTimesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateAttendeeBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AttendeeBaseable, len(val))
            for i, v := range val {
                res[i] = v.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AttendeeBaseable)
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["isOrganizerOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizerOptional(val)
        }
        return nil
    }
    res["locationConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateLocationConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationConstraint(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LocationConstraintable))
        }
        return nil
    }
    res["maxCandidates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCandidates(val)
        }
        return nil
    }
    res["meetingDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingDuration(val)
        }
        return nil
    }
    res["minimumAttendeePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumAttendeePercentage(val)
        }
        return nil
    }
    res["returnSuggestionReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnSuggestionReasons(val)
        }
        return nil
    }
    res["timeConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTimeConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeConstraint(val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeConstraintable))
        }
        return nil
    }
    return res
}
// GetIsOrganizerOptional gets the isOrganizerOptional property value. 
func (m *FindMeetingTimesRequestBody) GetIsOrganizerOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizerOptional
    }
}
// GetLocationConstraint gets the locationConstraint property value. 
func (m *FindMeetingTimesRequestBody) GetLocationConstraint()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LocationConstraintable) {
    if m == nil {
        return nil
    } else {
        return m.locationConstraint
    }
}
// GetMaxCandidates gets the maxCandidates property value. 
func (m *FindMeetingTimesRequestBody) GetMaxCandidates()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxCandidates
    }
}
// GetMeetingDuration gets the meetingDuration property value. 
func (m *FindMeetingTimesRequestBody) GetMeetingDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.meetingDuration
    }
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. 
func (m *FindMeetingTimesRequestBody) GetMinimumAttendeePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimumAttendeePercentage
    }
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. 
func (m *FindMeetingTimesRequestBody) GetReturnSuggestionReasons()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.returnSuggestionReasons
    }
}
// GetTimeConstraint gets the timeConstraint property value. 
func (m *FindMeetingTimesRequestBody) GetTimeConstraint()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeConstraintable) {
    if m == nil {
        return nil
    } else {
        return m.timeConstraint
    }
}
// Serialize serializes information the current object
func (m *FindMeetingTimesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOrganizerOptional", m.GetIsOrganizerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locationConstraint", m.GetLocationConstraint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxCandidates", m.GetMaxCandidates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("meetingDuration", m.GetMeetingDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimumAttendeePercentage", m.GetMinimumAttendeePercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("returnSuggestionReasons", m.GetReturnSuggestionReasons())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeConstraint", m.GetTimeConstraint())
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
func (m *FindMeetingTimesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttendees sets the attendees property value. 
func (m *FindMeetingTimesRequestBody) SetAttendees(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AttendeeBaseable)() {
    if m != nil {
        m.attendees = value
    }
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. 
func (m *FindMeetingTimesRequestBody) SetIsOrganizerOptional(value *bool)() {
    if m != nil {
        m.isOrganizerOptional = value
    }
}
// SetLocationConstraint sets the locationConstraint property value. 
func (m *FindMeetingTimesRequestBody) SetLocationConstraint(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.LocationConstraintable)() {
    if m != nil {
        m.locationConstraint = value
    }
}
// SetMaxCandidates sets the maxCandidates property value. 
func (m *FindMeetingTimesRequestBody) SetMaxCandidates(value *int32)() {
    if m != nil {
        m.maxCandidates = value
    }
}
// SetMeetingDuration sets the meetingDuration property value. 
func (m *FindMeetingTimesRequestBody) SetMeetingDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.meetingDuration = value
    }
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. 
func (m *FindMeetingTimesRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    if m != nil {
        m.minimumAttendeePercentage = value
    }
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. 
func (m *FindMeetingTimesRequestBody) SetReturnSuggestionReasons(value *bool)() {
    if m != nil {
        m.returnSuggestionReasons = value
    }
}
// SetTimeConstraint sets the timeConstraint property value. 
func (m *FindMeetingTimesRequestBody) SetTimeConstraint(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TimeConstraintable)() {
    if m != nil {
        m.timeConstraint = value
    }
}
