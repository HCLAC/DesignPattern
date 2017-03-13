/*
  Iterator 迭代器模式：
        提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示
 个人想法：
 作者：   HCLAC
 日期：   20170310
*/
package iterator

//"fmt"

type Book struct {
	name string
}

type Iterator interface {
	first() interface{}
	next() interface{}
}

type BookGroup struct {
	books []Book
}

func (b *BookGroup) add(newb Book) {
	if b == nil {
		return
	}
	b.books = append(b.books, newb)
}

func (b *BookGroup) createIterator() *BookIterator {
	if b == nil {
		return nil
	}
	return &BookIterator{b, 0}
}

type BookIterator struct {
	g     *BookGroup
	index int
}

func (b *BookIterator) first() interface{} {
	if b == nil {
		return nil
	}
	if len(b.g.books) > 0 {
		b.index = 0
		return b.g.books[b.index]
	}
	return nil
}

func (b *BookIterator) next() interface{} {
	if b == nil {
		return nil
	}
	if len(b.g.books) > b.index+1 {
		b.index++
		return b.g.books[b.index]
	}
	return nil
}
