package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteItem 
type RemoteItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identity of the user, device, and application which created the item. Read-only.
    createdBy IdentitySetable;
    // Date and time of item creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates that the remote item is a file. Read-only.
    file Fileable;
    // Information about the remote item from the local file system. Read-only.
    fileSystemInfo FileSystemInfoable;
    // Indicates that the remote item is a folder. Read-only.
    folder Folderable;
    // Unique identifier for the remote item in its drive. Read-only.
    id *string;
    // Image metadata, if the item is an image. Read-only.
    image Imageable;
    // Identity of the user, device, and application which last modified the item. Read-only.
    lastModifiedBy IdentitySetable;
    // Date and time the item was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional. Filename of the remote item. Read-only.
    name *string;
    // If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
    package_escaped Package_escapedable;
    // Properties of the parent of the remote item. Read-only.
    parentReference ItemReferenceable;
    // Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
    shared Sharedable;
    // Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
    sharepointIds SharepointIdsable;
    // Size of the remote item. Read-only.
    size *int64;
    // If the current item is also available as a special folder, this facet is returned. Read-only.
    specialFolder SpecialFolderable;
    // Video metadata, if the item is a video. Read-only.
    video Videoable;
    // DAV compatible URL for the item.
    webDavUrl *string;
    // URL that displays the resource in the browser. Read-only.
    webUrl *string;
}
// NewRemoteItem instantiates a new remoteItem and sets the default values.
func NewRemoteItem()(*RemoteItem) {
    m := &RemoteItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRemoteItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoteItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedBy gets the createdBy property value. Identity of the user, device, and application which created the item. Read-only.
func (m *RemoteItem) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *RemoteItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteItem) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdBy"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["file"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val.(Fileable))
        }
        return nil
    }
    res["fileSystemInfo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileSystemInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(FileSystemInfoable))
        }
        return nil
    }
    res["folder"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(Folderable))
        }
        return nil
    }
    res["id"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["image"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImage(val.(Imageable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["package"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackage_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackage(val.(Package_escapedable))
        }
        return nil
    }
    res["parentReference"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(ItemReferenceable))
        }
        return nil
    }
    res["shared"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShared(val.(Sharedable))
        }
        return nil
    }
    res["sharepointIds"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["size"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["specialFolder"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpecialFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialFolder(val.(SpecialFolderable))
        }
        return nil
    }
    res["video"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVideoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideo(val.(Videoable))
        }
        return nil
    }
    res["webDavUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebDavUrl(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetFile gets the file property value. Indicates that the remote item is a file. Read-only.
func (m *RemoteItem) GetFile()(Fileable) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFileSystemInfo gets the fileSystemInfo property value. Information about the remote item from the local file system. Read-only.
func (m *RemoteItem) GetFileSystemInfo()(FileSystemInfoable) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// GetFolder gets the folder property value. Indicates that the remote item is a folder. Read-only.
func (m *RemoteItem) GetFolder()(Folderable) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
// GetId gets the id property value. Unique identifier for the remote item in its drive. Read-only.
func (m *RemoteItem) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetImage gets the image property value. Image metadata, if the item is an image. Read-only.
func (m *RemoteItem) GetImage()(Imageable) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *RemoteItem) GetLastModifiedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *RemoteItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetName gets the name property value. Optional. Filename of the remote item. Read-only.
func (m *RemoteItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPackage gets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *RemoteItem) GetPackage()(Package_escapedable) {
    if m == nil {
        return nil
    } else {
        return m.package_escaped
    }
}
// GetParentReference gets the parentReference property value. Properties of the parent of the remote item. Read-only.
func (m *RemoteItem) GetParentReference()(ItemReferenceable) {
    if m == nil {
        return nil
    } else {
        return m.parentReference
    }
}
// GetShared gets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *RemoteItem) GetShared()(Sharedable) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// GetSharepointIds gets the sharepointIds property value. Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
func (m *RemoteItem) GetSharepointIds()(SharepointIdsable) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// GetSize gets the size property value. Size of the remote item. Read-only.
func (m *RemoteItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetSpecialFolder gets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *RemoteItem) GetSpecialFolder()(SpecialFolderable) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
// GetVideo gets the video property value. Video metadata, if the item is a video. Read-only.
func (m *RemoteItem) GetVideo()(Videoable) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// GetWebDavUrl gets the webDavUrl property value. DAV compatible URL for the item.
func (m *RemoteItem) GetWebDavUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webDavUrl
    }
}
// GetWebUrl gets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *RemoteItem) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *RemoteItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fileSystemInfo", m.GetFileSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("folder", m.GetFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("image", m.GetImage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("package", m.GetPackage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentReference", m.GetParentReference())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("shared", m.GetShared())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("specialFolder", m.GetSpecialFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("video", m.GetVideo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webDavUrl", m.GetWebDavUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoteItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user, device, and application which created the item. Read-only.
func (m *RemoteItem) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *RemoteItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetFile sets the file property value. Indicates that the remote item is a file. Read-only.
func (m *RemoteItem) SetFile(value Fileable)() {
    if m != nil {
        m.file = value
    }
}
// SetFileSystemInfo sets the fileSystemInfo property value. Information about the remote item from the local file system. Read-only.
func (m *RemoteItem) SetFileSystemInfo(value FileSystemInfoable)() {
    if m != nil {
        m.fileSystemInfo = value
    }
}
// SetFolder sets the folder property value. Indicates that the remote item is a folder. Read-only.
func (m *RemoteItem) SetFolder(value Folderable)() {
    if m != nil {
        m.folder = value
    }
}
// SetId sets the id property value. Unique identifier for the remote item in its drive. Read-only.
func (m *RemoteItem) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetImage sets the image property value. Image metadata, if the item is an image. Read-only.
func (m *RemoteItem) SetImage(value Imageable)() {
    if m != nil {
        m.image = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *RemoteItem) SetLastModifiedBy(value IdentitySetable)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *RemoteItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetName sets the name property value. Optional. Filename of the remote item. Read-only.
func (m *RemoteItem) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPackage sets the package property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *RemoteItem) SetPackage(value Package_escapedable)() {
    if m != nil {
        m.package_escaped = value
    }
}
// SetParentReference sets the parentReference property value. Properties of the parent of the remote item. Read-only.
func (m *RemoteItem) SetParentReference(value ItemReferenceable)() {
    if m != nil {
        m.parentReference = value
    }
}
// SetShared sets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *RemoteItem) SetShared(value Sharedable)() {
    if m != nil {
        m.shared = value
    }
}
// SetSharepointIds sets the sharepointIds property value. Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
func (m *RemoteItem) SetSharepointIds(value SharepointIdsable)() {
    if m != nil {
        m.sharepointIds = value
    }
}
// SetSize sets the size property value. Size of the remote item. Read-only.
func (m *RemoteItem) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
// SetSpecialFolder sets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *RemoteItem) SetSpecialFolder(value SpecialFolderable)() {
    if m != nil {
        m.specialFolder = value
    }
}
// SetVideo sets the video property value. Video metadata, if the item is a video. Read-only.
func (m *RemoteItem) SetVideo(value Videoable)() {
    if m != nil {
        m.video = value
    }
}
// SetWebDavUrl sets the webDavUrl property value. DAV compatible URL for the item.
func (m *RemoteItem) SetWebDavUrl(value *string)() {
    if m != nil {
        m.webDavUrl = value
    }
}
// SetWebUrl sets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *RemoteItem) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
