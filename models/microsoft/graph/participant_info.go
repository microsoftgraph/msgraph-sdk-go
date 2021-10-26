package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ParticipantInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
    countryCode *string;
    // The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
    endpointType *EndpointType;
    // 
    identity *IdentitySet;
    // The language culture string. Read-only.
    languageId *string;
    // The participant ID of the participant. Read-only.
    participantId *string;
    // The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
    region *string;
}
// Instantiates a new participantInfo and sets the default values.
func NewParticipantInfo()(*ParticipantInfo) {
    m := &ParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
func (m *ParticipantInfo) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
// Gets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
func (m *ParticipantInfo) GetEndpointType()(*EndpointType) {
    if m == nil {
        return nil
    } else {
        return m.endpointType
    }
}
// Gets the identity property value. 
func (m *ParticipantInfo) GetIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// Gets the languageId property value. The language culture string. Read-only.
func (m *ParticipantInfo) GetLanguageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageId
    }
}
// Gets the participantId property value. The participant ID of the participant. Read-only.
func (m *ParticipantInfo) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// Gets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
func (m *ParticipantInfo) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// The deserialization information for the current model
func (m *ParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["countryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryCode(val)
        return nil
    }
    res["endpointType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        cast := val.(EndpointType)
        m.SetEndpointType(&cast)
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetIdentity(val.(*IdentitySet))
        return nil
    }
    res["languageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguageId(val)
        return nil
    }
    res["participantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParticipantId(val)
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRegion(val)
        return nil
    }
    return res
}
func (m *ParticipantInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ParticipantInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    if m.GetEndpointType() != nil {
        cast := m.GetEndpointType().String()
        err := writer.WriteStringValue("endpointType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageId", m.GetLanguageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("participantId", m.GetParticipantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
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
func (m *ParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
// Parameters:
//  - value : Value to set for the countryCode property.
func (m *ParticipantInfo) SetCountryCode(value *string)() {
    m.countryCode = value
}
// Sets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
// Parameters:
//  - value : Value to set for the endpointType property.
func (m *ParticipantInfo) SetEndpointType(value *EndpointType)() {
    m.endpointType = value
}
// Sets the identity property value. 
// Parameters:
//  - value : Value to set for the identity property.
func (m *ParticipantInfo) SetIdentity(value *IdentitySet)() {
    m.identity = value
}
// Sets the languageId property value. The language culture string. Read-only.
// Parameters:
//  - value : Value to set for the languageId property.
func (m *ParticipantInfo) SetLanguageId(value *string)() {
    m.languageId = value
}
// Sets the participantId property value. The participant ID of the participant. Read-only.
// Parameters:
//  - value : Value to set for the participantId property.
func (m *ParticipantInfo) SetParticipantId(value *string)() {
    m.participantId = value
}
// Sets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
// Parameters:
//  - value : Value to set for the region property.
func (m *ParticipantInfo) SetRegion(value *string)() {
    m.region = value
}
