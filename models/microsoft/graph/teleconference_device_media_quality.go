package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeleconferenceDeviceMediaQuality struct {
    additionalData map[string]interface{};
    averageInboundJitter *string;
    averageInboundPacketLossRateInPercentage *float64;
    averageInboundRoundTripDelay *string;
    averageOutboundJitter *string;
    averageOutboundPacketLossRateInPercentage *float64;
    averageOutboundRoundTripDelay *string;
    channelIndex *int32;
    inboundPackets *int64;
    localIPAddress *string;
    localPort *int32;
    maximumInboundJitter *string;
    maximumInboundPacketLossRateInPercentage *float64;
    maximumInboundRoundTripDelay *string;
    maximumOutboundJitter *string;
    maximumOutboundPacketLossRateInPercentage *float64;
    maximumOutboundRoundTripDelay *string;
    mediaDuration *string;
    networkLinkSpeedInBytes *int64;
    outboundPackets *int64;
    remoteIPAddress *string;
    remotePort *int32;
}
func NewTeleconferenceDeviceMediaQuality()(*TeleconferenceDeviceMediaQuality) {
    m := &TeleconferenceDeviceMediaQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeleconferenceDeviceMediaQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundJitter
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundPacketLossRateInPercentage
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageInboundRoundTripDelay
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundJitter
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundPacketLossRateInPercentage
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.averageOutboundRoundTripDelay
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetChannelIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.channelIndex
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetInboundPackets()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.inboundPackets
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetLocalIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localIPAddress
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetLocalPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.localPort
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundJitter
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundPacketLossRateInPercentage
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumInboundRoundTripDelay
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundJitter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundJitter
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundPacketLossRateInPercentage
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelay()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumOutboundRoundTripDelay
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetMediaDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaDuration
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.networkLinkSpeedInBytes
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetOutboundPackets()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outboundPackets
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetRemoteIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteIPAddress
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetRemotePort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remotePort
    }
}
func (m *TeleconferenceDeviceMediaQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["averageInboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAverageInboundJitter(val)
        return nil
    }
    res["averageInboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAverageInboundPacketLossRateInPercentage(val)
        return nil
    }
    res["averageInboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAverageInboundRoundTripDelay(val)
        return nil
    }
    res["averageOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAverageOutboundJitter(val)
        return nil
    }
    res["averageOutboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAverageOutboundPacketLossRateInPercentage(val)
        return nil
    }
    res["averageOutboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAverageOutboundRoundTripDelay(val)
        return nil
    }
    res["channelIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetChannelIndex(val)
        return nil
    }
    res["inboundPackets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetInboundPackets(val)
        return nil
    }
    res["localIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocalIPAddress(val)
        return nil
    }
    res["localPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLocalPort(val)
        return nil
    }
    res["maximumInboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumInboundJitter(val)
        return nil
    }
    res["maximumInboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaximumInboundPacketLossRateInPercentage(val)
        return nil
    }
    res["maximumInboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumInboundRoundTripDelay(val)
        return nil
    }
    res["maximumOutboundJitter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumOutboundJitter(val)
        return nil
    }
    res["maximumOutboundPacketLossRateInPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaximumOutboundPacketLossRateInPercentage(val)
        return nil
    }
    res["maximumOutboundRoundTripDelay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaximumOutboundRoundTripDelay(val)
        return nil
    }
    res["mediaDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaDuration(val)
        return nil
    }
    res["networkLinkSpeedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetNetworkLinkSpeedInBytes(val)
        return nil
    }
    res["outboundPackets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutboundPackets(val)
        return nil
    }
    res["remoteIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoteIPAddress(val)
        return nil
    }
    res["remotePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemotePort(val)
        return nil
    }
    return res
}
func (m *TeleconferenceDeviceMediaQuality) IsNil()(bool) {
    return m == nil
}
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
func (m *TeleconferenceDeviceMediaQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundJitter(value *string)() {
    m.averageInboundJitter = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentage(value *float64)() {
    m.averageInboundPacketLossRateInPercentage = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelay(value *string)() {
    m.averageInboundRoundTripDelay = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundJitter(value *string)() {
    m.averageOutboundJitter = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentage(value *float64)() {
    m.averageOutboundPacketLossRateInPercentage = value
}
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelay(value *string)() {
    m.averageOutboundRoundTripDelay = value
}
func (m *TeleconferenceDeviceMediaQuality) SetChannelIndex(value *int32)() {
    m.channelIndex = value
}
func (m *TeleconferenceDeviceMediaQuality) SetInboundPackets(value *int64)() {
    m.inboundPackets = value
}
func (m *TeleconferenceDeviceMediaQuality) SetLocalIPAddress(value *string)() {
    m.localIPAddress = value
}
func (m *TeleconferenceDeviceMediaQuality) SetLocalPort(value *int32)() {
    m.localPort = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundJitter(value *string)() {
    m.maximumInboundJitter = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentage(value *float64)() {
    m.maximumInboundPacketLossRateInPercentage = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelay(value *string)() {
    m.maximumInboundRoundTripDelay = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundJitter(value *string)() {
    m.maximumOutboundJitter = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentage(value *float64)() {
    m.maximumOutboundPacketLossRateInPercentage = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelay(value *string)() {
    m.maximumOutboundRoundTripDelay = value
}
func (m *TeleconferenceDeviceMediaQuality) SetMediaDuration(value *string)() {
    m.mediaDuration = value
}
func (m *TeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytes(value *int64)() {
    m.networkLinkSpeedInBytes = value
}
func (m *TeleconferenceDeviceMediaQuality) SetOutboundPackets(value *int64)() {
    m.outboundPackets = value
}
func (m *TeleconferenceDeviceMediaQuality) SetRemoteIPAddress(value *string)() {
    m.remoteIPAddress = value
}
func (m *TeleconferenceDeviceMediaQuality) SetRemotePort(value *int32)() {
    m.remotePort = value
}
