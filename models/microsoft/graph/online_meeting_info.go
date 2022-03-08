package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnlineMeetingInfo provides operations to manage the drive singleton.
type OnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ID of the conference.
    conferenceId *string;
    // The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
    joinUrl *string;
    // All of the phone numbers associated with this conference.
    phones []Phoneable;
    // The pre-formatted quickdial for this call.
    quickDial *string;
    // The toll free numbers that can be used to join the conference.
    tollFreeNumbers []string;
    // The toll number that can be used to join the conference.
    tollNumber *string;
}
// NewOnlineMeetingInfo instantiates a new onlineMeetingInfo and sets the default values.
func NewOnlineMeetingInfo()(*OnlineMeetingInfo) {
    m := &OnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnlineMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnlineMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConferenceId gets the conferenceId property value. The ID of the conference.
func (m *OnlineMeetingInfo) GetConferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conferenceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConferenceId(val)
        }
        return nil
    }
    res["joinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinUrl(val)
        }
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                res[i] = v.(Phoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["quickDial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickDial(val)
        }
        return nil
    }
    res["tollFreeNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTollFreeNumbers(res)
        }
        return nil
    }
    res["tollNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTollNumber(val)
        }
        return nil
    }
    return res
}
// GetJoinUrl gets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
func (m *OnlineMeetingInfo) GetJoinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinUrl
    }
}
// GetPhones gets the phones property value. All of the phone numbers associated with this conference.
func (m *OnlineMeetingInfo) GetPhones()([]Phoneable) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// GetQuickDial gets the quickDial property value. The pre-formatted quickdial for this call.
func (m *OnlineMeetingInfo) GetQuickDial()(*string) {
    if m == nil {
        return nil
    } else {
        return m.quickDial
    }
}
// GetTollFreeNumbers gets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollFreeNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumbers
    }
}
// GetTollNumber gets the tollNumber property value. The toll number that can be used to join the conference.
func (m *OnlineMeetingInfo) GetTollNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumber
    }
}
func (m *OnlineMeetingInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetPhones() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetTollFreeNumbers() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnlineMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConferenceId sets the conferenceId property value. The ID of the conference.
func (m *OnlineMeetingInfo) SetConferenceId(value *string)() {
    if m != nil {
        m.conferenceId = value
    }
}
// SetJoinUrl sets the joinUrl property value. The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting.
func (m *OnlineMeetingInfo) SetJoinUrl(value *string)() {
    if m != nil {
        m.joinUrl = value
    }
}
// SetPhones sets the phones property value. All of the phone numbers associated with this conference.
func (m *OnlineMeetingInfo) SetPhones(value []Phoneable)() {
    if m != nil {
        m.phones = value
    }
}
// SetQuickDial sets the quickDial property value. The pre-formatted quickdial for this call.
func (m *OnlineMeetingInfo) SetQuickDial(value *string)() {
    if m != nil {
        m.quickDial = value
    }
}
// SetTollFreeNumbers sets the tollFreeNumbers property value. The toll free numbers that can be used to join the conference.
func (m *OnlineMeetingInfo) SetTollFreeNumbers(value []string)() {
    if m != nil {
        m.tollFreeNumbers = value
    }
}
// SetTollNumber sets the tollNumber property value. The toll number that can be used to join the conference.
func (m *OnlineMeetingInfo) SetTollNumber(value *string)() {
    if m != nil {
        m.tollNumber = value
    }
}
