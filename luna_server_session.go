package main

type LunaServerSession struct {
	Session LunaSocketSession
	_mother *LunaSockets
}

// Wird auf einer Serversitzung verwendet um Daten zu empfangen
func (obj *LunaServerSession) Serve() error {
	obj.Session._connected = true
	reval := obj._mother._wrappWS(&obj.Session)
	obj.Session._connected = false
	return reval
}

// Fügt dieser Sitzung einen Parameter hinzu
func (obj *LunaServerSession) AddOutParameter(data interface{}) {
	obj.Session.AddOutParameter(data)
}
