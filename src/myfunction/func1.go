package myfunction

import (
	"math"
	"os"
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

//func name(parameter-list)(result-list){

//}

/**
 * UNusedImport
 */

func  MyHypot ( x,y float64) float64{

	return math.Sqrt(x*x + y*y)
}

func f1(i,j,k int,s ,t string){

}

func f2(i int,j int, k int, s string, t string){

}


/**
 *实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实
 *参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的简介
 *引用被修改
 */

func  Add( x int , y int ) int{
	return x+y
}

func Sub(x,y int )(z int){

	z = x -y

	return z
}

func First( x int ,_ int) int{
	return x
}

func Zero(int,int) int{
	return 0
}

//没有函数体的函数声明，这表示该函数不是以Go实现的

func ManyResult(){

	for _,url :=range  os.Args[1:]{

	 links,error :=	findLinks(url)
	 if error != nil{
	 	fmt.Fprintf(os.Stderr,"findlinks2: %v\n",error)
	 	continue
	 }
	 for _,link := range links{
	 	fmt.Printf(link)
	 }

	}
}

func findLinks(url string)([]string,error){

	 resp ,error :=http.Get(url)

	 if error != nil{

	 	return nil,error
	 }

	 if resp.StatusCode != http.StatusOK{

	 	resp.Body.Close()
	 	return nil, fmt.Errorf("getting %s: %s",url,resp.Status)
	 }

	 doc ,error := html.Parse(resp.Body)
	 resp.Body.Close()
	 if error!= nil{
	 	return nil,fmt.Errorf("parsing %s as html: %v",url,error)
	 }
	 return visit(nil,doc),nil
}


func visit(links []string, n *html.Node) []string {

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
	return links
}

