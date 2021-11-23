package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// teleconferenceDeviceQuality 
type TeleconferenceDeviceQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
    callChainId *string;
    // A geo-region where the service is deployed, such as ProdNoam.
    cloudServiceDeploymentEnvironment *string;
    // A unique deployment identifier assigned by Azure.
    cloudServiceDeploymentId *string;
    // The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
    cloudServiceInstanceName *string;
    // The Azure deployed cloud service name, such as contoso.cloudapp.net.
    cloudServiceName *string;
    // Any additional description, such as VTC Bldg 30/21.
    deviceDescription *string;
    // The user media agent name, such as Cisco SX80.
    deviceName *string;
    // A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
    mediaLegId *string;
    // The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
    mediaQualityList []TeleconferenceDeviceMediaQuality;
    // A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
    participantId *string;
}
// NewTeleconferenceDeviceQuality instantiates a new teleconferenceDeviceQuality and sets the default values.
func NewTeleconferenceDeviceQuality()(*TeleconferenceDeviceQuality) {
    m := &TeleconferenceDeviceQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallChainId gets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *TeleconferenceDeviceQuality) GetCallChainId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callChainId
    }
}
// GetCloudServiceDeploymentEnvironment gets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentEnvironment
    }
}
// GetCloudServiceDeploymentId gets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentId
    }
}
// GetCloudServiceInstanceName gets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
func (m *TeleconferenceDeviceQuality) GetCloudServiceInstanceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceInstanceName
    }
}
// GetCloudServiceName gets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
func (m *TeleconferenceDeviceQuality) GetCloudServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceName
    }
}
// GetDeviceDescription gets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
func (m *TeleconferenceDeviceQuality) GetDeviceDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDescription
    }
}
// GetDeviceName gets the deviceName property value. The user media agent name, such as Cisco SX80.
func (m *TeleconferenceDeviceQuality) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetMediaLegId gets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
func (m *TeleconferenceDeviceQuality) GetMediaLegId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaLegId
    }
}
// GetMediaQualityList gets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
func (m *TeleconferenceDeviceQuality) GetMediaQualityList()([]TeleconferenceDeviceMediaQuality) {
    if m == nil {
        return nil
    } else {
        return m.mediaQualityList
    }
}
// GetParticipantId gets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
func (m *TeleconferenceDeviceQuality) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeleconferenceDeviceQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["callChainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallChainId(val)
        }
        return nil
    }
    res["cloudServiceDeploymentEnvironment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceDeploymentEnvironment(val)
        }
        return nil
    }
    res["cloudServiceDeploymentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceDeploymentId(val)
        }
        return nil
    }
    res["cloudServiceInstanceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceInstanceName(val)
        }
        return nil
    }
    res["cloudServiceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceName(val)
        }
        return nil
    }
    res["deviceDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDescription(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["mediaLegId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaLegId(val)
        }
        return nil
    }
    res["mediaQualityList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeleconferenceDeviceMediaQuality() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeleconferenceDeviceMediaQuality, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeleconferenceDeviceMediaQuality))
            }
            m.SetMediaQualityList(res)
        }
        return nil
    }
    res["participantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantId(val)
        }
        return nil
    }
    return res
}
func (m *TeleconferenceDeviceQuality) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallChainId sets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *TeleconferenceDeviceQuality) SetCallChainId(value *string)() {
    m.callChainId = value
}
// SetCloudServiceDeploymentEnvironment sets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironment(value *string)() {
    m.cloudServiceDeploymentEnvironment = value
}
// SetCloudServiceDeploymentId sets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentId(value *string)() {
    m.cloudServiceDeploymentId = value
}
// SetCloudServiceInstanceName sets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
func (m *TeleconferenceDeviceQuality) SetCloudServiceInstanceName(value *string)() {
    m.cloudServiceInstanceName = value
}
// SetCloudServiceName sets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
func (m *TeleconferenceDeviceQuality) SetCloudServiceName(value *string)() {
    m.cloudServiceName = value
}
// SetDeviceDescription sets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
func (m *TeleconferenceDeviceQuality) SetDeviceDescription(value *string)() {
    m.deviceDescription = value
}
// SetDeviceName sets the deviceName property value. The user media agent name, such as Cisco SX80.
func (m *TeleconferenceDeviceQuality) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetMediaLegId sets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
func (m *TeleconferenceDeviceQuality) SetMediaLegId(value *string)() {
    m.mediaLegId = value
}
// SetMediaQualityList sets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
func (m *TeleconferenceDeviceQuality) SetMediaQualityList(value []TeleconferenceDeviceMediaQuality)() {
    m.mediaQualityList = value
}
// SetParticipantId sets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
func (m *TeleconferenceDeviceQuality) SetParticipantId(value *string)() {
    m.participantId = value
}
