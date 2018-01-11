package main
import (
	"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
	"os"
	"io"
	"strings"
)

func httpGet (){
	resp, err := http.Get("http://www.jjxsw.cc/txt/list6-1.html")
	if err != nil {
        fmt.Println(err)
    }
   	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    topReg := regexp.MustCompile("<div class=\"toplist\">[\\w\\W]*?</div>") 
    toplist := topReg.FindString(string(body))

    listReg := regexp.MustCompile("http://www.jjxsw.cc/txt/[\\w\\W]*?.html\" title=\"[\\w\\W]*?\"" ) 
    list := listReg.FindAllString(toplist,-1)
     
    for i := range list {
    	index1 := strings.Index(list[i], "\"") //1
    	namefliter:= regexp.MustCompile("http://www.jjxsw.cc/txt/[\\w\\W]*?.html\" ") 
    	list_id := list[i][0:index1]
    	list_name := namefliter.ReplaceAllString(list[i],"")
    	fliter1 := regexp.MustCompile("http://www.jjxsw.cc/txt/") 
    	fliter2 := regexp.MustCompile(".html") 
   		list1 := fliter1.ReplaceAllString(list_id,"")
   		list2 := fliter2.ReplaceAllString(list1,"")
   		namefliter1 := regexp.MustCompile("title=\"") 
   		namefliter2 := regexp.MustCompile("\"") 
   		name1 := namefliter1.ReplaceAllString(list_name,"")
   		name2 := namefliter2.ReplaceAllString(name1,"")
        res, err := http.Get("http://www.jjxsw.cc/home/down/txt/id/"+list2)
        if err != nil {  
	        fmt.Println(err)
	    }  
	    defer res.Body.Close()
	    f, err := os.Create("./book/"+name2+".txt")  
	    if err != nil {  
	        fmt.Println(err)
	    }  
	    io.Copy(f, res.Body)
	}
}

func main(){
	httpGet()
}