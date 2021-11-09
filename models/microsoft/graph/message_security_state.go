package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    connectingIP *string;
    // 
    deliveryAction *string;
    // 
    deliveryLocation *string;
    // 
    directionality *string;
    // 
    internetMessageId *string;
    // 
    messageFingerprint *string;
    // 
    messageReceivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    messageSubject *string;
    // 
    networkMessageId *string;
}
// Instantiates a new messageSecurityState and sets the default values.
func NewMessageSecurityState()(*MessageSecurityState) {
    m := &MessageSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the connectingIP property value. 
func (m *MessageSecurityState) GetConnectingIP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectingIP
    }
}
// Gets the deliveryAction property value. 
func (m *MessageSecurityState) GetDeliveryAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryAction
    }
}
// Gets the deliveryLocation property value. 
func (m *MessageSecurityState) GetDeliveryLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryLocation
    }
}
// Gets the directionality property value. 
func (m *MessageSecurityState) GetDirectionality()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directionality
    }
}
// Gets the internetMessageId property value. 
func (m *MessageSecurityState) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
// Gets the messageFingerprint property value. 
func (m *MessageSecurityState) GetMessageFingerprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageFingerprint
    }
}
// Gets the messageReceivedDateTime property value. 
func (m *MessageSecurityState) GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.messageReceivedDateTime
    }
}
// Gets the messageSubject property value. 
func (m *MessageSecurityState) GetMessageSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageSubject
    }
}
// Gets the networkMessageId property value. 
func (m *MessageSecurityState) GetNetworkMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkMessageId
    }
}
// The deserialization information for the current model
func (m *MessageSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connectingIP"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectingIP(val)
        }
        return nil
    }
    res["deliveryAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryAction(val)
        }
        return nil
    }
    res["deliveryLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryLocation(val)
        }
        return nil
    }
    res["directionality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectionality(val)
        }
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["messageFingerprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageFingerprint(val)
        }
        return nil
    }
    res["messageReceivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageReceivedDateTime(val)
        }
        return nil
    }
    res["messageSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageSubject(val)
        }
        return nil
    }
    res["networkMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkMessageId(val)
        }
        return nil
    }
    return res
}
func (m *MessageSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("connectingIP", m.GetConnectingIP())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deliveryAction", m.GetDeliveryAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deliveryLocation", m.GetDeliveryLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("directionality", m.GetDirectionality())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageFingerprint", m.GetMessageFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("messageReceivedDateTime", m.GetMessageReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageSubject", m.GetMessageSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
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
func (m *MessageSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the connectingIP property value. 
// Parameters:
//  - value : Value to set for the connectingIP property.
func (m *MessageSecurityState) SetConnectingIP(value *string)() {
    m.connectingIP = value
}
// Sets the deliveryAction property value. 
// Parameters:
//  - value : Value to set for the deliveryAction property.
func (m *MessageSecurityState) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
// Sets the deliveryLocation property value. 
// Parameters:
//  - value : Value to set for the deliveryLocation property.
func (m *MessageSecurityState) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
// Sets the directionality property value. 
// Parameters:
//  - value : Value to set for the directionality property.
func (m *MessageSecurityState) SetDirectionality(value *string)() {
    m.directionality = value
}
// Sets the internetMessageId property value. 
// Parameters:
//  - value : Value to set for the internetMessageId property.
func (m *MessageSecurityState) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// Sets the messageFingerprint property value. 
// Parameters:
//  - value : Value to set for the messageFingerprint property.
func (m *MessageSecurityState) SetMessageFingerprint(value *string)() {
    m.messageFingerprint = value
}
// Sets the messageReceivedDateTime property value. 
// Parameters:
//  - value : Value to set for the messageReceivedDateTime property.
func (m *MessageSecurityState) SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.messageReceivedDateTime = value
}
// Sets the messageSubject property value. 
// Parameters:
//  - value : Value to set for the messageSubject property.
func (m *MessageSecurityState) SetMessageSubject(value *string)() {
    m.messageSubject = value
}
// Sets the networkMessageId property value. 
// Parameters:
//  - value : Value to set for the networkMessageId property.
func (m *MessageSecurityState) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
