package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Drive struct {
    BaseItem
    driveType *string;
    following []DriveItem;
    items []DriveItem;
    list *List;
    owner *IdentitySet;
    quota *Quota;
    root *DriveItem;
    sharePointIds *SharepointIds;
    special []DriveItem;
    system *SystemFacet;
}
func NewDrive()(*Drive) {
    m := &Drive{
        BaseItem: *NewBaseItem(),
    }
    return m
}
func (m *Drive) GetDriveType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.driveType
    }
}
func (m *Drive) GetFollowing()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.following
    }
}
func (m *Drive) GetItems()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
func (m *Drive) GetList()(*List) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
func (m *Drive) GetOwner()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
func (m *Drive) GetQuota()(*Quota) {
    if m == nil {
        return nil
    } else {
        return m.quota
    }
}
func (m *Drive) GetRoot()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
func (m *Drive) GetSharePointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharePointIds
    }
}
func (m *Drive) GetSpecial()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.special
    }
}
func (m *Drive) GetSystem()(*SystemFacet) {
    if m == nil {
        return nil
    } else {
        return m.system
    }
}
func (m *Drive) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["driveType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDriveType(val)
        return nil
    }
    res["following"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        res := make([]DriveItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*DriveItem))
        }
        m.SetFollowing(res)
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        res := make([]DriveItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*DriveItem))
        }
        m.SetItems(res)
        return nil
    }
    res["list"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewList() })
        if err != nil {
            return err
        }
        m.SetList(val.(*List))
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetOwner(val.(*IdentitySet))
        return nil
    }
    res["quota"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewQuota() })
        if err != nil {
            return err
        }
        m.SetQuota(val.(*Quota))
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        m.SetRoot(val.(*DriveItem))
        return nil
    }
    res["sharePointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharepointIds() })
        if err != nil {
            return err
        }
        m.SetSharePointIds(val.(*SharepointIds))
        return nil
    }
    res["special"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        res := make([]DriveItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*DriveItem))
        }
        m.SetSpecial(res)
        return nil
    }
    res["system"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSystemFacet() })
        if err != nil {
            return err
        }
        m.SetSystem(val.(*SystemFacet))
        return nil
    }
    return res
}
func (m *Drive) IsNil()(bool) {
    return m == nil
}
func (m *Drive) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
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
func (m *Drive) SetDriveType(value *string)() {
    m.driveType = value
}
func (m *Drive) SetFollowing(value []DriveItem)() {
    m.following = value
}
func (m *Drive) SetItems(value []DriveItem)() {
    m.items = value
}
func (m *Drive) SetList(value *List)() {
    m.list = value
}
func (m *Drive) SetOwner(value *IdentitySet)() {
    m.owner = value
}
func (m *Drive) SetQuota(value *Quota)() {
    m.quota = value
}
func (m *Drive) SetRoot(value *DriveItem)() {
    m.root = value
}
func (m *Drive) SetSharePointIds(value *SharepointIds)() {
    m.sharePointIds = value
}
func (m *Drive) SetSpecial(value []DriveItem)() {
    m.special = value
}
func (m *Drive) SetSystem(value *SystemFacet)() {
    m.system = value
}
