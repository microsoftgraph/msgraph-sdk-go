package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingTimeSuggestionsResult struct {
    additionalData map[string]interface{};
    emptySuggestionsReason *string;
    meetingTimeSuggestions []MeetingTimeSuggestion;
}
func NewMeetingTimeSuggestionsResult()(*MeetingTimeSuggestionsResult) {
    m := &MeetingTimeSuggestionsResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MeetingTimeSuggestionsResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MeetingTimeSuggestionsResult) GetEmptySuggestionsReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emptySuggestionsReason
    }
}
func (m *MeetingTimeSuggestionsResult) GetMeetingTimeSuggestions()([]MeetingTimeSuggestion) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSuggestions
    }
}
func (m *MeetingTimeSuggestionsResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["emptySuggestionsReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmptySuggestionsReason(val)
        return nil
    }
    res["meetingTimeSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingTimeSuggestion() })
        if err != nil {
            return err
        }
        res := make([]MeetingTimeSuggestion, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingTimeSuggestion))
        }
        m.SetMeetingTimeSuggestions(res)
        return nil
    }
    return res
}
func (m *MeetingTimeSuggestionsResult) IsNil()(bool) {
    return m == nil
}
func (m *MeetingTimeSuggestionsResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emptySuggestionsReason", m.GetEmptySuggestionsReason())
        if err != nil {
            return err
        }
    }
    {
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
func (m *MeetingTimeSuggestionsResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MeetingTimeSuggestionsResult) SetEmptySuggestionsReason(value *string)() {
    m.emptySuggestionsReason = value
}
func (m *MeetingTimeSuggestionsResult) SetMeetingTimeSuggestions(value []MeetingTimeSuggestion)() {
    m.meetingTimeSuggestions = value
}
