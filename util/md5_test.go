package util

import "testing"

func TestHash(t *testing.T) {

	str := `0123456789abcdefghijklmnopqrstuvwxyz`
	v1 := Hash(str)
	t.Logf("str:%s hash:%v", str, v1)

	v2 := Hash(str)

	if v1 != v2 {
		t.Errorf("expect:%v actual:%v", v1, v2)
		t.Fail()
	}

	str = `年底项目任务中,忙里偷闲更新一篇.
众所周知,PHP拥有一个万能的数据结构--Array.这货实在是太易用太强大了,导致PHP程序员在程序员界都混不下去了.使用PHP的Array几乎可以实现在大学数据结构课本里出现的所有数据结构.
实际上在PHP内部,Array是由Zend引擎中的Zend HashTable实现的.其实不光是Array,包括各种常量,变量,函数体,class等都是用它来组织的.
在PHP的源代码里找到zend_hash.c后不难发现,zend hashtable包含两个结构体,分别是`

	str2 := `年底项目任务中,忙里偷闲更新一篇.
众所周知,PHP拥有一个万能的数据结构--Array.这货实在是太易用太强大了,导致PHP程序员在程序员界都混不下去了.使用PHP的Array几乎可以实现在大学数据结构课本里出现的所有数据结构.
实际上在PHP内部,Array是由Zend引擎中的Zend HashTable实现的.其实不光是Array,包括各种常量,变量,函数体,class等都是用它来组织的.
在PHP的源代码里找到zend_hash.c后不难发现,zend hashtable包含两个结构体,分别是`

	v1 = Hash(str)
	v2 = Hash(str2)

	if v1 != v2 {
		t.Errorf("expect:%v actual:%v", v1, v2)
		t.Fail()
	} else {
		t.Logf("test pass, hash:%v", v1)
	}
}
