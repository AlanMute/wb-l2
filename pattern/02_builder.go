/*
Паттерн Builder является порождающим. Он предлагает разбить процесс конструирования объекта на отдельные шаги.
Это нужно для того, что бы не перегружать конструктор большим количеством параметров.

Причем у нас может быть несколько строителей, которые делают одни и те же действия по разному.
Например постройка дома. Дом может быть из кирпича, тогда все его компоненты будут из кирпича. Это один билдер.
А второй билдер будет делать дом из дерева.

Можно пойти дальше и создать класс деректора, который будет задавать порядок строитесльства а строитель - выполнять их

Плюсы:
	- Позволяет создавать сложные объекты, не загромождая основной код множеством конструкторов.
	- Паттерн позволяет создавать различные варианты объектов с помощью различных строителей, контролируя каждый шаг создания.
	- Позволяет использовать понятные имена методов для каждого шага конструирования, что упрощает чтение и понимание кода.

Минусы:
	- В случае простых объектов использование паттерна может показаться излишним и усложнить структуру кода.
 	- При создании паттерна "Строитель" для каждого типа объекта требуется отдельный класс строителя,
	  что может привести к увеличению количества классов в программе.

В качестве примера создадим builder для компьютеров. У нас будут два строителя. Один из них собирает игровой пк, а другой офисный.
*/

package pattern

type Computer struct {
	Monitor   string
	Mouse     string
	Keyboard  string
	Videocard string
}

type IBuilder interface { // интерфейс строителя
	setMonitor()
	setMouse()
	setKeyboard()
	getComputer() *Computer
}

func getBuilder(typePC string) IBuilder {
	if typePC == "base" {
		return newBasic()
	}

	if typePC == "top" {
		return newTop()
	}
	return nil
}

type BasicBuilder struct { // строитель для базовой комплектации
	c *Computer
}

func newBasic() *BasicBuilder {
	return &BasicBuilder{}
}

func (b *BasicBuilder) setKeyboard() {
	b.c.Keyboard = "Base Keyboard"
}

func (b *BasicBuilder) setMouse() {
	b.c.Mouse = "Base Mouse"
}

func (b *BasicBuilder) setMonitor() {
	b.c.Monitor = "Base Monitor"
}

func (b *BasicBuilder) getComputer() *Computer {
	return b.c
}

type TopBuilder struct { // топовой для топовой комплектации
	c *Computer
}

func newTop() *TopBuilder {
	return &TopBuilder{}
}

func (b *TopBuilder) setKeyboard() {
	b.c.Keyboard = "Top Keyboard"
}

func (b *TopBuilder) setMouse() {
	b.c.Mouse = "Top Mouse"
}

func (b *TopBuilder) setMonitor() {
	b.c.Monitor = "Top Monitor"
}

func (b *TopBuilder) getComputer() *Computer {
	return b.c
}

// так же можно создать директора

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) createPC() *Computer { // директор говорит строителю что и в каком порядке строить и тот строит
	d.builder.setKeyboard()
	d.builder.setMonitor()
	d.builder.setMouse()
	return d.builder.getComputer()
}
