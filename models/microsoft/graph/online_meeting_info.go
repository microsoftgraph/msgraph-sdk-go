package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ID of the conference.
    conferenceId *string;
    // The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
    joinUrl *string;
    // All of the phone numbers associated with this conference.
    phones []Phone;
    // The pre-formatted quickdial for this call.
    quickDial *string;
    // The toll free numbers that can be used to join the conference.
    tollFreeNumbers []string;
    // The toll number that can be used to join the conference.
    tollNumber *string;
}
// Instantiates a new onlineMeetingInfo and sets the default values.
func NewOnlineMeetingInfo()(*OnlineMeetingInfo) {
    m := &OnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the conferenceId property value. The ID of the conference.
func (m *OnlineMeetingInfo) GetConferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conferenceId
    }
}
// Gets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
func (m *OnlineMeetingInfo) GetJoinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinUrl
    }
}
// Gets the phones property value. All of the phone numbers associated with this conference.
func (m *OnlineMeetingInfo) GetPhones()([]Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Gets the quickDial property value. The pre-formatted quickdial for this call.
func (m *OnlineMeetingInfo) GetQuickDial()(*string) {
    if m == nil {
        return nil
    } else {
        return m.quickDial
    }
}
// Gets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollFreeNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumbers
    }
}
// Gets the tollNumber property value. The toll number that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumber
    }
}
// The deserialization information for the current model
func (m *OnlineMeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConferenceId(val)
        return nil
    }
    res["joinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoinUrl(val)
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhone() })
        if err != nil {
            return err
        }
        res := make([]Phone, len(val))
        for i, v := range val {
            res[i] = *(v.(*Phone))
        }
        m.SetPhones(res)
        return nil
    }
    res["quickDial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQuickDial(val)
        return nil
    }
    res["tollFreeNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTollFreeNumbers(res)
        return nil
    }
    res["tollNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTollNumber(val)
        return nil
    }
    return res
}
func (m *OnlineMeetingInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnlineMeetingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinUrl", m.GetJoinUrl())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quickDial", m.GetQuickDial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("tollFreeNumbers", m.GetTollFreeNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollNumber", m.GetTollNumber())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OnlineMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the conferenceId property value. The ID of the conference.
// Parameters:
//  - value : Value to set for the conferenceId property.
func (m *OnlineMeetingInfo) SetConferenceId(value *string)() {
    m.conferenceId = value
}
// Sets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
// Parameters:
//  - value : Value to set for the joinUrl property.
func (m *OnlineMeetingInfo) SetJoinUrl(value *string)() {
    m.joinUrl = value
}
// Sets the phones property value. All of the phone numbers associated with this conference.
// Parameters:
//  - value : Value to set for the phones property.
func (m *OnlineMeetingInfo) SetPhones(value []Phone)() {
    m.phones = value
}
// Sets the quickDial property value. The pre-formatted quickdial for this call.
// Parameters:
//  - value : Value to set for the quickDial property.
func (m *OnlineMeetingInfo) SetQuickDial(value *string)() {
    m.quickDial = value
}
// Sets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
// Parameters:
//  - value : Value to set for the tollFreeNumbers property.
func (m *OnlineMeetingInfo) SetTollFreeNumbers(value []string)() {
    m.tollFreeNumbers = value
}
// Sets the tollNumber property value. The toll number that can be used to join the conference.
// Parameters:
//  - value : Value to set for the tollNumber property.
func (m *OnlineMeetingInfo) SetTollNumber(value *string)() {
    m.tollNumber = value
}
