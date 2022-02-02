package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItem 
type DriveItem struct {
    BaseItem
    // Analytics about the view activities that took place on this item.
    analytics *ItemAnalytics;
    // Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
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
    // If present, indicates that indicates that one or more operations that may affect the state of the driveItem are pending completion. Read-only.
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
// NewDriveItem instantiates a new driveItem and sets the default values.
func NewDriveItem()(*DriveItem) {
    m := &DriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// GetAnalytics gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *DriveItem) GetAnalytics()(*ItemAnalytics) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// GetAudio gets the audio property value. Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
func (m *DriveItem) GetAudio()(*Audio) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
// GetBundle gets the bundle property value. 
func (m *DriveItem) GetBundle()(*Bundle) {
    if m == nil {
        return nil
    } else {
        return m.bundle
    }
}
// GetChildren gets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) GetChildren()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
// GetContent gets the content property value. The content stream, if the item represents a file.
func (m *DriveItem) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetCTag gets the cTag property value. An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
func (m *DriveItem) GetCTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cTag
    }
}
// GetDeleted gets the deleted property value. Information about the deleted state of the item. Read-only.
func (m *DriveItem) GetDeleted()(*Deleted) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
// GetFile gets the file property value. File metadata, if the item is a file. Read-only.
func (m *DriveItem) GetFile()(*File) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFileSystemInfo gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItem) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// GetFolder gets the folder property value. Folder metadata, if the item is a folder. Read-only.
func (m *DriveItem) GetFolder()(*Folder) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
// GetImage gets the image property value. Image metadata, if the item is an image. Read-only.
func (m *DriveItem) GetImage()(*Image) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
// GetListItem gets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *DriveItem) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetLocation gets the location property value. Location metadata, if the item has location data. Read-only.
func (m *DriveItem) GetLocation()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetMalware gets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
func (m *DriveItem) GetMalware()(*Malware) {
    if m == nil {
        return nil
    } else {
        return m.malware
    }
}
// GetPackage gets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *DriveItem) GetPackage()(*Package_escaped) {
    if m == nil {
        return nil
    } else {
        return m.package_escaped
    }
}
// GetPendingOperations gets the pendingOperations property value. If present, indicates that indicates that one or more operations that may affect the state of the driveItem are pending completion. Read-only.
func (m *DriveItem) GetPendingOperations()(*PendingOperations) {
    if m == nil {
        return nil
    } else {
        return m.pendingOperations
    }
}
// GetPermissions gets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) GetPermissions()([]Permission) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// GetPhoto gets the photo property value. Photo metadata, if the item is a photo. Read-only.
func (m *DriveItem) GetPhoto()(*Photo) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetPublication gets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
func (m *DriveItem) GetPublication()(*PublicationFacet) {
    if m == nil {
        return nil
    } else {
        return m.publication
    }
}
// GetRemoteItem gets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
func (m *DriveItem) GetRemoteItem()(*RemoteItem) {
    if m == nil {
        return nil
    } else {
        return m.remoteItem
    }
}
// GetRoot gets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
func (m *DriveItem) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSearchResult gets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
func (m *DriveItem) GetSearchResult()(*SearchResult) {
    if m == nil {
        return nil
    } else {
        return m.searchResult
    }
}
// GetShared gets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *DriveItem) GetShared()(*Shared) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *DriveItem) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// GetSize gets the size property value. Size of the item in bytes. Read-only.
func (m *DriveItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetSpecialFolder gets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *DriveItem) GetSpecialFolder()(*SpecialFolder) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) GetSubscriptions()([]Subscription) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// GetThumbnails gets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) GetThumbnails()([]ThumbnailSet) {
    if m == nil {
        return nil
    } else {
        return m.thumbnails
    }
}
// GetVersions gets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) GetVersions()([]DriveItemVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// GetVideo gets the video property value. Video metadata, if the item is a video. Read-only.
func (m *DriveItem) GetVideo()(*Video) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// GetWebDavUrl gets the webDavUrl property value. WebDAV compatible URL for the item.
func (m *DriveItem) GetWebDavUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webDavUrl
    }
}
// GetWorkbook gets the workbook property value. For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
func (m *DriveItem) GetWorkbook()(*Workbook) {
    if m == nil {
        return nil
    } else {
        return m.workbook
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAnalytics() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(*ItemAnalytics))
        }
        return nil
    }
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudio() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudio(val.(*Audio))
        }
        return nil
    }
    res["bundle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBundle() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundle(val.(*Bundle))
        }
        return nil
    }
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItem))
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["cTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCTag(val)
        }
        return nil
    }
    res["deleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeleted() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val.(*Deleted))
        }
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFile() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(*File))
        }
        return nil
    }
    res["fileSystemInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSystemInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(*FileSystemInfo))
        }
        return nil
    }
    res["folder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(*Folder))
        }
        return nil
    }
    res["image"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImage() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(*Image))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewListItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(*ListItem))
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeoCoordinates() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*GeoCoordinates))
        }
        return nil
    }
    res["malware"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMalware() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalware(val.(*Malware))
        }
        return nil
    }
    res["package"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPackage_escaped() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackage(val.(*Package_escaped))
        }
        return nil
    }
    res["pendingOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPendingOperations() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOperations(val.(*PendingOperations))
        }
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermission() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Permission, len(val))
            for i, v := range val {
                res[i] = *(v.(*Permission))
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*Photo))
        }
        return nil
    }
    res["publication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicationFacet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublication(val.(*PublicationFacet))
        }
        return nil
    }
    res["remoteItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteItem(val.(*RemoteItem))
        }
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(*Root))
        }
        return nil
    }
    res["searchResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchResult(val.(*SearchResult))
        }
        return nil
    }
    res["shared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShared() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShared(val.(*Shared))
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
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["specialFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSpecialFolder() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialFolder(val.(*SpecialFolder))
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
    res["thumbnails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnailSet() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThumbnailSet, len(val))
            for i, v := range val {
                res[i] = *(v.(*ThumbnailSet))
            }
            m.SetThumbnails(res)
        }
        return nil
    }
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItemVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemVersion, len(val))
            for i, v := range val {
                res[i] = *(v.(*DriveItemVersion))
            }
            m.SetVersions(res)
        }
        return nil
    }
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVideo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val.(*Video))
        }
        return nil
    }
    res["webDavUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebDavUrl(val)
        }
        return nil
    }
    res["workbook"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbook() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbook(val.(*Workbook))
        }
        return nil
    }
    return res
}
func (m *DriveItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetChildren() != nil {
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
        err = writer.WriteObjectValue("package", m.GetPackage())
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
    if m.GetPermissions() != nil {
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
    if m.GetThumbnails() != nil {
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
    if m.GetVersions() != nil {
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
// SetAnalytics sets the analytics property value. Analytics about the view activities that took place on this item.
func (m *DriveItem) SetAnalytics(value *ItemAnalytics)() {
    if m != nil {
        m.analytics = value
    }
}
// SetAudio sets the audio property value. Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
func (m *DriveItem) SetAudio(value *Audio)() {
    if m != nil {
        m.audio = value
    }
}
// SetBundle sets the bundle property value. 
func (m *DriveItem) SetBundle(value *Bundle)() {
    if m != nil {
        m.bundle = value
    }
}
// SetChildren sets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) SetChildren(value []DriveItem)() {
    if m != nil {
        m.children = value
    }
}
// SetContent sets the content property value. The content stream, if the item represents a file.
func (m *DriveItem) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetCTag sets the cTag property value. An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
func (m *DriveItem) SetCTag(value *string)() {
    if m != nil {
        m.cTag = value
    }
}
// SetDeleted sets the deleted property value. Information about the deleted state of the item. Read-only.
func (m *DriveItem) SetDeleted(value *Deleted)() {
    if m != nil {
        m.deleted = value
    }
}
// SetFile sets the file property value. File metadata, if the item is a file. Read-only.
func (m *DriveItem) SetFile(value *File)() {
    if m != nil {
        m.file = value
    }
}
// SetFileSystemInfo sets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItem) SetFileSystemInfo(value *FileSystemInfo)() {
    if m != nil {
        m.fileSystemInfo = value
    }
}
// SetFolder sets the folder property value. Folder metadata, if the item is a folder. Read-only.
func (m *DriveItem) SetFolder(value *Folder)() {
    if m != nil {
        m.folder = value
    }
}
// SetImage sets the image property value. Image metadata, if the item is an image. Read-only.
func (m *DriveItem) SetImage(value *Image)() {
    if m != nil {
        m.image = value
    }
}
// SetListItem sets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *DriveItem) SetListItem(value *ListItem)() {
    if m != nil {
        m.listItem = value
    }
}
// SetLocation sets the location property value. Location metadata, if the item has location data. Read-only.
func (m *DriveItem) SetLocation(value *GeoCoordinates)() {
    if m != nil {
        m.location = value
    }
}
// SetMalware sets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
func (m *DriveItem) SetMalware(value *Malware)() {
    if m != nil {
        m.malware = value
    }
}
// SetPackage sets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *DriveItem) SetPackage(value *Package_escaped)() {
    if m != nil {
        m.package_escaped = value
    }
}
// SetPendingOperations sets the pendingOperations property value. If present, indicates that indicates that one or more operations that may affect the state of the driveItem are pending completion. Read-only.
func (m *DriveItem) SetPendingOperations(value *PendingOperations)() {
    if m != nil {
        m.pendingOperations = value
    }
}
// SetPermissions sets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) SetPermissions(value []Permission)() {
    if m != nil {
        m.permissions = value
    }
}
// SetPhoto sets the photo property value. Photo metadata, if the item is a photo. Read-only.
func (m *DriveItem) SetPhoto(value *Photo)() {
    if m != nil {
        m.photo = value
    }
}
// SetPublication sets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
func (m *DriveItem) SetPublication(value *PublicationFacet)() {
    if m != nil {
        m.publication = value
    }
}
// SetRemoteItem sets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
func (m *DriveItem) SetRemoteItem(value *RemoteItem)() {
    if m != nil {
        m.remoteItem = value
    }
}
// SetRoot sets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
func (m *DriveItem) SetRoot(value *Root)() {
    if m != nil {
        m.root = value
    }
}
// SetSearchResult sets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
func (m *DriveItem) SetSearchResult(value *SearchResult)() {
    if m != nil {
        m.searchResult = value
    }
}
// SetShared sets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *DriveItem) SetShared(value *Shared)() {
    if m != nil {
        m.shared = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *DriveItem) SetSharepointIds(value *SharepointIds)() {
    if m != nil {
        m.sharepointIds = value
    }
}
// SetSize sets the size property value. Size of the item in bytes. Read-only.
func (m *DriveItem) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
// SetSpecialFolder sets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *DriveItem) SetSpecialFolder(value *SpecialFolder)() {
    if m != nil {
        m.specialFolder = value
    }
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) SetSubscriptions(value []Subscription)() {
    if m != nil {
        m.subscriptions = value
    }
}
// SetThumbnails sets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) SetThumbnails(value []ThumbnailSet)() {
    if m != nil {
        m.thumbnails = value
    }
}
// SetVersions sets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) SetVersions(value []DriveItemVersion)() {
    if m != nil {
        m.versions = value
    }
}
// SetVideo sets the video property value. Video metadata, if the item is a video. Read-only.
func (m *DriveItem) SetVideo(value *Video)() {
    if m != nil {
        m.video = value
    }
}
// SetWebDavUrl sets the webDavUrl property value. WebDAV compatible URL for the item.
func (m *DriveItem) SetWebDavUrl(value *string)() {
    if m != nil {
        m.webDavUrl = value
    }
}
// SetWorkbook sets the workbook property value. For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
func (m *DriveItem) SetWorkbook(value *Workbook)() {
    if m != nil {
        m.workbook = value
    }
}
