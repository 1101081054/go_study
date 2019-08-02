package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"os"
	"strconv"
	"time"
)

func imageList() [17]string {
	var images [17]string
	images[0] = "https://uploads.classba.cn/fengmian1.png"
	images[1] = "https://uploads.classba.cn/fengmian2.png"
	images[2] = "https://uploads.classba.cn/fengmian3.png"
	images[3] = "https://uploads.classba.cn/fengmian4.png"
	images[4] = "https://uploads.classba.cn/fengmian5.png"
	images[5] = "https://uploads.classba.cn/fengmian6.png"
	images[6] = "https://uploads.classba.cn/fengmian7.png"
	images[7] = "https://uploads.classba.cn/fengmian8.png"
	images[8] = "https://uploads.classba.cn/fengmian9.png"
	images[9] = "https://uploads.classba.cn/fengmian10.png"
	images[10] = "https://uploads.classba.cn/fengmian11.png"
	images[11] = "https://uploads.classba.cn/fengmian12.png"
	images[12] = "https://uploads.classba.cn/fengmian13.png"
	images[13] = "https://uploads.classba.cn/fengmian14.png"
	images[14] = "https://uploads.classba.cn/fengmian15.png"
	images[15] = "https://uploads.classba.cn/fengmian16.png"
	images[16] = "https://uploads.classba.cn/fengmian17.png"

	return images
}


func main() {
	var images [17]string  = imageList()

	xlsx, err := excelize.OpenFile("imports.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	time := time.Now().Unix()
	//os.Exit(1)
	fmt.Println(`<?xml version="1.0" encoding="UTF-8"?>`)

	fmt.Print("<dataset>")
	for i := 2; i <= 201; i++ {
		id, _ := xlsx.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		province, _ := xlsx.GetCellValue("Sheet1", "B"+strconv.Itoa(i))
		name, _ := xlsx.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		address, _ := xlsx.GetCellValue("Sheet1", "D"+strconv.Itoa(i))

		fmt.Print(`
	<data>
		<id>` + id + `</id>
		<name>` + name + `</name>
		<target>
			<province>
				<e>` + province + `</e>
			</province>
		</target>
		<image>
			<url>` + images[(i % 17)] + `</url>
			<width>600</width>
			<height>600</height>
		</image>
		<video>
			<url>https://uploads.classba.cn/invite_partner.mp4</url>
			<width>720</width>
			<height>1280</height>
		</video>
		<targetUrlMobile>https://www.chengzijianzhan.com/tetris/page/6711608914795823111/</targetUrlMobile>
		<address>` + address + `</address>
		<onlineTime>`, time, `</onlineTime>
		<stock>1</stock>
		<status>1</status>
	</data>
`)

	}
	fmt.Print("</dataset>")

}
