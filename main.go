package main

import(
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
)


func main(){

	var r RSS
	data:=readGoogleTrends()

	err:=xml.Unmarshal(data,&r)	// This will unmarshal data and structure it 

	if err!=nil{
		fmt.Println("Unable to unmarshal the xml Data")
		os.Exit(1)
	}

	// Displaying The data 

	fmt.Println("\n Google Trend 2022")
	fmt.Println("---------------------------------------------------------")

	for i:=range r.Channel.Title{
		
		rank:= (i+1)
		fmt.Println("#",rank)
		fmt.Println("Searched Term",r.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend",r.Channel.ItemList[i].Link)
		fmt.Println("Headline :",r.Channel.ItemList[i].NewsItems[i].Headline)
		fmt.Println("Link to Article",r.Channel.ItemList[i].NewsItems[i].HeadlineLink)

		fmt.Println("---------------------------------------------------------")
	}


	fmt.Println("---------------------------END----------------------------------")


}

func readGoogleTrends() []byte{
	res:=getgoogletrends()
	
	data,err:=ioutil.ReadAll(res.Body)		// For all Http request We should read the data 

	if err!=nil{
		fmt.Println("Unable to read the XML Data")
		os.Exit(1)
	}

	return data

}

// This Function is going to all the Url 
func getgoogletrends()	*http.Response{

	res,err:= http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=IN")
	if err!=nil{
		fmt.Println("An Error Occured while calling API")
		os.Exit(1)
	}
	return res
}


// Structs 

type Item struct{

	Title		string			`xml:"title"`
	Link		string			`xml:"link"`
	Traffic		string			`xml:"approx_traffic"`
	NewsItems	[]News			`xml:"news_item"`
}

type News struct{
	Headline 		string		`xml:"news_item_title"`
	HeadlineLink	string		`xml:"news_item_url"`
}

type RSS struct{
	XMLName			xml.Name 	`xml:"rss"`
	Channel			*Channel	`xml:"channel"`
}

type Channel struct{
	Title 			string		`xml:"title"`
	ItemList		[]Item		`xml:"item"`
}