package main

import "fmt"

//二叉树
//前序中序后序遍历
//根左右、左根右、左右根

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历[先输root结点，然后再输出左子树，然后再输出右子树]
func PreOrder(hero *Hero) {
	if hero != nil {
		fmt.Printf("no=%d name=%s\n", hero.No, hero.Name)
		PreOrder(hero.Left)
		PreOrder(hero.Right)
	}
}

// 中序遍历[先输出root的左子树，再输root结点，最后输出root的右子树]
func InfixOrder(hero *Hero) {
	if hero != nil {
		InfixOrder(hero.Left)
		fmt.Printf("no=%d name=%s\n", hero.No, hero.Name)
		InfixOrder(hero.Right)
	}
}

// 后续遍历
func PostOrder(hero *Hero) {
	if hero != nil {
		PostOrder(hero.Left)
		PostOrder(hero.Right)
		fmt.Printf("no=%d name=%s\n", hero.No, hero.Name)
	}
}
func main() {
	//构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2
	left2 := &Hero{
		No:   5,
		Name: "鲁智深",
	}
	right3 := &Hero{
		No:   6,
		Name: "杨志",
	}
	left1.Left = left2
	left1.Right = right3
	fmt.Println("前序遍历")
	PreOrder(root)
	fmt.Println("中序遍历")
	InfixOrder(root)
	fmt.Println("后序遍历")
	PostOrder(root)
}
