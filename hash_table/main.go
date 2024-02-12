package main

import (
	"errors"
	"fmt"
	"hash/fnv"
)

type data struct{
	key string
	value int
}

type hashTable struct{
	bucket [][]data
	size int
}

func (ht *hashTable)createHash(key string) int{
hash:=fnv.New32()
hash.Write([]byte(key))
hashCode:=hash.Sum32()
return int(hashCode)%ht.size
}

func (ht *hashTable)insert(key string,value int){
	index:=ht.createHash(key)
	fmt.Println(index)
	ht.bucket[index]=append(ht.bucket[index], data{key: key,value: value})
}

func (ht *hashTable)get(key string) (int,error){
index:=ht.createHash(key)
for _,entry:=range ht.bucket[index]{
	if entry.key==key{
		return entry.value,nil
	}
}
	return 0,errors.New("key not found")
}


func main(){
	hash:=hashTable{bucket: make([][]data, 6),size: 6}
	hash.insert("hello",12121)	
	hash.insert("hellpa",7878787)	
	// hash.insert("helloo",123331)	
	
	ret,err:=hash.get("hello")

	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(ret)
}