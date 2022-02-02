package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// List 
type List struct {
    BaseItem
    // The collection of field definitions for this list.
    columns []ColumnDefinition;
    // The collection of content types present in this list.
    contentTypes []ContentType;
    // The displayable title of the list.
    displayName *string;
    // Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
    drive *Drive;
    // All items contained in the list.
    items []ListItem;
    // Provides additional details about the list.
    list *ListInfo;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds *SharepointIds;
    // The set of subscriptions on the list.
    subscriptions []Subscription;
    // If present, indicates that this is a system-managed list. Read-only.
    system *SystemFacet;
}
// NewList instantiates a new list and sets the default values.
func NewList()(*List) {
    m := &List{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// GetColumns gets the columns property value. The collection of field definitions for this list.
func (m *List) GetColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// GetContentTypes gets the contentTypes property value. The collection of content types present in this list.
func (m *List) GetContentTypes()([]ContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
// GetDisplayName gets the displayName property value. The displayable title of the list.
func (m *List) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDrive gets the drive property value. Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
func (m *List) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// GetItems gets the items property value. All items contained in the list.
func (m *List) GetItems()([]ListItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. Provides additional details about the list.
func (m *List) GetList()(*ListInfo) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the list.
func (m *List) GetSubscriptions()([]Subscription) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) GetSystem()(*SystemFacet) {
    if m == nil {
        return nil
    } else {
        return m.system
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *List) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnDefinition))
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ContentType))
            }
            m.SetContentTypes(res)
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
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(*Drive))
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*ListItem))
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(*ListInfo))
        }
        return nil
    }
    res["sharepointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharepointIds() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(*SharepointIds))
        }
        return nil
    }
    res["subscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubscription() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Subscription, len(val))
            for i, v := range val {
                res[i] = *(v.(*Subscription))
            }
            m.SetSubscriptions(res)
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
func (m *List) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *List) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContentTypes()))
        for i, v := range m.GetContentTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("contentTypes", cast)
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
    {
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
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
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
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
// SetColumns sets the columns property value. The collection of field definitions for this list.
func (m *List) SetColumns(value []ColumnDefinition)() {
    if m != nil {
        m.columns = value
    }
}
// SetContentTypes sets the contentTypes property value. The collection of content types present in this list.
func (m *List) SetContentTypes(value []ContentType)() {
    if m != nil {
        m.contentTypes = value
    }
}
// SetDisplayName sets the displayName property value. The displayable title of the list.
func (m *List) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDrive sets the drive property value. Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
func (m *List) SetDrive(value *Drive)() {
    if m != nil {
        m.drive = value
    }
}
// SetItems sets the items property value. All items contained in the list.
func (m *List) SetItems(value []ListItem)() {
    if m != nil {
        m.items = value
    }
}
// SetList sets the list property value. Provides additional details about the list.
func (m *List) SetList(value *ListInfo)() {
    if m != nil {
        m.list = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) SetSharepointIds(value *SharepointIds)() {
    if m != nil {
        m.sharepointIds = value
    }
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the list.
func (m *List) SetSubscriptions(value []Subscription)() {
    if m != nil {
        m.subscriptions = value
    }
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) SetSystem(value *SystemFacet)() {
    if m != nil {
        m.system = value
    }
}
