package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceUpdateMessageable 
type ServiceUpdateMessageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ServiceAnnouncementBaseable
    GetActionRequiredByDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAttachments()([]ServiceAnnouncementAttachmentable)
    GetAttachmentsArchive()([]byte)
    GetBody()(ItemBodyable)
    GetCategory()(*ServiceUpdateCategory)
    GetHasAttachments()(*bool)
    GetIsMajorChange()(*bool)
    GetServices()([]string)
    GetSeverity()(*ServiceUpdateSeverity)
    GetTags()([]string)
    GetViewPoint()(ServiceUpdateMessageViewpointable)
    SetActionRequiredByDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAttachments(value []ServiceAnnouncementAttachmentable)()
    SetAttachmentsArchive(value []byte)()
    SetBody(value ItemBodyable)()
    SetCategory(value *ServiceUpdateCategory)()
    SetHasAttachments(value *bool)()
    SetIsMajorChange(value *bool)()
    SetServices(value []string)()
    SetSeverity(value *ServiceUpdateSeverity)()
    SetTags(value []string)()
    SetViewPoint(value ServiceUpdateMessageViewpointable)()
}
