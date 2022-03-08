package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedDriveItem provides operations to manage the collection of sharedDriveItem entities.
type SharedDriveItem struct {
    BaseItem
    // Used to access the underlying driveItem
    driveItem DriveItemable;
    // All driveItems contained in the sharing root. This collection cannot be enumerated.
    items []DriveItemable;
    // Used to access the underlying list
    list Listable;
    // Used to access the underlying listItem
    listItem ListItemable;
    // Information about the owner of the shared item being referenced.
    owner IdentitySetable;
    // Used to access the permission representing the underlying sharing link
    permission Permissionable;
    // Used to access the underlying driveItem. Deprecated -- use driveItem instead.
    root DriveItemable;
    // Used to access the underlying site
    site Siteable;
}
// NewSharedDriveItem instantiates a new sharedDriveItem and sets the default values.
func NewSharedDriveItem()(*SharedDriveItem) {
    m := &SharedDriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateSharedDriveItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedDriveItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSharedDriveItem(), nil
}
// GetDriveItem gets the driveItem property value. Used to access the underlying driveItem
func (m *SharedDriveItem) GetDriveItem()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedDriveItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(Listable))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(IdentitySetable))
        }
        return nil
    }
    res["permission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val.(Permissionable))
        }
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(DriveItemable))
        }
        return nil
    }
    res["site"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val.(Siteable))
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
func (m *SharedDriveItem) GetItems()([]DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. Used to access the underlying list
func (m *SharedDriveItem) GetList()(Listable) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetListItem gets the listItem property value. Used to access the underlying listItem
func (m *SharedDriveItem) GetListItem()(ListItemable) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetOwner gets the owner property value. Information about the owner of the shared item being referenced.
func (m *SharedDriveItem) GetOwner()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetPermission gets the permission property value. Used to access the permission representing the underlying sharing link
func (m *SharedDriveItem) GetPermission()(Permissionable) {
    if m == nil {
        return nil
    } else {
        return m.permission
    }
}
// GetRoot gets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
func (m *SharedDriveItem) GetRoot()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSite gets the site property value. Used to access the underlying site
func (m *SharedDriveItem) GetSite()(Siteable) {
    if m == nil {
        return nil
    } else {
        return m.site
    }
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
    if m.GetItems() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *SharedDriveItem) SetDriveItem(value DriveItemable)() {
    if m != nil {
        m.driveItem = value
    }
}
// SetItems sets the items property value. All driveItems contained in the sharing root. This collection cannot be enumerated.
func (m *SharedDriveItem) SetItems(value []DriveItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetList sets the list property value. Used to access the underlying list
func (m *SharedDriveItem) SetList(value Listable)() {
    if m != nil {
        m.list = value
    }
}
// SetListItem sets the listItem property value. Used to access the underlying listItem
func (m *SharedDriveItem) SetListItem(value ListItemable)() {
    if m != nil {
        m.listItem = value
    }
}
// SetOwner sets the owner property value. Information about the owner of the shared item being referenced.
func (m *SharedDriveItem) SetOwner(value IdentitySetable)() {
    if m != nil {
        m.owner = value
    }
}
// SetPermission sets the permission property value. Used to access the permission representing the underlying sharing link
func (m *SharedDriveItem) SetPermission(value Permissionable)() {
    if m != nil {
        m.permission = value
    }
}
// SetRoot sets the root property value. Used to access the underlying driveItem. Deprecated -- use driveItem instead.
func (m *SharedDriveItem) SetRoot(value DriveItemable)() {
    if m != nil {
        m.root = value
    }
}
// SetSite sets the site property value. Used to access the underlying site
func (m *SharedDriveItem) SetSite(value Siteable)() {
    if m != nil {
        m.site = value
    }
}
