package sessions

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func New() {
	store = session.New(session.Config{})
}

func Add(c *fiber.Ctx, key string, value interface{}) {
	var sess, _ = store.Get(c)
	sess.Set(key, value)
	sess.SetExpiry(time.Second * 60 * 60 * 24 * 365)
	sess.Save()
}

func Get(c *fiber.Ctx, key string) interface{} {
	var sess, _ = store.Get(c)
	var value = sess.Get(key)
	return value
}

func Delete(c *fiber.Ctx) {
	var sess, _ = store.Get(c)
	sess.Destroy()
	sess.Save()
}
