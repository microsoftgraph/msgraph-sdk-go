package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new cloudAppSecurityState and sets the default values.
func NewCloudAppSecurityState()(*CloudAppSecurityState) {
    m := &CloudAppSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudAppSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
func (m *CloudAppSecurityState) GetDestinationServiceIp()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceIp
    }
}
// Gets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
func (m *CloudAppSecurityState) GetDestinationServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationServiceName
    }
}
// Gets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
func (m *CloudAppSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// The deserialization information for the current model
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
func (m *CloudAppSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CloudAppSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the destinationServiceIp property value. Destination IP Address of the connection to the cloud application/service.
// Parameters:
//  - value : Value to set for the destinationServiceIp property.
func (m *CloudAppSecurityState) SetDestinationServiceIp(value *string)() {
    m.destinationServiceIp = value
}
// Sets the destinationServiceName property value. Cloud application/service name (for example 'Salesforce', 'DropBox', etc.).
// Parameters:
//  - value : Value to set for the destinationServiceName property.
func (m *CloudAppSecurityState) SetDestinationServiceName(value *string)() {
    m.destinationServiceName = value
}
// Sets the riskScore property value. Provider-generated/calculated risk score of the Cloud Application/Service. Recommended value range of 0-1, which equates to a percentage.
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *CloudAppSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
