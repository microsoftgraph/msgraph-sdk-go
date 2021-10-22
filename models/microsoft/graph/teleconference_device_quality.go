package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeleconferenceDeviceQuality struct {
    additionalData map[string]interface{};
    callChainId *string;
    cloudServiceDeploymentEnvironment *string;
    cloudServiceDeploymentId *string;
    cloudServiceInstanceName *string;
    cloudServiceName *string;
    deviceDescription *string;
    deviceName *string;
    mediaLegId *string;
    mediaQualityList []TeleconferenceDeviceMediaQuality;
    participantId *string;
}
func NewTeleconferenceDeviceQuality()(*TeleconferenceDeviceQuality) {
    m := &TeleconferenceDeviceQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeleconferenceDeviceQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeleconferenceDeviceQuality) GetCallChainId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callChainId
    }
}
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentEnvironment
    }
}
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentId
    }
}
func (m *TeleconferenceDeviceQuality) GetCloudServiceInstanceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceInstanceName
    }
}
func (m *TeleconferenceDeviceQuality) GetCloudServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceName
    }
}
func (m *TeleconferenceDeviceQuality) GetDeviceDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDescription
    }
}
func (m *TeleconferenceDeviceQuality) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *TeleconferenceDeviceQuality) GetMediaLegId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaLegId
    }
}
func (m *TeleconferenceDeviceQuality) GetMediaQualityList()([]TeleconferenceDeviceMediaQuality) {
    if m == nil {
        return nil
    } else {
        return m.mediaQualityList
    }
}
func (m *TeleconferenceDeviceQuality) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
func (m *TeleconferenceDeviceQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["callChainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCallChainId(val)
        return nil
    }
    res["cloudServiceDeploymentEnvironment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCloudServiceDeploymentEnvironment(val)
        return nil
    }
    res["cloudServiceDeploymentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCloudServiceDeploymentId(val)
        return nil
    }
    res["cloudServiceInstanceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCloudServiceInstanceName(val)
        return nil
    }
    res["cloudServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCloudServiceName(val)
        return nil
    }
    res["deviceDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceDescription(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["mediaLegId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaLegId(val)
        return nil
    }
    res["mediaQualityList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeleconferenceDeviceMediaQuality() })
        if err != nil {
            return err
        }
        res := make([]TeleconferenceDeviceMediaQuality, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeleconferenceDeviceMediaQuality))
        }
        m.SetMediaQualityList(res)
        return nil
    }
    res["participantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParticipantId(val)
        return nil
    }
    return res
}
func (m *TeleconferenceDeviceQuality) IsNil()(bool) {
    return m == nil
}
func (m *TeleconferenceDeviceQuality) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callChainId", m.GetCallChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceDeploymentEnvironment", m.GetCloudServiceDeploymentEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceDeploymentId", m.GetCloudServiceDeploymentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceInstanceName", m.GetCloudServiceInstanceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceName", m.GetCloudServiceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceDescription", m.GetDeviceDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaLegId", m.GetMediaLegId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMediaQualityList()))
        for i, v := range m.GetMediaQualityList() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("mediaQualityList", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("participantId", m.GetParticipantId())
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
func (m *TeleconferenceDeviceQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeleconferenceDeviceQuality) SetCallChainId(value *string)() {
    m.callChainId = value
}
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironment(value *string)() {
    m.cloudServiceDeploymentEnvironment = value
}
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentId(value *string)() {
    m.cloudServiceDeploymentId = value
}
func (m *TeleconferenceDeviceQuality) SetCloudServiceInstanceName(value *string)() {
    m.cloudServiceInstanceName = value
}
func (m *TeleconferenceDeviceQuality) SetCloudServiceName(value *string)() {
    m.cloudServiceName = value
}
func (m *TeleconferenceDeviceQuality) SetDeviceDescription(value *string)() {
    m.deviceDescription = value
}
func (m *TeleconferenceDeviceQuality) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *TeleconferenceDeviceQuality) SetMediaLegId(value *string)() {
    m.mediaLegId = value
}
func (m *TeleconferenceDeviceQuality) SetMediaQualityList(value []TeleconferenceDeviceMediaQuality)() {
    m.mediaQualityList = value
}
func (m *TeleconferenceDeviceQuality) SetParticipantId(value *string)() {
    m.participantId = value
}
