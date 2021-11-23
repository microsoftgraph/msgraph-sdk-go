package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// messageSecurityState 
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
// NewMessageSecurityState instantiates a new messageSecurityState and sets the default values.
func NewMessageSecurityState()(*MessageSecurityState) {
    m := &MessageSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnectingIP gets the connectingIP property value. 
func (m *MessageSecurityState) GetConnectingIP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectingIP
    }
}
// GetDeliveryAction gets the deliveryAction property value. 
func (m *MessageSecurityState) GetDeliveryAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryAction
    }
}
// GetDeliveryLocation gets the deliveryLocation property value. 
func (m *MessageSecurityState) GetDeliveryLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryLocation
    }
}
// GetDirectionality gets the directionality property value. 
func (m *MessageSecurityState) GetDirectionality()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directionality
    }
}
// GetInternetMessageId gets the internetMessageId property value. 
func (m *MessageSecurityState) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
// GetMessageFingerprint gets the messageFingerprint property value. 
func (m *MessageSecurityState) GetMessageFingerprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageFingerprint
    }
}
// GetMessageReceivedDateTime gets the messageReceivedDateTime property value. 
func (m *MessageSecurityState) GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.messageReceivedDateTime
    }
}
// GetMessageSubject gets the messageSubject property value. 
func (m *MessageSecurityState) GetMessageSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageSubject
    }
}
// GetNetworkMessageId gets the networkMessageId property value. 
func (m *MessageSecurityState) GetNetworkMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkMessageId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConnectingIP sets the connectingIP property value. 
func (m *MessageSecurityState) SetConnectingIP(value *string)() {
    m.connectingIP = value
}
// SetDeliveryAction sets the deliveryAction property value. 
func (m *MessageSecurityState) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
// SetDeliveryLocation sets the deliveryLocation property value. 
func (m *MessageSecurityState) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
// SetDirectionality sets the directionality property value. 
func (m *MessageSecurityState) SetDirectionality(value *string)() {
    m.directionality = value
}
// SetInternetMessageId sets the internetMessageId property value. 
func (m *MessageSecurityState) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetMessageFingerprint sets the messageFingerprint property value. 
func (m *MessageSecurityState) SetMessageFingerprint(value *string)() {
    m.messageFingerprint = value
}
// SetMessageReceivedDateTime sets the messageReceivedDateTime property value. 
func (m *MessageSecurityState) SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.messageReceivedDateTime = value
}
// SetMessageSubject sets the messageSubject property value. 
func (m *MessageSecurityState) SetMessageSubject(value *string)() {
    m.messageSubject = value
}
// SetNetworkMessageId sets the networkMessageId property value. 
func (m *MessageSecurityState) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
