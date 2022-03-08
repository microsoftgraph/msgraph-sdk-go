package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenotePageable 
type OnenotePageable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    OnenoteEntitySchemaObjectModelable
    GetContent()([]byte)
    GetContentUrl()(*string)
    GetCreatedByAppId()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLevel()(*int32)
    GetLinks()(PageLinksable)
    GetOrder()(*int32)
    GetParentNotebook()(Notebookable)
    GetParentSection()(OnenoteSectionable)
    GetTitle()(*string)
    GetUserTags()([]string)
    SetContent(value []byte)()
    SetContentUrl(value *string)()
    SetCreatedByAppId(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLevel(value *int32)()
    SetLinks(value PageLinksable)()
    SetOrder(value *int32)()
    SetParentNotebook(value Notebookable)()
    SetParentSection(value OnenoteSectionable)()
    SetTitle(value *string)()
    SetUserTags(value []string)()
}
