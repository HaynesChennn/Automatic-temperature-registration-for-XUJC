package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

var User_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36"
var cookie string

var u_name, u_no, u_time string

type Xujc struct {
	Data struct {
		ID    string `json:"id"`
		Owner struct {
			UserID                                 string      `json:"userId"`
			UserNo                                 string      `json:"userNo"`
			Account                                interface{} `json:"account"`
			Name                                   string      `json:"name"`
			UserType                               interface{} `json:"userType"`
			OrganizationID                         interface{} `json:"organizationId"`
			Grade                                  string      `json:"grade"`
			StudentType                            interface{} `json:"studentType"`
			CollegeCode                            string      `json:"collegeCode"`
			CollegeName                            interface{} `json:"collegeName"`
			EducationLevel                         string      `json:"educationLevel"`
			GraduationDestination                  interface{} `json:"graduationDestination"`
			GraduationDestinationName              interface{} `json:"graduationDestinationName"`
			UnitProperties                         interface{} `json:"unitProperties"`
			UnitPropertiesName                     interface{} `json:"unitPropertiesName"`
			UniTrade                               interface{} `json:"uniTrade"`
			UniTradeName                           interface{} `json:"uniTradeName"`
			UnitLocation                           interface{} `json:"unitLocation"`
			UnitLocationName                       interface{} `json:"unitLocationName"`
			JobCategory                            interface{} `json:"jobCategory"`
			JobCategoryName                        interface{} `json:"jobCategoryName"`
			RegistrationCertificateIssuingType     interface{} `json:"registrationCertificateIssuingType"`
			RegistrationCertificateIssuingTypeName interface{} `json:"registrationCertificateIssuingTypeName"`
			SignatureUnitLocation                  interface{} `json:"signatureUnitLocation"`
			SignatureUnitLocationName              interface{} `json:"signatureUnitLocationName"`
			Major                                  interface{} `json:"major"`
			MajorName                              interface{} `json:"majorName"`
			Sex                                    interface{} `json:"sex"`
			SexName                                interface{} `json:"sexName"`
			Nation                                 interface{} `json:"nation"`
			NationName                             interface{} `json:"nationName"`
			PoliticalStatus                        interface{} `json:"politicalStatus"`
			PoliticalStatusName                    interface{} `json:"politicalStatusName"`
			CultivateMode                          interface{} `json:"cultivateMode"`
			CultivateModeName                      interface{} `json:"cultivateModeName"`
			OriginPlace                            interface{} `json:"originPlace"`
			OriginPlaceName                        interface{} `json:"originPlaceName"`
			NormalSchoolStudentsType               interface{} `json:"normalSchoolStudentsType"`
			NormalSchoolStudentsTypeName           interface{} `json:"normalSchoolStudentsTypeName"`
			DifficultIdentificationLevel           interface{} `json:"difficultIdentificationLevel"`
			DifficultIdentificationLevelName       interface{} `json:"difficultIdentificationLevelName"`
			DepartmentID                           interface{} `json:"departmentId"`
		} `json:"owner"`
		FormData []struct {
			Name  string `json:"name"`
			Title string `json:"title"`
			Value struct {
				DataType        string        `json:"dataType"`
				DateFormat      interface{}   `json:"dateFormat"`
				IntegerValue    interface{}   `json:"integerValue"`
				FloatValue      interface{}   `json:"floatValue"`
				DateValue       interface{}   `json:"dateValue"`
				StringValue     string        `json:"stringValue"`
				ListValue       []interface{} `json:"listValue"`
				TableValue      []interface{} `json:"tableValue"`
				AttachmentValue []interface{} `json:"attachmentValue"`
				AddressValue    struct {
					Province  interface{} `json:"province"`
					City      interface{} `json:"city"`
					Area      interface{} `json:"area"`
					FullValue interface{} `json:"fullValue"`
				} `json:"addressValue"`
				LocationValue struct {
					Lat      interface{} `json:"lat"`
					Lng      interface{} `json:"lng"`
					Nation   interface{} `json:"nation"`
					Province interface{} `json:"province"`
					City     interface{} `json:"city"`
					Addr     interface{} `json:"addr"`
					Accuracy interface{} `json:"accuracy"`
					Type     interface{} `json:"type"`
					InRange  interface{} `json:"inRange"`
				} `json:"locationValue"`
				NameCardValue struct {
					Rows []interface{} `json:"rows"`
				} `json:"nameCardValue"`
				ChangeValue                string `json:"changeValue"`
				DateFormatWithDefaultValue string `json:"dateFormatWithDefaultValue"`
			} `json:"value"`
			Hide     bool `json:"hide"`
			Readonly bool `json:"readonly"`
		} `json:"formData"`
		Variable struct {
			AppID           int         `json:"appId"`
			BusinessID      int         `json:"businessId"`
			ApplicationType interface{} `json:"applicationType"`
		} `json:"variable"`
		UserRelationShip struct {
			InstructorID         int64       `json:"instructorId"`
			ClassID              int         `json:"classId"`
			OrganizationID       interface{} `json:"organizationId"`
			Tutor                interface{} `json:"tutor"`
			VolunteerUserID      interface{} `json:"volunteerUserId"`
			PartTimeInstructorID interface{} `json:"partTimeInstructorId"`
			DepartmentID         interface{} `json:"departmentId"`
			BuildingID           interface{} `json:"buildingId"`
		} `json:"userRelationShip"`
		Editable   bool   `json:"editable"`
		CreateTime string `json:"createTime"`
		UpdateTime string `json:"updateTime"`
		IsDelete   bool   `json:"isDelete"`
	} `json:"data"`
	State bool `json:"state"`
}

