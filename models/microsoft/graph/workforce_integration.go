package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkforceIntegration struct {
    ChangeTrackedEntity
    apiVersion *int32;
    displayName *string;
    encryption *WorkforceIntegrationEncryption;
    isActive *bool;
    supportedEntities *WorkforceIntegrationSupportedEntities;
    url *string;
}
func NewWorkforceIntegration()(*WorkforceIntegration) {
    m := &WorkforceIntegration{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
func (m *WorkforceIntegration) GetApiVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.apiVersion
    }
}
func (m *WorkforceIntegration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WorkforceIntegration) GetEncryption()(*WorkforceIntegrationEncryption) {
    if m == nil {
        return nil
    } else {
        return m.encryption
    }
}
func (m *WorkforceIntegration) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
func (m *WorkforceIntegration) GetSupportedEntities()(*WorkforceIntegrationSupportedEntities) {
    if m == nil {
        return nil
    } else {
        return m.supportedEntities
    }
}
func (m *WorkforceIntegration) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *WorkforceIntegration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["apiVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetApiVersion(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["encryption"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkforceIntegrationEncryption() })
        if err != nil {
            return err
        }
        m.SetEncryption(val.(*WorkforceIntegrationEncryption))
        return nil
    }
    res["isActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsActive(val)
        return nil
    }
    res["supportedEntities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkforceIntegrationSupportedEntities)
        if err != nil {
            return err
        }
        cast := val.(WorkforceIntegrationSupportedEntities)
        m.SetSupportedEntities(&cast)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *WorkforceIntegration) IsNil()(bool) {
    return m == nil
}
func (m *WorkforceIntegration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("apiVersion", m.GetApiVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("encryption", m.GetEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedEntities() != nil {
        cast := m.GetSupportedEntities().String()
        err = writer.WriteStringValue("supportedEntities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkforceIntegration) SetApiVersion(value *int32)() {
    m.apiVersion = value
}
func (m *WorkforceIntegration) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WorkforceIntegration) SetEncryption(value *WorkforceIntegrationEncryption)() {
    m.encryption = value
}
func (m *WorkforceIntegration) SetIsActive(value *bool)() {
    m.isActive = value
}
func (m *WorkforceIntegration) SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)() {
    m.supportedEntities = value
}
func (m *WorkforceIntegration) SetUrl(value *string)() {
    m.url = value
}
