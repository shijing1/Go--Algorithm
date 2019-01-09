package main

import "fmt"

//创建一个结构体 存储信息
type HeroNonde struct {
	no int
	name string
	nickName string
	pre *HeroNonde
	next *HeroNonde
}
func InsterNoder(head *HeroNonde,newHeroNode *HeroNonde){
	//创建一个 辅助节点
	temp:=head
	//插入链表
	flag:=true
	for  {
		if temp.next==nil{
			break
		}else if temp.next.no>=newHeroNode.no{
			break
		}else if temp.next.no==newHeroNode.no{
			flag=false
			break
		}
		temp=temp.next
	}
	if !flag{
		fmt.Println("你输入的no已经存在",newHeroNode)
		return
	}else {
		newHeroNode.next=temp.next
		newHeroNode.pre=temp
		if temp.next!=nil{
			temp.next.pre=newHeroNode
		}
		temp.next=newHeroNode
	}
}
//查找链表
func SearchHeroNode(head *HeroNonde){
	temp:=head
	if temp.next==nil{
		fmt.Println("链表为空" )
		return
	}
	for  {
		fmt.Printf("[%d ,%s,%s ]",temp.next.no,temp.next.name,temp.next.nickName)
		temp=temp.next
		if temp.next==nil{
			break
		}
	}

	}

//删除节点
func DeleteHeroNode(head *HeroNonde,id int){
	temp:=head
	flag:=false
	for  {
		if temp.next==nil{
				break
		}else if temp.next.no==id{
			flag=true
			break
		}
		temp=temp.next
	}
	if flag{
		temp.next=temp.next.next
		if temp.next!=nil{
			temp.next.pre=temp
		}else {
			fmt.Println("你要阿删除的节点不存在")
		}
	}
}

func main(){
	//创建一个头结点    不存放任何信息
head:=&HeroNonde{}
hero1:=&HeroNonde{
	no:1,
	name:"宋江",
	nickName:"及时雨",
}
hero2:=&HeroNonde{
	no:2,
	name:"卢俊义",
	nickName:"玉麒麟",
}
InsterNoder(head,hero1)
InsterNoder(head,hero2)
SearchHeroNode(head)
}
