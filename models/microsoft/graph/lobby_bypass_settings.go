package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LobbyBypassSettings struct {
    additionalData map[string]interface{};
    isDialInBypassEnabled *bool;
    scope *LobbyBypassScope;
}
func NewLobbyBypassSettings()(*LobbyBypassSettings) {
    m := &LobbyBypassSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LobbyBypassSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LobbyBypassSettings) GetIsDialInBypassEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDialInBypassEnabled
    }
}
func (m *LobbyBypassSettings) GetScope()(*LobbyBypassScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
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
func (m *LobbyBypassSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LobbyBypassSettings) SetIsDialInBypassEnabled(value *bool)() {
    m.isDialInBypassEnabled = value
}
func (m *LobbyBypassSettings) SetScope(value *LobbyBypassScope)() {
    m.scope = value
}
