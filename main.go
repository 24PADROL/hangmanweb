package main

import (
	h "hangmanweb/hangmanweb"
)

func main() {
	h.Init()
	h.Web()
	h.CheckWin()
}
