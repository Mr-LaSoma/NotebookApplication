package wsproto

import (
	"encoding/json"

	"github.com/team-condominio/skribio/htypes"
)

// WsPacket è la struttura che viene mandata tramite
// ws per la comunicazione client-server
type WsPacket struct {
	// Contiene il tipo di messaggio che
	// il client o il server deve interpretare
	Header htypes.HeaderType

	// Contiene dati aggiuntivi specifici per il tipo
	// di messaggio ricevuto, codificato in formato JSON.
	// 
	// In caso non si voglia mandare informazioni aggiuntive utilizzare:
	// 	wp := WsPacket{
	//		Header: TIPO_DA_INSERIRE
	//		Body: json.RawMessage(`null`)
	//	} 
	Body json.RawMessage
}
