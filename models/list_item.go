package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListItem 
type ListItem struct {
    BaseItem
    // Analytics about the view activities that took place on this item.
    analytics ItemAnalyticsable
    // The content type of this list item
    contentType ContentTypeInfoable
    // For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
    driveItem DriveItemable
    // The values of the columns set on this list item.
    fields FieldValueSetable
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds SharepointIdsable
    // The list of previous versions of the list item.
    versions []ListItemVersionable
}
// NewListItem instantiates a new listItem and sets the default values.
func NewListItem()(*ListItem) {
    m := &ListItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListItem(), nil
}
// GetAnalytics gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *ListItem) GetAnalytics()(ItemAnalyticsable) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// GetContentType gets the contentType property value. The content type of this list item
func (m *ListItem) GetContentType()(ContentTypeInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetDriveItem gets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
func (m *ListItem) GetDriveItem()(DriveItemable) {
    if m == nil {
        return nil
    } else {
        return m.driveItem
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(ItemAnalyticsable))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["driveItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFieldValueSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(FieldValueSetable))
        }
        return nil
    }
    res["sharepointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemVersionable, len(val))
            for i, v := range val {
                res[i] = v.(ListItemVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The values of the columns set on this list item.
func (m *ListItem) GetFields()(FieldValueSetable) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *ListItem) GetSharepointIds()(SharepointIdsable) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// GetVersions gets the versions property value. The list of previous versions of the list item.
func (m *ListItem) GetVersions()([]ListItemVersionable) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// Serialize serializes information the current object
func (m *ListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnalytics sets the analytics property value. Analytics about the view activities that took place on this item.
func (m *ListItem) SetAnalytics(value ItemAnalyticsable)() {
    if m != nil {
        m.analytics = value
    }
}
// SetContentType sets the contentType property value. The content type of this list item
func (m *ListItem) SetContentType(value ContentTypeInfoable)() {
    if m != nil {
        m.contentType = value
    }
}
// SetDriveItem sets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
func (m *ListItem) SetDriveItem(value DriveItemable)() {
    if m != nil {
        m.driveItem = value
    }
}
// SetFields sets the fields property value. The values of the columns set on this list item.
func (m *ListItem) SetFields(value FieldValueSetable)() {
    if m != nil {
        m.fields = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *ListItem) SetSharepointIds(value SharepointIdsable)() {
    if m != nil {
        m.sharepointIds = value
    }
}
// SetVersions sets the versions property value. The list of previous versions of the list item.
func (m *ListItem) SetVersions(value []ListItemVersionable)() {
    if m != nil {
        m.versions = value
    }
}
