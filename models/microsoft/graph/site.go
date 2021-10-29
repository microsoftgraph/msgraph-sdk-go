package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Site struct {
    BaseItem
    // Analytics about the view activities that took place in this site.
    analytics *ItemAnalytics;
    // The collection of column definitions reusable across lists under this site.
    columns []ColumnDefinition;
    // The collection of content types defined for this site.
    contentTypes []ContentType;
    // The full title for the site. Read-only.
    displayName *string;
    // The default drive (document library) for this site.
    drive *Drive;
    // The collection of drives (document libraries) under this site.
    drives []Drive;
    // 
    error *PublicError;
    // The collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
    externalColumns []ColumnDefinition;
    // Used to address any item contained in this site. This collection can't be enumerated.
    items []BaseItem;
    // The collection of lists under this site.
    lists []List;
    // Calls the OneNote service for notebook related operations.
    onenote *Onenote;
    // The permissions associated with the site. Nullable.
    permissions []Permission;
    // If present, indicates that this is the root site in the site collection. Read-only.
    root *Root;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds *SharepointIds;
    // Provides details about the site's site collection. Available only on the root site. Read-only.
    siteCollection *SiteCollection;
    // The collection of the sub-sites under this site.
    sites []Site;
    // The default termStore under this site.
    termStore *Store;
    // The collection of termStores under this site.
    termStores []Store;
}
// Instantiates a new site and sets the default values.
func NewSite()(*Site) {
    m := &Site{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// Gets the analytics property value. Analytics about the view activities that took place in this site.
func (m *Site) GetAnalytics()(*ItemAnalytics) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// Gets the columns property value. The collection of column definitions reusable across lists under this site.
func (m *Site) GetColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// Gets the contentTypes property value. The collection of content types defined for this site.
func (m *Site) GetContentTypes()([]ContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
// Gets the displayName property value. The full title for the site. Read-only.
func (m *Site) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the drive property value. The default drive (document library) for this site.
func (m *Site) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// Gets the drives property value. The collection of drives (document libraries) under this site.
func (m *Site) GetDrives()([]Drive) {
    if m == nil {
        return nil
    } else {
        return m.drives
    }
}
// Gets the error property value. 
func (m *Site) GetError()(*PublicError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the externalColumns property value. The collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
func (m *Site) GetExternalColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.externalColumns
    }
}
// Gets the items property value. Used to address any item contained in this site. This collection can't be enumerated.
func (m *Site) GetItems()([]BaseItem) {
    if m == nil {
        return nil
    } else {
        return m.items
    }
}
// Gets the lists property value. The collection of lists under this site.
func (m *Site) GetLists()([]List) {
    if m == nil {
        return nil
    } else {
        return m.lists
    }
}
// Gets the onenote property value. Calls the OneNote service for notebook related operations.
func (m *Site) GetOnenote()(*Onenote) {
    if m == nil {
        return nil
    } else {
        return m.onenote
    }
}
// Gets the permissions property value. The permissions associated with the site. Nullable.
func (m *Site) GetPermissions()([]Permission) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// Gets the root property value. If present, indicates that this is the root site in the site collection. Read-only.
func (m *Site) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// Gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *Site) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// Gets the siteCollection property value. Provides details about the site's site collection. Available only on the root site. Read-only.
func (m *Site) GetSiteCollection()(*SiteCollection) {
    if m == nil {
        return nil
    } else {
        return m.siteCollection
    }
}
// Gets the sites property value. The collection of the sub-sites under this site.
func (m *Site) GetSites()([]Site) {
    if m == nil {
        return nil
    } else {
        return m.sites
    }
}
// Gets the termStore property value. The default termStore under this site.
func (m *Site) GetTermStore()(*Store) {
    if m == nil {
        return nil
    } else {
        return m.termStore
    }
}
// Gets the termStores property value. The collection of termStores under this site.
func (m *Site) GetTermStores()([]Store) {
    if m == nil {
        return nil
    } else {
        return m.termStores
    }
}
// The deserialization information for the current model
func (m *Site) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAnalytics() })
        if err != nil {
            return err
        }
        m.SetAnalytics(val.(*ItemAnalytics))
        return nil
    }
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetColumns(res)
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        res := make([]ContentType, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContentType))
        }
        m.SetContentTypes(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        m.SetDrive(val.(*Drive))
        return nil
    }
    res["drives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        res := make([]Drive, len(val))
        for i, v := range val {
            res[i] = *(v.(*Drive))
        }
        m.SetDrives(res)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*PublicError))
        return nil
    }
    res["externalColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetExternalColumns(res)
        return nil
    }
    res["items"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBaseItem() })
        if err != nil {
            return err
        }
        res := make([]BaseItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*BaseItem))
        }
        m.SetItems(res)
        return nil
    }
    res["lists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewList() })
        if err != nil {
            return err
        }
        res := make([]List, len(val))
        for i, v := range val {
            res[i] = *(v.(*List))
        }
        m.SetLists(res)
        return nil
    }
    res["onenote"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenote() })
        if err != nil {
            return err
        }
        m.SetOnenote(val.(*Onenote))
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermission() })
        if err != nil {
            return err
        }
        res := make([]Permission, len(val))
        for i, v := range val {
            res[i] = *(v.(*Permission))
        }
        m.SetPermissions(res)
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoot() })
        if err != nil {
            return err
        }
        m.SetRoot(val.(*Root))
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
    res["siteCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSiteCollection() })
        if err != nil {
            return err
        }
        m.SetSiteCollection(val.(*SiteCollection))
        return nil
    }
    res["sites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        res := make([]Site, len(val))
        for i, v := range val {
            res[i] = *(v.(*Site))
        }
        m.SetSites(res)
        return nil
    }
    res["termStore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStore() })
        if err != nil {
            return err
        }
        m.SetTermStore(val.(*Store))
        return nil
    }
    res["termStores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStore() })
        if err != nil {
            return err
        }
        res := make([]Store, len(val))
        for i, v := range val {
            res[i] = *(v.(*Store))
        }
        m.SetTermStores(res)
        return nil
    }
    return res
}
func (m *Site) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Site) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExternalColumns()))
        for i, v := range m.GetExternalColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("externalColumns", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
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
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteCollection", m.GetSiteCollection())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termStore", m.GetTermStore())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTermStores()))
        for i, v := range m.GetTermStores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("termStores", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the analytics property value. Analytics about the view activities that took place in this site.
// Parameters:
//  - value : Value to set for the analytics property.
func (m *Site) SetAnalytics(value *ItemAnalytics)() {
    m.analytics = value
}
// Sets the columns property value. The collection of column definitions reusable across lists under this site.
// Parameters:
//  - value : Value to set for the columns property.
func (m *Site) SetColumns(value []ColumnDefinition)() {
    m.columns = value
}
// Sets the contentTypes property value. The collection of content types defined for this site.
// Parameters:
//  - value : Value to set for the contentTypes property.
func (m *Site) SetContentTypes(value []ContentType)() {
    m.contentTypes = value
}
// Sets the displayName property value. The full title for the site. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Site) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the drive property value. The default drive (document library) for this site.
// Parameters:
//  - value : Value to set for the drive property.
func (m *Site) SetDrive(value *Drive)() {
    m.drive = value
}
// Sets the drives property value. The collection of drives (document libraries) under this site.
// Parameters:
//  - value : Value to set for the drives property.
func (m *Site) SetDrives(value []Drive)() {
    m.drives = value
}
// Sets the error property value. 
// Parameters:
//  - value : Value to set for the error property.
func (m *Site) SetError(value *PublicError)() {
    m.error = value
}
// Sets the externalColumns property value. The collection of column definitions available in the site that are referenced from the sites in the parent hierarchy of the current site.
// Parameters:
//  - value : Value to set for the externalColumns property.
func (m *Site) SetExternalColumns(value []ColumnDefinition)() {
    m.externalColumns = value
}
// Sets the items property value. Used to address any item contained in this site. This collection can't be enumerated.
// Parameters:
//  - value : Value to set for the items property.
func (m *Site) SetItems(value []BaseItem)() {
    m.items = value
}
// Sets the lists property value. The collection of lists under this site.
// Parameters:
//  - value : Value to set for the lists property.
func (m *Site) SetLists(value []List)() {
    m.lists = value
}
// Sets the onenote property value. Calls the OneNote service for notebook related operations.
// Parameters:
//  - value : Value to set for the onenote property.
func (m *Site) SetOnenote(value *Onenote)() {
    m.onenote = value
}
// Sets the permissions property value. The permissions associated with the site. Nullable.
// Parameters:
//  - value : Value to set for the permissions property.
func (m *Site) SetPermissions(value []Permission)() {
    m.permissions = value
}
// Sets the root property value. If present, indicates that this is the root site in the site collection. Read-only.
// Parameters:
//  - value : Value to set for the root property.
func (m *Site) SetRoot(value *Root)() {
    m.root = value
}
// Sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
// Parameters:
//  - value : Value to set for the sharepointIds property.
func (m *Site) SetSharepointIds(value *SharepointIds)() {
    m.sharepointIds = value
}
// Sets the siteCollection property value. Provides details about the site's site collection. Available only on the root site. Read-only.
// Parameters:
//  - value : Value to set for the siteCollection property.
func (m *Site) SetSiteCollection(value *SiteCollection)() {
    m.siteCollection = value
}
// Sets the sites property value. The collection of the sub-sites under this site.
// Parameters:
//  - value : Value to set for the sites property.
func (m *Site) SetSites(value []Site)() {
    m.sites = value
}
// Sets the termStore property value. The default termStore under this site.
// Parameters:
//  - value : Value to set for the termStore property.
func (m *Site) SetTermStore(value *Store)() {
    m.termStore = value
}
// Sets the termStores property value. The collection of termStores under this site.
// Parameters:
//  - value : Value to set for the termStores property.
func (m *Site) SetTermStores(value []Store)() {
    m.termStores = value
}
