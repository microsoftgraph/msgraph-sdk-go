package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DriveItem struct {
    BaseItem
    analytics *ItemAnalytics;
    audio *Audio;
    children []DriveItem;
    content []byte;
    cTag *string;
    deleted *Deleted;
    file *File;
    fileSystemInfo *FileSystemInfo;
    folder *Folder;
    image *Image;
    listItem *ListItem;
    location *GeoCoordinates;
    package_escpaped *Package_escpaped;
    pendingOperations *PendingOperations;
    permissions []Permission;
    photo *Photo;
    publication *PublicationFacet;
    remoteItem *RemoteItem;
    root *Root;
    searchResult *SearchResult;
    shared *Shared;
    sharepointIds *SharepointIds;
    size *int64;
    specialFolder *SpecialFolder;
    subscriptions []Subscription;
    thumbnails []ThumbnailSet;
    versions []DriveItemVersion;
    video *Video;
    webDavUrl *string;
    workbook *Workbook;
}
func NewDriveItem()(*DriveItem) {
    m := &DriveItem{
        BaseItem: *NewBaseItem(),
    }
    return m
}
func (m *DriveItem) GetAnalytics()(*ItemAnalytics) {
    if m == nil {
        return nil
    } else {
        return m.analytics
    }
}
func (m *DriveItem) GetAudio()(*Audio) {
    if m == nil {
        return nil
    } else {
        return m.audio
    }
}
func (m *DriveItem) GetChildren()([]DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.children
    }
}
func (m *DriveItem) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *DriveItem) GetCTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cTag
    }
}
func (m *DriveItem) GetDeleted()(*Deleted) {
    if m == nil {
        return nil
    } else {
        return m.deleted
    }
}
func (m *DriveItem) GetFile()(*File) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
func (m *DriveItem) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
func (m *DriveItem) GetFolder()(*Folder) {
    if m == nil {
        return nil
    } else {
        return m.folder
    }
}
func (m *DriveItem) GetImage()(*Image) {
    if m == nil {
        return nil
    } else {
        return m.image
    }
}
func (m *DriveItem) GetListItem()(*ListItem) {
    if m == nil {
        return nil
    } else {
        return m.listItem
    }
}
func (m *DriveItem) GetLocation()(*GeoCoordinates) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
func (m *DriveItem) GetPackage_escpaped()(*Package_escpaped) {
    if m == nil {
        return nil
    } else {
        return m.package_escpaped
    }
}
func (m *DriveItem) GetPendingOperations()(*PendingOperations) {
    if m == nil {
        return nil
    } else {
        return m.pendingOperations
    }
}
func (m *DriveItem) GetPermissions()([]Permission) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
func (m *DriveItem) GetPhoto()(*Photo) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
func (m *DriveItem) GetPublication()(*PublicationFacet) {
    if m == nil {
        return nil
    } else {
        return m.publication
    }
}
func (m *DriveItem) GetRemoteItem()(*RemoteItem) {
    if m == nil {
        return nil
    } else {
        return m.remoteItem
    }
}
func (m *DriveItem) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
func (m *DriveItem) GetSearchResult()(*SearchResult) {
    if m == nil {
        return nil
    } else {
        return m.searchResult
    }
}
func (m *DriveItem) GetShared()(*Shared) {
    if m == nil {
        return nil
    } else {
        return m.shared
    }
}
func (m *DriveItem) GetSharepointIds()(*SharepointIds) {
    if m == nil {
        return nil
    } else {
        return m.sharepointIds
    }
}
func (m *DriveItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *DriveItem) GetSpecialFolder()(*SpecialFolder) {
    if m == nil {
        return nil
    } else {
        return m.specialFolder
    }
}
func (m *DriveItem) GetSubscriptions()([]Subscription) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
func (m *DriveItem) GetThumbnails()([]ThumbnailSet) {
    if m == nil {
        return nil
    } else {
        return m.thumbnails
    }
}
func (m *DriveItem) GetVersions()([]DriveItemVersion) {
    if m == nil {
        return nil
    } else {
        return m.versions
    }
}
func (m *DriveItem) GetVideo()(*Video) {
    if m == nil {
        return nil
    } else {
        return m.video
    }
}
func (m *DriveItem) GetWebDavUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webDavUrl
    }
}
func (m *DriveItem) GetWorkbook()(*Workbook) {
    if m == nil {
        return nil
    } else {
        return m.workbook
    }
}
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
    res["package_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPackage_escpaped() })
        if err != nil {
            return err
        }
        m.SetPackage_escpaped(val.(*Package_escpaped))
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
        err = writer.WriteObjectValue("package_escpaped", m.GetPackage_escpaped())
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
func (m *DriveItem) SetAnalytics(value *ItemAnalytics)() {
    m.analytics = value
}
func (m *DriveItem) SetAudio(value *Audio)() {
    m.audio = value
}
func (m *DriveItem) SetChildren(value []DriveItem)() {
    m.children = value
}
func (m *DriveItem) SetContent(value []byte)() {
    m.content = value
}
func (m *DriveItem) SetCTag(value *string)() {
    m.cTag = value
}
func (m *DriveItem) SetDeleted(value *Deleted)() {
    m.deleted = value
}
func (m *DriveItem) SetFile(value *File)() {
    m.file = value
}
func (m *DriveItem) SetFileSystemInfo(value *FileSystemInfo)() {
    m.fileSystemInfo = value
}
func (m *DriveItem) SetFolder(value *Folder)() {
    m.folder = value
}
func (m *DriveItem) SetImage(value *Image)() {
    m.image = value
}
func (m *DriveItem) SetListItem(value *ListItem)() {
    m.listItem = value
}
func (m *DriveItem) SetLocation(value *GeoCoordinates)() {
    m.location = value
}
func (m *DriveItem) SetPackage_escpaped(value *Package_escpaped)() {
    m.package_escpaped = value
}
func (m *DriveItem) SetPendingOperations(value *PendingOperations)() {
    m.pendingOperations = value
}
func (m *DriveItem) SetPermissions(value []Permission)() {
    m.permissions = value
}
func (m *DriveItem) SetPhoto(value *Photo)() {
    m.photo = value
}
func (m *DriveItem) SetPublication(value *PublicationFacet)() {
    m.publication = value
}
func (m *DriveItem) SetRemoteItem(value *RemoteItem)() {
    m.remoteItem = value
}
func (m *DriveItem) SetRoot(value *Root)() {
    m.root = value
}
func (m *DriveItem) SetSearchResult(value *SearchResult)() {
    m.searchResult = value
}
func (m *DriveItem) SetShared(value *Shared)() {
    m.shared = value
}
func (m *DriveItem) SetSharepointIds(value *SharepointIds)() {
    m.sharepointIds = value
}
func (m *DriveItem) SetSize(value *int64)() {
    m.size = value
}
func (m *DriveItem) SetSpecialFolder(value *SpecialFolder)() {
    m.specialFolder = value
}
func (m *DriveItem) SetSubscriptions(value []Subscription)() {
    m.subscriptions = value
}
func (m *DriveItem) SetThumbnails(value []ThumbnailSet)() {
    m.thumbnails = value
}
func (m *DriveItem) SetVersions(value []DriveItemVersion)() {
    m.versions = value
}
func (m *DriveItem) SetVideo(value *Video)() {
    m.video = value
}
func (m *DriveItem) SetWebDavUrl(value *string)() {
    m.webDavUrl = value
}
func (m *DriveItem) SetWorkbook(value *Workbook)() {
    m.workbook = value
}
