package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AdministrativeUnit struct {
    DirectoryObject
    // An optional description for the administrative unit.
    description *string;
    // Display name for the administrative unit.
    displayName *string;
    // The collection of open extensions defined for this Administrative Unit. Nullable.
    extensions []Extension;
    // Users and groups that are members of this Adminsitrative Unit. HTTP Methods: GET (list members), POST (add members), DELETE (remove members).
    members []DirectoryObject;
    // Scoped-role members of this Administrative Unit.  HTTP Methods: GET (list scopedRoleMemberships), POST (add scopedRoleMembership), DELETE (remove scopedRoleMembership).
    scopedRoleMembers []ScopedRoleMembership;
    // Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
    visibility *string;
}
// Instantiates a new administrativeUnit and sets the default values.
func NewAdministrativeUnit()(*AdministrativeUnit) {
    m := &AdministrativeUnit{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the description property value. An optional description for the administrative unit.
func (m *AdministrativeUnit) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name for the administrative unit.
func (m *AdministrativeUnit) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the extensions property value. The collection of open extensions defined for this Administrative Unit. Nullable.
func (m *AdministrativeUnit) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the members property value. Users and groups that are members of this Adminsitrative Unit. HTTP Methods: GET (list members), POST (add members), DELETE (remove members).
func (m *AdministrativeUnit) GetMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the scopedRoleMembers property value. Scoped-role members of this Administrative Unit.  HTTP Methods: GET (list scopedRoleMemberships), POST (add scopedRoleMembership), DELETE (remove scopedRoleMembership).
func (m *AdministrativeUnit) GetScopedRoleMembers()([]ScopedRoleMembership) {
    if m == nil {
        return nil
    } else {
        return m.scopedRoleMembers
    }
}
// Gets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
func (m *AdministrativeUnit) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// The deserialization information for the current model
func (m *AdministrativeUnit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        res := make([]Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetMembers(res)
        return nil
    }
    res["scopedRoleMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScopedRoleMembership() })
        if err != nil {
            return err
        }
        res := make([]ScopedRoleMembership, len(val))
        for i, v := range val {
            res[i] = *(v.(*ScopedRoleMembership))
        }
        m.SetScopedRoleMembers(res)
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVisibility(val)
        return nil
    }
    return res
}
func (m *AdministrativeUnit) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AdministrativeUnit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScopedRoleMembers()))
        for i, v := range m.GetScopedRoleMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scopedRoleMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. An optional description for the administrative unit.
// Parameters:
//  - value : Value to set for the description property.
func (m *AdministrativeUnit) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name for the administrative unit.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AdministrativeUnit) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the extensions property value. The collection of open extensions defined for this Administrative Unit. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *AdministrativeUnit) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the members property value. Users and groups that are members of this Adminsitrative Unit. HTTP Methods: GET (list members), POST (add members), DELETE (remove members).
// Parameters:
//  - value : Value to set for the members property.
func (m *AdministrativeUnit) SetMembers(value []DirectoryObject)() {
    m.members = value
}
// Sets the scopedRoleMembers property value. Scoped-role members of this Administrative Unit.  HTTP Methods: GET (list scopedRoleMemberships), POST (add scopedRoleMembership), DELETE (remove scopedRoleMembership).
// Parameters:
//  - value : Value to set for the scopedRoleMembers property.
func (m *AdministrativeUnit) SetScopedRoleMembers(value []ScopedRoleMembership)() {
    m.scopedRoleMembers = value
}
// Sets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
// Parameters:
//  - value : Value to set for the visibility property.
func (m *AdministrativeUnit) SetVisibility(value *string)() {
    m.visibility = value
}
