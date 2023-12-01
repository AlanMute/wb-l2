/*
	Цепочка обязанностей - поведенческий паттерн который позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Это помогает когда необходимо что бы запрос проходил кучу проверок перед тем как выполнится. И что бы не копить эти проверки в одном
	месте паттерн предлагает разделять их на отдельные объекты в котором будет хранится ссылка на соедующий.

	Плюсы:
		- Уменьшает зависимость между клиентом и обработчиками.
		- Реализует принцип единственной обязанности.
		- Реализует принцип открытости/закрытости.
	Минусы:
		- Запрос может остаться никем не обработанным.

	Пример:
	Клитент банка хочет перевести какую то сумму себе на счет, и в зависимоти от того, сколько именно он хочет перевести
	будет выполнять этот запрос определенный узел
*/

package pattern

import "fmt"

type Bank interface {
	execute(amount int)
	setNext(bank Bank)
}

type BankEmployee struct { // первый уровень - Банковский сотрудник
	next Bank
}

func (b *BankEmployee) setNext(bank Bank) {
	b.next = bank
}

func (b *BankEmployee) execute(amount int) {
	if amount <= 1000 {
		fmt.Printf("Обработка запроса на перевод %d рублей банковским сотрудником\n", amount)
	} else if b.next != nil {
		fmt.Printf("Передача запроса на перевод %d рублей дальше\n", amount)
		b.next.execute(amount)
	}
}

type Supervisor struct { // второй уровень - Банковский руководитель
	next Bank
}

func (s *Supervisor) setNext(bank Bank) {
	s.next = bank
}

func (s *Supervisor) execute(amount int) {
	if amount <= 10000 {
		fmt.Printf("Обработка запроса на перевод %d рублей банковским руководителем\n", amount)
	} else if s.next != nil {
		fmt.Printf("Передача запроса на перевод %d рублей дальше\n", amount)
		s.next.execute(amount)
	}
}

type Manager struct { // финальный босс - Банковский менеджер
	next Bank
}

func (m *Manager) setNext(bank Bank) {
	m.next = bank
}

func (m *Manager) execute(amount int) {
	fmt.Printf("Обработка запроса на перевод %d рублей менеджером\n", amount)
}

// как бы это выглядело

// func main() {
// 	// настраиваем цепочку
// 	bankEmployee := &BankEmployee{}
// 	supervisor := &Supervisor{}
// 	manager := &Manager{}

// 	supervisor.setNext(manager)
// 	bankEmployee.setNext(supervisor)

// 	// делаем запрос
// 	bankEmployee.execute(5002) // запрос будет обработан руководителем
// }
