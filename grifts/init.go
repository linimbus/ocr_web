package grifts

import (
	"github.com/linimbus/ocr_web/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
