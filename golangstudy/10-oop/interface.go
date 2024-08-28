package main

/*
本质是一个指针
*/
// type Animal interface {
// 	Sleep()
// 	GetColor() string
// 	GetType() string
// }
//
// /*
// 具体实现类，需要完全实现接口所有的方法
// */
// type Cat struct {
// 	Color string
// }
//
// func (cat *Cat) Sleep() {
// 	fmt.Println("Cat is sleep")
// }
//
// func (cat *Cat) GetColor() string {
// 	return cat.Color
// }
//
// func (cat *Cat) GetType() string {
// 	return "Cat"
// }
//
// /*
// 具体实现类
// */
// type Dog struct {
// 	Color string
// }
//
// func (dog *Dog) Sleep() {
// 	fmt.Println("Dog is sleep")
// }
//
// func (dog *Dog) GetColor() string {
// 	return dog.Color
// }
//
// func (dog *Dog) GetType() string {
// 	return "Dog"
// }
//
// func main() {
// 	var animal Animal
// 	animal = &Cat{"Yellow"}
// 	animal.Sleep()
// 	color := animal.GetColor()
// 	fmt.Println("Cat color is ", color)
//
// 	animal = &Dog{"Black"}
// 	animal.Sleep()
//
// }
