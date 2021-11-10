package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharedDriveItem struct {
    BaseItem
    // Used to access the underlying driveItem
    driveItem *DriveItem;
    // All driveItems contained in the sharing root. This collection cannot be enumerated.
    items []DriveItem;
    // Used to access the underlying list
    list *List;
    // Used to access the underlying listItem
    listItem *ListItem;
    // Information about the owner of the shared item being referenced.
    owner *IdentitySet;
    // Used to access the permission representing the underlying sharing link
    permission *Permission;
    // Used to access the underlying driveItem. Deprecated -- use driveItem instead.
    root *DriveItem;
    // Used to access the underlying site
    site *Site;
}
// Instantiates a new sharedDriveItem and sets the default values.
func NewSharedDriveItem()(*SharedDriveItem) {
    m := &SharedDriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// Gets the driveItem property value. Used to access the underlying driveItem
func (m *SharedDriveItem) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// Gets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
func (m *SharedDriveItem) GetItems()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// Gets the list property value. Used to access the underlying list
func (m *SharedDriveItem) GetList()(*List) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// Gets the listItem property value. Used to access the underlying listItem
func (m *SharedDriveItem) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// Gets the owner property value. Information about the owner of the shared item being referenced.
func (m *SharedDriveItem) GetOwner()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// Gets the permission property value. Used to access the permission representing the underlying sharing link
func (m *SharedDriveItem) GetPermission()(*Permission) {
    if m == nil {
        return nil
    } else {
        return m.permission
    }
}
// Gets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
func (m *SharedDriveItem) GetRoot()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// Gets the site property value. Used to access the underlying site
func (m *SharedDriveItem) GetSite()(*Site) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
// The deserialization information for the current model
func (m *SharedDriveItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(*DriveItem))
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItem))
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewList() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(*List))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(*ListItem))
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(*IdentitySet))
        }
        return nil
    }
    res["permission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermission() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val.(*Permission))
        }
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(*DriveItem))
        }
        return nil
    }
    res["site"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(*Site))
        }
        return nil
    }
    return res
}
func (m *SharedDriveItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SharedDriveItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("list", m.GetList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("listItem", m.GetListItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("root", m.GetRoot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the driveItem property value. Used to access the underlying driveItem
// Parameters:
//  - value : Value to set for the driveItem property.
func (m *SharedDriveItem) SetDriveItem(value *DriveItem)() {
    m.driveItem = value
}
// Sets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
// Parameters:
//  - value : Value to set for the items property.
func (m *SharedDriveItem) SetItems(value []DriveItem)() {
    m.items = value
}
// Sets the list property value. Used to access the underlying list
// Parameters:
//  - value : Value to set for the list property.
func (m *SharedDriveItem) SetList(value *List)() {
    m.list = value
}
// Sets the listItem property value. Used to access the underlying listItem
// Parameters:
//  - value : Value to set for the listItem property.
func (m *SharedDriveItem) SetListItem(value *ListItem)() {
    m.listItem = value
}
// Sets the owner property value. Information about the owner of the shared item being referenced.
// Parameters:
//  - value : Value to set for the owner property.
func (m *SharedDriveItem) SetOwner(value *IdentitySet)() {
    m.owner = value
}
// Sets the permission property value. Used to access the permission representing the underlying sharing link
// Parameters:
//  - value : Value to set for the permission property.
func (m *SharedDriveItem) SetPermission(value *Permission)() {
    m.permission = value
}
// Sets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
// Parameters:
//  - value : Value to set for the root property.
func (m *SharedDriveItem) SetRoot(value *DriveItem)() {
    m.root = value
}
// Sets the site property value. Used to access the underlying site
// Parameters:
//  - value : Value to set for the site property.
func (m *SharedDriveItem) SetSite(value *Site)() {
    m.site = value
}
