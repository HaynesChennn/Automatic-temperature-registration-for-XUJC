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
	req, err := http.NewRequest("GET", "http://ijg.xujc.com/api/formEngine/business/"+bussid+"/table/fields?playerId=owner", nil)
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
	//println("absd")
	//fmt.Printf("%s\n", bodyText)
	//fmt.Println("change:", len(bodyText))
	if len(bodyText) == 10605 {
		return true
	} else {
		return false
	}
}

func setbody(url string, st string, t string, is string) string {
	//mbody := "Name:" + u_name + "<br>" +
	//		"BusinessID:" + bussid + "<br>" +
	//		"URL:" + url + "<br>" +
	//		"State:" + st + "<br>" +
	//		"isChange:" + is + "<br>" +
	//		"UpdateTime:" + u_time
	rand.Seed(time.Now().UnixNano())
	food := strconv.Itoa(rand.Intn(33))
	photo := strconv.Itoa(rand.Intn(10))
	println("food:", food)
	println("photo:", photo)
	mbody := "<html xmlns=\"http://www.w3.org/1999/xhtml\">\n<head>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\">\n</head>\n<body marginwidth=\"10\" marginheight=\"10\" leftmargin=\"10\" topmargin=\"0\" style=\"height: 100% !important; margin: 0; padding: 0; width: 100% !important;min-width: 100%;\">\n<meta content=\"width=device-width, initial-scale=1.0\" name=\"viewport\">\n<style type=\"text/css\">    \nBODY {color: black; background: white;}\n.subhead{display:none;}\n.nonprint{display:none;}\n/*** BMEMBF Start ***/    \n[name=bmeMainBody]{min-height:1000px;}    \n@media only screen and (max-width: 480px){table.blk, table.tblText, .bmeHolder, .bmeHolder1, table.bmeMainColumn{width:100% !important;} }        \n@media only screen and (max-width: 480px){.bmeImageCard table.bmeCaptionTable td.tblCell{padding:0px 20px 20px 20px !important;} }        \n@media only screen and (max-width: 480px){.bmeImageCard table.bmeCaptionTable.bmeCaptionTableMobileTop td.tblCell{padding:20px 20px 0 20px !important;} }        \n@media only screen and (max-width: 480px){table.bmeCaptionTable td.tblCell{padding:10px !important;} }        \n@media only screen and (max-width: 480px){table.tblGtr{ padding-bottom:20px !important;} }        \n@media only screen and (max-width: 480px){td.blk_container, .blk_parent, .bmeLeftColumn, .bmeRightColumn, .bmeColumn1, .bmeColumn2, .bmeColumn3, .bmeBody{display:table !important;max-width:600px !important;width:100% !important;} }        \n@media only screen and (max-width: 480px){table.container-table, .bmeheadertext, .container-table { width: 95% !important; } }        \n@media only screen and (max-width: 480px){.mobile-footer, .mobile-footer a{ font-size: 13px !important; line-height: 18px !important; } .mobile-footer{ text-align: center !important; } table.share-tbl { padding-bottom: 15px; width: 100% !important; } table.share-tbl td { display: block !important; text-align: center !important; width: 100% !important; } }        \n@media only screen and (max-width: 480px){td.bmeShareTD, td.bmeSocialTD{width: 100% !important; } }        \n@media only screen and (max-width: 480px){td.tdBoxedTextBorder{width: auto !important;} }    \n@media only screen and (max-width: 480px){th.tdBoxedTextBorder{width: auto !important;} }    \n    \n@media only screen and (max-width: 480px){table.blk, table[name=tblText], .bmeHolder, .bmeHolder1, table[name=bmeMainColumn]{width:100% !important;} }    \n@media only screen and (max-width: 480px){.bmeImageCard table.bmeCaptionTable td[name=tblCell]{padding:0px 20px 20px 20px !important;} }    \n@media only screen and (max-width: 480px){.bmeImageCard table.bmeCaptionTable.bmeCaptionTableMobileTop td[name=tblCell]{padding:20px 20px 0 20px !important;} }    \n@media only screen and (max-width: 480px){table.bmeCaptionTable td[name=tblCell]{padding:10px !important;} }    \n@media only screen and (max-width: 480px){table[name=tblGtr]{ padding-bottom:20px !important;} }    \n@media only screen and (max-width: 480px){td.blk_container, .blk_parent, [name=bmeLeftColumn], [name=bmeRightColumn], [name=bmeColumn1], [name=bmeColumn2], [name=bmeColumn3], [name=bmeBody]{display:table !important;max-width:600px !important;width:100% !important;} }    \n@media only screen and (max-width: 480px){table[class=container-table], .bmeheadertext, .container-table { width: 95% !important; } }    \n@media only screen and (max-width: 480px){.mobile-footer, .mobile-footer a{ font-size: 13px !important; line-height: 18px !important; } .mobile-footer{ text-align: center !important; } table[class=\"share-tbl\"] { padding-bottom: 15px; width: 100% !important; } table[class=\"share-tbl\"] td { display: block !important; text-align: center !important; width: 100% !important; } }    \n@media only screen and (max-width: 480px){td[name=bmeShareTD], td[name=bmeSocialTD]{width: 100% !important; } }    \n@media only screen and (max-width: 480px){td[name=tdBoxedTextBorder]{width: auto !important;} }    \n@media only screen and (max-width: 480px){th[name=tdBoxedTextBorder]{width: auto !important;} }    \n               \n@media only screen and (max-width: 480px){.bmeImageCard table.bmeImageTable{height: auto !important; width:100% !important; padding:20px !important;clear:both; float:left !important; border-collapse: separate;} }                \n@media only screen and (max-width: 480px){.bmeMblInline table.bmeImageTable{height: auto !important; width:100% !important; padding:10px !important;clear:both;} }    \n@media only screen and (max-width: 480px){.bmeMblInline table.bmeCaptionTable{width:100% !important; clear:both;} }    \n@media only screen and (max-width: 480px){table.bmeImageTable{height: auto !important; width:100% !important; padding:10px !important;clear:both; } }    \n@media only screen and (max-width: 480px){table.bmeCaptionTable{width:100% !important;  clear:both;} }    \n@media only screen and (max-width: 480px){table.bmeImageContainer{width:100% !important; clear:both; float:left !important;} }                \n@media only screen and (max-width: 480px){table.bmeImageTable td{padding:0px !important; height: auto; } }                \n@media only screen and (max-width: 480px){img.mobile-img-large{width:100% !important; height:auto !important;} }    \n@media only screen and (max-width: 480px){img.bmeRSSImage{max-width:320px; height:auto !important;} }    \n@media only screen and (min-width: 640px){img.bmeRSSImage{max-width:600px !important; height:auto !important;} }                \n@media only screen and (max-width: 480px){.trMargin img{height:10px;} }              \n@media only screen and (max-width: 480px){div.bmefooter, div.bmeheader{ display:block !important;} }      \n@media only screen and (max-width: 480px){.tdPart{ width:100% !important; clear:both; float:left !important; } }    \n@media only screen and (max-width: 480px){table.blk_parent1, table.tblPart {width: 100% !important; } }                \n@media only screen and (max-width: 480px){.tblLine{min-width: 100% !important;} }                \n@media only screen and (max-width: 480px){.bmeMblCenter img { margin: 0 auto; } }       \n@media only screen and (max-width: 480px){.bmeMblCenter, .bmeMblCenter div, .bmeMblCenter span  { text-align: center !important; text-align: -webkit-center !important; } }    \n@media only screen and (max-width: 480px){.bmeNoBr br, .bmeImageGutterRow, .bmeMblStackCenter .bmeShareItem .tdMblHide { display: none !important; } }    \n@media only screen and (max-width: 480px){.bmeMblInline table.bmeImageTable, .bmeMblInline table.bmeCaptionTable, .bmeMblInline { clear: none !important; width:50% !important; } }    \n@media only screen and (max-width: 480px){.bmeMblInlineHide, .bmeShareItem .trMargin { display: none !important; } }    \n@media only screen and (max-width: 480px){.bmeMblInline table.bmeImageTable img, .bmeMblShareCenter.tblContainer.mblSocialContain, .bmeMblFollowCenter.tblContainer.mblSocialContain{width: 100% !important; } }    \n@media only screen and (max-width: 480px){.bmeMblStack> .bmeShareItem{width: 100% !important; clear: both !important;} }    \n@media only screen and (max-width: 480px){.bmeShareItem{padding-top: 10px !important;} }                \n@media only screen and (max-width: 480px){.tdPart.bmeMblStackCenter, .bmeMblStackCenter .bmeFollowItemIcon {padding:0px !important; text-align: center !important;} }    \n@media only screen and (max-width: 480px){.bmeMblStackCenter> .bmeShareItem{width: 100% !important;} }    \n@media only screen and (max-width: 480px){ td.bmeMblCenter {border: 0 none transparent !important;} }    \n@media only screen and (max-width: 480px){.bmeLinkTable.tdPart td{padding-left:0px !important; padding-right:0px !important; border:0px none transparent !important;padding-bottom:15px !important;height: auto !important;} }    \n@media only screen and (max-width: 480px){.tdMblHide{width:10px !important;} }                \n@media only screen and (max-width: 480px){.bmeShareItemBtn{display:table !important;} }    \n@media only screen and (max-width: 480px){.bmeMblStack td {text-align: left !important;} }    \n@media only screen and (max-width: 480px){.bmeMblStack .bmeFollowItem{clear:both !important; padding-top: 10px !important;} }    \n@media only screen and (max-width: 480px){.bmeMblStackCenter .bmeFollowItemText{padding-left: 5px !important;} }    \n@media only screen and (max-width: 480px){.bmeMblStackCenter .bmeFollowItem{clear:both !important;align-self:center; float:none !important; padding-top:10px;margin: 0 auto;} }                \n@media only screen and (max-width: 480px){    \n.tdPart> table{width:100% !important;}    \n}    \n@media only screen and (max-width: 480px){.tdPart>table.bmeLinkContainer{ width:auto !important; } }    \n@media only screen and (max-width: 480px){.tdPart.mblStackCenter>table.bmeLinkContainer{ width:100% !important;} }     \n.blk_parent:first-child, .blk_parent{float:left;}                \n.blk_parent:last-child{float:right;}    \n    \n</style>\n<table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" name=\"bmeMainBody\" style=\"background-color: rgb(255, 255, 255);\" bgcolor=\"#ffffff\">\n  <tbody>\n    <tr>\n      <td width=\"100%\" valign=\"top\" align=\"center\"><table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" name=\"bmeMainColumnParentTable\" style=\"border-collapse: separate; border-spacing: 0px;\">\n          <tbody>\n            <tr>\n              <td name=\"bmeMainColumnParent\" style=\"border: 0px none transparent; border-radius: 0px; border-collapse: separate; border-spacing: 0px; overflow: hidden;\"><table name=\"bmeMainColumn\" class=\"bmeMainColumn bmeHolder\" style=\"max-width: 600px; overflow: visible; border-radius: 0px; border-collapse: separate; border-spacing: 0px; border-color: transparent; border-width: 0px; border-style: none;\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" align=\"center\">\n                  <tbody>\n                    <tr id=\"section_1\" class=\"flexible-section\" data-columns=\"1\" data-section-type=\"bmePreHeader\">\n                      <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmePreHeader\" valign=\"top\" align=\"center\" style=\"color: rgb(102, 102, 102); border: 0px none transparent; background-color: rgb(253, 198, 137);\" bgcolor=\"#fdc689\"></td>\n                    </tr>\n                    <tr>\n                      <td width=\"100%\" class=\"bmeHolder\" valign=\"top\" align=\"center\" name=\"bmeMainContentParent\" style=\"border: 0px none transparent; border-radius: 0px; border-collapse: separate; border-spacing: 0px; overflow: visible;\"><table name=\"bmeMainContent\" style=\"border-radius: 0px; border-collapse: separate; border-spacing: 0px; border: 0px none transparent; overflow: visible;\" width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" align=\"center\">\n                          <tbody>\n                            <tr id=\"section_2\" class=\"flexible-section\" data-columns=\"1\">\n                              <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmeHeader\" valign=\"top\" align=\"center\" style=\"color: rgb(56, 56, 56); border: 0px none transparent; background-color: rgb(0, 118, 163);\" bgcolor=\"#0076a3\"><div id=\"dv_11\" class=\"blk_wrapper\">\n                                  <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"blk\" name=\"blk_imagecaption\" style=\"width: 600px;\">\n                                    <tbody>\n                                      <tr>\n                                        <td><table cellspacing=\"0\" cellpadding=\"0\" class=\"bmeCaptionContainer\" width=\"100%\" style=\"padding-left:20px; padding-right:20px; padding-top:20px; padding-bottom:20px;border-collapse:separate;\">\n                                            <tbody>\n                                              <tr>\n                                                <td class=\"bmeImageContainerRow\" valign=\"top\" gutter=\"10\"><table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" border=\"0\">\n                                                    <tbody>\n                                                      <tr>\n                                                        <td class=\"tdPart\" valign=\"top\"><table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"bmeImageContainer\" width=\"100%\" align=\"left\" style=\"float:left;\">\n                                                            <tbody>\n                                                              <tr>\n                                                                <td valign=\"top\" name=\"tdContainer\"><table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"bmeImageTable\" dimension=\"50%\" imgid=\"1\" style=\"float: left; width: 270px;\" align=\"left\" width=\"270\" height=\"63\">\n                                                                    <tbody>\n                                                                      <tr>\n                                                                        <td name=\"bmeImgHolder\" width=\"280\" align=\"left\" valign=\"middle\" class=\"bmeMblCenter\" height=\"63\">" +
		"<img src=\"https://hayneschen.oss-cn-shenzhen.aliyuncs.com/icon/" + food + ".png\" width=\"63.2\" style=\"max-width: 63.2px; display: block; border-radius: 0px;\" alt=\"\" data-radius=\"109\" data-original-max-width=\"234\" class=\"keep-custom-width\" data-customwidth=\"27\"></td>\n                                                                      </tr>\n                                                                    </tbody>\n                                                                  </table>\n                                                                  <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"bmeCaptionTable\" style=\"float: right; width: 270px; word-break: break-word;\" align=\"right\" width=\"270\">\n                                                                    <tbody>\n                                                                      <tr>\n                                                                        <td name=\"tblCell\" valign=\"top\" align=\"left\" style=\"font-family: Arial, Helvetica, sans-serif; font-size: 14px; font-weight: normal; color: rgb(56, 56, 56); text-align: left; word-break: break-word;\" class=\"bmeMblCenter tblCell\"><div style=\"line-height: 150%; text-align: right;\"><span style=\"font-family: Helvetica, Arial, sans-serif; font-size: 14px; color: #dfdfdf; line-height: 150%;\">" +
		u_name + "</span> <br>\n                                                                            <span style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #ab966c; line-height: 150%;\"><strong>" +
		u_no + "</strong></span> <br>\n                                                                            <span style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #dedede; line-height: 150%;\">" +
		u_time + "</span></div></td>\n                                                                      </tr>\n                                                                    </tbody>\n                                                                  </table></td>\n                                                              </tr>\n                                                            </tbody>\n                                                          </table></td>\n                                                      </tr>\n                                                    </tbody>\n                                                  </table></td>\n                                              </tr>\n                                            </tbody>\n                                          </table></td>\n                                      </tr>\n                                    </tbody>\n                                  </table>\n                                </div></td>\n                            </tr>\n                            <tr id=\"section_3\" class=\"flexible-section\" data-columns=\"1\">\n                              <td width=\"100%\" class=\"blk_container bmeHolder bmeBody\" name=\"bmeBody\" valign=\"top\" align=\"center\" style=\"color: rgb(56, 56, 56); border: 0px none transparent; background-color: rgb(219, 231, 246);\" bgcolor=\"#dbe7f6\"><div id=\"dv_3\" class=\"blk_wrapper\">\n                                  <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"blk\" name=\"blk_image\">\n                                    <tbody>\n                                      <tr>\n                                        <td><table width=\"100%\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\">\n                                            <tbody>\n                                              <tr>\n                                                <td align=\"center\" class=\"bmeImage\" style=\"border-collapse: collapse; padding: 0px;\">" +
		"<img src=\"https://hayneschen.oss-cn-shenzhen.aliyuncs.com/postcard/" + photo + ".jpg\" class=\"mobile-img-large\" width=\"600\" style=\"max-width: 1196px; display: block;\" alt=\"\" border=\"0\"></td>\n                                              </tr>\n                                            </tbody>\n                                          </table></td>\n                                      </tr>\n                                    </tbody>\n                                  </table>\n                                </div>\n                                <div id=\"dv_1\" class=\"blk_wrapper\" style=\"\">\n                                  <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"blk\" name=\"blk_text\">\n                                    <tbody>\n                                      <tr>\n                                        <td><table cellpadding=\"0\" cellspacing=\"0\" border=\"0\" width=\"100%\" class=\"bmeContainerRow\">\n                                            <tbody>\n                                              <tr>\n                                                <th valign=\"top\" align=\"center\" class=\"tdPart\"> <table cellspacing=\"0\" cellpadding=\"0\" border=\"0\" width=\"600\" name=\"tblText\" style=\"float:left; background-color:transparent;\" align=\"left\" class=\"tblText\">\n                                                    <tbody>\n                                                      <tr>\n                                                        <td valign=\"top\" align=\"left\" name=\"tblCell\" style=\"padding: 20px; font-family: Arial, Helvetica, sans-serif; font-size: 14px; font-weight: 400; color: rgb(56, 56, 56); text-align: left; word-break: break-word;\" class=\"tblCell\"><div style=\"line-height: 200%;\"><span style=\"font-size: 20px; font-family: Helvetica, Arial, sans-serif; color: #343434; line-height: 200%;\"><strong>" +
		t + "</strong></span> <br>\n                                                            <span style=\"font-size: 14px; font-family: Helvetica, Arial, sans-serif; color: #b0b0b0; line-height: 200%;\">" +
		"URL:" + url + "<br>" + "BusinessID:" + bussid + "<br>" + "State:" + st + "<br>" + "isChange:" + is + "<br>" + "UpdateTime:" + u_time + "</span></div></td>\n                                                      </tr>\n                                                    </tbody>\n                                                  </table></th>\n                                              </tr>\n                                            </tbody>\n                                          </table></td>\n                                      </tr>\n                                    </tbody>\n                                  </table>\n                                </div></td>\n                            </tr>\n                          </tbody>\n                        </table></td>\n                    </tr>\n                    <tr id=\"section_5\" class=\"flexible-section\" data-columns=\"1\" data-section-type=\"bmeFooter\">\n                      <td width=\"100%\" class=\"blk_container bmeHolder\" name=\"bmeFooter\" valign=\"top\" align=\"center\" style=\"color: rgb(102, 102, 102); border: 0px none transparent; background-color: rgb(125, 167, 217);\" bgcolor=\"#7da7d9\"><div id=\"dv_2\" class=\"blk_wrapper\" style=\"\">\n                          <table width=\"600\" cellspacing=\"0\" cellpadding=\"0\" border=\"0\" class=\"blk\" name=\"blk_footer\" style=\"\">\n                            <tbody>\n                              <tr>\n                                <td name=\"tblCell\" class=\"tblCell\" style=\"padding: 20px; word-break: break-word;\" valign=\"top\" align=\"left\"><table cellpadding=\"0\" cellspacing=\"0\" border=\"0\" width=\"100%\">\n                                    <tbody>\n                                      <tr align=\"center\">\n                                        <td style=\"color: #EFC10C\">Made with ❤ by HaynesChen.</td>\n                                      </tr>\n                                    </tbody>\n                                  </table></td>\n                              </tr>\n                            </tbody>\n                          </table>\n                        </div></td>\n                    </tr>\n                  </tbody>\n                </table></td>\n            </tr>\n          </tbody>\n        </table></td>\n    </tr>\n  </tbody>\n</table>\n</body>\n</html>"
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
	fmt.Println("Automatic temperature registration by Haynes v1.6.5")
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
