/*
Package wsproto definisce la struttura del protocollo utilizzato da Skribio
per la comunicazione tra client e server.

Il protocollo si basa sullo scambio di messaggi in formato JSON, suddivisi
in due sezioni principali: Header e Body.

	    +-------------------------------+
	    |             HEADER            |
	    |-------------------------------|
	    |                               |
	    |   Contiene il tipo di         |
	    |   messaggio che il            |
	    |   client o il server          |
	    |   deve interpretare           |
	    |                               |
	    +-------------------------------+
	    |              BODY             |
	    |-------------------------------|
	    |                               |
	    |   Contiene dati aggiuntivi    |
	    |   specifici per il tipo di    |
	    |   messaggio ricevuto,         |
	    |   codificati in formato JSON  |
	    |                               |
	    +-------------------------------+

L'Header viene utilizzato per determinare la logica di gestione del messaggio,
mentre il Body trasporta le informazioni necessarie all'elaborazione dello stesso.
*/
package wsproto