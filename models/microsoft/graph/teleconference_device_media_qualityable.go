package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeleconferenceDeviceMediaQualityable 
type TeleconferenceDeviceMediaQualityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAverageInboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAverageInboundPacketLossRateInPercentage()(*float64)
    GetAverageInboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAverageOutboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAverageOutboundPacketLossRateInPercentage()(*float64)
    GetAverageOutboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetChannelIndex()(*int32)
    GetInboundPackets()(*int32)
    GetLocalIPAddress()(*string)
    GetLocalPort()(*int32)
    GetMaximumInboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMaximumInboundPacketLossRateInPercentage()(*float64)
    GetMaximumInboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMaximumOutboundJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMaximumOutboundPacketLossRateInPercentage()(*float64)
    GetMaximumOutboundRoundTripDelay()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMediaDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetNetworkLinkSpeedInBytes()(*int32)
    GetOutboundPackets()(*int32)
    GetRemoteIPAddress()(*string)
    GetRemotePort()(*int32)
    SetAverageInboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAverageInboundPacketLossRateInPercentage(value *float64)()
    SetAverageInboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAverageOutboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAverageOutboundPacketLossRateInPercentage(value *float64)()
    SetAverageOutboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetChannelIndex(value *int32)()
    SetInboundPackets(value *int32)()
    SetLocalIPAddress(value *string)()
    SetLocalPort(value *int32)()
    SetMaximumInboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMaximumInboundPacketLossRateInPercentage(value *float64)()
    SetMaximumInboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMaximumOutboundJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMaximumOutboundPacketLossRateInPercentage(value *float64)()
    SetMaximumOutboundRoundTripDelay(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMediaDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetNetworkLinkSpeedInBytes(value *int32)()
    SetOutboundPackets(value *int32)()
    SetRemoteIPAddress(value *string)()
    SetRemotePort(value *int32)()
}
