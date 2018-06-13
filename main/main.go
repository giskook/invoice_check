// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	c, err := chromedp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var res string
	//err = c.Run(ctxt, check("123", "111", "1", "2", "3", "4", "5", "6", &res))
	//err = c.Run(ctxt, open_youshan())
	file, ferr := os.Open("D:\\docs\\work\\pf\\hb.csv")
	if ferr != nil {
		log.Println(ferr.Error())
	}
	defer func() {
		file.Close()
	}()
	scanner := bufio.NewScanner(file)
	if err_ := c.Run(ctxt, chromedp.Navigate(`https://ybs.he-n-tax.gov.cn:8888/fpzx-web/apps/views/fpcx/fpcx.html`)); err != nil {
		log.Println(err_.Error())
	}
	var flag string
	for scanner.Scan() {
		item := scanner.Text()
		err, res = check(ctxt, c, item)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Scanln(&flag) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
		// fmt.Scanf("%s %s", &firstName, &lastName)    //Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
		if err := c.Run(ctxt, chromedp.Click(`fplxcz`, chromedp.ByID)); err != nil {
		}

	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("overview: %s", res)
}

func check(ctxt context.Context, c *chromedp.CDP, item string) (error, string) {
	ss := strings.Split(item, ",")
	var code, number, id string
	code = ss[8]
	number = ss[9]
	id = ss[13]

	var cancel func()
	ctxt, cancel = context.WithTimeout(ctxt, 25*time.Second)
	defer cancel()
	if err := c.Run(ctxt, chromedp.WaitVisible(`#fplx`, chromedp.ByID)); err != nil {
		return err, ""
	}

	if err := c.Run(ctxt, chromedp.SendKeys(`#fpdm\24 text`, code+"\n", chromedp.ByID)); err != nil {
		return err, ""
	}

	if err := c.Run(ctxt, chromedp.SendKeys(`#fphm\24 text`, number+"\n", chromedp.ByID)); err != nil {
		return err, ""
	}
	if err := c.Run(ctxt, chromedp.SendKeys(`#sbm\24 text`, id+"\n", chromedp.ByID)); err != nil {
		return err, ""
	}
	//	if err := c.Run(ctxt, chromedp.Click(`#fplx > option:nth-child(2)`, chromedp.ByID)); err != nil {
	//		return err, ""
	//	}

	return nil, ""
}
