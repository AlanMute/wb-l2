/*
	Посетитель - это шаблон проектирования, который позволяет внедрять новые функциональности в программу,
	не внося изменения в классы объектов, для которых предназначены эти функциональности.

	Плюсы:
		- Упрощает добавление операций, работающих со сложными структурами объектов.
		- Объединяет родственные операции в одном классе.
		- Посетитель может накапливать состояние при обходе структуры элементов.

	Минусы:
		- Паттерн не оправдан, если иерархия элементов часто меняется.
		- Может привести к нарушению инкапсуляции элементов.

	В качестве примера:
	У нас есть классы предоставляющие собой разные файлы для картинок. И нужно для каждого формата добавить возможность сжать изображение,
	а так же улучшить его качества(с помощью ИИ например)
*/

package pattern

import "fmt"

type visitor interface {
	visitorForPng(*Png)
	visitorForJpg(*Jpg)
	visitorForSwg(*Swg)
}

type IPicture interface {
	GetExt() string
	accept(visitor) // Для того что бы использовать паттерн посетитель необходимо внести только одно изменение, а именно добавить метод accept()
}

type Png struct {
	weight int
}

type Jpg struct {
	weight int
}

type Swg struct {
	weight int
}

func (p *Png) GetExt() string {
	return ".png"
}

func (j *Jpg) GetExt() string {
	return ".jpg"
}

func (s *Swg) GetExt() string {
	return ".swg"
}

func (p *Png) accept(v visitor) {
	v.visitorForPng(p)
}

func (j *Jpg) accept(v visitor) {
	v.visitorForJpg(j)
}

func (s *Swg) accept(v visitor) {
	v.visitorForSwg(s)
}

// реализуем конкретного посетителя, который позволит сжимать фото всех форматов
type Zip struct {
	level int
}

func (z *Zip) visitorForPng(p *Png) {
	// процесс сжатия png
	fmt.Println("Png сжали")
}

func (z *Zip) visitorForJpg(j *Jpg) {
	// процесс сжатия jpg
	fmt.Println("Jpg сжали")
}

func (z *Zip) visitorForSwg(s *Swg) {
	// процесс сжатия svg
	fmt.Println("Swg сжали")
}

// реализуем конкретного посетителя, который позволит улучшать фото всех форматов

type Improve struct {
	level int
}

func (i *Improve) visitorForPng(p *Png) {
	// процесс улучшения png
	fmt.Println("Png улучшили. Спасибо ИИ!")
}

func (i *Improve) visitorForJpg(j *Jpg) {
	// процесс улучшения jpg
	fmt.Println("Jpg улучшили. Спасибо ИИ!")
}

func (i *Improve) visitorForSwg(s *Swg) {
	// процесс улучшения svg
	fmt.Println("Swg улучшили. Спасибо ИИ!")
}
