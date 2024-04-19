package main

import (
	"fmt"
	"os"
	"time"
)

type UserRole int

const (
	Admin UserRole = iota
	Normal
)

type Usuario struct {
	Username string
	Role     UserRole
}

var listaUsuarios = []Usuario{
	{Username: "admin", Role: Admin},
	{Username: "usuario", Role: Normal},
}

var activeSessions = map[string]Session{}

type Session struct {
	User      Usuario
	StartTime time.Time
}

type Tarea struct {
	ID          int
	Descripcion string
	Estado      string
}

var tareas = []Tarea{
	{ID: 1, Descripcion: "Implementar la función X", Estado: "Pendiente"},
	{ID: 2, Descripcion: "Resolver el problema Y", Estado: "En progreso"},
	{ID: 3, Descripcion: "Actualizar la documentación", Estado: "Completado"},
}

func iniciarSesion(username string) (*Session, error) {
	var usuario Usuario
	for _, u := range listaUsuarios {
		if u.Username == username {
			usuario = u
			break
		}
	}

	if usuario.Username == "" {
		return nil, fmt.Errorf("Usuario no encontrado")
	}

	session := Session{
		User:      usuario,
		StartTime: time.Now(),
	}

	activeSessions[username] = session

	return &session, nil
}

func cerrarSesion(username string) {
	delete(activeSessions, username)
}

func verificarRol(username string, requiredRole UserRole) bool {
	session, ok := activeSessions[username]
	if !ok {
		return false
	}
	return session.User.Role == requiredRole
}

func obtenerTareas() []Tarea {
	return tareas
}

func generarInforme() {

	fmt.Println("Generando informe...")
	time.Sleep(2 * time.Second)
	fmt.Println("Informe generado exitosamente!")
}

func comunicacionTiempoReal() {

	fmt.Println("Iniciando comunicación en tiempo real...")
	time.Sleep(3 * time.Second)
	fmt.Println("Comunicación en tiempo real establecida.")
}

func crearArchivoTexto(username string) {
	fmt.Printf("%s está creando un archivo de texto...\n", username)
	time.Sleep(2 * time.Second)

	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	if verificarRol(username, Admin) {
		fmt.Println("Registrando cantidad de teclas pulsadas...")
		//robotgo.EventStart()
		//defer robotgo.EventEnd()
	}
	fmt.Println("Archivo de texto creado exitosamente.")
}

func main() {

	var username string
	fmt.Println("Usuarios disponibles:")
	for _, u := range listaUsuarios {
		fmt.Println("-", u.Username)
	}
	fmt.Print("\nIngrese el nombre de usuario: ")
	fmt.Scanln(&username)

	session, err := iniciarSesion(username)
	if err != nil {
		fmt.Println("Error al iniciar sesión:", err)
		return
	}
	defer cerrarSesion(session.User.Username)

	if verificarRol(session.User.Username, Admin) {
		fmt.Println("El usuario es un administrador")
		fmt.Println("Generando informe como administrador...")
		generarInforme()
	} else {
		fmt.Println("El usuario no es un administrador")
	}

	fmt.Println("\nLista de tareas:")
	tareas := obtenerTareas()
	for _, tarea := range tareas {
		fmt.Printf("ID: %d - Descripción: %s - Estado: %s\n", tarea.ID, tarea.Descripcion, tarea.Estado)
	}

	crearArchivoTexto(session.User.Username)

	fmt.Println("\nComunicación en tiempo real:")
	comunicacionTiempoReal()

	fmt.Println("\nPresione Enter para salir...")
	fmt.Scanln()
}
