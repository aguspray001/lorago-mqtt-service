package entity

type GPSLora struct {
	ID              int     `json:"id"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	GPSSpeed        float32 `json:"gps_speed"`
	TransmitterRSSI int32   `json:"transmitter_rssi"`
	ReceiverRSSI    int32   `json:"receiver_rssi"`
}
