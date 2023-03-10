package wtf

import (
	"github.com/wtfutil/wtf/cfg"

	"github.com/rivo/tview"
)

// Wtfable is the interface that enforces WTF system capabilities on a module
type Wtfable interface {
	Enablable
	Schedulable
	Stoppable

	BorderColor() string
	ConfigText() string
	FocusChar() string
	Focusable() bool
	HelpText() string
	Initialize()
	Name() string
	QuitChan() chan bool
	SetFocusChar(string)
	TextView() *tview.TextView

	CommonSettings() *cfg.Common
}
