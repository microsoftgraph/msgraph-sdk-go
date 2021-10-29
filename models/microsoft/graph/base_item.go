package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type BaseItem struct {
    Entity
    // Identity of the user, device, or application which created the item. Read-only.
    createdBy *IdentitySet;
    // Identity of the user who created the item. Read-only.
    createdByUser *User;
    // Date and time of item creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Provides a user-visible description of the item. Optional.
    description *string;
    // ETag for the item. Read-only.
    eTag *string;
    // Identity of the user, device, and application which last modified the item. Read-only.
    lastModifiedBy *IdentitySet;
    // Identity of the user who last modified the item. Read-only.
    lastModifiedByUser *User;
    // Date and time the item was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the item. Read-write.
    name *string;
    // Parent information, if the item has a parent. Read-write.
    parentReference *ItemReference;
    // URL that displays the resource in the browser. Read-only.
    webUrl *string;
}
// Instantiates a new baseItem and sets the default values.
func NewBaseItem()(*BaseItem) {
    m := &BaseItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdBy property value. Identity of the user, device, or application which created the item. Read-only.
func (m *BaseItem) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdByUser property value. Identity of the user who created the item. Read-only.
func (m *BaseItem) GetCreatedByUser()(*User) {
    if m == nil {
        return nil
    } else {
        return m.createdByUser
    }
}
// Gets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *BaseItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Provides a user-visible description of the item. Optional.
func (m *BaseItem) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the eTag property value. ETag for the item. Read-only.
func (m *BaseItem) GetETag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eTag
    }
}
// Gets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *BaseItem) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedByUser property value. Identity of the user who last modified the item. Read-only.
func (m *BaseItem) GetLastModifiedByUser()(*User) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedByUser
    }
}
// Gets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *BaseItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the name property value. The name of the item. Read-write.
func (m *BaseItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parentReference property value. Parent information, if the item has a parent. Read-write.
func (m *BaseItem) GetParentReference()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.parentReference
    }
}
// Gets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *BaseItem) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *BaseItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
        return nil
    }
    res["createdByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        m.SetCreatedByUser(val.(*User))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["eTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetETag(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*IdentitySet))
        return nil
    }
    res["lastModifiedByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        m.SetLastModifiedByUser(val.(*User))
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["parentReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        m.SetParentReference(val.(*ItemReference))
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *BaseItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *BaseItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdByUser", m.GetCreatedByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedByUser", m.GetLastModifiedByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentReference", m.GetParentReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdBy property value. Identity of the user, device, or application which created the item. Read-only.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *BaseItem) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdByUser property value. Identity of the user who created the item. Read-only.
// Parameters:
//  - value : Value to set for the createdByUser property.
func (m *BaseItem) SetCreatedByUser(value *User)() {
    m.createdByUser = value
}
// Sets the createdDateTime property value. Date and time of item creation. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *BaseItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Provides a user-visible description of the item. Optional.
// Parameters:
//  - value : Value to set for the description property.
func (m *BaseItem) SetDescription(value *string)() {
    m.description = value
}
// Sets the eTag property value. ETag for the item. Read-only.
// Parameters:
//  - value : Value to set for the eTag property.
func (m *BaseItem) SetETag(value *string)() {
    m.eTag = value
}
// Sets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *BaseItem) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedByUser property value. Identity of the user who last modified the item. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedByUser property.
func (m *BaseItem) SetLastModifiedByUser(value *User)() {
    m.lastModifiedByUser = value
}
// Sets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *BaseItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the name property value. The name of the item. Read-write.
// Parameters:
//  - value : Value to set for the name property.
func (m *BaseItem) SetName(value *string)() {
    m.name = value
}
// Sets the parentReference property value. Parent information, if the item has a parent. Read-write.
// Parameters:
//  - value : Value to set for the parentReference property.
func (m *BaseItem) SetParentReference(value *ItemReference)() {
    m.parentReference = value
}
// Sets the webUrl property value. URL that displays the resource in the browser. Read-only.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *BaseItem) SetWebUrl(value *string)() {
    m.webUrl = value
}
