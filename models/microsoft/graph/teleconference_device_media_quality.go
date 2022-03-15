package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeleconferenceDeviceMediaQuality provides operations to call the logTeleconferenceDeviceQuality method.
type TeleconferenceDeviceMediaQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The average inbound stream network jitter.
    averageInboundJitter *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    averageInboundPacketLossRateInPercentage *float64;
    // The average inbound stream network round trip delay.
    averageInboundRoundTripDelay *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The average outbound stream network jitter.
    averageOutboundJitter *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    averageOutboundPacketLossRateInPercentage *float64;
    // The average outbound stream network round trip delay.
    averageOutboundRoundTripDelay *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
    channelIndex *int32;
    // The total number of the inbound packets.
    inboundPackets *int32;
    // the local IP address for the media session.
    localIPAddress *string;
    // The local media port.
    localPort *int32;
    // The maximum inbound stream network jitter.
    maximumInboundJitter *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    maximumInboundPacketLossRateInPercentage *float64;
    // The maximum inbound stream network round trip delay.
    maximumInboundRoundTripDelay *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The maximum outbound stream network jitter.
    maximumOutboundJitter *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
    maximumOutboundPacketLossRateInPercentage *float64;
    // The maximum outbound stream network round trip delay.
    maximumOutboundRoundTripDelay *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
    mediaDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The network link speed in bytes
    networkLinkSpeedInBytes *int32;
    // The total number of the outbound packets.
    outboundPackets *int32;
    // The remote IP address for the media session.
    remoteIPAddress *string;
    // The remote media port.
    remotePort *int32;
}
// NewTeleconferenceDeviceMediaQuality instantiates a new teleconferenceDeviceMediaQuality and sets the default values.
func NewTeleconferenceDeviceMediaQuality()(*TeleconferenceDeviceMediaQuality) {
    m := &TeleconferenceDeviceMediaQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeleconferenceDeviceMediaQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeleconferenceDeviceMediaQualityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeleconferenceDeviceMediaQuality(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceMediaQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAverageInboundJitter gets the averageInboundJitter property value. The average inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundJitter
    }
}
// GetAverageInboundPacketLossRateInPercentage gets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundPacketLossRateInPercentage
    }
}
// GetAverageInboundRoundTripDelay gets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundRoundTripDelay
    }
}
// GetAverageOutboundJitter gets the averageOutboundJitter property value. The average outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundJitter
    }
}
// GetAverageOutboundPacketLossRateInPercentage gets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundPacketLossRateInPercentage
    }
}
// GetAverageOutboundRoundTripDelay gets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundRoundTripDelay
    }
}
// GetChannelIndex gets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
func (m *TeleconferenceDeviceMediaQuality) GetChannelIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.channelIndex
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeleconferenceDeviceMediaQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["averageInboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
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
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundRoundTripDelay(val)
        }
        return nil
    }
    res["averageOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
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
        val, err := n.GetISODurationValue()
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
        val, err := n.GetInt32Value()
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
        val, err := n.GetISODurationValue()
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
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundRoundTripDelay(val)
        }
        return nil
    }
    res["maximumOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
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
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundRoundTripDelay(val)
        }
        return nil
    }
    res["mediaDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaDuration(val)
        }
        return nil
    }
    res["networkLinkSpeedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkLinkSpeedInBytes(val)
        }
        return nil
    }
    res["outboundPackets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetInboundPackets gets the inboundPackets property value. The total number of the inbound packets.
func (m *TeleconferenceDeviceMediaQuality) GetInboundPackets()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inboundPackets
    }
}
// GetLocalIPAddress gets the localIPAddress property value. the local IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) GetLocalIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIPAddress
    }
}
// GetLocalPort gets the localPort property value. The local media port.
func (m *TeleconferenceDeviceMediaQuality) GetLocalPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.localPort
    }
}
// GetMaximumInboundJitter gets the maximumInboundJitter property value. The maximum inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundJitter
    }
}
// GetMaximumInboundPacketLossRateInPercentage gets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundPacketLossRateInPercentage
    }
}
// GetMaximumInboundRoundTripDelay gets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundRoundTripDelay
    }
}
// GetMaximumOutboundJitter gets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundJitter
    }
}
// GetMaximumOutboundPacketLossRateInPercentage gets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundPacketLossRateInPercentage
    }
}
// GetMaximumOutboundRoundTripDelay gets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundRoundTripDelay
    }
}
// GetMediaDuration gets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
func (m *TeleconferenceDeviceMediaQuality) GetMediaDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.mediaDuration
    }
}
// GetNetworkLinkSpeedInBytes gets the networkLinkSpeedInBytes property value. The network link speed in bytes
func (m *TeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.networkLinkSpeedInBytes
    }
}
// GetOutboundPackets gets the outboundPackets property value. The total number of the outbound packets.
func (m *TeleconferenceDeviceMediaQuality) GetOutboundPackets()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.outboundPackets
    }
}
// GetRemoteIPAddress gets the remoteIPAddress property value. The remote IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) GetRemoteIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteIPAddress
    }
}
// GetRemotePort gets the remotePort property value. The remote media port.
func (m *TeleconferenceDeviceMediaQuality) GetRemotePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remotePort
    }
}
func (m *TeleconferenceDeviceMediaQuality) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceMediaQuality) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("averageInboundJitter", m.GetAverageInboundJitter())
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
        err := writer.WriteISODurationValue("averageInboundRoundTripDelay", m.GetAverageInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageOutboundJitter", m.GetAverageOutboundJitter())
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
        err := writer.WriteISODurationValue("averageOutboundRoundTripDelay", m.GetAverageOutboundRoundTripDelay())
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
        err := writer.WriteInt32Value("inboundPackets", m.GetInboundPackets())
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
        err := writer.WriteISODurationValue("maximumInboundJitter", m.GetMaximumInboundJitter())
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
        err := writer.WriteISODurationValue("maximumInboundRoundTripDelay", m.GetMaximumInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumOutboundJitter", m.GetMaximumOutboundJitter())
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
        err := writer.WriteISODurationValue("maximumOutboundRoundTripDelay", m.GetMaximumOutboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("mediaDuration", m.GetMediaDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("networkLinkSpeedInBytes", m.GetNetworkLinkSpeedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("outboundPackets", m.GetOutboundPackets())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceMediaQuality) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAverageInboundJitter sets the averageInboundJitter property value. The average inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.averageInboundJitter = value
    }
}
// SetAverageInboundPacketLossRateInPercentage sets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentage(value *float64)() {
    if m != nil {
        m.averageInboundPacketLossRateInPercentage = value
    }
}
// SetAverageInboundRoundTripDelay sets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.averageInboundRoundTripDelay = value
    }
}
// SetAverageOutboundJitter sets the averageOutboundJitter property value. The average outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.averageOutboundJitter = value
    }
}
// SetAverageOutboundPacketLossRateInPercentage sets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentage(value *float64)() {
    if m != nil {
        m.averageOutboundPacketLossRateInPercentage = value
    }
}
// SetAverageOutboundRoundTripDelay sets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.averageOutboundRoundTripDelay = value
    }
}
// SetChannelIndex sets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
func (m *TeleconferenceDeviceMediaQuality) SetChannelIndex(value *int32)() {
    if m != nil {
        m.channelIndex = value
    }
}
// SetInboundPackets sets the inboundPackets property value. The total number of the inbound packets.
func (m *TeleconferenceDeviceMediaQuality) SetInboundPackets(value *int32)() {
    if m != nil {
        m.inboundPackets = value
    }
}
// SetLocalIPAddress sets the localIPAddress property value. the local IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) SetLocalIPAddress(value *string)() {
    if m != nil {
        m.localIPAddress = value
    }
}
// SetLocalPort sets the localPort property value. The local media port.
func (m *TeleconferenceDeviceMediaQuality) SetLocalPort(value *int32)() {
    if m != nil {
        m.localPort = value
    }
}
// SetMaximumInboundJitter sets the maximumInboundJitter property value. The maximum inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maximumInboundJitter = value
    }
}
// SetMaximumInboundPacketLossRateInPercentage sets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentage(value *float64)() {
    if m != nil {
        m.maximumInboundPacketLossRateInPercentage = value
    }
}
// SetMaximumInboundRoundTripDelay sets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maximumInboundRoundTripDelay = value
    }
}
// SetMaximumOutboundJitter sets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maximumOutboundJitter = value
    }
}
// SetMaximumOutboundPacketLossRateInPercentage sets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentage(value *float64)() {
    if m != nil {
        m.maximumOutboundPacketLossRateInPercentage = value
    }
}
// SetMaximumOutboundRoundTripDelay sets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.maximumOutboundRoundTripDelay = value
    }
}
// SetMediaDuration sets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
func (m *TeleconferenceDeviceMediaQuality) SetMediaDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.mediaDuration = value
    }
}
// SetNetworkLinkSpeedInBytes sets the networkLinkSpeedInBytes property value. The network link speed in bytes
func (m *TeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytes(value *int32)() {
    if m != nil {
        m.networkLinkSpeedInBytes = value
    }
}
// SetOutboundPackets sets the outboundPackets property value. The total number of the outbound packets.
func (m *TeleconferenceDeviceMediaQuality) SetOutboundPackets(value *int32)() {
    if m != nil {
        m.outboundPackets = value
    }
}
// SetRemoteIPAddress sets the remoteIPAddress property value. The remote IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) SetRemoteIPAddress(value *string)() {
    if m != nil {
        m.remoteIPAddress = value
    }
}
// SetRemotePort sets the remotePort property value. The remote media port.
func (m *TeleconferenceDeviceMediaQuality) SetRemotePort(value *int32)() {
    if m != nil {
        m.remotePort = value
    }
}
