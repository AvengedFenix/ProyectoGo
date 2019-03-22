package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Book struct {
	//..
	Isbn   int
	Name   string
	Editor string
	Copies int
}

func (b Book) restar() int {
	b.Copies--
	return b.Copies
}

func landing(w http.ResponseWriter, r *http.Request) {
	landing := "Hola ingeniero"
	t, _ := template.ParseFiles("index.html")
	fmt.Println(t.Execute(w, landing))
}

func showCarrito(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("Carrito.html")
	fmt.Println(t.Execute(w, Carrito))
}

var book1 = Book{8478290451, "El Lenguaje de Programación Java", "Pearson Education", 50}
var book2 = Book{9688802050, "El Lenguaje de Programación C", "Prentice Hall", 40}
var book3 = Book{9702602548, "Cómo Programar en C++", "Pearson Education", 45}
var book4 = Book{9688800767, "Lenguajes de Programación, Diseño e Implementación", "Prentice Hall", 60}
var book5 = Book{9701047567, "Lenguajes de Programación: Principios y Paradigmas", "McGraw-Hill", 40}

func main() {
	//http.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/index/", landing)
	http.HandleFunc("/carrito/", showCarrito)
	http.HandleFunc("/callme", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You called me!")
	})
	http.HandleFunc("/addJava", handleJava)
	http.HandleFunc("/addC", handleC)
	http.HandleFunc("/addCpp", handleCpp)
	http.HandleFunc("/addDiseno", handleDiseno)
	http.HandleFunc("/addPrincipios", handlePrincipios)

	http.ListenAndServe(":8000", nil)

	/*book1 := Book{Isbn: 8478290451, Name: "El Lenguaje de Programación Java", Editor: "Pearson Education", Copies: 50}
	book2 := Book{9688802050, "El Lenguaje de Programación C", "Prentice Hall", 40}
	book3 := Book{9702602548, "Cómo Programar en C++", "Pearson Education", 45}
	book5 := Book{9701047567, "Lenguajes de Programación: Principios y Paradigmas", "McGraw-Hill", 40}
	book4 := Book{9688800767, "Lenguajes de Programación, Diseño e Implementación", "Prentice Hall", 60}*/

	var books [5]Book
	books[0] = book1
	books[1] = book2
	books[2] = book3
	books[3] = book4
	books[4] = book5

	fmt.Println("Listado de libros: \n")
	printBooks(books)

	books = shopBook(books)

	fmt.Println("Listado de libros: \n")
	printBooks(books)
}

var myB []Book

func AddBook(b Book) {
	if b.Copies > 0 {
		b.Copies--
		myB = append(myB, b)
	} else {
		fmt.Println("No hay copias")
	}
}

var Carrito []Book
var subBookJava = Book{8478290451, "El Lenguaje de Programación Java", "Pearson Education", 1}
var subBookC = Book{9688802050, "El Lenguaje de Programación C", "Prentice Hall", 1}
var subBookCpp = Book{9702602548, "Cómo Programar en C++", "Pearson Education", 1}
var subBookDiseno = Book{9688800767, "Lenguajes de Programación, Diseño e Implementación", "Prentice Hall", 1}
var subBookPrincipios = Book{9701047567, "Lenguajes de Programación: Principios y Paradigmas", "McGraw-Hill", 1}

func handleJava(w http.ResponseWriter, r *http.Request) {
	//myBook := &subBook
	myBool := false
	for i := 0; i < len(Carrito); i++ {
		if Carrito[i].Isbn == subBookJava.Isbn {
			myBool = true

			//myBook.Copies++
			if book1.Copies-1 <= 0 {
				fmt.Println("No hay copias restantes para este libro")
			} else {
				Carrito[i].Copies++
				book1.Copies--
				break

			}
		}
	}
	if !myBool {
		Carrito = append(Carrito, subBookJava)
	}
	fmt.Println("Copias actualizadas: ", subBookJava.Copies)
}

func handleC(w http.ResponseWriter, r *http.Request) {
	//myBook := &subBook
	myBool := false
	for i := 0; i < len(Carrito); i++ {
		if Carrito[i].Isbn == subBookC.Isbn {
			myBool = true

			//myBook.Copies++
			if book2.Copies-1 <= 0 {
				fmt.Println("No hay copias restantes para este libro")
			} else {
				Carrito[i].Copies++
				break
			}
		}
	}
	if !myBool {
		Carrito = append(Carrito, subBookC)
	}
	fmt.Println("Copias actualizadas: ", subBookC.Copies)
}

func handleCpp(w http.ResponseWriter, r *http.Request) {
	//myBook := &subBook
	myBool := false
	for i := 0; i < len(Carrito); i++ {
		if Carrito[i].Isbn == subBookCpp.Isbn {
			myBool = true

			//myBook.Copies++
			if book3.Copies-1 <= 0 {
				fmt.Println("No hay copias restantes para este libro")
			} else {
				Carrito[i].Copies++
				break
			}
		}
	}
	if !myBool {
		Carrito = append(Carrito, subBookCpp)
	}
	fmt.Println("Copias actualizadas: ", subBookCpp.Copies)
}

func handleDiseno(w http.ResponseWriter, r *http.Request) {
	//myBook := &subBook
	myBool := false
	for i := 0; i < len(Carrito); i++ {
		if Carrito[i].Isbn == subBookDiseno.Isbn {
			myBool = true

			//myBook.Copies++
			if book4.Copies-1 <= 0 {
				fmt.Println("No hay copias restantes para este libro")
			} else {
				Carrito[i].Copies++
				break
			}
		}
	}
	if !myBool {
		Carrito = append(Carrito, subBookDiseno)
	}
	fmt.Println("Copias actualizadas: ", subBookDiseno.Copies)
}

func handlePrincipios(w http.ResponseWriter, r *http.Request) {
	//myBook := &subBook
	myBool := false
	for i := 0; i < len(Carrito); i++ {
		if Carrito[i].Isbn == subBookPrincipios.Isbn {
			myBool = true

			//myBook.Copies++
			if book5.Copies-1 <= 0 {
				fmt.Println("No hay copias restantes para este libro")
			} else {
				Carrito[i].Copies++
				break
			}
		}
	}
	if !myBool {
		Carrito = append(Carrito, subBookPrincipios)
	}
	fmt.Println("Copias actualizadas: ", subBookPrincipios.Copies)
}

func shopBook(books [5]Book) [5]Book {
	Isbn := 8478290451
	Copies := 5

	for i := 0; i < 5; i++ {
		book := &books[i]

		if book.Isbn == Isbn {
			if book.Copies-Copies >= 0 {
				book.Copies -= Copies
				fmt.Printf("Número de copias de %s: %d\n", book.Name, book.Copies)
			} else {
				fmt.Printf("No se puede vender esa cantidad de copias")
			}
		}
	}

	return books
}

func printBooks(books [5]Book) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Isbn: %d\nNombre: %s\nEditor: %s\nNúmero de Copias Disponibles: %d\n", books[i].Isbn, books[i].Name, books[i].Editor, books[i].Copies)
	}
	fmt.Println()
}
