package main
import (
	"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
	"os"
	"io"
)

func httpGet (){
	resp, err := http.Get("http://www.jjxsw.cc/txt/list2-1.html")
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

    listReg := regexp.MustCompile("http://www.jjxsw.cc/txt/[\\w\\W]*?.html") 
    list := listReg.FindAllString(toplist,-1)

    for i := range list {
    	fliter1 := regexp.MustCompile("http://www.jjxsw.cc/txt/") 
    	fliter2 := regexp.MustCompile(".html") 
   		list1 := fliter1.ReplaceAllString(list[i],"")
   		list2 := fliter2.ReplaceAllString(list1,"")
        res, err := http.Get("http://www.jjxsw.cc/home/down/txt/id/"+list2)
        if err != nil {  
	        fmt.Println(err)
	    }  
	    fmt.Println(res.Body)
	    defer res.Body.Close()
	    f, err := os.Create(list2+".txt")  
	    if err != nil {  
	        fmt.Println(err)
	    }  
	    io.Copy(f, res.Body)
	    }
}

func main(){
	httpGet()
}