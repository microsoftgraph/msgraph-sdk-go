package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DelegatedPermissionClassification struct {
    Entity
    classification *PermissionClassificationType;
    permissionId *string;
    permissionName *string;
}
func NewDelegatedPermissionClassification()(*DelegatedPermissionClassification) {
    m := &DelegatedPermissionClassification{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DelegatedPermissionClassification) GetClassification()(*PermissionClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
func (m *DelegatedPermissionClassification) GetPermissionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionId
    }
}
func (m *DelegatedPermissionClassification) GetPermissionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionName
    }
}
func (m *DelegatedPermissionClassification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionClassificationType)
        if err != nil {
            return err
        }
        cast := val.(PermissionClassificationType)
        m.SetClassification(&cast)
        return nil
    }
    res["permissionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionId(val)
        return nil
    }
    res["permissionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionName(val)
        return nil
    }
    return res
}
func (m *DelegatedPermissionClassification) IsNil()(bool) {
    return m == nil
}
func (m *DelegatedPermissionClassification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := m.GetClassification().String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionId", m.GetPermissionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionName", m.GetPermissionName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DelegatedPermissionClassification) SetClassification(value *PermissionClassificationType)() {
    m.classification = value
}
func (m *DelegatedPermissionClassification) SetPermissionId(value *string)() {
    m.permissionId = value
}
func (m *DelegatedPermissionClassification) SetPermissionName(value *string)() {
    m.permissionName = value
}
