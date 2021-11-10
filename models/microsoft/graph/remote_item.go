package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RemoteItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identity of the user, device, and application which created the item. Read-only.
    createdBy *IdentitySet;
    // Date and time of item creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates that the remote item is a file. Read-only.
    file *File;
    // Information about the remote item from the local file system. Read-only.
    fileSystemInfo *FileSystemInfo;
    // Indicates that the remote item is a folder. Read-only.
    folder *Folder;
    // Unique identifier for the remote item in its drive. Read-only.
    id *string;
    // Image metadata, if the item is an image. Read-only.
    image *Image;
    // Identity of the user, device, and application which last modified the item. Read-only.
    lastModifiedBy *IdentitySet;
    // Date and time the item was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional. Filename of the remote item. Read-only.
    name *string;
    // If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
    package_escaped *Package_escaped;
    // Properties of the parent of the remote item. Read-only.
    parentReference *ItemReference;
    // Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
    shared *Shared;
    // Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
    sharepointIds *SharepointIds;
    // Size of the remote item. Read-only.
    size *int64;
    // If the current item is also available as a special folder, this facet is returned. Read-only.
    specialFolder *SpecialFolder;
    // Video metadata, if the item is a video. Read-only.
    video *Video;
    // DAV compatible URL for the item.
    webDavUrl *string;
    // URL that displays the resource in the browser. Read-only.
    webUrl *string;
}
// Instantiates a new remoteItem and sets the default values.
func NewRemoteItem()(*RemoteItem) {
    m := &RemoteItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RemoteItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the createdBy property value. Identity of the user, device, and application which created the item. Read-only.
func (m *RemoteItem) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *RemoteItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the file property value. Indicates that the remote item is a file. Read-only.
func (m *RemoteItem) GetFile()(*File) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// Gets the fileSystemInfo property value. Information about the remote item from the local file system. Read-only.
func (m *RemoteItem) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// Gets the folder property value. Indicates that the remote item is a folder. Read-only.
func (m *RemoteItem) GetFolder()(*Folder) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
// Gets the id property value. Unique identifier for the remote item in its drive. Read-only.
func (m *RemoteItem) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the image property value. Image metadata, if the item is an image. Read-only.
func (m *RemoteItem) GetImage()(*Image) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
// Gets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *RemoteItem) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *RemoteItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the name property value. Optional. Filename of the remote item. Read-only.
func (m *RemoteItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the package_escaped property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
func (m *RemoteItem) GetPackage_escaped()(*Package_escaped) {
    if m == nil {
        return nil
    } else {
        return m.package_escaped
    }
}
// Gets the parentReference property value. Properties of the parent of the remote item. Read-only.
func (m *RemoteItem) GetParentReference()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.parentReference
    }
}
// Gets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
func (m *RemoteItem) GetShared()(*Shared) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
// Gets the sharepointIds property value. Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
func (m *RemoteItem) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
// Gets the size property value. Size of the remote item. Read-only.
func (m *RemoteItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Gets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
func (m *RemoteItem) GetSpecialFolder()(*SpecialFolder) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
// Gets the video property value. Video metadata, if the item is a video. Read-only.
func (m *RemoteItem) GetVideo()(*Video) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
// Gets the webDavUrl property value. DAV compatible URL for the item.
func (m *RemoteItem) GetWebDavUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webDavUrl
    }
}
// Gets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *RemoteItem) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *RemoteItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["package_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPackage_escaped() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackage_escaped(val.(*Package_escaped))
        }
        return nil
    }
    res["parentReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(*ItemReference))
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
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *RemoteItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RemoteItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("package_escaped", m.GetPackage_escaped())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RemoteItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the createdBy property value. Identity of the user, device, and application which created the item. Read-only.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *RemoteItem) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. Date and time of item creation. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *RemoteItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the file property value. Indicates that the remote item is a file. Read-only.
// Parameters:
//  - value : Value to set for the file property.
func (m *RemoteItem) SetFile(value *File)() {
    m.file = value
}
// Sets the fileSystemInfo property value. Information about the remote item from the local file system. Read-only.
// Parameters:
//  - value : Value to set for the fileSystemInfo property.
func (m *RemoteItem) SetFileSystemInfo(value *FileSystemInfo)() {
    m.fileSystemInfo = value
}
// Sets the folder property value. Indicates that the remote item is a folder. Read-only.
// Parameters:
//  - value : Value to set for the folder property.
func (m *RemoteItem) SetFolder(value *Folder)() {
    m.folder = value
}
// Sets the id property value. Unique identifier for the remote item in its drive. Read-only.
// Parameters:
//  - value : Value to set for the id property.
func (m *RemoteItem) SetId(value *string)() {
    m.id = value
}
// Sets the image property value. Image metadata, if the item is an image. Read-only.
// Parameters:
//  - value : Value to set for the image property.
func (m *RemoteItem) SetImage(value *Image)() {
    m.image = value
}
// Sets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *RemoteItem) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *RemoteItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the name property value. Optional. Filename of the remote item. Read-only.
// Parameters:
//  - value : Value to set for the name property.
func (m *RemoteItem) SetName(value *string)() {
    m.name = value
}
// Sets the package_escaped property value. If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in some contexts and folders in others. Read-only.
// Parameters:
//  - value : Value to set for the package_escaped property.
func (m *RemoteItem) SetPackage_escaped(value *Package_escaped)() {
    m.package_escaped = value
}
// Sets the parentReference property value. Properties of the parent of the remote item. Read-only.
// Parameters:
//  - value : Value to set for the parentReference property.
func (m *RemoteItem) SetParentReference(value *ItemReference)() {
    m.parentReference = value
}
// Sets the shared property value. Indicates that the item has been shared with others and provides information about the shared state of the item. Read-only.
// Parameters:
//  - value : Value to set for the shared property.
func (m *RemoteItem) SetShared(value *Shared)() {
    m.shared = value
}
// Sets the sharepointIds property value. Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers. Read-only.
// Parameters:
//  - value : Value to set for the sharepointIds property.
func (m *RemoteItem) SetSharepointIds(value *SharepointIds)() {
    m.sharepointIds = value
}
// Sets the size property value. Size of the remote item. Read-only.
// Parameters:
//  - value : Value to set for the size property.
func (m *RemoteItem) SetSize(value *int64)() {
    m.size = value
}
// Sets the specialFolder property value. If the current item is also available as a special folder, this facet is returned. Read-only.
// Parameters:
//  - value : Value to set for the specialFolder property.
func (m *RemoteItem) SetSpecialFolder(value *SpecialFolder)() {
    m.specialFolder = value
}
// Sets the video property value. Video metadata, if the item is a video. Read-only.
// Parameters:
//  - value : Value to set for the video property.
func (m *RemoteItem) SetVideo(value *Video)() {
    m.video = value
}
// Sets the webDavUrl property value. DAV compatible URL for the item.
// Parameters:
//  - value : Value to set for the webDavUrl property.
func (m *RemoteItem) SetWebDavUrl(value *string)() {
    m.webDavUrl = value
}
// Sets the webUrl property value. URL that displays the resource in the browser. Read-only.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *RemoteItem) SetWebUrl(value *string)() {
    m.webUrl = value
}
