package main

import (
	"errors"
	"fmt"
)

// PriorityQueue-representación de una estructura de datos de cola
type PriorityQueue struct {
	Alta []int
	Baja []int
}

// Enqueue - agregando un elemento de tipo int al final de nuestra cola
func (q *PriorityQueue) Enqueue(element int, altaPrioridad bool) {
	if altaPrioridad {
		q.Alta = append(q.Alta, element)
	} else {
		q.Baja = append(q.Baja, element)
	}
}

// Dequeue - devuelve el primer elemento de nuestra cola
func (q *PriorityQueue) Dequeue() (int, error) {
	// SI la longitud de la alta prioridad!=0 => devulve el primer elemento de la cola de alta prioridad
	// de lo contrario, si la longitud de la prioridad baja!=0 => devuelve el primer elemento de la cola de baja prioridad

	if len(q.Alta) != 0 {
		var primerElemt int
		primerElemt, q.Alta = q.Alta[0], q.Alta[1:]
		return primerElemt, nil
	}

	if len(q.Baja) != 0 {
		var primerElemt int
		primerElemt, q.Baja = q.Baja[0], q.Baja[1:]
		return primerElemt, nil
	}

	return 0, errors.New("Cola vacia")
}

// Length - devuelve la longitud de la cola
func (q *PriorityQueue) Length() int {
	return len(q.Alta) + len(q.Baja)
}

// Peek - devuelve el primer elemento de nuestra cola sin actualizar la cola
func (q *PriorityQueue) Peek() (int, error) {
	if len(q.Alta) != 0 {
		return q.Alta[0], nil
	}
	if len(q.Baja) != 0 {
		return q.Baja[0], nil
	}
	return 0, errors.New("Cola vacia")
}

// IsEmpty - Devuelve verdadero si la cola está vacía
func (q *PriorityQueue) IsEmpty() bool {
	return q.Length() == 0
}

func main() {
	fmt.Println("Sección de colas")

	cola := PriorityQueue{}

	fmt.Println(cola)
	cola.Enqueue(1, true)
	fmt.Println(cola)
	cola.Enqueue(10, false)
	fmt.Println(cola)

	element, _ := cola.Dequeue()
	fmt.Println(element)
	fmt.Println(cola)

	cola.Enqueue(2, true)
	fmt.Println(cola)

	element, _ = cola.Dequeue()
	fmt.Println(element)
	fmt.Println(cola)

	element, _ = cola.Dequeue()
	fmt.Println(element)
	fmt.Println(cola)

}
