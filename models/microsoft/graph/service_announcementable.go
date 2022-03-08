package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceAnnouncementable 
type ServiceAnnouncementable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetHealthOverviews()([]ServiceHealthable)
    GetIssues()([]ServiceHealthIssueable)
    GetMessages()([]ServiceUpdateMessageable)
    SetHealthOverviews(value []ServiceHealthable)()
    SetIssues(value []ServiceHealthIssueable)()
    SetMessages(value []ServiceUpdateMessageable)()
}
