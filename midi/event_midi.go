package midi

import "fmt"

type MidiEvent struct {
	deltaTime int
	status    byte
	data      []byte
}

func (e *MidiEvent) DeltaTime() int {
	return e.deltaTime
}

func (e *MidiEvent) Data() []byte {
	return e.data
}

func (e *MidiEvent) Status() byte {
	return e.status
}

func (e *MidiEvent) String() string {
	return fmt.Sprintf("MidiEvent [Status=%X, DeltaTime=%d, Data=%d]",
		e.Status(),
		e.DeltaTime(),
		e.Data())
}

func NewMidiEvent(t int, s byte, d []byte) *MidiEvent {
	return &MidiEvent{t, s, d}
}
