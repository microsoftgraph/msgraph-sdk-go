package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityGovernance struct {
    accessReviews *AccessReviewSet;
    additionalData map[string]interface{};
    appConsent *AppConsentApprovalRoute;
    entitlementManagement *EntitlementManagement;
    termsOfUse *TermsOfUseContainer;
}
func NewIdentityGovernance()(*IdentityGovernance) {
    m := &IdentityGovernance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *IdentityGovernance) GetAccessReviews()(*AccessReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.accessReviews
    }
}
func (m *IdentityGovernance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *IdentityGovernance) GetAppConsent()(*AppConsentApprovalRoute) {
    if m == nil {
        return nil
    } else {
        return m.appConsent
    }
}
func (m *IdentityGovernance) GetEntitlementManagement()(*EntitlementManagement) {
    if m == nil {
        return nil
    } else {
        return m.entitlementManagement
    }
}
func (m *IdentityGovernance) GetTermsOfUse()(*TermsOfUseContainer) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
func (m *IdentityGovernance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessReviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewSet() })
        if err != nil {
            return err
        }
        m.SetAccessReviews(val.(*AccessReviewSet))
        return nil
    }
    res["appConsent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppConsentApprovalRoute() })
        if err != nil {
            return err
        }
        m.SetAppConsent(val.(*AppConsentApprovalRoute))
        return nil
    }
    res["entitlementManagement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntitlementManagement() })
        if err != nil {
            return err
        }
        m.SetEntitlementManagement(val.(*EntitlementManagement))
        return nil
    }
    res["termsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsOfUseContainer() })
        if err != nil {
            return err
        }
        m.SetTermsOfUse(val.(*TermsOfUseContainer))
        return nil
    }
    return res
}
func (m *IdentityGovernance) IsNil()(bool) {
    return m == nil
}
func (m *IdentityGovernance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accessReviews", m.GetAccessReviews())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("appConsent", m.GetAppConsent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("entitlementManagement", m.GetEntitlementManagement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("termsOfUse", m.GetTermsOfUse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *IdentityGovernance) SetAccessReviews(value *AccessReviewSet)() {
    m.accessReviews = value
}
func (m *IdentityGovernance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *IdentityGovernance) SetAppConsent(value *AppConsentApprovalRoute)() {
    m.appConsent = value
}
func (m *IdentityGovernance) SetEntitlementManagement(value *EntitlementManagement)() {
    m.entitlementManagement = value
}
func (m *IdentityGovernance) SetTermsOfUse(value *TermsOfUseContainer)() {
    m.termsOfUse = value
}
