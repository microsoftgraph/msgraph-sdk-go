package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LobbyBypassSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
    isDialInBypassEnabled *bool;
    // Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
    scope *LobbyBypassScope;
}
// Instantiates a new lobbyBypassSettings and sets the default values.
func NewLobbyBypassSettings()(*LobbyBypassSettings) {
    m := &LobbyBypassSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LobbyBypassSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
func (m *LobbyBypassSettings) GetIsDialInBypassEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDialInBypassEnabled
    }
}
// Gets the scope property value. Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
func (m *LobbyBypassSettings) GetScope()(*LobbyBypassScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// The deserialization information for the current model
func (m *LobbyBypassSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDialInBypassEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDialInBypassEnabled(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLobbyBypassScope)
        if err != nil {
            return err
        }
        cast := val.(LobbyBypassScope)
        m.SetScope(&cast)
        return nil
    }
    return res
}
func (m *LobbyBypassSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LobbyBypassSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDialInBypassEnabled", m.GetIsDialInBypassEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := m.GetScope().String()
        err := writer.WriteStringValue("scope", &cast)
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
func (m *LobbyBypassSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
// Parameters:
//  - value : Value to set for the isDialInBypassEnabled property.
func (m *LobbyBypassSettings) SetIsDialInBypassEnabled(value *bool)() {
    m.isDialInBypassEnabled = value
}
// Sets the scope property value. Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
// Parameters:
//  - value : Value to set for the scope property.
func (m *LobbyBypassSettings) SetScope(value *LobbyBypassScope)() {
    m.scope = value
}
