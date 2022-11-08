package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type listTodo struct {
	lista []*todo
}

func (lt *listTodo) mostrarLista() []*todo {
	return lt.lista
}

func (lt *listTodo) mostrarTarea(Id string) *todo {
	if Id != "" {
		for _, todo := range lt.lista {
			if todo.Id == Id {
				return todo
			}
		}
	}
	return nil
}

func (lt *listTodo) agregarALista(todo *todo) {
	lt.lista = append(lt.lista, todo)
}

func (lt *listTodo) actualizarEstado(Id string) {
	for _, todo := range lt.lista {
		if todo.Id == Id {
			todo.Estado = true
		}
	}
}

func (lt *listTodo) removerTarea(Id string) {
	for index, todo := range lt.lista {
		if todo.Id == Id {
			lt.lista = append(lt.lista[:index], lt.lista[index+1:]...)
		}
	}
}

type todo struct {
	Id     string
	Nombre string
	Estado bool
}

func main() {

	lt := &listTodo{}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(lt.mostrarLista())
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		p := c.Params("id")
		todo := lt.mostrarTarea(p)
		return c.Status(fiber.StatusOK).JSON(todo)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		p := &todo{}
		if err := c.BodyParser(p); err != nil {
			return err
		}

		p.Id = uuid.New().String()
		lt.agregarALista(p)
		return c.Status(fiber.StatusOK).JSON(p)
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		p := c.Params("id")
		lt.actualizarEstado(p)
		return c.SendString("Estado actualizado")
	})

	app.Delete("/:id", func(c *fiber.Ctx) error {
		p := c.Params("id")
		lt.removerTarea(p)
		return c.SendString("Tarea eliminada")
	})

	app.Listen(":3000")
}
