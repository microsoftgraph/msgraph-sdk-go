package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedDriveItem 
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
// NewSharedDriveItem instantiates a new sharedDriveItem and sets the default values.
func NewSharedDriveItem()(*SharedDriveItem) {
    m := &SharedDriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// GetDriveItem gets the driveItem property value. Used to access the underlying driveItem
func (m *SharedDriveItem) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetItems gets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
func (m *SharedDriveItem) GetItems()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. Used to access the underlying list
func (m *SharedDriveItem) GetList()(*List) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetListItem gets the listItem property value. Used to access the underlying listItem
func (m *SharedDriveItem) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetOwner gets the owner property value. Information about the owner of the shared item being referenced.
func (m *SharedDriveItem) GetOwner()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetPermission gets the permission property value. Used to access the permission representing the underlying sharing link
func (m *SharedDriveItem) GetPermission()(*Permission) {
    if m == nil {
        return nil
    } else {
        return m.permission
    }
}
// GetRoot gets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
func (m *SharedDriveItem) GetRoot()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSite gets the site property value. Used to access the underlying site
func (m *SharedDriveItem) GetSite()(*Site) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDriveItem sets the driveItem property value. Used to access the underlying driveItem
func (m *SharedDriveItem) SetDriveItem(value *DriveItem)() {
    if m != nil {
        m.driveItem = value
    }
}
// SetItems sets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
func (m *SharedDriveItem) SetItems(value []DriveItem)() {
    if m != nil {
        m.items = value
    }
}
// SetList sets the list property value. Used to access the underlying list
func (m *SharedDriveItem) SetList(value *List)() {
    if m != nil {
        m.list = value
    }
}
// SetListItem sets the listItem property value. Used to access the underlying listItem
func (m *SharedDriveItem) SetListItem(value *ListItem)() {
    if m != nil {
        m.listItem = value
    }
}
// SetOwner sets the owner property value. Information about the owner of the shared item being referenced.
func (m *SharedDriveItem) SetOwner(value *IdentitySet)() {
    if m != nil {
        m.owner = value
    }
}
// SetPermission sets the permission property value. Used to access the permission representing the underlying sharing link
func (m *SharedDriveItem) SetPermission(value *Permission)() {
    if m != nil {
        m.permission = value
    }
}
// SetRoot sets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
func (m *SharedDriveItem) SetRoot(value *DriveItem)() {
    if m != nil {
        m.root = value
    }
}
// SetSite sets the site property value. Used to access the underlying site
func (m *SharedDriveItem) SetSite(value *Site)() {
    if m != nil {
        m.site = value
    }
}
