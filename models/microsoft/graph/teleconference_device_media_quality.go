package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeleconferenceDeviceMediaQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The average inbound stream network jitter.
    averageInboundJitter *string;
    // The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    averageInboundPacketLossRateInPercentage *float64;
    // The average inbound stream network round trip delay.
    averageInboundRoundTripDelay *string;
    // The average outbound stream network jitter.
    averageOutboundJitter *string;
    // The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    averageOutboundPacketLossRateInPercentage *float64;
    // The average outbound stream network round trip delay.
    averageOutboundRoundTripDelay *string;
    // The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
    channelIndex *int32;
    // The total number of the inbound packets.
    inboundPackets *int64;
    // the local IP address for the media session.
    localIPAddress *string;
    // The local media port.
    localPort *int32;
    // The maximum inbound stream network jitter.
    maximumInboundJitter *string;
    // The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    maximumInboundPacketLossRateInPercentage *float64;
    // The maximum inbound stream network round trip delay.
    maximumInboundRoundTripDelay *string;
    // The maximum outbound stream network jitter.
    maximumOutboundJitter *string;
    // The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    maximumOutboundPacketLossRateInPercentage *float64;
    // The maximum outbound stream network round trip delay.
    maximumOutboundRoundTripDelay *string;
    // The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
    mediaDuration *string;
    // The network link speed in bytes
    networkLinkSpeedInBytes *int64;
    // The total number of the outbound packets.
    outboundPackets *int64;
    // The remote IP address for the media session.
    remoteIPAddress *string;
    // The remote media port.
    remotePort *int32;
}
// Instantiates a new teleconferenceDeviceMediaQuality and sets the default values.
func NewTeleconferenceDeviceMediaQuality()(*TeleconferenceDeviceMediaQuality) {
    m := &TeleconferenceDeviceMediaQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceMediaQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the averageInboundJitter property value. The average inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundJitter
    }
}
// Gets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundPacketLossRateInPercentage
    }
}
// Gets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundRoundTripDelay
    }
}
// Gets the averageOutboundJitter property value. The average outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundJitter
    }
}
// Gets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundPacketLossRateInPercentage
    }
}
// Gets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundRoundTripDelay
    }
}
// Gets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
func (m *TeleconferenceDeviceMediaQuality) GetChannelIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.channelIndex
    }
}
// Gets the inboundPackets property value. The total number of the inbound packets.
func (m *TeleconferenceDeviceMediaQuality) GetInboundPackets()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.inboundPackets
    }
}
// Gets the localIPAddress property value. the local IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) GetLocalIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIPAddress
    }
}
// Gets the localPort property value. The local media port.
func (m *TeleconferenceDeviceMediaQuality) GetLocalPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.localPort
    }
}
// Gets the maximumInboundJitter property value. The maximum inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundJitter
    }
}
// Gets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundPacketLossRateInPercentage
    }
}
// Gets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundRoundTripDelay
    }
}
// Gets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundJitter
    }
}
// Gets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundPacketLossRateInPercentage
    }
}
// Gets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundRoundTripDelay
    }
}
// Gets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
func (m *TeleconferenceDeviceMediaQuality) GetMediaDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaDuration
    }
}
// Gets the networkLinkSpeedInBytes property value. The network link speed in bytes
func (m *TeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkLinkSpeedInBytes
    }
}
// Gets the outboundPackets property value. The total number of the outbound packets.
func (m *TeleconferenceDeviceMediaQuality) GetOutboundPackets()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outboundPackets
    }
}
// Gets the remoteIPAddress property value. The remote IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) GetRemoteIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteIPAddress
    }
}
// Gets the remotePort property value. The remote media port.
func (m *TeleconferenceDeviceMediaQuality) GetRemotePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remotePort
    }
}
// The deserialization information for the current model
func (m *TeleconferenceDeviceMediaQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["averageInboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundJitter(val)
        }
        return nil
    }
    res["averageInboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundPacketLossRateInPercentage(val)
        }
        return nil
    }
    res["averageInboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundRoundTripDelay(val)
        }
        return nil
    }
    res["averageOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundJitter(val)
        }
        return nil
    }
    res["averageOutboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundPacketLossRateInPercentage(val)
        }
        return nil
    }
    res["averageOutboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundRoundTripDelay(val)
        }
        return nil
    }
    res["channelIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelIndex(val)
        }
        return nil
    }
    res["inboundPackets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundPackets(val)
        }
        return nil
    }
    res["localIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIPAddress(val)
        }
        return nil
    }
    res["localPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalPort(val)
        }
        return nil
    }
    res["maximumInboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundJitter(val)
        }
        return nil
    }
    res["maximumInboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundPacketLossRateInPercentage(val)
        }
        return nil
    }
    res["maximumInboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundRoundTripDelay(val)
        }
        return nil
    }
    res["maximumOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundJitter(val)
        }
        return nil
    }
    res["maximumOutboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundPacketLossRateInPercentage(val)
        }
        return nil
    }
    res["maximumOutboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundRoundTripDelay(val)
        }
        return nil
    }
    res["mediaDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaDuration(val)
        }
        return nil
    }
    res["networkLinkSpeedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkLinkSpeedInBytes(val)
        }
        return nil
    }
    res["outboundPackets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundPackets(val)
        }
        return nil
    }
    res["remoteIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteIPAddress(val)
        }
        return nil
    }
    res["remotePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemotePort(val)
        }
        return nil
    }
    return res
}
func (m *TeleconferenceDeviceMediaQuality) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeleconferenceDeviceMediaQuality) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("averageInboundJitter", m.GetAverageInboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("averageInboundPacketLossRateInPercentage", m.GetAverageInboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("averageInboundRoundTripDelay", m.GetAverageInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("averageOutboundJitter", m.GetAverageOutboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("averageOutboundPacketLossRateInPercentage", m.GetAverageOutboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("averageOutboundRoundTripDelay", m.GetAverageOutboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("channelIndex", m.GetChannelIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("inboundPackets", m.GetInboundPackets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localIPAddress", m.GetLocalIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("localPort", m.GetLocalPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maximumInboundJitter", m.GetMaximumInboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximumInboundPacketLossRateInPercentage", m.GetMaximumInboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maximumInboundRoundTripDelay", m.GetMaximumInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maximumOutboundJitter", m.GetMaximumOutboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximumOutboundPacketLossRateInPercentage", m.GetMaximumOutboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maximumOutboundRoundTripDelay", m.GetMaximumOutboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaDuration", m.GetMediaDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("networkLinkSpeedInBytes", m.GetNetworkLinkSpeedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("outboundPackets", m.GetOutboundPackets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteIPAddress", m.GetRemoteIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remotePort", m.GetRemotePort())
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
func (m *TeleconferenceDeviceMediaQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the averageInboundJitter property value. The average inbound stream network jitter.
// Parameters:
//  - value : Value to set for the averageInboundJitter property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundJitter(value *string)() {
    m.averageInboundJitter = value
}
// Sets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// Parameters:
//  - value : Value to set for the averageInboundPacketLossRateInPercentage property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentage(value *float64)() {
    m.averageInboundPacketLossRateInPercentage = value
}
// Sets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
// Parameters:
//  - value : Value to set for the averageInboundRoundTripDelay property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelay(value *string)() {
    m.averageInboundRoundTripDelay = value
}
// Sets the averageOutboundJitter property value. The average outbound stream network jitter.
// Parameters:
//  - value : Value to set for the averageOutboundJitter property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundJitter(value *string)() {
    m.averageOutboundJitter = value
}
// Sets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// Parameters:
//  - value : Value to set for the averageOutboundPacketLossRateInPercentage property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentage(value *float64)() {
    m.averageOutboundPacketLossRateInPercentage = value
}
// Sets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
// Parameters:
//  - value : Value to set for the averageOutboundRoundTripDelay property.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelay(value *string)() {
    m.averageOutboundRoundTripDelay = value
}
// Sets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
// Parameters:
//  - value : Value to set for the channelIndex property.
func (m *TeleconferenceDeviceMediaQuality) SetChannelIndex(value *int32)() {
    m.channelIndex = value
}
// Sets the inboundPackets property value. The total number of the inbound packets.
// Parameters:
//  - value : Value to set for the inboundPackets property.
func (m *TeleconferenceDeviceMediaQuality) SetInboundPackets(value *int64)() {
    m.inboundPackets = value
}
// Sets the localIPAddress property value. the local IP address for the media session.
// Parameters:
//  - value : Value to set for the localIPAddress property.
func (m *TeleconferenceDeviceMediaQuality) SetLocalIPAddress(value *string)() {
    m.localIPAddress = value
}
// Sets the localPort property value. The local media port.
// Parameters:
//  - value : Value to set for the localPort property.
func (m *TeleconferenceDeviceMediaQuality) SetLocalPort(value *int32)() {
    m.localPort = value
}
// Sets the maximumInboundJitter property value. The maximum inbound stream network jitter.
// Parameters:
//  - value : Value to set for the maximumInboundJitter property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundJitter(value *string)() {
    m.maximumInboundJitter = value
}
// Sets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// Parameters:
//  - value : Value to set for the maximumInboundPacketLossRateInPercentage property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentage(value *float64)() {
    m.maximumInboundPacketLossRateInPercentage = value
}
// Sets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
// Parameters:
//  - value : Value to set for the maximumInboundRoundTripDelay property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelay(value *string)() {
    m.maximumInboundRoundTripDelay = value
}
// Sets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
// Parameters:
//  - value : Value to set for the maximumOutboundJitter property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundJitter(value *string)() {
    m.maximumOutboundJitter = value
}
// Sets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// Parameters:
//  - value : Value to set for the maximumOutboundPacketLossRateInPercentage property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentage(value *float64)() {
    m.maximumOutboundPacketLossRateInPercentage = value
}
// Sets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
// Parameters:
//  - value : Value to set for the maximumOutboundRoundTripDelay property.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelay(value *string)() {
    m.maximumOutboundRoundTripDelay = value
}
// Sets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
// Parameters:
//  - value : Value to set for the mediaDuration property.
func (m *TeleconferenceDeviceMediaQuality) SetMediaDuration(value *string)() {
    m.mediaDuration = value
}
// Sets the networkLinkSpeedInBytes property value. The network link speed in bytes
// Parameters:
//  - value : Value to set for the networkLinkSpeedInBytes property.
func (m *TeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytes(value *int64)() {
    m.networkLinkSpeedInBytes = value
}
// Sets the outboundPackets property value. The total number of the outbound packets.
// Parameters:
//  - value : Value to set for the outboundPackets property.
func (m *TeleconferenceDeviceMediaQuality) SetOutboundPackets(value *int64)() {
    m.outboundPackets = value
}
// Sets the remoteIPAddress property value. The remote IP address for the media session.
// Parameters:
//  - value : Value to set for the remoteIPAddress property.
func (m *TeleconferenceDeviceMediaQuality) SetRemoteIPAddress(value *string)() {
    m.remoteIPAddress = value
}
// Sets the remotePort property value. The remote media port.
// Parameters:
//  - value : Value to set for the remotePort property.
func (m *TeleconferenceDeviceMediaQuality) SetRemotePort(value *int32)() {
    m.remotePort = value
}
