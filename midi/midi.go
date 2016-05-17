package midi

import "fmt"

type Midi struct {
	Header *Header
	Tracks []Track
}

func (m *Midi) NumberOfTracks() int {

	return int(m.Header.NumberOfTracks)
}

func (m *Midi) Division() int {

	return int(m.Header.Division)
}

func NewMidi(h *Header, t []Track) *Midi {

	return &Midi{h, t}
}

func (m *Midi) String() string {

	return fmt.Sprintf("Midi [Header=%v, Tracks=%v]",
		m.Header,
		m.Tracks)
}