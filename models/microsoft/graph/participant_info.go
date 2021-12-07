package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ParticipantInfo 
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
// NewParticipantInfo instantiates a new participantInfo and sets the default values.
func NewParticipantInfo()(*ParticipantInfo) {
    m := &ParticipantInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParticipantInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCountryCode gets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
func (m *ParticipantInfo) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
// GetEndpointType gets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
func (m *ParticipantInfo) GetEndpointType()(*EndpointType) {
    if m == nil {
        return nil
    } else {
        return m.endpointType
    }
}
// GetIdentity gets the identity property value. 
func (m *ParticipantInfo) GetIdentity()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// GetLanguageId gets the languageId property value. The language culture string. Read-only.
func (m *ParticipantInfo) GetLanguageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageId
    }
}
// GetParticipantId gets the participantId property value. The participant ID of the participant. Read-only.
func (m *ParticipantInfo) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// GetRegion gets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
func (m *ParticipantInfo) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["countryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["endpointType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EndpointType)
            m.SetEndpointType(&cast)
        }
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(*IdentitySet))
        }
        return nil
    }
    res["languageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageId(val)
        }
        return nil
    }
    res["participantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    return res
}
func (m *ParticipantInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParticipantInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCountryCode sets the countryCode property value. The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
func (m *ParticipantInfo) SetCountryCode(value *string)() {
    if m != nil {
        m.countryCode = value
    }
}
// SetEndpointType sets the endpointType property value. The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
func (m *ParticipantInfo) SetEndpointType(value *EndpointType)() {
    if m != nil {
        m.endpointType = value
    }
}
// SetIdentity sets the identity property value. 
func (m *ParticipantInfo) SetIdentity(value *IdentitySet)() {
    if m != nil {
        m.identity = value
    }
}
// SetLanguageId sets the languageId property value. The language culture string. Read-only.
func (m *ParticipantInfo) SetLanguageId(value *string)() {
    if m != nil {
        m.languageId = value
    }
}
// SetParticipantId sets the participantId property value. The participant ID of the participant. Read-only.
func (m *ParticipantInfo) SetParticipantId(value *string)() {
    if m != nil {
        m.participantId = value
    }
}
// SetRegion sets the region property value. The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
func (m *ParticipantInfo) SetRegion(value *string)() {
    if m != nil {
        m.region = value
    }
}
