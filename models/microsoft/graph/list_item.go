package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ListItem struct {
    BaseItem
    // Analytics about the view activities that took place on this item.
    analytics *ItemAnalytics;
    // The content type of this list item
    contentType *ContentTypeInfo;
    // For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
    driveItem *DriveItem;
    // The values of the columns set on this list item.
    fields *FieldValueSet;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds *SharepointIds;
    // The list of previous versions of the list item.
    versions []ListItemVersion;
}
// Instantiates a new listItem and sets the default values.
func NewListItem()(*ListItem) {
    m := &ListItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// Gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *ListItem) GetAnalytics()(*ItemAnalytics) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// Gets the contentType property value. The content type of this list item
func (m *ListItem) GetContentType()(*ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
func (m *ListItem) GetDriveItem()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// Gets the fields property value. The values of the columns set on this list item.
func (m *ListItem) GetFields()(*FieldValueSet) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// Gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *ListItem) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// Gets the versions property value. The list of previous versions of the list item.
func (m *ListItem) GetVersions()([]ListItemVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// The deserialization information for the current model
func (m *ListItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAnalytics() })
        if err != nil {
            return err
        }
        m.SetAnalytics(val.(*ItemAnalytics))
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        m.SetContentType(val.(*ContentTypeInfo))
        return nil
    }
    res["driveItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        m.SetDriveItem(val.(*DriveItem))
        return nil
    }
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFieldValueSet() })
        if err != nil {
            return err
        }
        m.SetFields(val.(*FieldValueSet))
        return nil
    }
    res["sharepointIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharepointIds() })
        if err != nil {
            return err
        }
        m.SetSharepointIds(val.(*SharepointIds))
        return nil
    }
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItemVersion() })
        if err != nil {
            return err
        }
        res := make([]ListItemVersion, len(val))
        for i, v := range val {
            res[i] = *(v.(*ListItemVersion))
        }
        m.SetVersions(res)
        return nil
    }
    return res
}
func (m *ListItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ListItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("analytics", m.GetAnalytics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fields", m.GetFields())
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the analytics property value. Analytics about the view activities that took place on this item.
// Parameters:
//  - value : Value to set for the analytics property.
func (m *ListItem) SetAnalytics(value *ItemAnalytics)() {
    m.analytics = value
}
// Sets the contentType property value. The content type of this list item
// Parameters:
//  - value : Value to set for the contentType property.
func (m *ListItem) SetContentType(value *ContentTypeInfo)() {
    m.contentType = value
}
// Sets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
// Parameters:
//  - value : Value to set for the driveItem property.
func (m *ListItem) SetDriveItem(value *DriveItem)() {
    m.driveItem = value
}
// Sets the fields property value. The values of the columns set on this list item.
// Parameters:
//  - value : Value to set for the fields property.
func (m *ListItem) SetFields(value *FieldValueSet)() {
    m.fields = value
}
// Sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
// Parameters:
//  - value : Value to set for the sharepointIds property.
func (m *ListItem) SetSharepointIds(value *SharepointIds)() {
    m.sharepointIds = value
}
// Sets the versions property value. The list of previous versions of the list item.
// Parameters:
//  - value : Value to set for the versions property.
func (m *ListItem) SetVersions(value []ListItemVersion)() {
    m.versions = value
}
