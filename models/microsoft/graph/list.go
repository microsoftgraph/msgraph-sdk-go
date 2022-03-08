package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// List provides operations to manage the drive singleton.
type List struct {
    BaseItem
    // The collection of field definitions for this list.
    columns []ColumnDefinitionable;
    // The collection of content types present in this list.
    contentTypes []ContentTypeable;
    // The displayable title of the list.
    displayName *string;
    // Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
    drive Driveable;
    // All items contained in the list.
    items []ListItemable;
    // Provides additional details about the list.
    list ListInfoable;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds SharepointIdsable;
    // The set of subscriptions on the list.
    subscriptions []Subscriptionable;
    // If present, indicates that this is a system-managed list. Read-only.
    system SystemFacetable;
}
// NewList instantiates a new list and sets the default values.
func NewList()(*List) {
    m := &List{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewList(), nil
}
// GetColumns gets the columns property value. The collection of field definitions for this list.
func (m *List) GetColumns()([]ColumnDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// GetContentTypes gets the contentTypes property value. The collection of content types present in this list.
func (m *List) GetContentTypes()([]ContentTypeable) {
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
func (m *List) GetDrive()(Driveable) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *List) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateColumnDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ColumnDefinitionable)
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContentTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ContentTypeable)
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
        val, err := n.GetObjectValue(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(Driveable))
        }
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemable, len(val))
            for i, v := range val {
                res[i] = v.(ListItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateListInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(ListInfoable))
        }
        return nil
    }
    res["sharepointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["subscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubscriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Subscriptionable, len(val))
            for i, v := range val {
                res[i] = v.(Subscriptionable)
            }
            m.SetSubscriptions(res)
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
// GetItems gets the items property value. All items contained in the list.
func (m *List) GetItems()([]ListItemable) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// GetList gets the list property value. Provides additional details about the list.
func (m *List) GetList()(ListInfoable) {
    if m == nil {
        return nil
    } else {
        return m.list
    }
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) GetSharepointIds()(SharepointIdsable) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the list.
func (m *List) GetSubscriptions()([]Subscriptionable) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) GetSystem()(SystemFacetable) {
    if m == nil {
        return nil
    } else {
        return m.system
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContentTypes()))
        for i, v := range m.GetContentTypes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *List) SetColumns(value []ColumnDefinitionable)() {
    if m != nil {
        m.columns = value
    }
}
// SetContentTypes sets the contentTypes property value. The collection of content types present in this list.
func (m *List) SetContentTypes(value []ContentTypeable)() {
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
func (m *List) SetDrive(value Driveable)() {
    if m != nil {
        m.drive = value
    }
}
// SetItems sets the items property value. All items contained in the list.
func (m *List) SetItems(value []ListItemable)() {
    if m != nil {
        m.items = value
    }
}
// SetList sets the list property value. Provides additional details about the list.
func (m *List) SetList(value ListInfoable)() {
    if m != nil {
        m.list = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) SetSharepointIds(value SharepointIdsable)() {
    if m != nil {
        m.sharepointIds = value
    }
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the list.
func (m *List) SetSubscriptions(value []Subscriptionable)() {
    if m != nil {
        m.subscriptions = value
    }
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) SetSystem(value SystemFacetable)() {
    if m != nil {
        m.system = value
    }
}
