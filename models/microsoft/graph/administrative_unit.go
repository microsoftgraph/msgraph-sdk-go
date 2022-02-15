package graph

import (
    id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50 "github.com/microsoftgraph/msgraph-sdk-go/devices/getbyids"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdministrativeUnit 
type AdministrativeUnit struct {
    id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject
    // An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
    description *string;
    // Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string;
    // The collection of open extensions defined for this administrative unit. Nullable.
    extensions []Extension;
    // Users and groups that are members of this administrative unit. Supports $expand.
    members []id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject;
    // Scoped-role members of this administrative unit.
    scopedRoleMembers []ScopedRoleMembership;
    // Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
    visibility *string;
}
// NewAdministrativeUnit instantiates a new administrativeUnit and sets the default values.
func NewAdministrativeUnit()(*AdministrativeUnit) {
    m := &AdministrativeUnit{
        DirectoryObject: *id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.NewDirectoryObject(),
    }
    return m
}
// GetDescription gets the description property value. An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
func (m *AdministrativeUnit) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *AdministrativeUnit) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for this administrative unit. Nullable.
func (m *AdministrativeUnit) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetMembers gets the members property value. Users and groups that are members of this administrative unit. Supports $expand.
func (m *AdministrativeUnit) GetMembers()([]id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetScopedRoleMembers gets the scopedRoleMembers property value. Scoped-role members of this administrative unit.
func (m *AdministrativeUnit) GetScopedRoleMembers()([]ScopedRoleMembership) {
    if m == nil {
        return nil
    } else {
        return m.scopedRoleMembers
    }
}
// GetVisibility gets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
func (m *AdministrativeUnit) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdministrativeUnit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["scopedRoleMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScopedRoleMembership() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembership, len(val))
            for i, v := range val {
                res[i] = *(v.(*ScopedRoleMembership))
            }
            m.SetScopedRoleMembers(res)
        }
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
func (m *AdministrativeUnit) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetExtensions() != nil {
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
    if m.GetMembers() != nil {
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
    if m.GetScopedRoleMembers() != nil {
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
// SetDescription sets the description property value. An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
func (m *AdministrativeUnit) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *AdministrativeUnit) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for this administrative unit. Nullable.
func (m *AdministrativeUnit) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetMembers sets the members property value. Users and groups that are members of this administrative unit. Supports $expand.
func (m *AdministrativeUnit) SetMembers(value []id548309b7f062162722aee435eac927347015bc798cf5b757340af2dc3770b50.DirectoryObject)() {
    if m != nil {
        m.members = value
    }
}
// SetScopedRoleMembers sets the scopedRoleMembers property value. Scoped-role members of this administrative unit.
func (m *AdministrativeUnit) SetScopedRoleMembers(value []ScopedRoleMembership)() {
    if m != nil {
        m.scopedRoleMembers = value
    }
}
// SetVisibility sets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership or Public. If not set, default behavior is Public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
func (m *AdministrativeUnit) SetVisibility(value *string)() {
    if m != nil {
        m.visibility = value
    }
}
