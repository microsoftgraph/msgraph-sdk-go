package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppConsentApprovalRoute provides operations to manage the identityGovernance singleton.
type AppConsentApprovalRoute struct {
    Entity
    // A collection of userConsentRequest objects for a specific application.
    appConsentRequests []AppConsentRequestable;
}
// NewAppConsentApprovalRoute instantiates a new appConsentApprovalRoute and sets the default values.
func NewAppConsentApprovalRoute()(*AppConsentApprovalRoute) {
    m := &AppConsentApprovalRoute{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppConsentApprovalRouteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppConsentApprovalRouteFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppConsentApprovalRoute(), nil
}
// GetAppConsentRequests gets the appConsentRequests property value. A collection of userConsentRequest objects for a specific application.
func (m *AppConsentApprovalRoute) GetAppConsentRequests()([]AppConsentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.appConsentRequests
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppConsentApprovalRoute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appConsentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppConsentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppConsentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AppConsentRequestable)
            }
            m.SetAppConsentRequests(res)
        }
        return nil
    }
    return res
}
func (m *AppConsentApprovalRoute) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppConsentApprovalRoute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppConsentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppConsentRequests()))
        for i, v := range m.GetAppConsentRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appConsentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppConsentRequests sets the appConsentRequests property value. A collection of userConsentRequest objects for a specific application.
func (m *AppConsentApprovalRoute) SetAppConsentRequests(value []AppConsentRequestable)() {
    if m != nil {
        m.appConsentRequests = value
    }
}
