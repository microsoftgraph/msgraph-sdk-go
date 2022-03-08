package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyNotebookModelable 
type CopyNotebookModelable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCreatedBy()(*string)
    GetCreatedByIdentity()(IdentitySetable)
    GetCreatedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*string)
    GetIsDefault()(*bool)
    GetIsShared()(*bool)
    GetLastModifiedBy()(*string)
    GetLastModifiedByIdentity()(IdentitySetable)
    GetLastModifiedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinks()(NotebookLinksable)
    GetName()(*string)
    GetSectionGroupsUrl()(*string)
    GetSectionsUrl()(*string)
    GetSelf()(*string)
    GetUserRole()(*OnenoteUserRole)
    SetCreatedBy(value *string)()
    SetCreatedByIdentity(value IdentitySetable)()
    SetCreatedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *string)()
    SetIsDefault(value *bool)()
    SetIsShared(value *bool)()
    SetLastModifiedBy(value *string)()
    SetLastModifiedByIdentity(value IdentitySetable)()
    SetLastModifiedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinks(value NotebookLinksable)()
    SetName(value *string)()
    SetSectionGroupsUrl(value *string)()
    SetSectionsUrl(value *string)()
    SetSelf(value *string)()
    SetUserRole(value *OnenoteUserRole)()
}
