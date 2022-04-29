package session

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
	"bufio"
	"github.com/eniehack/nikkiamev2/internal/handler"
)

type sessionOption struct {

}

func (h *Handler) initial(b []byte, option *sessionOption) {
	str := base32.StdEncoding.EncodeToString(b)
	h.
}
