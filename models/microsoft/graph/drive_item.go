package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DriveItem struct {
    BaseItem
    // Analytics about the view activities that took place on this item.
    analytics *ItemAnalytics;
    // Audio metadata, if the item is an audio file. Read-only.
    audio *Audio;
    // 
    bundle *Bundle;
    // Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
    children []DriveItem;
    // The content stream, if the item represents a file.
    content []byte;
    // An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
    cTag *string;
    // Information about the deleted state of the item. Read-only.
    deleted *Deleted;
    // File metadata, if the item is a file. Read-only.
    file *File;
    // File system information on client. Read-write.
    fileSystemInfo *FileSystemInfo;
    // Folder metadata, if the item is a folder. Read-only.
    folder *Folder;
    // Image metadata, if the item is an image. Read-only.
    image *Image;
    // For drives in SharePoint, the associated document library list item. Read-only. Nullable.
    listItem *ListItem;
    // Location metadata, if the item has location data. Read-only.
    location *GeoCoordinates;
    // Malware metadata, if the item was detected to contain malware. Read-only.
    malware *Malware;
    // If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
    package_escaped *Package_escaped;
    // If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
    pendingOperations *PendingOperations;
    // The set of permissions for the item. Read-only. Nullable.
    permissions []Permission;
    // Photo metadata, if the item is a photo. Read-only.
    photo *Photo;
    // Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
    publication *PublicationFacet;
    // Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
    remoteItem *RemoteItem;
    // If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
    root *Root;
    // Search metadata, if the item is from a search result. Read-only.
    searchResult *SearchResult;
    // Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
    shared *Shared;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds *SharepointIds;
    // Size of the item in bytes. Read-only.
    size *int64;
    // If the current item is also available as a special folder, this facet is returned. Read-only.
    specialFolder *SpecialFolder;
    // The set of subscriptions on the item. Only supported on the root of a drive.
    subscriptions []Subscription;
    // Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
    thumbnails []ThumbnailSet;
    // The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
    versions []DriveItemVersion;
    // Video metadata, if the item is a video. Read-only.
    video *Video;
    // WebDAV compatible URL for the item.
    webDavUrl *string;
    // For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
    workbook *Workbook;
}
// Instantiates a new driveItem and sets the default values.
func NewDriveItem()(*DriveItem) {
    m := &DriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// Gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *DriveItem) GetAnalytics()(*ItemAnalytics) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// Gets the audio property value. Audio metadata, if the item is an audio file. Read-only.
func (m *DriveItem) GetAudio()(*Audio) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
// Gets the bundle property value. 
func (m *DriveItem) GetBundle()(*Bundle) {
    if m == nil {
        return nil
    } else {
        return m.bundle
    }
}
// Gets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) GetChildren()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// Gets the content property value. The content stream, if the item represents a file.
func (m *DriveItem) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the cTag property value. An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
func (m *DriveItem) GetCTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cTag
    }
}
// Gets the deleted property value. Information about the deleted state of the item. Read-only.
func (m *DriveItem) GetDeleted()(*Deleted) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
// Gets the file property value. File metadata, if the item is a file. Read-only.
func (m *DriveItem) GetFile()(*File) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// Gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItem) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// Gets the folder property value. Folder metadata, if the item is a folder. Read-only.
func (m *DriveItem) GetFolder()(*Folder) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
// Gets the image property value. Image metadata, if the item is an image. Read-only.
func (m *DriveItem) GetImage()(*Image) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
// Gets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *DriveItem) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// Gets the location property value. Location metadata, if the item has location data. Read-only.
func (m *DriveItem) GetLocation()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
func (m *DriveItem) GetMalware()(*Malware) {
    if m == nil {
        return nil
    } else {
        return m.malware
    }
}
// Gets the package_escaped property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *DriveItem) GetPackage_escaped()(*Package_escaped) {
    if m == nil {
        return nil
    } else {
        return m.package_escaped
    }
}
// Gets the pendingOperations property value. If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
func (m *DriveItem) GetPendingOperations()(*PendingOperations) {
    if m == nil {
        return nil
    } else {
        return m.pendingOperations
    }
}
// Gets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) GetPermissions()([]Permission) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// Gets the photo property value. Photo metadata, if the item is a photo. Read-only.
func (m *DriveItem) GetPhoto()(*Photo) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// Gets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
func (m *DriveItem) GetPublication()(*PublicationFacet) {
    if m == nil {
        return nil
    } else {
        return m.publication
    }
}
// Gets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
func (m *DriveItem) GetRemoteItem()(*RemoteItem) {
    if m == nil {
        return nil
    } else {
        return m.remoteItem
    }
}
// Gets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
func (m *DriveItem) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// Gets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
func (m *DriveItem) GetSearchResult()(*SearchResult) {
    if m == nil {
        return nil
    } else {
        return m.searchResult
    }
}
// Gets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *DriveItem) GetShared()(*Shared) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// Gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *DriveItem) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// Gets the size property value. Size of the item in bytes. Read-only.
func (m *DriveItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Gets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *DriveItem) GetSpecialFolder()(*SpecialFolder) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
// Gets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) GetSubscriptions()([]Subscription) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// Gets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) GetThumbnails()([]ThumbnailSet) {
    if m == nil {
        return nil
    } else {
        return m.thumbnails
    }
}
// Gets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) GetVersions()([]DriveItemVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// Gets the video property value. Video metadata, if the item is a video. Read-only.
func (m *DriveItem) GetVideo()(*Video) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// Gets the webDavUrl property value. WebDAV compatible URL for the item.
func (m *DriveItem) GetWebDavUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webDavUrl
    }
}
// Gets the workbook property value. For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
func (m *DriveItem) GetWorkbook()(*Workbook) {
    if m == nil {
        return nil
    } else {
        return m.workbook
    }
}
// The deserialization information for the current model
func (m *DriveItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAnalytics() })
        if err != nil {
            return err
        }
        m.SetAnalytics(val.(*ItemAnalytics))
        return nil
    }
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudio() })
        if err != nil {
            return err
        }
        m.SetAudio(val.(*Audio))
        return nil
    }
    res["bundle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBundle() })
        if err != nil {
            return err
        }
        m.SetBundle(val.(*Bundle))
        return nil
    }
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        res := make([]DriveItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*DriveItem))
        }
        m.SetChildren(res)
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["cTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCTag(val)
        return nil
    }
    res["deleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeleted() })
        if err != nil {
            return err
        }
        m.SetDeleted(val.(*Deleted))
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFile() })
        if err != nil {
            return err
        }
        m.SetFile(val.(*File))
        return nil
    }
    res["fileSystemInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSystemInfo() })
        if err != nil {
            return err
        }
        m.SetFileSystemInfo(val.(*FileSystemInfo))
        return nil
    }
    res["folder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFolder() })
        if err != nil {
            return err
        }
        m.SetFolder(val.(*Folder))
        return nil
    }
    res["image"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImage() })
        if err != nil {
            return err
        }
        m.SetImage(val.(*Image))
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        m.SetListItem(val.(*ListItem))
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeoCoordinates() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*GeoCoordinates))
        return nil
    }
    res["malware"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMalware() })
        if err != nil {
            return err
        }
        m.SetMalware(val.(*Malware))
        return nil
    }
    res["package_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPackage_escaped() })
        if err != nil {
            return err
        }
        m.SetPackage_escaped(val.(*Package_escaped))
        return nil
    }
    res["pendingOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPendingOperations() })
        if err != nil {
            return err
        }
        m.SetPendingOperations(val.(*PendingOperations))
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
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhoto() })
        if err != nil {
            return err
        }
        m.SetPhoto(val.(*Photo))
        return nil
    }
    res["publication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicationFacet() })
        if err != nil {
            return err
        }
        m.SetPublication(val.(*PublicationFacet))
        return nil
    }
    res["remoteItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteItem() })
        if err != nil {
            return err
        }
        m.SetRemoteItem(val.(*RemoteItem))
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
    res["searchResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchResult() })
        if err != nil {
            return err
        }
        m.SetSearchResult(val.(*SearchResult))
        return nil
    }
    res["shared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShared() })
        if err != nil {
            return err
        }
        m.SetShared(val.(*Shared))
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
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    res["specialFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSpecialFolder() })
        if err != nil {
            return err
        }
        m.SetSpecialFolder(val.(*SpecialFolder))
        return nil
    }
    res["subscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubscription() })
        if err != nil {
            return err
        }
        res := make([]Subscription, len(val))
        for i, v := range val {
            res[i] = *(v.(*Subscription))
        }
        m.SetSubscriptions(res)
        return nil
    }
    res["thumbnails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnailSet() })
        if err != nil {
            return err
        }
        res := make([]ThumbnailSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*ThumbnailSet))
        }
        m.SetThumbnails(res)
        return nil
    }
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItemVersion() })
        if err != nil {
            return err
        }
        res := make([]DriveItemVersion, len(val))
        for i, v := range val {
            res[i] = *(v.(*DriveItemVersion))
        }
        m.SetVersions(res)
        return nil
    }
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVideo() })
        if err != nil {
            return err
        }
        m.SetVideo(val.(*Video))
        return nil
    }
    res["webDavUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebDavUrl(val)
        return nil
    }
    res["workbook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbook() })
        if err != nil {
            return err
        }
        m.SetWorkbook(val.(*Workbook))
        return nil
    }
    return res
}
func (m *DriveItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DriveItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("audio", m.GetAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bundle", m.GetBundle())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cTag", m.GetCTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deleted", m.GetDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fileSystemInfo", m.GetFileSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("folder", m.GetFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("image", m.GetImage())
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
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("malware", m.GetMalware())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("package_escaped", m.GetPackage_escaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("pendingOperations", m.GetPendingOperations())
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
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publication", m.GetPublication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("remoteItem", m.GetRemoteItem())
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
        err = writer.WriteObjectValue("searchResult", m.GetSearchResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shared", m.GetShared())
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
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("specialFolder", m.GetSpecialFolder())
        if err != nil {
            return err
        }
    }
    {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThumbnails()))
        for i, v := range m.GetThumbnails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("thumbnails", cast)
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
    {
        err = writer.WriteObjectValue("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webDavUrl", m.GetWebDavUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workbook", m.GetWorkbook())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the analytics property value. Analytics about the view activities that took place on this item.
// Parameters:
//  - value : Value to set for the analytics property.
func (m *DriveItem) SetAnalytics(value *ItemAnalytics)() {
    m.analytics = value
}
// Sets the audio property value. Audio metadata, if the item is an audio file. Read-only.
// Parameters:
//  - value : Value to set for the audio property.
func (m *DriveItem) SetAudio(value *Audio)() {
    m.audio = value
}
// Sets the bundle property value. 
// Parameters:
//  - value : Value to set for the bundle property.
func (m *DriveItem) SetBundle(value *Bundle)() {
    m.bundle = value
}
// Sets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the children property.
func (m *DriveItem) SetChildren(value []DriveItem)() {
    m.children = value
}
// Sets the content property value. The content stream, if the item represents a file.
// Parameters:
//  - value : Value to set for the content property.
func (m *DriveItem) SetContent(value []byte)() {
    m.content = value
}
// Sets the cTag property value. An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
// Parameters:
//  - value : Value to set for the cTag property.
func (m *DriveItem) SetCTag(value *string)() {
    m.cTag = value
}
// Sets the deleted property value. Information about the deleted state of the item. Read-only.
// Parameters:
//  - value : Value to set for the deleted property.
func (m *DriveItem) SetDeleted(value *Deleted)() {
    m.deleted = value
}
// Sets the file property value. File metadata, if the item is a file. Read-only.
// Parameters:
//  - value : Value to set for the file property.
func (m *DriveItem) SetFile(value *File)() {
    m.file = value
}
// Sets the fileSystemInfo property value. File system information on client. Read-write.
// Parameters:
//  - value : Value to set for the fileSystemInfo property.
func (m *DriveItem) SetFileSystemInfo(value *FileSystemInfo)() {
    m.fileSystemInfo = value
}
// Sets the folder property value. Folder metadata, if the item is a folder. Read-only.
// Parameters:
//  - value : Value to set for the folder property.
func (m *DriveItem) SetFolder(value *Folder)() {
    m.folder = value
}
// Sets the image property value. Image metadata, if the item is an image. Read-only.
// Parameters:
//  - value : Value to set for the image property.
func (m *DriveItem) SetImage(value *Image)() {
    m.image = value
}
// Sets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the listItem property.
func (m *DriveItem) SetListItem(value *ListItem)() {
    m.listItem = value
}
// Sets the location property value. Location metadata, if the item has location data. Read-only.
// Parameters:
//  - value : Value to set for the location property.
func (m *DriveItem) SetLocation(value *GeoCoordinates)() {
    m.location = value
}
// Sets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
// Parameters:
//  - value : Value to set for the malware property.
func (m *DriveItem) SetMalware(value *Malware)() {
    m.malware = value
}
// Sets the package_escaped property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
// Parameters:
//  - value : Value to set for the package_escaped property.
func (m *DriveItem) SetPackage_escaped(value *Package_escaped)() {
    m.package_escaped = value
}
// Sets the pendingOperations property value. If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
// Parameters:
//  - value : Value to set for the pendingOperations property.
func (m *DriveItem) SetPendingOperations(value *PendingOperations)() {
    m.pendingOperations = value
}
// Sets the permissions property value. The set of permissions for the item. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the permissions property.
func (m *DriveItem) SetPermissions(value []Permission)() {
    m.permissions = value
}
// Sets the photo property value. Photo metadata, if the item is a photo. Read-only.
// Parameters:
//  - value : Value to set for the photo property.
func (m *DriveItem) SetPhoto(value *Photo)() {
    m.photo = value
}
// Sets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
// Parameters:
//  - value : Value to set for the publication property.
func (m *DriveItem) SetPublication(value *PublicationFacet)() {
    m.publication = value
}
// Sets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
// Parameters:
//  - value : Value to set for the remoteItem property.
func (m *DriveItem) SetRemoteItem(value *RemoteItem)() {
    m.remoteItem = value
}
// Sets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
// Parameters:
//  - value : Value to set for the root property.
func (m *DriveItem) SetRoot(value *Root)() {
    m.root = value
}
// Sets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
// Parameters:
//  - value : Value to set for the searchResult property.
func (m *DriveItem) SetSearchResult(value *SearchResult)() {
    m.searchResult = value
}
// Sets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
// Parameters:
//  - value : Value to set for the shared property.
func (m *DriveItem) SetShared(value *Shared)() {
    m.shared = value
}
// Sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
// Parameters:
//  - value : Value to set for the sharepointIds property.
func (m *DriveItem) SetSharepointIds(value *SharepointIds)() {
    m.sharepointIds = value
}
// Sets the size property value. Size of the item in bytes. Read-only.
// Parameters:
//  - value : Value to set for the size property.
func (m *DriveItem) SetSize(value *int64)() {
    m.size = value
}
// Sets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
// Parameters:
//  - value : Value to set for the specialFolder property.
func (m *DriveItem) SetSpecialFolder(value *SpecialFolder)() {
    m.specialFolder = value
}
// Sets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
// Parameters:
//  - value : Value to set for the subscriptions property.
func (m *DriveItem) SetSubscriptions(value []Subscription)() {
    m.subscriptions = value
}
// Sets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the thumbnails property.
func (m *DriveItem) SetThumbnails(value []ThumbnailSet)() {
    m.thumbnails = value
}
// Sets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the versions property.
func (m *DriveItem) SetVersions(value []DriveItemVersion)() {
    m.versions = value
}
// Sets the video property value. Video metadata, if the item is a video. Read-only.
// Parameters:
//  - value : Value to set for the video property.
func (m *DriveItem) SetVideo(value *Video)() {
    m.video = value
}
// Sets the webDavUrl property value. WebDAV compatible URL for the item.
// Parameters:
//  - value : Value to set for the webDavUrl property.
func (m *DriveItem) SetWebDavUrl(value *string)() {
    m.webDavUrl = value
}
// Sets the workbook property value. For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
// Parameters:
//  - value : Value to set for the workbook property.
func (m *DriveItem) SetWorkbook(value *Workbook)() {
    m.workbook = value
}
