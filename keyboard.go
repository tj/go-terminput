package terminput

// Mod represents a key modifier such as pressing alt or ctrl.
type Mod int

// Mod flags available.
const (
	ModNone Mod = 1 << iota
	ModMeta
	ModCtrl
	ModAlt
	ModShift
)

// Key represents the key pressed.
type Key rune

// Keys supported.
const (
	KeyNUL Key = iota
	KeySOH
	KeySTX
	KeyETX
	KeyEOT
	KeyENQ
	KeyACK
	KeyBEL
	KeyBS
	KeyTAB
	KeyLF
	KeyVT
	KeyFF
	KeyCR
	KeySO
	KeySI
	KeyDLE
	KeyDC1
	KeyDC2
	KeyDC3
	KeyDC4
	KeyNAK
	KeySYN
	KeyETB
	KeyCAN
	KeyEM
	KeySUB
	KeyESC
	KeyFS
	KeyGS
	KeyRS
	KeyUS

	// others
	KeyRune
	KeyUp
	KeyDown
	KeyRight
	KeyLeft
	KeyInsert
	KeyBacktab
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgUp
	KeyPgDn
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20

	// variants
	keyShiftLeft
	keyShiftRight
	keyAltLeft
	keyAltRight
)

// Aliases
const (
	KeyBackspace = KeyBS
	KeyTab       = KeyTAB
	KeyEscape    = KeyESC
	KeyEnter     = KeyCR
)

var keySequences = map[string]Key{
	"\x1b[A":     KeyUp,
	"\x1b[B":     KeyDown,
	"\x1b[C":     KeyRight,
	"\x1b[D":     KeyLeft,
	"\x1b[2~":    KeyInsert,
	"\x1b[3~":    KeyDelete,
	"\u007f":     KeyBackspace,
	"\x1b[Z":     KeyBacktab,
	"\x1bOH":     KeyHome,
	"\x1bOF":     KeyEnd,
	"\x1b[5~":    KeyPgUp,
	"\x1b[6~":    KeyPgDn,
	"\x1bOP":     KeyF1,
	"\x1bOQ":     KeyF2,
	"\x1bOR":     KeyF3,
	"\x1bOS":     KeyF4,
	"\x1b[15~":   KeyF5,
	"\x1b[17~":   KeyF6,
	"\x1b[18~":   KeyF7,
	"\x1b[19~":   KeyF8,
	"\x1b[20~":   KeyF9,
	"\x1b[21~":   KeyF10,
	"\x1b[23~":   KeyF11,
	"\x1b[24~":   KeyF12,
	"\x1b[1;2P":  KeyF13,
	"\x1b[1;2Q":  KeyF14,
	"\x1b[1;2R":  KeyF15,
	"\x1b[1;2S":  KeyF16,
	"\x1b[15;2~": KeyF17,
	"\x1b[17;2~": KeyF18,
	"\x1b[18;2~": KeyF19,
	"\x1b[19;2~": KeyF20,
	"\x1b[1;2D":  keyShiftLeft,
	"\x1b[1;2C":  keyShiftRight,
	"\x1bb":      keyAltLeft,
	"\x1bf":      keyAltRight,
}

// keyNames is a map of keys to a human-friendly representation.
var keyNames = map[Key]string{
	KeyNUL:     "NUL",
	KeySOH:     "SOH",
	KeySTX:     "STX",
	KeyETX:     "ETX",
	KeyEOT:     "EOT",
	KeyENQ:     "ENQ",
	KeyACK:     "ACK",
	KeyBEL:     "BEL",
	KeyBS:      "Backspace",
	KeyTAB:     "Tab",
	KeyLF:      "LF",
	KeyVT:      "VT",
	KeyFF:      "FF",
	KeyCR:      "CR",
	KeySO:      "SO",
	KeySI:      "SI",
	KeyDLE:     "DLE",
	KeyDC1:     "DC1",
	KeyDC2:     "DC2",
	KeyDC3:     "DC3",
	KeyDC4:     "DC4",
	KeyNAK:     "NAK",
	KeySYN:     "SYN",
	KeyETB:     "ETB",
	KeyCAN:     "CAN",
	KeyEM:      "EM",
	KeySUB:     "SUB",
	KeyESC:     "Escape",
	KeyFS:      "FS",
	KeyGS:      "GS",
	KeyRS:      "RS",
	KeyUS:      "US",
	KeyUp:      "Up",
	KeyDown:    "Down",
	KeyRight:   "Right",
	KeyLeft:    "Left",
	KeyInsert:  "Insert",
	KeyDelete:  "Delete",
	KeyBacktab: "Backtab",
	KeyHome:    "Home",
	KeyEnd:     "End",
	KeyPgUp:    "PgUp",
	KeyPgDn:    "PgDn",
	KeyF1:      "F1",
	KeyF2:      "F2",
	KeyF3:      "F3",
	KeyF4:      "F4",
	KeyF5:      "F5",
	KeyF6:      "F6",
	KeyF7:      "F7",
	KeyF8:      "F8",
	KeyF9:      "F9",
	KeyF10:     "F10",
	KeyF11:     "F11",
	KeyF12:     "F12",
	KeyF13:     "F13",
	KeyF14:     "F14",
	KeyF15:     "F15",
	KeyF16:     "F16",
	KeyF17:     "F17",
	KeyF18:     "F18",
	KeyF19:     "F19",
	KeyF20:     "F20",
}

// KeyboardInput represents a key press.
type KeyboardInput struct {
	k Key
	r rune
	m Mod
}

// String returns a human-friendly string representing the key.
func (k *KeyboardInput) String() string {
	if s, ok := keyNames[k.Key()]; ok {
		switch {
		case k.Alt():
			s = "Alt+" + s
		case k.Ctrl():
			s = "Ctrl+" + s
		case k.Shift():
			s = "Shift+" + s
		case k.Meta():
			s = "Meta+" + s
		}
		return s
	}
	return string(k.Rune())
}

// Key returns the key pressed. In the case that the
// key returned is KeyRune, the literal value can be
// accessed via Rune().
func (k *KeyboardInput) Key() Key {
	return k.k
}

// Rune returns the rune corresponding to the key press.
// This value is only defined when Key() is KeyRune.
func (k *KeyboardInput) Rune() rune {
	return k.r
}

// Mod returns the modifier flags.
func (k *KeyboardInput) Mod() Mod {
	return k.m
}

// Ctrl returns true if ctrl was pressed.
func (k *KeyboardInput) Ctrl() bool {
	return k.Mod()&ModCtrl != 0
}

// Alt returns true if alt was pressed.
func (k *KeyboardInput) Alt() bool {
	return k.Mod()&ModAlt != 0
}

// Meta returns true if meta was pressed.
func (k *KeyboardInput) Meta() bool {
	return k.Mod()&ModMeta != 0
}

// Shift returns true if shift was pressed.
func (k *KeyboardInput) Shift() bool {
	return k.Mod()&ModShift != 0
}

// newKeyboardInput returns a keyboard input event.
func newKeyboardInput(k Key, r rune, m Mod) *KeyboardInput {
	return &KeyboardInput{k, r, m}
}
