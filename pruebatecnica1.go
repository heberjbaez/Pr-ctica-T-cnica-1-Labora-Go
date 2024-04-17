package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

type User struct {
	Name       string
	Password   string
	Role       string
	IsLoggedIn bool
}

type Session struct {
	ID        string
	User      User
	ExpiresAt time.Time
}

func authenticateUser() (User, error) {

	var username, password string
	fmt.Print("Ingrese su nombre de usuario: ")
	fmt.Scanln(&username)
	fmt.Print("Ingrese su contraseña: ")
	fmt.Scanln(&password)

	if isValidCredentials {
		sessionID := generateSessionID()
		session := Session{
			ID:        sessionID,
			User:      User{Name: username, Password: password, Role: "admin"}, // Asignar rol de administrador
			ExpiresAt: time.Now().Add(time.Minute * 10),                        // Sesión válida por 10 minutos
		}
		return session, nil
	}

	return User{}, fmt.Errorf("Credenciales inválidas")
}

func generateSessionID() string {

	buffer := make([]byte, 16)

	_, err := rand.Read(buffer)
	if err != nil {
		fmt.Println("Error al generar identificador de sesión:", err)
		return ""
	}

	sessionID := hex.EncodeToString(buffer)

	return sessionID
}

func createFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	robotgo.StartCapture()
	defer robotgo.EndCapture()

	fmt.Print("Ingrese el texto para el archivo: ")
	var text string
	fmt.Scanln(&text)

	for _, char := range text {
		robotgo.KeyStroke(string(char))
		file.Write([]byte(string(char)))
	}

	events := robotgo.GetEvents()

	for _, event := range events {

		fmt.Println("Pulsación:", event.Type, event.Key, event.X, event.Y)
	}

}

func main() {

	session, err := authenticateUser()
	if err != nil {
		fmt.Println("Error de autenticación:", err)
		os.Exit(1)
	}

	for {

		fmt.Println("\nMenú:")
		if session.User.Role == "admin" {
			fmt.Println("1. Crear archivo de texto")
			fmt.Println("2. Generar informe")
			fmt.Println("3. Salir")
		} else {
			fmt.Println("1. Visualizar archivos")
			fmt.Println("2. Gestionar tareas")
			fmt.Println("3. Comunicarse con otros usuarios")
			fmt.Println("4. Salir")
		}

		var option int
		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			if session.User.Role == "admin" {
				var fileName string
				fmt.Print("Ingrese el nombre del archivo: ")
				fmt.Scanln(&fileName)
				createFile(fileName)
			} else {
				fmt.Println("Opción no disponible para este rol")
			}
		case 2:
			if session.User.Role == "admin" {

			} else {
				fmt.Println("Opción no disponible para este rol")
			}
		case 3:
			fmt.Println("Saliendo...")
			break
		default:
			fmt.Println("Opción inválida")
		}

		if time.Now().After(session.ExpiresAt) {
			fmt.Println("Sesión expirada. Inicie sesión nuevamente")
			break
		}
	}
}