type bid struct {
	Data []struct {
		AppSettingType string `json:"appSettingType"`
		AppName        string `json:"appName"`
		Business       struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Year             string `json:"year"`
			LinkURL          string `json:"linkUrl"`
			Guide            string `json:"guide"`
			SettingID        string `json:"settingId"`
			LimitType        string `json:"limitType"`
			BeginTime        string `json:"beginTime"`
			EndTime          string `json:"endTime"`
			IsEnable         bool   `json:"isEnable"`
			IsPublic         bool   `json:"isPublic"`
			IsPublishing     bool   `json:"isPublishing"`
			PublishCode      string `json:"publishCode"`
			PublishTime      string `json:"publishTime"`
			BusinessTimeList []struct {
				ID        int         `json:"id"`
				NodeID    string      `json:"nodeId"`
				StartDate interface{} `json:"startDate"`
				EndDate   string      `json:"endDate"`
			} `json:"businessTimeList"`
			BusinessLimit interface{} `json:"businessLimit"`
			Key           interface{} `json:"key"`
			ReadyPublish  bool        `json:"readyPublish"`
		} `json:"business"`
		CanApplyOnAuthority bool          `json:"canApplyOnAuthority"`
		CanApplyOnTime      bool          `json:"canApplyOnTime"`
		CarbonCopyAuthority interface{}   `json:"carbonCopyAuthority"`
		ProxyAuthority      interface{}   `json:"proxyAuthority"`
		IsGuideShow         bool          `json:"isGuideShow"`
		MultiForm           bool          `json:"multiForm"`
		NodeList            interface{}   `json:"nodeList"`
		Players             []interface{} `json:"players"`
		IsImmediately       interface{}   `json:"isImmediately"`
		AppQuotaList        interface{}   `json:"appQuotaList"`
	} `json:"data"`
	State bool `json:"state"`
}

type xujc_response struct {
	Data  string `json:"data"`
	State bool   `json:"state"`
}

func getCookie() string {
	f, err := os.Open("cookie.txt")
	if err != nil {
		log.Printf("cookie open failed : %#v", err)
		return ""
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("cookie read failed :%#v", err)
		return ""
	}
	return string(fd)
}

func getformdata() string {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Printf("formdata open failed : %#v", err)
		return ""
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("formdata read failed :%#v", err)
		return ""
	}
	return string(fd)
}

