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

var u_name, u_no, u_time, bussid string

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
	req, err := http.NewRequest("GET", "http://ijg.xujc.com/api/app/229/business/now?getFirst=true", nil)
	if err != nil {
		send_email("GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", getCookie())
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", User_Agent)
	resp, err := client.Do(req)
	if err != nil {
		send_email("GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		send_email("GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", resp.StatusCode)
	//fmt.Printf("%s\n", bodyText)
	var getb bid
	err = json.Unmarshal(bodyText, &getb) //反序列化
	if err != nil {
		send_email("GetbussinessID出现了一些问题.....  "+u_time, err.Error())
		log.Fatal(err)
	}
	bid := getb.Data[0].Business.ID
	//println(getb.Data[0].Business.ID)
	return strconv.Itoa(bid)
}

func getMailAdress() string {
	f, err := os.Open("mail.txt")
	if err != nil {
		log.Printf("mailadress open failed : %#v", err)
		return ""
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("mailadress read failed :%#v", err)
		return ""
	}
	return string(fd)
}

func getID() string {
	client := &http.Client{}
	bussid = getbussinessID()
	fmt.Println("businessID:", bussid)
	url := "http://ijg.xujc.com/api/formEngine/business/" + getbussinessID() + "/myFormInstance"
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
	req.Header.Set("Cookie", getCookie())
	req.Header.Set("If-None-Match", `W/"4138-V3qtR8Z1sFVcaru9at8wciTNIFM"`)
	req.Header.Set("Referer", "http://ijg.xujc.com/app/229")
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
	req, err := http.NewRequest("GET", "http://ijg.xujc.com/api/formEngine/business/1574/table/fields?playerId=owner", nil)
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
	req.Header.Set("Referer", "http://ijg.xujc.com/app/229")
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

func setbody(url string, st string, t string, is string) string {
	mbody := "<!DOCTYPE html>\n<html lang=\"en\" dir=\"ltr\">\n<head>\n<meta charset=\"utf-8\">\n<title></title>\n<style>\n* {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n\tfont-family: \"Open Sans\", sans-serif;\n}\nbody {\n\tbackground: #786fa6;\n\theight: 100vh;\n\tdisplay: flex;\n\talign-items: center;\n\tjustify-content: center;\n}\n.newsletter {\n\tbackground: linear-gradient(125deg, #778beb, #f8a5c2);\n\tpadding: 70px;\n\ttext-align: center;\n\tbox-shadow: 0 0 20px #00000060;\n}\n.newsletter h1 {\n\ttext-transform: uppercase;\n\tcolor: #fff;\n\tfont-size: 30px;\n\tline-height: 40px;\n}\n.newsletter h1 span {\n\tdisplay: block;\n\tfont-size: 38px;\n}\n.newsletter p {\n\tcolor: #fff;\n\tfont-size: 14px;\n\tmargin: 10px 0;\n}\n.txtb {\n\twidth: 100%;\n\theight: 70px;\n\tbackground: #f1f1f199;\n\tborder-radius: 40px;\n\tposition: relative;\n\tmargin-top: 40px;\n}\n.txtb input {\n\twidth: 100%;\n\theight: 70px;\n\tborder-radius: 40px;\n\tborder: 0;\n\tbackground: none;\n\tpadding: 0 30px;\n\toutline: none;\n\tfont-size: 15px;\n\tpadding-right: 80px;\n}\n.txtb button {\n\tbackground: #574b90;\n\tborder: 0;\n\twidth: 50px;\n\theight: 50px;\n\tborder-radius: 50%;\n\tposition: absolute;\n\tright: 10px;\n\ttop: 10px;\n\toutline: none;\n\tcursor: pointer;\n\tcolor: #fff;\n\ttransition: 0.3s linear;\n}\n.txtb button:hover {\n\topacity: .5;\n}\n.s {\n\tcolor:#0084FF;\n\tfont-size: 15px;\n}\n</style>\n<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.11.2/css/all.css\">\n</head>\n<body>\n<div class=\"newsletter\">\n  <h1> Temperature Sign In</h1><br>\n  \n " +
		" <span class=\"s\">" + "NO:" + u_no + "<br>" +
		"Name:" + u_name + "<br>" +
		"BusinessID:" + bussid + "<br>" +
		"URL:" + url + "<br>" +
		"State:" + st + "<br>" +
		"isChange:" + is + "<br>" +
		"UpdateTime:" + u_time +
		"</span> <br>\n\t" +
		"<p>Made with ❤ by HaynesChen.</p>\n\t</div>\n</body>\n</html>"
	return mbody
}

func send_email(mailtitle string, mailbody string) {
	//简单设置log参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	//设置sender发送方的邮箱，此处可以填写自己的邮箱
	em.From = "ian-chen0713@qq.com"

	mail := getMailAdress()
	if mail == "" {
		return
	}

	//设置receiver接收方的邮箱此处也可以填写自己的邮箱，就是自己发邮件给自己
	em.To = []string{mail}

	//设置主题
	em.Subject = mailtitle

	//简单设置文件发送的内容
	em.HTML = []byte(mailbody)

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
	fmt.Println("Automatic temperature registration by Haynes v1.4")
	t := time.Now() //当前时间
	timeLayoutStr := "2006-01-02 15:04:05"
	u_time = t.Format(timeLayoutStr)
	id := getID()
	client := &http.Client{}
	var data = strings.NewReader(getformdata())
	url := "http://ijg.xujc.com/api/formEngine/formInstance/" + id
	//println("url:", url)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		send_email("Post出现了一些问题..... "+u_time, err.Error())
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0. 8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", getCookie())
	req.Header.Set("Origin", "http://ijg.xujc.com")
	req.Header.Set("Referer", "http://ijg.xujc.com/app/229")
	req.Header.Set("User-Agent", User_Agent)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
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
	if !isChange() {
		setbody(url, st, u_time, "True")
		fmt.Println("!_! 表单数据改变了,快来看看")
		send_email("!_! 表单数据改变了,快来看看", setbody(url, st, u_time, "True"))
	}
	fmt.Println("Status:", st)
	fmt.Println("isChange:", isch)
	fmt.Println("Time:", u_time)
	if xujc.State != true {
		send_email("QAQ 体温打卡失败啦,快来检查一下  "+u_time, setbody(url, st, u_time, "False"))
	} else {
		send_email("0.0 体温打卡成功!   "+u_time, setbody(url, st, u_time, "False"))
	}
	pause()
}
