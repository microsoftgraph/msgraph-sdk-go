package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageSecurityState struct {
    additionalData map[string]interface{};
    connectingIP *string;
    deliveryAction *string;
    deliveryLocation *string;
    directionality *string;
    internetMessageId *string;
    messageFingerprint *string;
    messageReceivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    messageSubject *string;
    networkMessageId *string;
}
func NewMessageSecurityState()(*MessageSecurityState) {
    m := &MessageSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MessageSecurityState) GetConnectingIP()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectingIP
    }
}
func (m *MessageSecurityState) GetDeliveryAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryAction
    }
}
func (m *MessageSecurityState) GetDeliveryLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deliveryLocation
    }
}
func (m *MessageSecurityState) GetDirectionality()(*string) {
    if m == nil {
        return nil
    } else {
        return m.directionality
    }
}
func (m *MessageSecurityState) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
func (m *MessageSecurityState) GetMessageFingerprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageFingerprint
    }
}
func (m *MessageSecurityState) GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.messageReceivedDateTime
    }
}
func (m *MessageSecurityState) GetMessageSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageSubject
    }
}
func (m *MessageSecurityState) GetNetworkMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkMessageId
    }
}
func (m *MessageSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connectingIP"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConnectingIP(val)
        return nil
    }
    res["deliveryAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeliveryAction(val)
        return nil
    }
    res["deliveryLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeliveryLocation(val)
        return nil
    }
    res["directionality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDirectionality(val)
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternetMessageId(val)
        return nil
    }
    res["messageFingerprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessageFingerprint(val)
        return nil
    }
    res["messageReceivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMessageReceivedDateTime(val)
        return nil
    }
    res["messageSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessageSubject(val)
        return nil
    }
    res["networkMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkMessageId(val)
        return nil
    }
    return res
}
func (m *MessageSecurityState) IsNil()(bool) {
    return m == nil
}
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
func (m *MessageSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageSecurityState) SetConnectingIP(value *string)() {
    m.connectingIP = value
}
func (m *MessageSecurityState) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
func (m *MessageSecurityState) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
func (m *MessageSecurityState) SetDirectionality(value *string)() {
    m.directionality = value
}
func (m *MessageSecurityState) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
func (m *MessageSecurityState) SetMessageFingerprint(value *string)() {
    m.messageFingerprint = value
}
func (m *MessageSecurityState) SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.messageReceivedDateTime = value
}
func (m *MessageSecurityState) SetMessageSubject(value *string)() {
    m.messageSubject = value
}
func (m *MessageSecurityState) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
