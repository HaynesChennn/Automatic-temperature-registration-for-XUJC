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

//var cookie = "_dx_uzZo5y=a389b56dee2bbc800696affa8164893f72a5651fc24098911bbaf3f2bd5ac83fc2f6370f; Hm_lvt_d4b4fe5895335a64dc71a1e3d97ecaae=1650788542,1650966832,1651718592; SAAS_S_ID=jgxy; JSESSIONID=93A8F358CC96D982F91AEE0B173324F5; _dx_app_5c7bafe274b534f13ec3b614135a362e=627a69fcopGU37T5P1TYiyGssFahg9ISpJaGmtL1; _dx_captcha_vid=0D84355393FB38CCB0E29E0771352C0AEA4F7D37B0B1AEDF8E848B01DBFEA1F615AF05286A5DD11B52A65E1DA4D628AAAE729001ABD62ABA2B21B6EADD17B3826036ADAD99C617C968DDB3C3FAC9D847; SAAS_U=1455cd5698302b35041285c57a960dfc995be15fd444a2a77d85911fa04055a9"
var User_Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36"

var u_name, u_no string

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
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", resp.StatusCode)
	//fmt.Printf("%s\n", bodyText)
	var getb bid
	err = json.Unmarshal(bodyText, &getb) //反序列化
	if err != nil {
		log.Fatal(err)
	}
	bid := getb.Data[0].Business.ID
	//println(getb.Data[0].Business.ID)
	return strconv.Itoa(bid)
}

func getID() string {
	client := &http.Client{}
	fmt.Println("businessID:", getbussinessID())
	url := "http://ijg.xujc.com/api/formEngine/business/" + getbussinessID() + "/myFormInstance"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
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
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var get Xujc
	err = json.Unmarshal(bodyText, &get) //反序列化
	if err != nil {
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
	fmt.Println("Automatic temperature registration by Haynes v1.3")
	id := getID()
	client := &http.Client{}
	//formdata := `{"formData":[{"name":"label_1582537738348","title":"文本","value":{},"hide":false,"readonly":true},{"name":"input_1582537793424","title":"学号","value":{"stringValue":"ROB21026"},"hide":false,"readonly":true},{"name":"input_1582537796856","title":"姓名","value":{"stringValue":"陈鸿壹"},"hide":false,"readonly":true},{"name":"input_1582537799026","title":"学院","value":{"stringValue":"信息科学与技术学院"},"hide":false,"readonly":true},{"name":"select_1582639884820","title":"性别","value":{"stringValue":"男"},"hide":false,"readonly":true},{"name":"input_1582639887112","title":"年级","value":{"stringValue":"2021"},"hide":false,"readonly":true},{"name":"input_1582639888256","title":"班级","value":{"stringValue":"计算机21(3)"},"hide":false,"readonly":true},{"name":"input_1582639889478","title":"辅导员","value":{"stringValue":"高新琪"},"hide":false,"readonly":true},{"name":"select_1631714040062","title":"完成情况","value":{"stringValue":"已打卡"},"hide":false,"readonly":false},{"name":"label_1582538217569","title":"文本","value":{},"hide":false,"readonly":true},{"name":"select_1582538214785","title":"当前所在的地理位置","value":{"stringValue":"校内"},"hide":false,"readonly":false},{"name":"input_1582538157713","title":"当前所在国家","value":{"stringValue":""},"hide":true,"readonly":false},{"name":"address_1582538163410","title":"当前所在的省市区","value":{"addressValue":{"province":"福建省","city":"漳州市","area":"龙海市","fullValue":"福建省漳州市龙海市"}},"hide":false,"readonly":false},{"name":"select_1648711313366","title":"目前是否居住在招商局漳州开发区","value":{"stringValue":"是"},"hide":false,"readonly":false},{"name":"label_1582537910151","title":"文本","value":{},"hide":false,"readonly":true},{"name":"select_1631790340241","title":"今日体温","value":{"stringValue":"37.3以下"},"hide":false,"readonly":false},{"name":"select_1640686551031","title":"是否已接种疫苗","value":{"stringValue":"已接种3针"},"hide":false,"readonly":false},{"name":"datetime_1640686554409","title":"最后一针疫苗接种日期","value":{"dateValue":"2022-02-27 00:00:00"},"hide":false,"readonly":false},{"name":"label_1644481853234","title":"文本","value":{},"hide":false,"readonly":true},{"name":"select_1641522783266","title":"今日是否有中高风险【所在城市】旅居史","value":{"stringValue":"否"},"hide":false,"readonly":false},{"name":"label_1641522839410","title":"中高风险地区查询","value":{},"hide":false,"readonly":true},{"name":"select_1641522890815","title":"今日是否有中高风险地区旅居史","value":{"stringValue":null},"hide":true,"readonly":false},{"name":"input_1641522901563","title":"中高风险地区旅居史详细地址","value":{"stringValue":null},"hide":true,"readonly":false},{"name":"select_1641523063583","title":"今日个人疫情管控情况","value":{"stringValue":"无"},"hide":false,"readonly":false},{"name":"select_1641523103608","title":"今日是否有境外旅居史","value":{"stringValue":"无"},"hide":false,"readonly":false},{"name":"select_1588863625331","title":"今日本人及共同居住人员是否与中高风险地区或境外回国人员接触？","value":{"stringValue":"否"},"hide":false,"readonly":false},{"name":"select_1582538846920","title":"今日是否出现发热、咳嗽、胸闷、呼吸困难等症状？","value":{"stringValue":"否"},"hide":false,"readonly":false},{"name":"select_1582538869774","title":"是否就诊或住院","value":{"stringValue":""},"hide":true,"readonly":false},{"name":"table_1588863652072","title":"与中高风险地区或境外回国人员接触情况登记（需点击左下角“+新增”）","value":{"tableValue":[]},"hide":true,"readonly":false},{"name":"label_1582538416593","title":"文本","value":{},"hide":false,"readonly":true},{"name":"input_1582538924486","title":"备注","value":{"stringValue":null},"hide":false,"readonly":false},{"name":"select_1582538939790","title":"本人是否承诺以上所填报的全部内容均属实、准确，不存在任何隐瞒和不实的情况，更无遗漏之处。","value":{"stringValue":"是"},"hide":false,"readonly":false}],"playerId":"owner"}`
	var data = strings.NewReader(getformdata())
	url := "http://ijg.xujc.com/api/formEngine/formInstance/" + id
	//println("url:", url)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
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
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	var xujc xujc_response
	err = json.Unmarshal(bodyText, &xujc) //反序列化
	if err != nil {
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
	t := time.Now() //当前时间
	timeLayoutStr := "2006-01-02 15:04:05"
	u_time := t.Format(timeLayoutStr)
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
