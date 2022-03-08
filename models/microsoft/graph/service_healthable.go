package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceHealthable 
type ServiceHealthable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIssues()([]ServiceHealthIssueable)
    GetService()(*string)
    GetStatus()(*ServiceHealthStatus)
    SetIssues(value []ServiceHealthIssueable)()
    SetService(value *string)()
    SetStatus(value *ServiceHealthStatus)()
}
