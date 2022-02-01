package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceHealth 
type ServiceHealth struct {
    Entity
    // A collection of issues that happened on the service, with detailed information for each issue.
    issues []ServiceHealthIssue;
    // The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
    service *string;
    // Show the overall service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. For more details, see serviceHealthStatus values.
    status *ServiceHealthStatus;
}
// NewServiceHealth instantiates a new serviceHealth and sets the default values.
func NewServiceHealth()(*ServiceHealth) {
    m := &ServiceHealth{
        Entity: *NewEntity(),
    }
    return m
}
// GetIssues gets the issues property value. A collection of issues that happened on the service, with detailed information for each issue.
func (m *ServiceHealth) GetIssues()([]ServiceHealthIssue) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
// GetService gets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
func (m *ServiceHealth) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetStatus gets the status property value. Show the overall service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. For more details, see serviceHealthStatus values.
func (m *ServiceHealth) GetStatus()(*ServiceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *ServiceHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIssues() != nil {
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
// SetIssues sets the issues property value. A collection of issues that happened on the service, with detailed information for each issue.
func (m *ServiceHealth) SetIssues(value []ServiceHealthIssue)() {
    if m != nil {
        m.issues = value
    }
}
// SetService sets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
func (m *ServiceHealth) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetStatus sets the status property value. Show the overall service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. For more details, see serviceHealthStatus values.
func (m *ServiceHealth) SetStatus(value *ServiceHealthStatus)() {
    if m != nil {
        m.status = value
    }
}
