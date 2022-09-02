package main

import (
	"encoding/json"
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"log"
	"math/rand"
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
	//mbody := "Name:" + u_name + "<br>" +
	//		"BusinessID:" + bussid + "<br>" +
	//		"URL:" + url + "<br>" +
	//		"State:" + st + "<br>" +
	//		"isChange:" + is + "<br>" +
	//		"UpdateTime:" + u_time
	rand.Seed(time.Now().UnixNano())
	food := strconv.Itoa(rand.Intn(20))
	photo := strconv.Itoa(rand.Intn(10))
	println("food:", food)
	println("photo:", photo)
	mbody := "<!-- saved from url=(0078)https://ui.benchmarkemail.com/Emails/Print?email_id=21929284&client_id=1477270 -->\n<html xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\">\n\n</head>\n<body marginwidth=\"10\" marginheight=\"10\" leftmargin=\"10\" topmargin=\"0\"\n      style=\"height: 100% !important; margin: 0; padding: 0; width: 100% !important;min-width: 100%;\">\n<link rel=\"stylesheet\" href=\"./Print_files/print.css\" type=\"text/css\" media=\"print\">\n<br>\n\n\n<meta content=\"width=device-width, initial-scale=1.0\" name=\"viewport\">\n<style type=\"text/css\">\n    /*** BMEMBF Start ***/\n    /* CMS V4 Editor Test */\n    [name=bmeMainBody] {\n        min-height: 1000px;\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.blk, table.tblText, .bmeHolder, .bmeHolder1, table.bmeMainColumn {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeImageCard table.bmeCaptionTable td.tblCell {\n            padding: 0px 20px 20px 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeImageCard table.bmeCaptionTable.bmeCaptionTableMobileTop td.tblCell {\n            padding: 20px 20px 0 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeCaptionTable td.tblCell {\n            padding: 10px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.tblGtr {\n            padding-bottom: 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td.blk_container, .blk_parent, .bmeLeftColumn, .bmeRightColumn, .bmeColumn1, .bmeColumn2, .bmeColumn3, .bmeBody {\n            display: table !important;\n            max-width: 600px !important;\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.container-table, .bmeheadertext, .container-table {\n            width: 95% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .mobile-footer, .mobile-footer a {\n            font-size: 13px !important;\n            line-height: 18px !important;\n        }\n\n        .mobile-footer {\n            text-align: center !important;\n        }\n\n        table.share-tbl {\n            padding-bottom: 15px;\n            width: 100% !important;\n        }\n\n        table.share-tbl td {\n            display: block !important;\n            text-align: center !important;\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td.bmeShareTD, td.bmeSocialTD {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td.tdBoxedTextBorder {\n            width: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        th.tdBoxedTextBorder {\n            width: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.blk, table[name=tblText], .bmeHolder, .bmeHolder1, table[name=bmeMainColumn] {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeImageCard table.bmeCaptionTable td[name=tblCell] {\n            padding: 0px 20px 20px 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeImageCard table.bmeCaptionTable.bmeCaptionTableMobileTop td[name=tblCell] {\n            padding: 20px 20px 0 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeCaptionTable td[name=tblCell] {\n            padding: 10px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table[name=tblGtr] {\n            padding-bottom: 20px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td.blk_container, .blk_parent, [name=bmeLeftColumn], [name=bmeRightColumn], [name=bmeColumn1], [name=bmeColumn2], [name=bmeColumn3], [name=bmeBody] {\n            display: table !important;\n            max-width: 600px !important;\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table[class=container-table], .bmeheadertext, .container-table {\n            width: 95% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .mobile-footer, .mobile-footer a {\n            font-size: 13px !important;\n            line-height: 18px !important;\n        }\n\n        .mobile-footer {\n            text-align: center !important;\n        }\n\n        table[class=\"share-tbl\"] {\n            padding-bottom: 15px;\n            width: 100% !important;\n        }\n\n        table[class=\"share-tbl\"] td {\n            display: block !important;\n            text-align: center !important;\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td[name=bmeShareTD], td[name=bmeSocialTD] {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td[name=tdBoxedTextBorder] {\n            width: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        th[name=tdBoxedTextBorder] {\n            width: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeImageCard table.bmeImageTable {\n            height: auto !important;\n            width: 100% !important;\n            padding: 20px !important;\n            clear: both;\n            float: left !important;\n            border-collapse: separate;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblInline table.bmeImageTable {\n            height: auto !important;\n            width: 100% !important;\n            padding: 10px !important;\n            clear: both;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblInline table.bmeCaptionTable {\n            width: 100% !important;\n            clear: both;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeImageTable {\n            height: auto !important;\n            width: 100% !important;\n            padding: 10px !important;\n            clear: both;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeCaptionTable {\n            width: 100% !important;\n            clear: both;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeImageContainer {\n            width: 100% !important;\n            clear: both;\n            float: left !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.bmeImageTable td {\n            padding: 0px !important;\n            height: auto;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        img.mobile-img-large {\n            width: 100% !important;\n            height: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        img.bmeRSSImage {\n            max-width: 320px;\n            height: auto !important;\n        }\n    }\n\n    @media only screen and (min-width: 640px) {\n        img.bmeRSSImage {\n            max-width: 600px !important;\n            height: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .trMargin img {\n            height: 10px;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        div.bmefooter, div.bmeheader {\n            display: block !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdPart {\n            width: 100% !important;\n            clear: both;\n            float: left !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        table.blk_parent1, table.tblPart {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tblLine {\n            min-width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblCenter img {\n            margin: 0 auto;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblCenter, .bmeMblCenter div, .bmeMblCenter span {\n            text-align: center !important;\n            text-align: -webkit-center !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeNoBr br, .bmeImageGutterRow, .bmeMblStackCenter .bmeShareItem .tdMblHide {\n            display: none !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblInline table.bmeImageTable, .bmeMblInline table.bmeCaptionTable, .bmeMblInline {\n            clear: none !important;\n            width: 50% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblInlineHide, .bmeShareItem .trMargin {\n            display: none !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblInline table.bmeImageTable img, .bmeMblShareCenter.tblContainer.mblSocialContain, .bmeMblFollowCenter.tblContainer.mblSocialContain {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStack > .bmeShareItem {\n            width: 100% !important;\n            clear: both !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeShareItem {\n            padding-top: 10px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdPart.bmeMblStackCenter, .bmeMblStackCenter .bmeFollowItemIcon {\n            padding: 0px !important;\n            text-align: center !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStackCenter > .bmeShareItem {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        td.bmeMblCenter {\n            border: 0 none transparent !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeLinkTable.tdPart td {\n            padding-left: 0px !important;\n            padding-right: 0px !important;\n            border: 0px none transparent !important;\n            padding-bottom: 15px !important;\n            height: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdMblHide {\n            width: 10px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeShareItemBtn {\n            display: table !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStack td {\n            text-align: left !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStack .bmeFollowItem {\n            clear: both !important;\n            padding-top: 10px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStackCenter .bmeFollowItemText {\n            padding-left: 5px !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .bmeMblStackCenter .bmeFollowItem {\n            clear: both !important;\n            align-self: center;\n            float: none !important;\n            padding-top: 10px;\n            margin: 0 auto;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdPart > table {\n            width: 100% !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdPart > table.bmeLinkContainer {\n            width: auto !important;\n        }\n    }\n\n    @media only screen and (max-width: 480px) {\n        .tdPart.mblStackCenter > table.bmeLinkContainer {\n            width: 100% !important;\n        }\n    }\n\n    .blk_parent:first-child, .blk_parent {\n        float: left;\n    }\n\n    .blk_parent:last-child {\n        float: right;\n    }\n\n    BODY {\n        color: black;\n        background: white;\n    }\n\n    /*A:link, A:visited {background: white; color: black; text-decoration: underline;font-weight: bold;}*/\n    /*H1, H2, H3 {background: white; color: black; padding-bottom: 1px;border-bottom: 1px solid gray;}*/\n    .subhead {\n        display: none;\n    }\n\n    .nonprint {\n        display: none;\n    }\n\n    /*** BMEMBF END ***/\n\n</style>\n<!--[if gte mso 9]>\n<xml>\n    <o:OfficeDocumentSettings>\n        <o:AllowPNG/>\n        <o:PixelsPerInch>96</o:PixelsPerInch>\n    </o:OfficeDocumentSettings>\n</xml>\n<![endif]-->\n\n\n<table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" name=\"bmeMainBody\" bgcolor=\"#ededed\"\n       style=\"background-color: rgb(237, 237, 237);\">\n    <tbody>\n    <tr>\n        <td width=\"100%\" valign=\"top\" align=\"center\">\n            <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" name=\"bmeMainColumnParentTable\">\n                <tbody>\n                <tr>\n                    <td name=\"bmeMainColumnParent\"\n                        style=\"border: 0px none transparent; border-radius: 0px; border-collapse: separate;\">\n                        <table name=\"bmeMainColumn\" class=\"bmeHolder bmeMainColumn\"\n                               style=\"max-width: 600px; overflow: visible; border-radius: 0px; border-collapse: separate; border-spacing: 0px;\"\n                               cellspacing=\"0\" cellpadding=\"0\" border=\"0\" align=\"center\">\n                            <tbody>\n                            <tr id=\"section_1\" class=\"flexible-section\" data-columns=\"1\"\n                                data-section-type=\"bmePreHeader\">\n                                <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmePreHeader\" valign=\"top\"\n                                    align=\"center\" style=\"color: rgb(102, 102, 102); border: 0px none transparent;\"\n                                    bgcolor=\"\"></td>\n                            </tr>\n                            <tr>\n                                <td width=\"100%\" class=\"bmeHolder\" valign=\"top\" align=\"center\"\n                                    name=\"bmeMainContentParent\"\n                                    style=\"border: 0px none rgb(102, 102, 102); border-radius: 0px; border-collapse: separate; border-spacing: 0px; overflow: hidden;\">\n                                    <table name=\"bmeMainContent\"\n                                           style=\"border-radius: 0px; border-collapse: separate; border-spacing: 0px; border: 0px none transparent;\"\n                                           width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" align=\"center\">\n                                        <tbody>\n                                        <tr id=\"section_2\" class=\"flexible-section\" data-columns=\"1\">\n                                            <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmeHeader\"\n                                                valign=\"top\" align=\"center\"\n                                                style=\"color: rgb(56, 56, 56); border: 0px none transparent; background-color: rgb(52, 52, 52);\"\n                                                bgcolor=\"#343434\">\n                                                <div id=\"dv_11\" class=\"blk_wrapper\" style=\"\">\n                                                    <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"\n                                                           class=\"blk\" name=\"blk_imagecaption\" style=\"width: 600px;\">\n                                                        <tbody>\n                                                        <tr>\n                                                            <td>\n                                                                <table cellspacing=\"0\" cellpadding=\"0\"\n                                                                       class=\"bmeCaptionContainer\" width=\"100%\"\n                                                                       style=\"padding-left:20px; padding-right:20px; padding-top:20px; padding-bottom:20px;border-collapse:separate;\">\n                                                                    <tbody>\n                                                                    <tr>\n                                                                        <td class=\"bmeImageContainerRow\" valign=\"top\"\n                                                                            gutter=\"10\">\n                                                                            <table width=\"100%\" cellpadding=\"0\"\n                                                                                   cellspacing=\"0\" border=\"0\">\n                                                                                <tbody>\n                                                                                <tr>\n                                                                                    <td class=\"tdPart\" valign=\"top\">\n                                                                                        <table cellspacing=\"0\"\n                                                                                               cellpadding=\"0\"\n                                                                                               border=\"0\"\n                                                                                               class=\"bmeImageContainer\"\n                                                                                               width=\"100%\" align=\"left\"\n                                                                                               style=\"float:left;\">\n                                                                                            <tbody>\n                                                                                            <tr>\n                                                                                                <td valign=\"top\"\n                                                                                                    name=\"tdContainer\">\n                                                                                                    <table cellspacing=\"0\"\n                                                                                                           cellpadding=\"0\"\n                                                                                                           border=\"0\"\n                                                                                                           class=\"bmeImageTable\"\n                                                                                                           dimension=\"50%\"\n                                                                                                           imgid=\"1\"\n                                                                                                           style=\"float: left; width: 270px;\"\n                                                                                                           align=\"left\"\n                                                                                                           width=\"270\"\n                                                                                                           height=\"63\">\n                                                                                                        <tbody>\n                                                                                                        <tr>\n                                                                                                            <td name=\"bmeImgHolder\"\n                                                                                                                width=\"280\"\n                                                                                                                align=\"left\"\n                                                                                                                valign=\"middle\"\n                                                                                                                class=\"bmeMblCenter\"\n                                                                                                                height=\"63\">\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t" +
		"<img src=\"http://chenhy.tk/images/icon/" + food + ".png\" width=\"63.2\" style=\"max-width: 63.2px; display: block; border-radius: 109px;\" alt=\"\" data-radius=\"109\" data-original-max-width=\"234\" class=\"keep-custom-width\" data-customwidth=\"27\">\n                                                                                                            </td>\n                                                                                                        </tr>\n                                                                                                        </tbody>\n                                                                                                    </table>\n                                                                                                    <table cellspacing=\"0\"\n                                                                                                           cellpadding=\"0\"\n                                                                                                           border=\"0\"\n                                                                                                           class=\"bmeCaptionTable\"\n                                                                                                           style=\"float: right; width: 270px; word-break: break-word;\"\n                                                                                                           align=\"right\"\n                                                                                                           width=\"270\">\n                                                                                                        <tbody>\n                                                                                                        <tr>\n                                                                                                            <td name=\"tblCell\"\n                                                                                                                valign=\"top\"\n                                                                                                                align=\"left\"\n                                                                                                                style=\"font-family: Arial, Helvetica, sans-serif; font-size: 14px; font-weight: normal; color: rgb(56, 56, 56); text-align: left; word-break: break-word;\"\n                                                                                                                class=\"bmeMblCenter tblCell\">\n                                                                                                                <div style=\"line-height: 150%; text-align: right;\">\n                                                                                                                    <span style=\"font-family: Helvetica, Arial, sans-serif; font-size: 14px; color: #dfdfdf; line-height: 150%;\">" +
		u_name + "</span>\n                                                                                                                    <br><span\n                                                                                                                        style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #ab966c; line-height: 150%;\"><strong>" +
		u_no + "</strong></span>\n                                                                                                                    <br><span\n                                                                                                                        style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #dedede; line-height: 150%;\">" +
		u_time + "</span>\n                                                                                                                </div>\n                                                                                                            </td>\n                                                                                                        </tr>\n                                                                                                        </tbody>\n                                                                                                    </table>\n                                                                                                </td>\n                                                                                            </tr>\n                                                                                            </tbody>\n                                                                                        </table>\n                                                                                    </td>\n                                                                                </tr>\n                                                                                </tbody>\n                                                                            </table>\n                                                                        </td>\n                                                                    </tr>\n                                                                    </tbody>\n                                                                </table>\n                                                            </td>\n                                                        </tr>\n                                                        </tbody>\n                                                    </table>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                        <tr id=\"section_3\" class=\"flexible-section\" data-columns=\"1\">\n                                            <td width=\"100%\" class=\"blk_container bmeHolder bmeBody\" name=\"bmeBody\"\n                                                valign=\"top\" align=\"center\"\n                                                style=\"color: rgb(56, 56, 56); border: 0px none transparent; background-color: rgb(255, 255, 255);\"\n                                                bgcolor=\"#ffffff\">\n                                                <div id=\"dv_3\" class=\"blk_wrapper\">\n                                                    <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"\n                                                           class=\"blk\" name=\"blk_image\">\n                                                        <tbody>\n                                                        <tr>\n                                                            <td>\n                                                                <table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\"\n                                                                       border=\"0\">\n                                                                    <tbody>\n                                                                    <tr>\n                                                                        <td align=\"center\" class=\"bmeImage\"\n                                                                            style=\"border-collapse: collapse; padding: 0px;\">\n                                                                            " +
		"<img src=\"http://chenhy.tk/images/" + photo + ".jpg\"\n                                                                                 class=\"mobile-img-large\" width=\"600\"\n                                                                                 style=\"max-width: 1196px; display: block;\"\n                                                                                 alt=\"\" border=\"0\"></td>\n                                                                    </tr>\n                                                                    </tbody>\n                                                                </table>\n                                                            </td>\n                                                        </tr>\n                                                        </tbody>\n                                                    </table>\n                                                </div>\n                                                <div id=\"dv_6\" class=\"blk_wrapper\">\n                                                    <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"\n                                                           class=\"blk\" name=\"blk_divider\" style=\"\">\n                                                        <tbody>\n                                                        <tr>\n                                                            <td class=\"tblCellMain\" style=\"padding: 10px 20px;\">\n                                                                <table class=\"tblLine\" cellspacing=\"0\" cellpadding=\"0\"\n                                                                       border=\"0\" width=\"100%\"\n                                                                       style=\"border-top-width: 0px; border-top-style: none; min-width: 1px;\">\n                                                                    <tbody>\n                                                                    <tr>\n                                                                        <td><span></span></td>\n                                                                    </tr>\n                                                                    </tbody>\n                                                                </table>\n                                                            </td>\n                                                        </tr>\n                                                        </tbody>\n                                                    </table>\n                                                </div>\n                                                <div id=\"dv_1\" class=\"blk_wrapper\">\n                                                    <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"\n                                                           class=\"blk\" name=\"blk_text\">\n                                                        <tbody>\n                                                        <tr>\n                                                            <td>\n                                                                <table cellpadding=\"0\" cellspacing=\"0\" border=\"0\"\n                                                                       width=\"100%\" class=\"bmeContainerRow\">\n                                                                    <tbody>\n                                                                    <tr>\n                                                                        <th valign=\"top\" align=\"center\" class=\"tdPart\">\n                                                                            <table cellspacing=\"0\" cellpadding=\"0\"\n                                                                                   border=\"0\" width=\"600\" name=\"tblText\"\n                                                                                   style=\"float:left; background-color:transparent;\"\n                                                                                   align=\"left\" class=\"tblText\">\n                                                                                <tbody>\n                                                                                <tr>\n                                                                                    <td valign=\"top\" align=\"left\"\n                                                                                        name=\"tblCell\"\n                                                                                        style=\"padding: 20px; font-family: Arial, Helvetica, sans-serif; font-size: 14px; font-weight: 400; color: rgb(56, 56, 56); text-align: left; word-break: break-word;\"\n                                                                                        class=\"tblCell\">\n                                                                                        <div style=\"line-height: 200%;\">\n                                                                                            <span style=\"font-size: 20px; font-family: Helvetica, Arial, sans-serif; color: #343434; line-height: 200%;\"><strong>" +
		t+"</strong></span>\n                                                                                            <br>\n                                                                                            <br><span\n                                                                                                style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #b0b0b0; line-height: 200%;\">" +
		"URL:" + url + "<br>" + "BusinessID:" + bussid + "<br>" + "State:" + st + "<br>" + "isChange:" + is + "<br>" + "UpdateTime:" + u_time +
		"</span>\n                                                                                            \n                                                                                        <span style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #b0b0b0; line-height: 200%;\">&nbsp;</span> </div>\n                                                                                    </td>\n                                                                                </tr>\n                                                                                </tbody>\n                                                                            </table>\n                                                                        </th>\n                                                                    </tr>\n                                                                    </tbody>\n                                                                </table>\n                                                            </td>\n                                                        </tr>\n                                                        </tbody>\n                                                    </table>\n                                                </div>\n\n\n                                            </td>\n                                        </tr>\n                                        <tr id=\"section_4\" class=\"flexible-section\" data-columns=\"1\">\n                                            <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmePreFooter\"\n                                                valign=\"top\" align=\"center\"\n                                                style=\"border: 0px none transparent; background-color: rgb(255, 255, 255);\"\n                                                bgcolor=\"#ffffff\">\n                                                <div id=\"dv_5\" class=\"blk_wrapper\">\n                                                    <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\"\n                                                           class=\"blk\" name=\"blk_divider\" style=\"\">\n                                                        <tbody>\n                                                        <tr></tr>\n                                                        </tbody>\n                                                    </table>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                        </tbody>\n                                    </table>\n                                </td>\n                            </tr>\n                            <tr id=\"section_5\" class=\"flexible-section\" data-columns=\"1\" data-section-type=\"bmeFooter\">\n                                <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmeFooter\" valign=\"top\"\n                                    align=\"center\" style=\"color: rgb(102, 102, 102); border: 0px none transparent;\"\n                                    bgcolor=\"\">\n                                    <div id=\"dv_2\" class=\"blk_wrapper\" style=\"\">\n                                        <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"blk\"\n                                               name=\"blk_footer\" style=\"\">\n                                            <tbody>\n                                            <tr></tr>\n                                            </tbody>\n                                        </table>\n                                    </div>\n                                </td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </td>\n                </tr>\n                </tbody>\n            </table>\n        </td>\n    </tr>\n    </tbody>\n</table>\n<!-- /Test Path -->\n\n\n</body>\n</html>"
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
	fmt.Println("Automatic temperature registration by Haynes v1.6")
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
		isch = "True"
		setbody(url, st, u_time, "True")
		fmt.Println("!_! 表单数据改变了,快来看看")
		send_email("!_! 表单数据改变了,快来看看", setbody(url, st, "!_! 表单数据改变了,快来看看", "True"))
	} else {
		isch = "None"
	}
	fmt.Println("Status:", st)
	fmt.Println("isChange:", isch)
	fmt.Println("Time:", u_time)
	if xujc.State != true {
		send_email("QAQ 体温打卡失败啦,快来检查一下  "+u_time, setbody(url, st, "QAQ 体温打卡失败啦,快来检查一下", "None"))
	} else {
		send_email("0.0 体温打卡成功!   "+u_time, setbody(url, st, "0.0 体温打卡成功!", "None"))
	}
	pause()
}
