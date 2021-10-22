package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ScopedRoleMembership struct {
    Entity
    administrativeUnitId *string;
    roleId *string;
    roleMemberInfo *Identity;
}
func NewScopedRoleMembership()(*ScopedRoleMembership) {
    m := &ScopedRoleMembership{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ScopedRoleMembership) GetAdministrativeUnitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnitId
    }
}
func (m *ScopedRoleMembership) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
func (m *ScopedRoleMembership) GetRoleMemberInfo()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.roleMemberInfo
    }
}
func (m *ScopedRoleMembership) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnitId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdministrativeUnitId(val)
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleId(val)
        return nil
    }
    res["roleMemberInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetRoleMemberInfo(val.(*Identity))
        return nil
    }
    return res
}
func (m *ScopedRoleMembership) IsNil()(bool) {
    return m == nil
}
func (m *ScopedRoleMembership) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("administrativeUnitId", m.GetAdministrativeUnitId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleMemberInfo", m.GetRoleMemberInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ScopedRoleMembership) SetAdministrativeUnitId(value *string)() {
    m.administrativeUnitId = value
}
func (m *ScopedRoleMembership) SetRoleId(value *string)() {
    m.roleId = value
}
func (m *ScopedRoleMembership) SetRoleMemberInfo(value *Identity)() {
    m.roleMemberInfo = value
}
