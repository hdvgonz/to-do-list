package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct{
	Nombre string;
	Descripcion string;
	completado bool;
}

type ListaTareas struct{
	tareas []Tarea;
}

/*Metodo para agregar Tarea*/
func (l *ListaTareas) agregarTarea (t Tarea){

	l.tareas = append(l.tareas, t);
	fmt.Println("¡Tarea Agregada!");
}

/*Metodo para marcar como completado tarea*/

func (l *ListaTareas) marcarCompletado(index int){

	l.tareas[index].completado = true;
	fmt.Println("¡Tarea Completada!");
}

func (l *ListaTareas) editarTareas(index int, t Tarea){
	
	l.tareas[index] = t
	fmt.Println("¡Tarea Editada!");
}

//Eliminar Tarea

func (l *ListaTareas) eliminarTarea(index int){

	l.tareas =  append(l.tareas[:index], l.tareas[index+1:]... );
	fmt.Println("¡Tarea Eliminada!");
}

//Listar Tareas

func (l *ListaTareas) listarTareas(){
	fmt.Println("Lista de Tareas");
	fmt.Println("------------------------------------");

	for index, tarea := range l.tareas{

		fmt.Printf("%d - %s - %s - Completado: %t\n", index, tarea.Nombre, tarea.Descripcion, tarea.completado);
	}
	fmt.Println("------------------------------------");
}

func main() {
	
	//Instancia de la lista de tareas
	lista := ListaTareas{};

	leer := bufio.NewReader(os.Stdin);
	for {
		var opcion int;
		fmt.Println("Seleccione una opción:\n" ,
		"1. Agregar Tarea\n",
		"2. Marcar Tarea Completada\n",
		"3. Editar Tarea\n",
		"4. Eliminar Tarea\n",
		"5. Listar Tareas\n",
		"6. Salir");
		
		fmt.Println("Ingrese Opción: ");
		fmt.Scanln(&opcion);

		switch opcion{
		case 1:
			var t Tarea;
			fmt.Println("Ingrese nombre de la tarea: ");
			t.Nombre, _ = leer.ReadString('\n');
			fmt.Println("Ingrese Descripcion de la tarea");
			t.Descripcion, _ = leer.ReadString('\n');
			lista.agregarTarea(t);
		case 2:
			var index int;
			fmt.Println("Digite # de la tarea Completada");
			fmt.Scanln(&index);
			lista.marcarCompletado(index);
		
		case 3:
			var index int;
			var t Tarea;

		fmt.Println("Digite el # de la tarea a editar: ");
			fmt.Scanln(&index);

			fmt.Println("Escriba nuevo nombre de la Tarea");
			fmt.Scanln(t.Nombre);

			fmt.Println("Escriba la nueva descripción de la tarea");
			fmt.Scanln(t.Descripcion);

			lista.editarTareas(index, t);
		
		case 4:
			var index int;
			fmt.Println("Digite el # de la tarea a eliminar: ");
			fmt.Scanln(&index);

			lista.eliminarTarea(index);
		
		case 5:
			lista.listarTareas();
		case 6:
			fmt.Println("Saliendo del programa...");
			return;
			// os.Exit(0)
		
		default:
			fmt.Println("Has ingresado una opción incorrecta")
		}
	}

	
}
