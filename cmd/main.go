package main 

import ( 
  "text/template" 
  "io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
  templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

type Contact struct {
  Name string
  Email string
}

func newContact(name string, email string) Contact {
  return Contact{
    Name: name,
    Email: email,
  }
}

type Contacts = []Contact

type Data struct {
  Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
  for _, contact := range d.Contacts {
    if contact.Email == email {
      return true
    }
  }
  return false
}

func newData() Data {
  return Data{
    Contacts: []Contact{
      newContact("John", "jd@gmail.com"),
      newContact("Clara", "cd@gmail.com"),
    },
  }
}

type Count struct {
  Count int
}

type FormData struct {
  Values map[string]string
  Errors map[string]string
}

func newFormData() FormData {
  return FormData{
    Values: make(map[string]string),
    Errors: make(map[string]string),
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  data := newData()
  e.Renderer = newTemplate()

  e.GET("/", func(c echo.Context) error {
    return c.Render(200, "index", data)
  })

  e.POST("/contacts", func(c echo.Context) error {
    name := c.FormValue("name")
    email := c.FormValue("email")

    if data.hasEmail(email) {


    data.Contacts = append(data.Contacts, newContact(name, email))
    return c.Render(200, "display", data)
  })
  
  e.Logger.Fatal(e.Start(":42069"))
}
