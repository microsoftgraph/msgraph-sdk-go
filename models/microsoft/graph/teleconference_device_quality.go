package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new teleconferenceDeviceQuality and sets the default values.
func NewTeleconferenceDeviceQuality()(*TeleconferenceDeviceQuality) {
    m := &TeleconferenceDeviceQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *TeleconferenceDeviceQuality) GetCallChainId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callChainId
    }
}
// Gets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentEnvironment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentEnvironment
    }
}
// Gets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
func (m *TeleconferenceDeviceQuality) GetCloudServiceDeploymentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceDeploymentId
    }
}
// Gets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
func (m *TeleconferenceDeviceQuality) GetCloudServiceInstanceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceInstanceName
    }
}
// Gets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
func (m *TeleconferenceDeviceQuality) GetCloudServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudServiceName
    }
}
// Gets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
func (m *TeleconferenceDeviceQuality) GetDeviceDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceDescription
    }
}
// Gets the deviceName property value. The user media agent name, such as Cisco SX80.
func (m *TeleconferenceDeviceQuality) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
func (m *TeleconferenceDeviceQuality) GetMediaLegId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaLegId
    }
}
// Gets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
func (m *TeleconferenceDeviceQuality) GetMediaQualityList()([]TeleconferenceDeviceMediaQuality) {
    if m == nil {
        return nil
    } else {
        return m.mediaQualityList
    }
}
// Gets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
func (m *TeleconferenceDeviceQuality) GetParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participantId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TeleconferenceDeviceQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the callChainId property value. A unique identifier for all  the participant calls in a conference or a unique identifier for two participant calls in P2P call. This needs to be copied over from Microsoft.Graph.Call.CallChainId.
// Parameters:
//  - value : Value to set for the callChainId property.
func (m *TeleconferenceDeviceQuality) SetCallChainId(value *string)() {
    m.callChainId = value
}
// Sets the cloudServiceDeploymentEnvironment property value. A geo-region where the service is deployed, such as ProdNoam.
// Parameters:
//  - value : Value to set for the cloudServiceDeploymentEnvironment property.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentEnvironment(value *string)() {
    m.cloudServiceDeploymentEnvironment = value
}
// Sets the cloudServiceDeploymentId property value. A unique deployment identifier assigned by Azure.
// Parameters:
//  - value : Value to set for the cloudServiceDeploymentId property.
func (m *TeleconferenceDeviceQuality) SetCloudServiceDeploymentId(value *string)() {
    m.cloudServiceDeploymentId = value
}
// Sets the cloudServiceInstanceName property value. The Azure deployed cloud service instance name, such as FrontEnd_IN_3.
// Parameters:
//  - value : Value to set for the cloudServiceInstanceName property.
func (m *TeleconferenceDeviceQuality) SetCloudServiceInstanceName(value *string)() {
    m.cloudServiceInstanceName = value
}
// Sets the cloudServiceName property value. The Azure deployed cloud service name, such as contoso.cloudapp.net.
// Parameters:
//  - value : Value to set for the cloudServiceName property.
func (m *TeleconferenceDeviceQuality) SetCloudServiceName(value *string)() {
    m.cloudServiceName = value
}
// Sets the deviceDescription property value. Any additional description, such as VTC Bldg 30/21.
// Parameters:
//  - value : Value to set for the deviceDescription property.
func (m *TeleconferenceDeviceQuality) SetDeviceDescription(value *string)() {
    m.deviceDescription = value
}
// Sets the deviceName property value. The user media agent name, such as Cisco SX80.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *TeleconferenceDeviceQuality) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the mediaLegId property value. A unique identifier for a specific media leg of a participant in a conference.  One participant can have multiple media leg identifiers if retargeting happens. CVI partner assigns this value.
// Parameters:
//  - value : Value to set for the mediaLegId property.
func (m *TeleconferenceDeviceQuality) SetMediaLegId(value *string)() {
    m.mediaLegId = value
}
// Sets the mediaQualityList property value. The list of media qualities in a media session (call), such as audio quality, video quality, and/or screen sharing quality.
// Parameters:
//  - value : Value to set for the mediaQualityList property.
func (m *TeleconferenceDeviceQuality) SetMediaQualityList(value []TeleconferenceDeviceMediaQuality)() {
    m.mediaQualityList = value
}
// Sets the participantId property value. A unique identifier for a specific participant in a conference. The CVI partner needs to copy over Call.MyParticipantId to this property.
// Parameters:
//  - value : Value to set for the participantId property.
func (m *TeleconferenceDeviceQuality) SetParticipantId(value *string)() {
    m.participantId = value
}
