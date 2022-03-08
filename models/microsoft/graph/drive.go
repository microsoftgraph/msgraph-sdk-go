package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Drive provides operations to manage the drive singleton.
type Drive struct {
    BaseItem
    // Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
    bundles []DriveItemable;
    // Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
    driveType *string;
    // The list of items the user is following. Only in OneDrive for Business.
    following []DriveItemable;
    // All items contained in the drive. Read-only. Nullable.
    items []DriveItemable;
    // For drives in SharePoint, the underlying document library list. Read-only. Nullable.
    list Listable;
    // Optional. The user account that owns the drive. Read-only.
    owner IdentitySetable;
    // Optional. Information about the drive's storage space quota. Read-only.
    quota Quotaable;
    // The root folder of the drive. Read-only.
    root DriveItemable;
    // 
    sharePointIds SharepointIdsable;
    // Collection of common folders available in OneDrive. Read-only. Nullable.
    special []DriveItemable;
    // If present, indicates that this is a system-managed drive. Read-only.
    system SystemFacetable;
}
// NewDrive instantiates a new drive and sets the default values.
func NewDrive()(*Drive) {
    m := &Drive{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateDriveFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDrive(), nil
}
// GetBundles gets the bundles property value. Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *Drive) GetBundles()([]DriveItemable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *Drive) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["bundles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
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
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetFollowing(res)
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
    res["quota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateQuotaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val.(Quotaable))
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
    res["sharePointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["special"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetSpecial(res)
        }
        return nil
    }
    res["system"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSystemFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(SystemFacetable))
        }
        return nil
    }
    return res
}
// GetFollowing gets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) GetFollowing()([]DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.following
    }
}
// GetItems gets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) GetItems()([]DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) GetList()(Listable) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetOwner gets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) GetOwner()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetQuota gets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) GetQuota()(Quotaable) {
    if m == nil {
        return nil
    } else {
        return m.quota
    }
}
// GetRoot gets the root property value. The root folder of the drive. Read-only.
func (m *Drive) GetRoot()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSharePointIds gets the sharePointIds property value. 
func (m *Drive) GetSharePointIds()(SharepointIdsable) {
    if m == nil {
        return nil
    } else {
        return m.sharePointIds
    }
}
// GetSpecial gets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) GetSpecial()([]DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.special
    }
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) GetSystem()(SystemFacetable) {
    if m == nil {
        return nil
    } else {
        return m.system
    }
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
    if m.GetBundles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBundles()))
        for i, v := range m.GetBundles() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetFollowing() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFollowing()))
        for i, v := range m.GetFollowing() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("following", cast)
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
    if m.GetSpecial() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSpecial()))
        for i, v := range m.GetSpecial() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *Drive) SetBundles(value []DriveItemable)() {
    if m != nil {
        m.bundles = value
    }
}
// SetDriveType sets the driveType property value. Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
func (m *Drive) SetDriveType(value *string)() {
    if m != nil {
        m.driveType = value
    }
}
// SetFollowing sets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) SetFollowing(value []DriveItemable)() {
    if m != nil {
        m.following = value
    }
}
// SetItems sets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) SetItems(value []DriveItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetList sets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) SetList(value Listable)() {
    if m != nil {
        m.list = value
    }
}
// SetOwner sets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) SetOwner(value IdentitySetable)() {
    if m != nil {
        m.owner = value
    }
}
// SetQuota sets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) SetQuota(value Quotaable)() {
    if m != nil {
        m.quota = value
    }
}
// SetRoot sets the root property value. The root folder of the drive. Read-only.
func (m *Drive) SetRoot(value DriveItemable)() {
    if m != nil {
        m.root = value
    }
}
// SetSharePointIds sets the sharePointIds property value. 
func (m *Drive) SetSharePointIds(value SharepointIdsable)() {
    if m != nil {
        m.sharePointIds = value
    }
}
// SetSpecial sets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) SetSpecial(value []DriveItemable)() {
    if m != nil {
        m.special = value
    }
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) SetSystem(value SystemFacetable)() {
    if m != nil {
        m.system = value
    }
}
