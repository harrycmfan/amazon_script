package main
import (
        "bytes"
        "fmt"
        "io/ioutil"
        "net/http"
        "strconv"
        "strings"
        "time"
)
func main() {
        defer func() {
                if r := recover(); r != nil {
                        var jsonStr = []byte(`{
                                                                        "text": {
                                                                                "content": "库存通知机器人异常了"
                                                                        },
                                                                        "msgtype": "text"
                                                                }`)
                        _, _ = http.Post("https://oapi.dingtalk.com/robot/send?access_token=your token", "application/json", bytes.NewBuffer(jsonStr))
                }
        }()
        count := 448439
        for {
                now := time.Now()
                if (now.Hour() == 15 || now.Hour() == 3) && now.Minute() == 0 && now.Second() == 0 {
                        var jsonStr = []byte(`{
                                                                        "text": {
                                                                                "content": "库存检查已经运行` + strconv.Itoa(count) + `次"
                                                                        },
                                                                        "msgtype": "text"
                                                                }`)
                        _, _ = http.Post("https://oapi.dingtalk.com/robot/send?access_token=d5d3573928d71671cca8174e90b29200fe06628b2f357c5663163c6a0127517e", "application/json", bytes.NewBuffer(jsonStr))
                        //fmt.Println(a)
                }
                resp, err := http.Get("https://www.amazon.sg/Sony-PlayStation-5-DualSense-Wireless-Controller/dp/B08HNRSVQP/ref=sr_1_1?crid=H0ICLGPS30O3&dchild=1&keywords=playstation%2B5&qid=1605700549&sprefix=plays%2Caps%2C343&s
r=8-1&th=1")
                if err != nil {
                        fmt.Println(err)
                        time.Sleep(time.Millisecond * 500)
                        continue
                }
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                        fmt.Println(err)
                        time.Sleep(time.Millisecond * 500)
                        continue
                }
                //fmt.Println(string(body), err)
                if strings.Contains(string(body), "add-to-cart-button") {
                        var jsonStr = []byte(`{
                                                                        "text": {
                                                                                "content": "光驱版有库存了,https://www.amazon.sg/Sony-PlayStation-5-DualSense-Wireless-Controller/dp/B08HNRSVQP/ref=sr_1_1?crid=H0ICLGPS30O3&dchild=1
&keywords=playstation%2B5&qid=1605700549&sprefix=plays%2Caps%2C343&sr=8-1&th=1"
                                                                        },
                                                                        "msgtype": "text"
                                                                }`)
                        _, _ = http.Post("https://oapi.dingtalk.com/robot/send?access_token=your toke", "application/json", bytes.NewBuffer(jsonStr))
                }
                
                               resp, err = http.Get("https://www.amazon.sg/Sony-PlayStation-5-DualSense-Wireless-Controller/dp/B08HNSWWT7/ref=sr_1_1?crid=H0ICLGPS30O3&dchild=1&keywords=playstation%2B5&qid=1605700549&sprefix=plays%2Caps%2C343&sr
=8-1&th=1")
                if err != nil {
                        fmt.Println(err)
                        time.Sleep(time.Millisecond * 500)
                        continue
                }
                body, err = ioutil.ReadAll(resp.Body)
                if err != nil {
                        fmt.Println(err)
                        time.Sleep(time.Millisecond * 500)
                        continue
                }
                _ = resp.Body.Close()
                //fmt.Println(string(body), err)
                if strings.Contains(string(body), "add-to-cart-button") {
                        var jsonStr = []byte(`{
                                                                        "text": {
                                                                                "content": "数字版有库存了,https://www.amazon.sg/Sony-PlayStation-5-DualSense-Wireless-Controller/dp/B08HNSWWT7/ref=sr_1_1?crid=H0ICLGPS30O3&dchild=1
&keywords=playstation%2B5&qid=1605700549&sprefix=plays%2Caps%2C343&sr=8-1&th=1"
                                                                        },
                                                                        "msgtype": "text"
                                                                }`)
                        _, _ = http.Post("https://oapi.dingtalk.com/robot/send?access_token=your token", "application/json", bytes.NewBuffer(jsonStr))
                }
                count++
                fmt.Println("已运行", count, "次")
                time.Sleep(time.Millisecond * 500)
                //break;
        }
        //resp, err := http.Get("https://www.amazon.sg/Sony-PlayStation-5-DualSense-Wireless-Controller/dp/B08HNSHD7K/ref=sr_1_1?crid=H0ICLGPS30O3&dchild=1&keywords=playstation%2B5&qid=1605700549&sprefix=plays%2Caps%2C343&sr=8-1&
th=1")
        //if err != nil {
        //      fmt.Println(err)
        //      time.Sleep(time.Millisecond * 500)
        //}
        //body, err := ioutil.ReadAll(resp.Body)
        //if err != nil {
        //      fmt.Println(err)
        //      time.Sleep(time.Millisecond * 500)
        //}
        ////fmt.Println(string(body), err)
        //if strings.Contains(string(body), "add-to-cart-button") {
        //      fmt.Println(111)
        //}
}
                                         
