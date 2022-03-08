package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceHealthIssueable 
type ServiceHealthIssueable interface {
    ServiceAnnouncementBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClassification()(*ServiceHealthClassificationType)
    GetFeature()(*string)
    GetFeatureGroup()(*string)
    GetImpactDescription()(*string)
    GetIsResolved()(*bool)
    GetOrigin()(*ServiceHealthOrigin)
    GetPosts()([]ServiceHealthIssuePostable)
    GetService()(*string)
    GetStatus()(*ServiceHealthStatus)
    SetClassification(value *ServiceHealthClassificationType)()
    SetFeature(value *string)()
    SetFeatureGroup(value *string)()
    SetImpactDescription(value *string)()
    SetIsResolved(value *bool)()
    SetOrigin(value *ServiceHealthOrigin)()
    SetPosts(value []ServiceHealthIssuePostable)()
    SetService(value *string)()
    SetStatus(value *ServiceHealthStatus)()
}
