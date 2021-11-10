package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServiceHealth struct {
    Entity
    // A collection of issues happened on the service, with detailed information for each issue.
    issues []ServiceHealthIssue;
    // The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
    service *string;
    // Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
    status *ServiceHealthStatus;
}
// Instantiates a new serviceHealth and sets the default values.
func NewServiceHealth()(*ServiceHealth) {
    m := &ServiceHealth{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the issues property value. A collection of issues happened on the service, with detailed information for each issue.
func (m *ServiceHealth) GetIssues()([]ServiceHealthIssue) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
// Gets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
func (m *ServiceHealth) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// Gets the status property value. Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
func (m *ServiceHealth) GetStatus()(*ServiceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *ServiceHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["issues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssue, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceHealthIssue))
            }
            m.SetIssues(res)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ServiceHealthStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *ServiceHealth) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServiceHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIssues()))
        for i, v := range m.GetIssues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("issues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the issues property value. A collection of issues happened on the service, with detailed information for each issue.
// Parameters:
//  - value : Value to set for the issues property.
func (m *ServiceHealth) SetIssues(value []ServiceHealthIssue)() {
    m.issues = value
}
// Sets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
// Parameters:
//  - value : Value to set for the service property.
func (m *ServiceHealth) SetService(value *string)() {
    m.service = value
}
// Sets the status property value. Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *ServiceHealth) SetStatus(value *ServiceHealthStatus)() {
    m.status = value
}
