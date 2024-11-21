package hangmanweb

import (
	h "hangmanweb/hangmanweb"
)

const port = ":8080"

func main() {
	h.Init()
	h.Web()
}
