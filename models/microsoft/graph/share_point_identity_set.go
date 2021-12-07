package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharePointIdentitySet 
type SharePointIdentitySet struct {
    IdentitySet
    // The group associated with this action. Optional.
    group *Identity;
    // The SharePoint group associated with this action. Optional.
    siteGroup *SharePointIdentity;
    // The SharePoint user associated with this action. Optional.
    siteUser *SharePointIdentity;
}
// NewSharePointIdentitySet instantiates a new sharePointIdentitySet and sets the default values.
func NewSharePointIdentitySet()(*SharePointIdentitySet) {
    m := &SharePointIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// GetGroup gets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) GetGroup()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// GetSiteGroup gets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteGroup()(*SharePointIdentity) {
    if m == nil {
        return nil
    } else {
        return m.siteGroup
    }
}
// GetSiteUser gets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteUser()(*SharePointIdentity) {
    if m == nil {
        return nil
    } else {
        return m.siteUser
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(*Identity))
        }
        return nil
    }
    res["siteGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharePointIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteGroup(val.(*SharePointIdentity))
        }
        return nil
    }
    res["siteUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharePointIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUser(val.(*SharePointIdentity))
        }
        return nil
    }
    return res
}
func (m *SharePointIdentitySet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharePointIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteGroup", m.GetSiteGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteUser", m.GetSiteUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) SetGroup(value *Identity)() {
    if m != nil {
        m.group = value
    }
}
// SetSiteGroup sets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteGroup(value *SharePointIdentity)() {
    if m != nil {
        m.siteGroup = value
    }
}
// SetSiteUser sets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteUser(value *SharePointIdentity)() {
    if m != nil {
        m.siteUser = value
    }
}
