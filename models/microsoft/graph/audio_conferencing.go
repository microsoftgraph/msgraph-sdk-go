package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AudioConferencing struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The conference id of the online meeting.
    conferenceId *string;
    // A URL to the externally-accessible web page that contains dial-in information.
    dialinUrl *string;
    // 
    tollFreeNumber *string;
    // List of toll-free numbers that are displayed in the meeting invite.
    tollFreeNumbers []string;
    // 
    tollNumber *string;
    // List of toll numbers that are displayed in the meeting invite.
    tollNumbers []string;
}
// Instantiates a new audioConferencing and sets the default values.
func NewAudioConferencing()(*AudioConferencing) {
    m := &AudioConferencing{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AudioConferencing) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the conferenceId property value. The conference id of the online meeting.
func (m *AudioConferencing) GetConferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conferenceId
    }
}
// Gets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
func (m *AudioConferencing) GetDialinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dialinUrl
    }
}
// Gets the tollFreeNumber property value. 
func (m *AudioConferencing) GetTollFreeNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumber
    }
}
// Gets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollFreeNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollFreeNumbers
    }
}
// Gets the tollNumber property value. 
func (m *AudioConferencing) GetTollNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumber
    }
}
// Gets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
func (m *AudioConferencing) GetTollNumbers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tollNumbers
    }
}
// The deserialization information for the current model
func (m *AudioConferencing) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["conferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConferenceId(val)
        return nil
    }
    res["dialinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDialinUrl(val)
        return nil
    }
    res["tollFreeNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTollFreeNumber(val)
        return nil
    }
    res["tollFreeNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
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
    res["tollNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTollNumbers(res)
        return nil
    }
    return res
}
func (m *AudioConferencing) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AudioConferencing) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dialinUrl", m.GetDialinUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tollFreeNumber", m.GetTollFreeNumber())
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
        err := writer.WriteCollectionOfStringValues("tollNumbers", m.GetTollNumbers())
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
func (m *AudioConferencing) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the conferenceId property value. The conference id of the online meeting.
// Parameters:
//  - value : Value to set for the conferenceId property.
func (m *AudioConferencing) SetConferenceId(value *string)() {
    m.conferenceId = value
}
// Sets the dialinUrl property value. A URL to the externally-accessible web page that contains dial-in information.
// Parameters:
//  - value : Value to set for the dialinUrl property.
func (m *AudioConferencing) SetDialinUrl(value *string)() {
    m.dialinUrl = value
}
// Sets the tollFreeNumber property value. 
// Parameters:
//  - value : Value to set for the tollFreeNumber property.
func (m *AudioConferencing) SetTollFreeNumber(value *string)() {
    m.tollFreeNumber = value
}
// Sets the tollFreeNumbers property value. List of toll-free numbers that are displayed in the meeting invite.
// Parameters:
//  - value : Value to set for the tollFreeNumbers property.
func (m *AudioConferencing) SetTollFreeNumbers(value []string)() {
    m.tollFreeNumbers = value
}
// Sets the tollNumber property value. 
// Parameters:
//  - value : Value to set for the tollNumber property.
func (m *AudioConferencing) SetTollNumber(value *string)() {
    m.tollNumber = value
}
// Sets the tollNumbers property value. List of toll numbers that are displayed in the meeting invite.
// Parameters:
//  - value : Value to set for the tollNumbers property.
func (m *AudioConferencing) SetTollNumbers(value []string)() {
    m.tollNumbers = value
}
