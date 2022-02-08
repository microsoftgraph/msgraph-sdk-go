package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingTimeSuggestionsResult 
type MeetingTimeSuggestionsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
    emptySuggestionsReason *string;
    // An array of meeting suggestions.
    meetingTimeSuggestions []MeetingTimeSuggestion;
}
// NewMeetingTimeSuggestionsResult instantiates a new meetingTimeSuggestionsResult and sets the default values.
func NewMeetingTimeSuggestionsResult()(*MeetingTimeSuggestionsResult) {
    m := &MeetingTimeSuggestionsResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestionsResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmptySuggestionsReason gets the emptySuggestionsReason property value. A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
func (m *MeetingTimeSuggestionsResult) GetEmptySuggestionsReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emptySuggestionsReason
    }
}
// GetMeetingTimeSuggestions gets the meetingTimeSuggestions property value. An array of meeting suggestions.
func (m *MeetingTimeSuggestionsResult) GetMeetingTimeSuggestions()([]MeetingTimeSuggestion) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSuggestions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingTimeSuggestionsResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["emptySuggestionsReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmptySuggestionsReason(val)
        }
        return nil
    }
    res["meetingTimeSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingTimeSuggestion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingTimeSuggestion, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingTimeSuggestion))
            }
            m.SetMeetingTimeSuggestions(res)
        }
        return nil
    }
    return res
}
func (m *MeetingTimeSuggestionsResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingTimeSuggestionsResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emptySuggestionsReason", m.GetEmptySuggestionsReason())
        if err != nil {
            return err
        }
    }
    if m.GetMeetingTimeSuggestions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMeetingTimeSuggestions()))
        for i, v := range m.GetMeetingTimeSuggestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("meetingTimeSuggestions", cast)
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
func (m *MeetingTimeSuggestionsResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmptySuggestionsReason sets the emptySuggestionsReason property value. A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
func (m *MeetingTimeSuggestionsResult) SetEmptySuggestionsReason(value *string)() {
    if m != nil {
        m.emptySuggestionsReason = value
    }
}
// SetMeetingTimeSuggestions sets the meetingTimeSuggestions property value. An array of meeting suggestions.
func (m *MeetingTimeSuggestionsResult) SetMeetingTimeSuggestions(value []MeetingTimeSuggestion)() {
    if m != nil {
        m.meetingTimeSuggestions = value
    }
}
