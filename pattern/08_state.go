package pattern

import "fmt"

/*
	Состояние — это поведенческий паттерн проектирования,
	который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта

	Все очень просто. Допустим у нас есть документ. Он может находится в трех состояних: Черновик, Модерация, Опубликован.
	В самом документе есть кнопка "опубликовать", которая, в зависимости от состояния делает какие то действия.
	Например в состоянии Черновик, он отправит документ на модерацию. В состоянии модерация он опубликует доку в сеть.
	А в состоянии Опубликован, кнопка работать не будет.
	Что бы это красиво реализовать и не городить условных операторов, можно использовать Состояние.
	Паттерн Состояние предлагает создать отдельные классы для каждого состояния, в котором может пребывать объект, а затем вынести туда поведения,
	соответствующие этим состояниям. Вместо того, чтобы хранить код всех состояний, первоначальный объект, называемый контекстом, будет содержать
	ссылку на один из объектов-состояний и делегировать ему работу, зависящую от состояния.

	Плюсы:
	- Избавляет от множества больших условных операторов машины состояний.
 	- Концентрирует в одном месте код, связанный с определённым состоянием.
 	- Упрощает код контекста.

	Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type Document struct {
	draft      State
	moderation State
	published  State

	currentState State
}

func (d *Document) setState(s State) {
	d.currentState = s
}

func newDoc() *Document {
	d := &Document{}

	d.draft = &DraftState{document: d}
	d.moderation = &ModerState{document: d}
	d.published = &ModerState{document: d}

	d.currentState = d.draft

	return d
}

type State interface {
	public()
}

type DraftState struct {
	document *Document
}

func (d *DraftState) public() {
	fmt.Println("Отправляем доку на модерацию")
	d.document.setState(d.document.moderation) // смена состояния
}

type ModerState struct {
	document *Document
}

func (m *ModerState) public() {
	fmt.Println("Отправляем доку в сеть")
	m.document.setState(m.document.moderation) // смена состояния
}

type PublishState struct {
	document *Document
}

func (p *PublishState) public() {
	fmt.Println("Документ уже отправлен")
}

// func main() {
// 	d := newDoc()

// 	d.currentState.public()
// 	d.currentState.public()
// 	d.currentState.public()
// }
