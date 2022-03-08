package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudAppSecurityState provides operations to manage the security singleton.
type CloudAppSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Destination IP Address of the connection to the cloud application/service.
    destinationServiceIp *string;
    // Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
    destinationServiceName *string;
    // Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
}
// NewCloudAppSecurityState instantiates a new cloudAppSecurityState and sets the default values.
func NewCloudAppSecurityState()(*CloudAppSecurityState) {
    m := &CloudAppSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudAppSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudAppSecurityStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCloudAppSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudAppSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDestinationServiceIp gets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
func (m *CloudAppSecurityState) GetDestinationServiceIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceIp
    }
}
// GetDestinationServiceName gets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
func (m *CloudAppSecurityState) GetDestinationServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudAppSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationServiceIp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationServiceIp(val)
        }
        return nil
    }
    res["destinationServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationServiceName(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    return res
}
// GetRiskScore gets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
func (m *CloudAppSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *CloudAppSecurityState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudAppSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationServiceIp", m.GetDestinationServiceIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationServiceName", m.GetDestinationServiceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
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
func (m *CloudAppSecurityState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDestinationServiceIp sets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
func (m *CloudAppSecurityState) SetDestinationServiceIp(value *string)() {
    if m != nil {
        m.destinationServiceIp = value
    }
}
// SetDestinationServiceName sets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
func (m *CloudAppSecurityState) SetDestinationServiceName(value *string)() {
    if m != nil {
        m.destinationServiceName = value
    }
}
// SetRiskScore sets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
func (m *CloudAppSecurityState) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
