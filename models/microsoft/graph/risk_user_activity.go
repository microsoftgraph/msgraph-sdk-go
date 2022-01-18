package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskUserActivity 
type RiskUserActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
    detail *RiskDetail;
    // The type of risk event detected.
    riskEventTypes []string;
}
// NewRiskUserActivity instantiates a new riskUserActivity and sets the default values.
func NewRiskUserActivity()(*RiskUserActivity) {
    m := &RiskUserActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskUserActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetail gets the detail property value. The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskUserActivity) GetDetail()(*RiskDetail) {
    if m == nil {
        return nil
    } else {
        return m.detail
    }
}
// GetRiskEventTypes gets the riskEventTypes property value. The type of risk event detected.
func (m *RiskUserActivity) GetRiskEventTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.riskEventTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskUserActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["detail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RiskDetail)
            m.SetDetail(&cast)
        }
        return nil
    }
    res["riskEventTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypes(res)
        }
        return nil
    }
    return res
}
func (m *RiskUserActivity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RiskUserActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := m.GetDetail().String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("riskEventTypes", m.GetRiskEventTypes())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskUserActivity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetail sets the detail property value. The possible values are none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue.
func (m *RiskUserActivity) SetDetail(value *RiskDetail)() {
    if m != nil {
        m.detail = value
    }
}
// SetRiskEventTypes sets the riskEventTypes property value. The type of risk event detected.
func (m *RiskUserActivity) SetRiskEventTypes(value []string)() {
    if m != nil {
        m.riskEventTypes = value
    }
}
