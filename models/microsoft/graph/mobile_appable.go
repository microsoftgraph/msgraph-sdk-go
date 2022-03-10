package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppable 
type MobileAppable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]MobileAppAssignmentable)
    GetCategories()([]MobileAppCategoryable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeveloper()(*string)
    GetDisplayName()(*string)
    GetInformationUrl()(*string)
    GetIsFeatured()(*bool)
    GetLargeIcon()(MimeContentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()(*string)
    GetOwner()(*string)
    GetPrivacyInformationUrl()(*string)
    GetPublisher()(*string)
    GetPublishingState()(*MobileAppPublishingState)
    SetAssignments(value []MobileAppAssignmentable)()
    SetCategories(value []MobileAppCategoryable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeveloper(value *string)()
    SetDisplayName(value *string)()
    SetInformationUrl(value *string)()
    SetIsFeatured(value *bool)()
    SetLargeIcon(value MimeContentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value *string)()
    SetOwner(value *string)()
    SetPrivacyInformationUrl(value *string)()
    SetPublisher(value *string)()
    SetPublishingState(value *MobileAppPublishingState)()
}
