package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItem provides operations to manage the drive singleton.
type DriveItem struct {
    BaseItem
    // Analytics about the view activities that took place on this item.
    analytics ItemAnalyticsable;
    // Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
    audio Audioable;
    // 
    bundle Bundleable;
    // Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
    children []DriveItemable;
    // The content stream, if the item represents a file.
    content []byte;
    // An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
    cTag *string;
    // Information about the deleted state of the item. Read-only.
    deleted Deletedable;
    // File metadata, if the item is a file. Read-only.
    file Fileable;
    // File system information on client. Read-write.
    fileSystemInfo FileSystemInfoable;
    // Folder metadata, if the item is a folder. Read-only.
    folder Folderable;
    // Image metadata, if the item is an image. Read-only.
    image Imageable;
    // For drives in SharePoint, the associated document library list item. Read-only. Nullable.
    listItem ListItemable;
    // Location metadata, if the item has location data. Read-only.
    location GeoCoordinatesable;
    // Malware metadata, if the item was detected to contain malware. Read-only.
    malware Malwareable;
    // If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
    package_escaped Package_escapedable;
    // If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
    pendingOperations PendingOperationsable;
    // The set of permissions for the item. Read-only. Nullable.
    permissions []Permissionable;
    // Photo metadata, if the item is a photo. Read-only.
    photo Photoable;
    // Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
    publication PublicationFacetable;
    // Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
    remoteItem RemoteItemable;
    // If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
    root Rootable;
    // Search metadata, if the item is from a search result. Read-only.
    searchResult SearchResultable;
    // Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
    shared Sharedable;
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds SharepointIdsable;
    // Size of the item in bytes. Read-only.
    size *int64;
    // If the current item is also available as a special folder, this facet is returned. Read-only.
    specialFolder SpecialFolderable;
    // The set of subscriptions on the item. Only supported on the root of a drive.
    subscriptions []Subscriptionable;
    // Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
    thumbnails []ThumbnailSetable;
    // The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
    versions []DriveItemVersionable;
    // Video metadata, if the item is a video. Read-only.
    video Videoable;
    // WebDAV compatible URL for the item.
    webDavUrl *string;
    // For files that are Excel spreadsheets, accesses the workbook API to work with the spreadsheet's contents. Nullable.
    workbook Workbookable;
}
// NewDriveItem instantiates a new driveItem and sets the default values.
func NewDriveItem()(*DriveItem) {
    m := &DriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
// CreateDriveItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDriveItem(), nil
}
// GetAnalytics gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *DriveItem) GetAnalytics()(ItemAnalyticsable) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
// GetAudio gets the audio property value. Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
func (m *DriveItem) GetAudio()(Audioable) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
// GetBundle gets the bundle property value. 
func (m *DriveItem) GetBundle()(Bundleable) {
    if m == nil {
        return nil
    } else {
        return m.bundle
    }
}
// GetChildren gets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) GetChildren()([]DriveItemable) {
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
func (m *DriveItem) GetDeleted()(Deletedable) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(ItemAnalyticsable))
        }
        return nil
    }
    res["audio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudioFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudio(val.(Audioable))
        }
        return nil
    }
    res["bundle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateBundleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundle(val.(Bundleable))
        }
        return nil
    }
    res["children"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
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
        val, err := n.GetObjectValue(CreateDeletedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val.(Deletedable))
        }
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(Fileable))
        }
        return nil
    }
    res["fileSystemInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(FileSystemInfoable))
        }
        return nil
    }
    res["folder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(Folderable))
        }
        return nil
    }
    res["image"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(Imageable))
        }
        return nil
    }
    res["listItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(GeoCoordinatesable))
        }
        return nil
    }
    res["malware"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMalwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalware(val.(Malwareable))
        }
        return nil
    }
    res["package"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackage_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackage(val.(Package_escapedable))
        }
        return nil
    }
    res["pendingOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePendingOperationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOperations(val.(PendingOperationsable))
        }
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Permissionable, len(val))
            for i, v := range val {
                res[i] = v.(Permissionable)
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(Photoable))
        }
        return nil
    }
    res["publication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublication(val.(PublicationFacetable))
        }
        return nil
    }
    res["remoteItem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteItem(val.(RemoteItemable))
        }
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(Rootable))
        }
        return nil
    }
    res["searchResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchResult(val.(SearchResultable))
        }
        return nil
    }
    res["shared"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShared(val.(Sharedable))
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
        val, err := n.GetObjectValue(CreateSpecialFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialFolder(val.(SpecialFolderable))
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
    res["thumbnails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateThumbnailSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ThumbnailSetable, len(val))
            for i, v := range val {
                res[i] = v.(ThumbnailSetable)
            }
            m.SetThumbnails(res)
        }
        return nil
    }
    res["versions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemVersionable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    res["video"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateVideoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val.(Videoable))
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
        val, err := n.GetObjectValue(CreateWorkbookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkbook(val.(Workbookable))
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. File metadata, if the item is a file. Read-only.
func (m *DriveItem) GetFile()(Fileable) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFileSystemInfo gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItem) GetFileSystemInfo()(FileSystemInfoable) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// GetFolder gets the folder property value. Folder metadata, if the item is a folder. Read-only.
func (m *DriveItem) GetFolder()(Folderable) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
// GetImage gets the image property value. Image metadata, if the item is an image. Read-only.
func (m *DriveItem) GetImage()(Imageable) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
// GetListItem gets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *DriveItem) GetListItem()(ListItemable) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
// GetLocation gets the location property value. Location metadata, if the item has location data. Read-only.
func (m *DriveItem) GetLocation()(GeoCoordinatesable) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetMalware gets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
func (m *DriveItem) GetMalware()(Malwareable) {
    if m == nil {
        return nil
    } else {
        return m.malware
    }
}
// GetPackage gets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *DriveItem) GetPackage()(Package_escapedable) {
    if m == nil {
        return nil
    } else {
        return m.package_escaped
    }
}
// GetPendingOperations gets the pendingOperations property value. If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
func (m *DriveItem) GetPendingOperations()(PendingOperationsable) {
    if m == nil {
        return nil
    } else {
        return m.pendingOperations
    }
}
// GetPermissions gets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) GetPermissions()([]Permissionable) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// GetPhoto gets the photo property value. Photo metadata, if the item is a photo. Read-only.
func (m *DriveItem) GetPhoto()(Photoable) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetPublication gets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
func (m *DriveItem) GetPublication()(PublicationFacetable) {
    if m == nil {
        return nil
    } else {
        return m.publication
    }
}
// GetRemoteItem gets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
func (m *DriveItem) GetRemoteItem()(RemoteItemable) {
    if m == nil {
        return nil
    } else {
        return m.remoteItem
    }
}
// GetRoot gets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
func (m *DriveItem) GetRoot()(Rootable) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// GetSearchResult gets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
func (m *DriveItem) GetSearchResult()(SearchResultable) {
    if m == nil {
        return nil
    } else {
        return m.searchResult
    }
}
// GetShared gets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *DriveItem) GetShared()(Sharedable) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *DriveItem) GetSharepointIds()(SharepointIdsable) {
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
func (m *DriveItem) GetSpecialFolder()(SpecialFolderable) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) GetSubscriptions()([]Subscriptionable) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// GetThumbnails gets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) GetThumbnails()([]ThumbnailSetable) {
    if m == nil {
        return nil
    } else {
        return m.thumbnails
    }
}
// GetVersions gets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) GetVersions()([]DriveItemVersionable) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
// GetVideo gets the video property value. Video metadata, if the item is a video. Read-only.
func (m *DriveItem) GetVideo()(Videoable) {
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
func (m *DriveItem) GetWorkbook()(Workbookable) {
    if m == nil {
        return nil
    } else {
        return m.workbook
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetThumbnails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThumbnails()))
        for i, v := range m.GetThumbnails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("thumbnails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *DriveItem) SetAnalytics(value ItemAnalyticsable)() {
    if m != nil {
        m.analytics = value
    }
}
// SetAudio sets the audio property value. Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
func (m *DriveItem) SetAudio(value Audioable)() {
    if m != nil {
        m.audio = value
    }
}
// SetBundle sets the bundle property value. 
func (m *DriveItem) SetBundle(value Bundleable)() {
    if m != nil {
        m.bundle = value
    }
}
// SetChildren sets the children property value. Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable.
func (m *DriveItem) SetChildren(value []DriveItemable)() {
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
func (m *DriveItem) SetDeleted(value Deletedable)() {
    if m != nil {
        m.deleted = value
    }
}
// SetFile sets the file property value. File metadata, if the item is a file. Read-only.
func (m *DriveItem) SetFile(value Fileable)() {
    if m != nil {
        m.file = value
    }
}
// SetFileSystemInfo sets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItem) SetFileSystemInfo(value FileSystemInfoable)() {
    if m != nil {
        m.fileSystemInfo = value
    }
}
// SetFolder sets the folder property value. Folder metadata, if the item is a folder. Read-only.
func (m *DriveItem) SetFolder(value Folderable)() {
    if m != nil {
        m.folder = value
    }
}
// SetImage sets the image property value. Image metadata, if the item is an image. Read-only.
func (m *DriveItem) SetImage(value Imageable)() {
    if m != nil {
        m.image = value
    }
}
// SetListItem sets the listItem property value. For drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *DriveItem) SetListItem(value ListItemable)() {
    if m != nil {
        m.listItem = value
    }
}
// SetLocation sets the location property value. Location metadata, if the item has location data. Read-only.
func (m *DriveItem) SetLocation(value GeoCoordinatesable)() {
    if m != nil {
        m.location = value
    }
}
// SetMalware sets the malware property value. Malware metadata, if the item was detected to contain malware. Read-only.
func (m *DriveItem) SetMalware(value Malwareable)() {
    if m != nil {
        m.malware = value
    }
}
// SetPackage sets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *DriveItem) SetPackage(value Package_escapedable)() {
    if m != nil {
        m.package_escaped = value
    }
}
// SetPendingOperations sets the pendingOperations property value. If present, indicates that one or more operations that might affect the state of the driveItem are pending completion. Read-only.
func (m *DriveItem) SetPendingOperations(value PendingOperationsable)() {
    if m != nil {
        m.pendingOperations = value
    }
}
// SetPermissions sets the permissions property value. The set of permissions for the item. Read-only. Nullable.
func (m *DriveItem) SetPermissions(value []Permissionable)() {
    if m != nil {
        m.permissions = value
    }
}
// SetPhoto sets the photo property value. Photo metadata, if the item is a photo. Read-only.
func (m *DriveItem) SetPhoto(value Photoable)() {
    if m != nil {
        m.photo = value
    }
}
// SetPublication sets the publication property value. Provides information about the published or checked-out state of an item, in locations that support such actions. This property is not returned by default. Read-only.
func (m *DriveItem) SetPublication(value PublicationFacetable)() {
    if m != nil {
        m.publication = value
    }
}
// SetRemoteItem sets the remoteItem property value. Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
func (m *DriveItem) SetRemoteItem(value RemoteItemable)() {
    if m != nil {
        m.remoteItem = value
    }
}
// SetRoot sets the root property value. If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
func (m *DriveItem) SetRoot(value Rootable)() {
    if m != nil {
        m.root = value
    }
}
// SetSearchResult sets the searchResult property value. Search metadata, if the item is from a search result. Read-only.
func (m *DriveItem) SetSearchResult(value SearchResultable)() {
    if m != nil {
        m.searchResult = value
    }
}
// SetShared sets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *DriveItem) SetShared(value Sharedable)() {
    if m != nil {
        m.shared = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *DriveItem) SetSharepointIds(value SharepointIdsable)() {
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
func (m *DriveItem) SetSpecialFolder(value SpecialFolderable)() {
    if m != nil {
        m.specialFolder = value
    }
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the item. Only supported on the root of a drive.
func (m *DriveItem) SetSubscriptions(value []Subscriptionable)() {
    if m != nil {
        m.subscriptions = value
    }
}
// SetThumbnails sets the thumbnails property value. Collection containing [ThumbnailSet][] objects associated with the item. For more info, see [getting thumbnails][]. Read-only. Nullable.
func (m *DriveItem) SetThumbnails(value []ThumbnailSetable)() {
    if m != nil {
        m.thumbnails = value
    }
}
// SetVersions sets the versions property value. The list of previous versions of the item. For more info, see [getting previous versions][]. Read-only. Nullable.
func (m *DriveItem) SetVersions(value []DriveItemVersionable)() {
    if m != nil {
        m.versions = value
    }
}
// SetVideo sets the video property value. Video metadata, if the item is a video. Read-only.
func (m *DriveItem) SetVideo(value Videoable)() {
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
func (m *DriveItem) SetWorkbook(value Workbookable)() {
    if m != nil {
        m.workbook = value
    }
}
