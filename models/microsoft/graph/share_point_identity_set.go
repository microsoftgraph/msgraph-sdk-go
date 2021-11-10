package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharePointIdentitySet struct {
    IdentitySet
    // 
    group *Identity;
    // 
    siteGroup *SharePointIdentity;
    // 
    siteUser *SharePointIdentity;
}
// Instantiates a new sharePointIdentitySet and sets the default values.
func NewSharePointIdentitySet()(*SharePointIdentitySet) {
    m := &SharePointIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// Gets the group property value. 
func (m *SharePointIdentitySet) GetGroup()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// Gets the siteGroup property value. 
func (m *SharePointIdentitySet) GetSiteGroup()(*SharePointIdentity) {
    if m == nil {
        return nil
    } else {
        return m.siteGroup
    }
}
// Gets the siteUser property value. 
func (m *SharePointIdentitySet) GetSiteUser()(*SharePointIdentity) {
    if m == nil {
        return nil
    } else {
        return m.siteUser
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the group property value. 
// Parameters:
//  - value : Value to set for the group property.
func (m *SharePointIdentitySet) SetGroup(value *Identity)() {
    m.group = value
}
// Sets the siteGroup property value. 
// Parameters:
//  - value : Value to set for the siteGroup property.
func (m *SharePointIdentitySet) SetSiteGroup(value *SharePointIdentity)() {
    m.siteGroup = value
}
// Sets the siteUser property value. 
// Parameters:
//  - value : Value to set for the siteUser property.
func (m *SharePointIdentitySet) SetSiteUser(value *SharePointIdentity)() {
    m.siteUser = value
}
