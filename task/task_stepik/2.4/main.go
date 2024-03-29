package main

/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

Закрывающая фигурная скобка в конце main() вам не видна, но она есть.

Пакет main объявлять не нужно!
*/

type First struct {
	On    bool
	Ammo  int
	Power int
}

func (s *First) Shoot() bool {
	if s.On == false {
		return false
	}
	s.Ammo--
	return true

}
func (r *First) RideBike() bool {
	if r.On == false {
		return false
	}
	r.Power--
	return true
}

func main() {

	// testStruct :=
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */
}
