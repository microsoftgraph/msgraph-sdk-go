package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServiceHealthIssue struct {
    ServiceAnnouncementBase
    // The type of service health issue. Possible values are: advisory, incident, unknownFutureValue.
    classification *ServiceHealthClassificationType;
    // The feature name of the service issue.
    feature *string;
    // The feature group name of the service issue.
    featureGroup *string;
    // The description of the service issue impact.
    impactDescription *string;
    // Indicates whether the issue is resolved.
    isResolved *bool;
    // Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue.
    origin *ServiceHealthOrigin;
    // Collection of historical posts for the service issue.
    posts []ServiceHealthIssuePost;
    // Indicates the service affected by the issue.
    service *string;
    // The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
    status *ServiceHealthStatus;
}
// Instantiates a new serviceHealthIssue and sets the default values.
func NewServiceHealthIssue()(*ServiceHealthIssue) {
    m := &ServiceHealthIssue{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    return m
}
// Gets the classification property value. The type of service health issue. Possible values are: advisory, incident, unknownFutureValue.
func (m *ServiceHealthIssue) GetClassification()(*ServiceHealthClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// Gets the feature property value. The feature name of the service issue.
func (m *ServiceHealthIssue) GetFeature()(*string) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// Gets the featureGroup property value. The feature group name of the service issue.
func (m *ServiceHealthIssue) GetFeatureGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureGroup
    }
}
// Gets the impactDescription property value. The description of the service issue impact.
func (m *ServiceHealthIssue) GetImpactDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.impactDescription
    }
}
// Gets the isResolved property value. Indicates whether the issue is resolved.
func (m *ServiceHealthIssue) GetIsResolved()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isResolved
    }
}
// Gets the origin property value. Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue.
func (m *ServiceHealthIssue) GetOrigin()(*ServiceHealthOrigin) {
    if m == nil {
        return nil
    } else {
        return m.origin
    }
}
// Gets the posts property value. Collection of historical posts for the service issue.
func (m *ServiceHealthIssue) GetPosts()([]ServiceHealthIssuePost) {
    if m == nil {
        return nil
    } else {
        return m.posts
    }
}
// Gets the service property value. Indicates the service affected by the issue.
func (m *ServiceHealthIssue) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// Gets the status property value. The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
func (m *ServiceHealthIssue) GetStatus()(*ServiceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *ServiceHealthIssue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthClassificationType)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthClassificationType)
        m.SetClassification(&cast)
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFeature(val)
        return nil
    }
    res["featureGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFeatureGroup(val)
        return nil
    }
    res["impactDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImpactDescription(val)
        return nil
    }
    res["isResolved"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsResolved(val)
        return nil
    }
    res["origin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthOrigin)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthOrigin)
        m.SetOrigin(&cast)
        return nil
    }
    res["posts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssuePost() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealthIssuePost, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealthIssuePost))
        }
        m.SetPosts(res)
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetService(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthStatus)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *ServiceHealthIssue) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServiceHealthIssue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := m.GetClassification().String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("feature", m.GetFeature())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("featureGroup", m.GetFeatureGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactDescription", m.GetImpactDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResolved", m.GetIsResolved())
        if err != nil {
            return err
        }
    }
    if m.GetOrigin() != nil {
        cast := m.GetOrigin().String()
        err = writer.WriteStringValue("origin", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("posts", cast)
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
// Sets the classification property value. The type of service health issue. Possible values are: advisory, incident, unknownFutureValue.
// Parameters:
//  - value : Value to set for the classification property.
func (m *ServiceHealthIssue) SetClassification(value *ServiceHealthClassificationType)() {
    m.classification = value
}
// Sets the feature property value. The feature name of the service issue.
// Parameters:
//  - value : Value to set for the feature property.
func (m *ServiceHealthIssue) SetFeature(value *string)() {
    m.feature = value
}
// Sets the featureGroup property value. The feature group name of the service issue.
// Parameters:
//  - value : Value to set for the featureGroup property.
func (m *ServiceHealthIssue) SetFeatureGroup(value *string)() {
    m.featureGroup = value
}
// Sets the impactDescription property value. The description of the service issue impact.
// Parameters:
//  - value : Value to set for the impactDescription property.
func (m *ServiceHealthIssue) SetImpactDescription(value *string)() {
    m.impactDescription = value
}
// Sets the isResolved property value. Indicates whether the issue is resolved.
// Parameters:
//  - value : Value to set for the isResolved property.
func (m *ServiceHealthIssue) SetIsResolved(value *bool)() {
    m.isResolved = value
}
// Sets the origin property value. Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue.
// Parameters:
//  - value : Value to set for the origin property.
func (m *ServiceHealthIssue) SetOrigin(value *ServiceHealthOrigin)() {
    m.origin = value
}
// Sets the posts property value. Collection of historical posts for the service issue.
// Parameters:
//  - value : Value to set for the posts property.
func (m *ServiceHealthIssue) SetPosts(value []ServiceHealthIssuePost)() {
    m.posts = value
}
// Sets the service property value. Indicates the service affected by the issue.
// Parameters:
//  - value : Value to set for the service property.
func (m *ServiceHealthIssue) SetService(value *string)() {
    m.service = value
}
// Sets the status property value. The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *ServiceHealthIssue) SetStatus(value *ServiceHealthStatus)() {
    m.status = value
}
