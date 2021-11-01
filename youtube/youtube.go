package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	View        string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
}

func GetSubscribers() (Item, error) {

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	q := req.URL.Query()
	q.Add("key", "AIzaSyBdscTBpXiDP_Es3L-5XJGBS-wVXh1_H3g")
	q.Add("id", "UCwFl9Y49sWChrddQTD9QhRA")
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Reponse Status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}
	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		return Item{}, err
	}

	return response.Items[0], nil

}
