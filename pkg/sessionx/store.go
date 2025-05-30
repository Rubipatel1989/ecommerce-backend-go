package sessionx

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
)

var Store *session.Store

func InitSessionStore() {
	storage := memory.New()
	Store = session.New(session.Config{
		Storage: storage,
	})
}
