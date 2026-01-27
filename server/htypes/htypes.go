package htypes

// HeaderType serve per determinare la logica di gestione
// di un messaggio mandato tramite websocket
type HeaderType int

const (
	err_beg HeaderType = iota // Determina l'inizio degli errori
	err_end                   // Determina la fine degli errori

	msg_beg // Determina l'inizio dei messaggi
	msg_end // Determina la fine degli messaggi
)

// IsErr è una funzione helper per capire se
// un headertype rietra tra gli errori
func (ht HeaderType) IsErr() bool { return err_beg < ht && ht < err_end }

// IsMsg è una funzione helper per capire se
// un headertype rietra tra i messaggi normali
func (ht HeaderType) IsMsg() bool { return msg_beg < ht && ht < msg_end }
