package main

// // 定义一个结构体
// type Book struct {
// 	title  string
// 	author string
// }
//
// func change(book Book) {
// 	// 传递的是一个副本
// 	book.author = "lisi"
// }
//
// func changePointer(book *Book) {
// 	// 传递的是一个副本
// 	book.author = "wangwu"
// }
//
// // 如果类名首字母大写，表示其他包也可以访问
// type Hero struct {
// 	// 属性名大写，对外也是可以访问的，否则只能类的内部访问
// 	name  string
// 	ad    int
// 	level int
// }
//
// /*func (this Hero) show() {
// 	fmt.Println("name = ", this.name)
// 	fmt.Println("ad = ", this.ad)
// 	fmt.Println("level = ", this.level)
// }
//
// func (this Hero) getName() string {
// 	fmt.Println("name = ", this.name)
// 	return this.name
// }
//
// func (this Hero) setName(name string) {
// 	// 地址的拷贝
// 	this.name = name
// }*/
//
// func (this *Hero) show() {
// 	fmt.Println("name = ", this.name)
// 	fmt.Println("ad = ", this.ad)
// 	fmt.Println("level = ", this.level)
// }
//
// func (this *Hero) getName() string {
// 	fmt.Println("name = ", this.name)
// 	return this.name
// }
//
// func (this *Hero) setName(name string) {
// 	this.name = name
// }
//
// func main() {
// 	var book Book
// 	book.title = "Golang"
// 	book.author = "zhangsan"
// 	fmt.Printf("%v\n", book)
// 	change(book)
// 	fmt.Printf("%v\n", book)
// 	changePointer(&book)
// 	fmt.Printf("%v\n", book)
// 	// 创建一个对象
// 	hero := Hero{name: "SpiderMan", ad: 2, level: 1}
// 	hero.show()
// 	hero.setName("BatMan")
// 	hero.show()
// }
