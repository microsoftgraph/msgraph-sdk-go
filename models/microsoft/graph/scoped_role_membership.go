package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ScopedRoleMembership struct {
    Entity
    // Unique identifier for the administrative unit that the directory role is scoped to
    administrativeUnitId *string;
    // Unique identifier for the directory role that the member is in.
    roleId *string;
    // 
    roleMemberInfo *Identity;
}
// Instantiates a new scopedRoleMembership and sets the default values.
func NewScopedRoleMembership()(*ScopedRoleMembership) {
    m := &ScopedRoleMembership{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the administrativeUnitId property value. Unique identifier for the administrative unit that the directory role is scoped to
func (m *ScopedRoleMembership) GetAdministrativeUnitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.administrativeUnitId
    }
}
// Gets the roleId property value. Unique identifier for the directory role that the member is in.
func (m *ScopedRoleMembership) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// Gets the roleMemberInfo property value. 
func (m *ScopedRoleMembership) GetRoleMemberInfo()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.roleMemberInfo
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the administrativeUnitId property value. Unique identifier for the administrative unit that the directory role is scoped to
// Parameters:
//  - value : Value to set for the administrativeUnitId property.
func (m *ScopedRoleMembership) SetAdministrativeUnitId(value *string)() {
    m.administrativeUnitId = value
}
// Sets the roleId property value. Unique identifier for the directory role that the member is in.
// Parameters:
//  - value : Value to set for the roleId property.
func (m *ScopedRoleMembership) SetRoleId(value *string)() {
    m.roleId = value
}
// Sets the roleMemberInfo property value. 
// Parameters:
//  - value : Value to set for the roleMemberInfo property.
func (m *ScopedRoleMembership) SetRoleMemberInfo(value *Identity)() {
    m.roleMemberInfo = value
}
