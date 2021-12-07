package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LobbyBypassSettings 
type LobbyBypassSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
    isDialInBypassEnabled *bool;
    // Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
    scope *LobbyBypassScope;
}
// NewLobbyBypassSettings instantiates a new lobbyBypassSettings and sets the default values.
func NewLobbyBypassSettings()(*LobbyBypassSettings) {
    m := &LobbyBypassSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LobbyBypassSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsDialInBypassEnabled gets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
func (m *LobbyBypassSettings) GetIsDialInBypassEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDialInBypassEnabled
    }
}
// GetScope gets the scope property value. Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
func (m *LobbyBypassSettings) GetScope()(*LobbyBypassScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LobbyBypassSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDialInBypassEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDialInBypassEnabled(val)
        }
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLobbyBypassScope)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(LobbyBypassScope)
            m.SetScope(&cast)
        }
        return nil
    }
    return res
}
func (m *LobbyBypassSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LobbyBypassSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsDialInBypassEnabled sets the isDialInBypassEnabled property value. Specifies whether or not to always let dial-in callers bypass the lobby. Optional.
func (m *LobbyBypassSettings) SetIsDialInBypassEnabled(value *bool)() {
    if m != nil {
        m.isDialInBypassEnabled = value
    }
}
// SetScope sets the scope property value. Specifies the type of participants that are automatically admitted into a meeting, bypassing the lobby. Optional.
func (m *LobbyBypassSettings) SetScope(value *LobbyBypassScope)() {
    if m != nil {
        m.scope = value
    }
}
