package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// drive 
type Drive struct {
    BaseItem
    // Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
    bundles []DriveItem;
    // Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
    driveType *string;
    // The list of items the user is following. Only in OneDrive for Business.
    following []DriveItem;
    // All items contained in the drive. Read-only. Nullable.
    items []DriveItem;
    // For drives in SharePoint, the underlying document library list. Read-only. Nullable.
    list *List;
    // Optional. The user account that owns the drive. Read-only.
    owner *IdentitySet;
    // Optional. Information about the drive's storage space quota. Read-only.
    quota *Quota;
    // The root folder of the drive. Read-only.
    root *DriveItem;
    // 
    sharePointIds *SharepointIds;
    // Collection of common folders available in OneDrive. Read-only. Nullable.
    special []DriveItem;
    // If present, indicates that this is a system-managed drive. Read-only.
    system *SystemFacet;
}
// NewDrive instantiates a new drive and sets the default values.
func NewDrive()(*Drive) {
    m := &Drive{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// GetBundles gets the bundles property value. Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *Drive) GetBundles()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.bundles
    }
}
// GetDriveType gets the driveType property value. Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
func (m *Drive) GetDriveType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.driveType
    }
}
// GetFollowing gets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) GetFollowing()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.following
    }
}
// GetItems gets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) GetItems()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) GetList()(*List) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetOwner gets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) GetOwner()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetQuota gets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) GetQuota()(*Quota) {
    if m == nil {
        return nil
    } else {
        return m.quota
    }
}
// GetRoot gets the root property value. The root folder of the drive. Read-only.
func (m *Drive) GetRoot()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSharePointIds gets the sharePointIds property value. 
func (m *Drive) GetSharePointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharePointIds
    }
}
// GetSpecial gets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) GetSpecial()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.special
    }
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) GetSystem()(*SystemFacet) {
    if m == nil {
        return nil
    } else {
        return m.system
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Drive) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["bundles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItem))
            }
            m.SetBundles(res)
        }
        return nil
    }
    res["driveType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveType(val)
        }
        return nil
    }
    res["following"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItem))
            }
            m.SetFollowing(res)
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
    res["quota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewQuota() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val.(*Quota))
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
    res["sharePointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharepointIds() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointIds(val.(*SharepointIds))
        }
        return nil
    }
    res["special"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItem))
            }
            m.SetSpecial(res)
        }
        return nil
    }
    res["system"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSystemFacet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(*SystemFacet))
        }
        return nil
    }
    return res
}
func (m *Drive) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Drive) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBundles()))
        for i, v := range m.GetBundles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("bundles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("driveType", m.GetDriveType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFollowing()))
        for i, v := range m.GetFollowing() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("following", cast)
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
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("quota", m.GetQuota())
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
        err = writer.WriteObjectValue("sharePointIds", m.GetSharePointIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSpecial()))
        for i, v := range m.GetSpecial() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("special", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundles sets the bundles property value. Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *Drive) SetBundles(value []DriveItem)() {
    m.bundles = value
}
// SetDriveType sets the driveType property value. Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
func (m *Drive) SetDriveType(value *string)() {
    m.driveType = value
}
// SetFollowing sets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) SetFollowing(value []DriveItem)() {
    m.following = value
}
// SetItems sets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) SetItems(value []DriveItem)() {
    m.items = value
}
// SetList sets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) SetList(value *List)() {
    m.list = value
}
// SetOwner sets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) SetOwner(value *IdentitySet)() {
    m.owner = value
}
// SetQuota sets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) SetQuota(value *Quota)() {
    m.quota = value
}
// SetRoot sets the root property value. The root folder of the drive. Read-only.
func (m *Drive) SetRoot(value *DriveItem)() {
    m.root = value
}
// SetSharePointIds sets the sharePointIds property value. 
func (m *Drive) SetSharePointIds(value *SharepointIds)() {
    m.sharePointIds = value
}
// SetSpecial sets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) SetSpecial(value []DriveItem)() {
    m.special = value
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) SetSystem(value *SystemFacet)() {
    m.system = value
}
