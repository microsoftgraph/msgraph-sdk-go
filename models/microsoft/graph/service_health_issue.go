package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceHealthIssue 
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
    // The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. See more in the table below.
    status *ServiceHealthStatus;
}
// NewServiceHealthIssue instantiates a new serviceHealthIssue and sets the default values.
func NewServiceHealthIssue()(*ServiceHealthIssue) {
    m := &ServiceHealthIssue{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    return m
}
// GetClassification gets the classification property value. The type of service health issue. Possible values are: advisory, incident, unknownFutureValue.
func (m *ServiceHealthIssue) GetClassification()(*ServiceHealthClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetFeature gets the feature property value. The feature name of the service issue.
func (m *ServiceHealthIssue) GetFeature()(*string) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// GetFeatureGroup gets the featureGroup property value. The feature group name of the service issue.
func (m *ServiceHealthIssue) GetFeatureGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureGroup
    }
}
// GetImpactDescription gets the impactDescription property value. The description of the service issue impact.
func (m *ServiceHealthIssue) GetImpactDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.impactDescription
    }
}
// GetIsResolved gets the isResolved property value. Indicates whether the issue is resolved.
func (m *ServiceHealthIssue) GetIsResolved()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isResolved
    }
}
// GetOrigin gets the origin property value. Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue.
func (m *ServiceHealthIssue) GetOrigin()(*ServiceHealthOrigin) {
    if m == nil {
        return nil
    } else {
        return m.origin
    }
}
// GetPosts gets the posts property value. Collection of historical posts for the service issue.
func (m *ServiceHealthIssue) GetPosts()([]ServiceHealthIssuePost) {
    if m == nil {
        return nil
    } else {
        return m.posts
    }
}
// GetService gets the service property value. Indicates the service affected by the issue.
func (m *ServiceHealthIssue) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetStatus gets the status property value. The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. See more in the table below.
func (m *ServiceHealthIssue) GetStatus()(*ServiceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceHealthIssue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*ServiceHealthClassificationType))
        }
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val)
        }
        return nil
    }
    res["featureGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureGroup(val)
        }
        return nil
    }
    res["impactDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactDescription(val)
        }
        return nil
    }
    res["isResolved"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsResolved(val)
        }
        return nil
    }
    res["origin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthOrigin)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrigin(val.(*ServiceHealthOrigin))
        }
        return nil
    }
    res["posts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssuePost() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssuePost, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceHealthIssuePost))
            }
            m.SetPosts(res)
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
            m.SetStatus(val.(*ServiceHealthStatus))
        }
        return nil
    }
    return res
}
func (m *ServiceHealthIssue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ServiceHealthIssue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
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
        cast := (*m.GetOrigin()).String()
        err = writer.WriteStringValue("origin", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPosts() != nil {
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
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. The type of service health issue. Possible values are: advisory, incident, unknownFutureValue.
func (m *ServiceHealthIssue) SetClassification(value *ServiceHealthClassificationType)() {
    if m != nil {
        m.classification = value
    }
}
// SetFeature sets the feature property value. The feature name of the service issue.
func (m *ServiceHealthIssue) SetFeature(value *string)() {
    if m != nil {
        m.feature = value
    }
}
// SetFeatureGroup sets the featureGroup property value. The feature group name of the service issue.
func (m *ServiceHealthIssue) SetFeatureGroup(value *string)() {
    if m != nil {
        m.featureGroup = value
    }
}
// SetImpactDescription sets the impactDescription property value. The description of the service issue impact.
func (m *ServiceHealthIssue) SetImpactDescription(value *string)() {
    if m != nil {
        m.impactDescription = value
    }
}
// SetIsResolved sets the isResolved property value. Indicates whether the issue is resolved.
func (m *ServiceHealthIssue) SetIsResolved(value *bool)() {
    if m != nil {
        m.isResolved = value
    }
}
// SetOrigin sets the origin property value. Indicates the origin of the service issue. Possible values are: microsoft, thirdParty, customer, unknownFutureValue.
func (m *ServiceHealthIssue) SetOrigin(value *ServiceHealthOrigin)() {
    if m != nil {
        m.origin = value
    }
}
// SetPosts sets the posts property value. Collection of historical posts for the service issue.
func (m *ServiceHealthIssue) SetPosts(value []ServiceHealthIssuePost)() {
    if m != nil {
        m.posts = value
    }
}
// SetService sets the service property value. Indicates the service affected by the issue.
func (m *ServiceHealthIssue) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetStatus sets the status property value. The status of the service issue. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue. See more in the table below.
func (m *ServiceHealthIssue) SetStatus(value *ServiceHealthStatus)() {
    if m != nil {
        m.status = value
    }
}
