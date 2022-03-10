package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RemoteItemable 
type RemoteItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFile()(Fileable)
    GetFileSystemInfo()(FileSystemInfoable)
    GetFolder()(Folderable)
    GetId()(*string)
    GetImage()(Imageable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetPackage()(Package_escapedable)
    GetParentReference()(ItemReferenceable)
    GetShared()(Sharedable)
    GetSharepointIds()(SharepointIdsable)
    GetSize()(*int64)
    GetSpecialFolder()(SpecialFolderable)
    GetVideo()(Videoable)
    GetWebDavUrl()(*string)
    GetWebUrl()(*string)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFile(value Fileable)()
    SetFileSystemInfo(value FileSystemInfoable)()
    SetFolder(value Folderable)()
    SetId(value *string)()
    SetImage(value Imageable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetPackage(value Package_escapedable)()
    SetParentReference(value ItemReferenceable)()
    SetShared(value Sharedable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSize(value *int64)()
    SetSpecialFolder(value SpecialFolderable)()
    SetVideo(value Videoable)()
    SetWebDavUrl(value *string)()
    SetWebUrl(value *string)()
}
