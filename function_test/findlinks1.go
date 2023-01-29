package main

import (
	"bufio"
	"fmt"
	fetchtest "go_test_init/fetch_test"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("函数")
	fmt.Println("基础定义", f1(1, 2))
	// testRecursion()
	// testPanic()
	// testValue()
	testAnonymous()
}
func f1(a int, b int) int {
	return a + b
}

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32
type Attribute struct {
	Key, Val string
}

// 枚举
const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

// func Parse(r io.Reader) (*Node, error)

func testRecursion() {
	urls := []string{"https://juejin.cn"}
	docStr := fetchtest.Fetch(urls)
	doc, err := html.Parse(strings.NewReader(docStr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks1:%v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	outline(nil, doc)
}
func visit(links []string, n *html.Node) []string {
	fmt.Println(" testRecursion visit:", n.Type)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	fmt.Println(" testRecursion visit res:", links)
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func testPanic() error {
	log.Printf("log Printf,不中断")
	// log.Fatal("log fatal,会中断并打印")
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read fail:%v", err)
		}
		fmt.Println("read file r", r)
	}
	return nil
}

// 函数是一等公民
func testValue() {
	var f func(int) int
	fmt.Println(" 声明函数类型", f == nil)
	urls := []string{"https://juejin.cn"}
	docStr := fetchtest.Fetch(urls)
	doc, err := html.Parse(strings.NewReader(docStr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks1:%v\n", err)
	}
	forEachNode(doc, pre, post)

}

var depth int

func pre(node *html.Node) {
	// fmt.Println("pre")
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}
func post(node *html.Node) {
	// fmt.Println("post")
	if node.Type == html.ElementNode {
		depth++
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
	}

}

// 遍历树节点
func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if pre != nil {
		pre(node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(node)
	}
}

func testAnonymous() {
	fmt.Println("命名函数只能在包级中声明，而不能在函数内声明")
	f := func(s string, others ...string) {
		fmt.Println("匿名函数", s)
	}
	f("anonymous")
	fmt.Println("定义函数时，在参数类型前加省略号，表示可变参数")
	fmt.Println("调用函数时，在参数后加省略号，表示对切片拆解")
}