func getbussinessID() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/api/app/229/business/now?vpn-12-o1-ijg.xujc.com&getFirst=true", nil)
	if err != nil {
		send_email("1GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/app/229")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102", "Microsoft Edge";v="102"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("User-Agent", User_Agent)
	resp, err := client.Do(req)
	if err != nil {
		send_email("2GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		send_email("3GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", resp.StatusCode)
	//fmt.Printf("%s\n", bodyText)
	var getb bid
	err = json.Unmarshal(bodyText, &getb) //反序列化
	if err != nil {
		send_email("4GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	bid := getb.Data[0].Business.ID
	//println(getb.Data[0].Business.ID)
	return strconv.Itoa(bid)
}

func getID() string {
	client := &http.Client{}
	fmt.Println("businessID:", getbussinessID())
	url := "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/api/formEngine/business/" + getbussinessID() + "/myFormInstance"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		send_email("GetID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("If-None-Match", `W/"4138-V3qtR8Z1sFVcaru9at8wciTNIFM"`)
	req.Header.Set("Referer", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/app/229")
	req.Header.Set("User-Agent", User_Agent)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		send_email("GetID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		send_email("GetID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var get Xujc
	err = json.Unmarshal(bodyText, &get) //反序列化
	if err != nil {
		send_email("GetID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	fmt.Println("No:", get.Data.Owner.UserNo)
	fmt.Println("Name:", get.Data.Owner.Name)
	u_name = get.Data.Owner.Name
	u_no = get.Data.Owner.UserNo
	return get.Data.ID
}

func isChange() bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/api/formEngine/business/1574/table/fields?playerId=owner", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", getCookie())
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/app/229")
	req.Header.Set("User-Agent", User_Agent)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if len(bodyText) == 8425 {
		return true
	} else {
		return false
	}
	//fmt.Println("change:", len(bodyText))
	//fmt.Printf("%s\n", bodyText)
}

func send_email(mailtitle string, mailbody string) {
	//简单设置log参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	//设置sender发送方的邮箱，此处可以填写自己的邮箱
	em.From = "ian-chen0713@qq.com"

	//设置receiver接收方的邮箱此处也可以填写自己的邮箱，就是自己发邮件给自己
	em.To = []string{"hayneschen@163.com"}

	//设置主题
	em.Subject = mailtitle

	//简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(mailbody)

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "ian-chen0713@qq.com", "cwgnwjtzwcjmbcbb", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sendsuccessfully...")
}

func pause() {
	fmt.Printf("按任意键退出...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func main() {
	cookie = getCookie()
	fmt.Println("Automatic temperature registration by Haynes v1.5")
	t := time.Now() //当前时间
	timeLayoutStr := "2006-01-02 15:04:05"
	u_time = t.Format(timeLayoutStr)
	id := getID()
	client := &http.Client{}
	var data = strings.NewReader(getformdata())
	url := "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/api/formEngine/formInstance/" + id + "?vpn-12-o1-ijg.xujc.com"
	println("url:", url)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		send_email("Post出现了一些问题..... "+u_time, err.Error())
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0. 8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", User_Agent)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://webvpn.xmu.edu.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://webvpn.xmu.edu.cn/http/77726476706e69737468656265737421f9fd46d23f256253300b86a1/app/229")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		send_email("Post出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		send_email("Post出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var xujc xujc_response
	err = json.Unmarshal(bodyText, &xujc) //反序列化
	if err != nil {
		send_email("Post出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", dictResponse)

	var st, isch string
	//var isch string
	if xujc.State == true {
		st = "Success"
	} else {
		st = "Fail"
	}
	mbody := "Temperature Sign In Details" + "\n" +
		"URL:" + url + "\n" +
		"NO:" + u_no + "\n" +
		"Name:" + u_name + "\n" +
		"State:" + st + "\n" +
		"isChange:"
	if isChange() {
		isch = "False"
	} else {
		isch = "True"
		fmt.Println("!_! 表单数据改变了,快来看看")
		send_email("!_! 表单数据改变了,快来看看", mbody)
	}
	fmt.Println("Status:", st)
	fmt.Println("isChange:", isch)
	fmt.Println("Time:", u_time)
	mbody += isch + "\n" + "UpdateTime:" + u_time
	if xujc.State != true {
		send_email("QAQ 体温打卡失败啦,快来检查一下  "+u_time, mbody)
	} else {
		send_email("0.0 体温打卡成功!   "+u_time, mbody)
	}
	pause()
}
