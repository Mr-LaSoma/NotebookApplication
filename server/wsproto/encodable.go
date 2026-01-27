package wsproto

// Encodable è l'interfaccia che tutte le strutture di dati
// aggiunti devono implementare per essere mandate
// tramite websocket sotto forma di [WsPacket]
type Encodable interface {
	// Encode trasforma una struttura di dati in un
	// [WsPacket] mandabile tramite websocket
	Encode() WsPacket
}