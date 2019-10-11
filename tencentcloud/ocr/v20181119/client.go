// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20181119

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-11-19"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewArithmeticOCRRequest() (request *ArithmeticOCRRequest) {
    request = &ArithmeticOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "ArithmeticOCR")
    return
}

func NewArithmeticOCRResponse() (response *ArithmeticOCRResponse) {
    response = &ArithmeticOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持作业算式题目的自动识别，目前覆盖 K12 学力范围内的 14 种题型，包括加减乘除四则运算、分数四则运算、竖式四则运算、脱式计算等。
func (c *Client) ArithmeticOCR(request *ArithmeticOCRRequest) (response *ArithmeticOCRResponse, err error) {
    if request == nil {
        request = NewArithmeticOCRRequest()
    }
    response = NewArithmeticOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBankCardOCRRequest() (request *BankCardOCRRequest) {
    request = &BankCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BankCardOCR")
    return
}

func NewBankCardOCRResponse() (response *BankCardOCRResponse) {
    response = &BankCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对中国大陆主流银行卡的卡号、银行信息、有效期等关键字段的检测与识别。
func (c *Client) BankCardOCR(request *BankCardOCRRequest) (response *BankCardOCRResponse, err error) {
    if request == nil {
        request = NewBankCardOCRRequest()
    }
    response = NewBankCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBizLicenseOCRRequest() (request *BizLicenseOCRRequest) {
    request = &BizLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BizLicenseOCR")
    return
}

func NewBizLicenseOCRResponse() (response *BizLicenseOCRResponse) {
    response = &BizLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持快速精准识别营业执照上的字段，包括注册号、公司名称、经营场所、主体类型、法定代表人、注册资金、组成形式、成立日期、营业期限和经营范围等字段。
func (c *Client) BizLicenseOCR(request *BizLicenseOCRRequest) (response *BizLicenseOCRResponse, err error) {
    if request == nil {
        request = NewBizLicenseOCRRequest()
    }
    response = NewBizLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBusInvoiceOCRRequest() (request *BusInvoiceOCRRequest) {
    request = &BusInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BusInvoiceOCR")
    return
}

func NewBusInvoiceOCRResponse() (response *BusInvoiceOCRResponse) {
    response = &BusInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持识别公路汽车客票的发票代码、发票号码、日期、姓名、票价等字段。
func (c *Client) BusInvoiceOCR(request *BusInvoiceOCRRequest) (response *BusInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewBusInvoiceOCRRequest()
    }
    response = NewBusInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewBusinessCardOCRRequest() (request *BusinessCardOCRRequest) {
    request = &BusinessCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "BusinessCardOCR")
    return
}

func NewBusinessCardOCRResponse() (response *BusinessCardOCRResponse) {
    response = &BusinessCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持名片各字段的自动定位与识别，包含姓名、电话、手机号、邮箱、公司、部门、职位、网址、地址、QQ、微信、MSN等。
func (c *Client) BusinessCardOCR(request *BusinessCardOCRRequest) (response *BusinessCardOCRResponse, err error) {
    if request == nil {
        request = NewBusinessCardOCRRequest()
    }
    response = NewBusinessCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewCarInvoiceOCRRequest() (request *CarInvoiceOCRRequest) {
    request = &CarInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "CarInvoiceOCR")
    return
}

func NewCarInvoiceOCRResponse() (response *CarInvoiceOCRResponse) {
    response = &CarInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持机动车销售统一发票和二手车销售统一发票的识别，包括发票号码、发票代码、合计金额、合计税额等二十多个字段。
func (c *Client) CarInvoiceOCR(request *CarInvoiceOCRRequest) (response *CarInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewCarInvoiceOCRRequest()
    }
    response = NewCarInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewDriverLicenseOCRRequest() (request *DriverLicenseOCRRequest) {
    request = &DriverLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "DriverLicenseOCR")
    return
}

func NewDriverLicenseOCRResponse() (response *DriverLicenseOCRResponse) {
    response = &DriverLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对驾驶证主页所有字段的自动定位与识别，包含证号、姓名、性别、国籍、住址、出生日期、初次领证日期、准驾车型、有效期限等。
func (c *Client) DriverLicenseOCR(request *DriverLicenseOCRRequest) (response *DriverLicenseOCRResponse, err error) {
    if request == nil {
        request = NewDriverLicenseOCRRequest()
    }
    response = NewDriverLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewDutyPaidProofOCRRequest() (request *DutyPaidProofOCRRequest) {
    request = &DutyPaidProofOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "DutyPaidProofOCR")
    return
}

func NewDutyPaidProofOCRResponse() (response *DutyPaidProofOCRResponse) {
    response = &DutyPaidProofOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对完税证明的税号、纳税人识别号、纳税人名称、金额合计大写、金额合计小写、填发日期、税务机关、填票人等关键字段的识别。
func (c *Client) DutyPaidProofOCR(request *DutyPaidProofOCRRequest) (response *DutyPaidProofOCRResponse, err error) {
    if request == nil {
        request = NewDutyPaidProofOCRRequest()
    }
    response = NewDutyPaidProofOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEduPaperOCRRequest() (request *EduPaperOCRRequest) {
    request = &EduPaperOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "EduPaperOCR")
    return
}

func NewEduPaperOCRResponse() (response *EduPaperOCRResponse) {
    response = &EduPaperOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持数学试题内容的识别和结构化输出，包括通用文本解析和小学/初中/高中数学公式解析能力（包括91种题型，180种符号）。
func (c *Client) EduPaperOCR(request *EduPaperOCRRequest) (response *EduPaperOCRResponse, err error) {
    if request == nil {
        request = NewEduPaperOCRRequest()
    }
    response = NewEduPaperOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEnglishOCRRequest() (request *EnglishOCRRequest) {
    request = &EnglishOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "EnglishOCR")
    return
}

func NewEnglishOCRResponse() (response *EnglishOCRResponse) {
    response = &EnglishOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图像英文文字的检测和识别，返回文字框位置与文字内容。支持多场景、任意版面下的英文、字母、数字和常见字符的识别，同时覆盖英文印刷体和英文手写体识别。
func (c *Client) EnglishOCR(request *EnglishOCRRequest) (response *EnglishOCRResponse, err error) {
    if request == nil {
        request = NewEnglishOCRRequest()
    }
    response = NewEnglishOCRResponse()
    err = c.Send(request, response)
    return
}

func NewEstateCertOCRRequest() (request *EstateCertOCRRequest) {
    request = &EstateCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "EstateCertOCR")
    return
}

func NewEstateCertOCRResponse() (response *EstateCertOCRResponse) {
    response = &EstateCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持不动产权证关键字段的识别，包括使用期限、面积、用途、权利性质、权利类型、坐落、共有情况、权利人、权利其他状况等。
// 
// 
func (c *Client) EstateCertOCR(request *EstateCertOCRRequest) (response *EstateCertOCRResponse, err error) {
    if request == nil {
        request = NewEstateCertOCRRequest()
    }
    response = NewEstateCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFinanBillOCRRequest() (request *FinanBillOCRRequest) {
    request = &FinanBillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "FinanBillOCR")
    return
}

func NewFinanBillOCRResponse() (response *FinanBillOCRResponse) {
    response = &FinanBillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持常见银行票据的自动分类和识别。整单识别包括支票（含现金支票、普通支票、转账支票），承兑汇票（含银行承兑汇票、商业承兑汇票）以及进账单等，适用于中国人民银行印发的 2010 版银行票据凭证版式（银发[2010]299 号）。
func (c *Client) FinanBillOCR(request *FinanBillOCRRequest) (response *FinanBillOCRResponse, err error) {
    if request == nil {
        request = NewFinanBillOCRRequest()
    }
    response = NewFinanBillOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFinanBillSliceOCRRequest() (request *FinanBillSliceOCRRequest) {
    request = &FinanBillSliceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "FinanBillSliceOCR")
    return
}

func NewFinanBillSliceOCRResponse() (response *FinanBillSliceOCRResponse) {
    response = &FinanBillSliceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持常见银行票据的自动分类和识别。切片识别包括金融行业常见票据的重要切片字段识别，包括金额、账号、日期、凭证号码等。（金融票据切片：金融票据中待识别字段及其周围局部区域的裁剪图像。）
func (c *Client) FinanBillSliceOCR(request *FinanBillSliceOCRRequest) (response *FinanBillSliceOCRResponse, err error) {
    if request == nil {
        request = NewFinanBillSliceOCRRequest()
    }
    response = NewFinanBillSliceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFlightInvoiceOCRRequest() (request *FlightInvoiceOCRRequest) {
    request = &FlightInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "FlightInvoiceOCR")
    return
}

func NewFlightInvoiceOCRResponse() (response *FlightInvoiceOCRResponse) {
    response = &FlightInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持机票行程单关键字段的识别，包括姓名、身份证件号码、航班号、票价 、合计、电子客票号码、填开日期等。
func (c *Client) FlightInvoiceOCR(request *FlightInvoiceOCRRequest) (response *FlightInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewFlightInvoiceOCRRequest()
    }
    response = NewFlightInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewFormulaOCRRequest() (request *FormulaOCRRequest) {
    request = &FormulaOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "FormulaOCR")
    return
}

func NewFormulaOCRResponse() (response *FormulaOCRResponse) {
    response = &FormulaOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持识别主流初高中数学符号和公式，返回公式的 Latex 格式文本。
func (c *Client) FormulaOCR(request *FormulaOCRRequest) (response *FormulaOCRResponse, err error) {
    if request == nil {
        request = NewFormulaOCRRequest()
    }
    response = NewFormulaOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralAccurateOCRRequest() (request *GeneralAccurateOCRRequest) {
    request = &GeneralAccurateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralAccurateOCR")
    return
}

func NewGeneralAccurateOCRResponse() (response *GeneralAccurateOCRResponse) {
    response = &GeneralAccurateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图像整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，准确率和召回率更高。
func (c *Client) GeneralAccurateOCR(request *GeneralAccurateOCRRequest) (response *GeneralAccurateOCRResponse, err error) {
    if request == nil {
        request = NewGeneralAccurateOCRRequest()
    }
    response = NewGeneralAccurateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralBasicOCRRequest() (request *GeneralBasicOCRRequest) {
    request = &GeneralBasicOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralBasicOCR")
    return
}

func NewGeneralBasicOCRResponse() (response *GeneralBasicOCRResponse) {
    response = &GeneralBasicOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持多场景、任意版面下整图文字的识别。支持自动识别语言类型，同时支持自选语言种类（推荐），除中英文外，支持日语、韩语、西班牙语、法语、德语、葡萄牙语、越南语、马来语、俄语、意大利语、荷兰语、瑞典语、芬兰语、丹麦语、挪威语、匈牙利语、泰语等多种语言。应用场景包括：印刷文档识别、网络图片识别、广告图文字识别、街景店招识别、菜单识别、视频标题识别、头像文字识别等。
func (c *Client) GeneralBasicOCR(request *GeneralBasicOCRRequest) (response *GeneralBasicOCRResponse, err error) {
    if request == nil {
        request = NewGeneralBasicOCRRequest()
    }
    response = NewGeneralBasicOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralEfficientOCRRequest() (request *GeneralEfficientOCRRequest) {
    request = &GeneralEfficientOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralEfficientOCR")
    return
}

func NewGeneralEfficientOCRResponse() (response *GeneralEfficientOCRResponse) {
    response = &GeneralEfficientOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持多场景、任意版面下整图文字的识别。相较于“通用印刷体识别”接口，精简版接口在准召率有一定损失的情况下，耗时更短。适用于对接口耗时较为敏感的客户。
func (c *Client) GeneralEfficientOCR(request *GeneralEfficientOCRRequest) (response *GeneralEfficientOCRResponse, err error) {
    if request == nil {
        request = NewGeneralEfficientOCRRequest()
    }
    response = NewGeneralEfficientOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralFastOCRRequest() (request *GeneralFastOCRRequest) {
    request = &GeneralFastOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralFastOCR")
    return
}

func NewGeneralFastOCRResponse() (response *GeneralFastOCRResponse) {
    response = &GeneralFastOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片中整体文字的检测和识别，返回文字框位置与文字内容。相比通用印刷体识别接口，识别速度更快、支持的 QPS 更高。
func (c *Client) GeneralFastOCR(request *GeneralFastOCRRequest) (response *GeneralFastOCRResponse, err error) {
    if request == nil {
        request = NewGeneralFastOCRRequest()
    }
    response = NewGeneralFastOCRResponse()
    err = c.Send(request, response)
    return
}

func NewGeneralHandwritingOCRRequest() (request *GeneralHandwritingOCRRequest) {
    request = &GeneralHandwritingOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "GeneralHandwritingOCR")
    return
}

func NewGeneralHandwritingOCRResponse() (response *GeneralHandwritingOCRResponse) {
    response = &GeneralHandwritingOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内手写体文字的检测和识别，针对手写字体无规则、字迹潦草、模糊等特点进行了识别能力的增强。
func (c *Client) GeneralHandwritingOCR(request *GeneralHandwritingOCRRequest) (response *GeneralHandwritingOCRResponse, err error) {
    if request == nil {
        request = NewGeneralHandwritingOCRRequest()
    }
    response = NewGeneralHandwritingOCRResponse()
    err = c.Send(request, response)
    return
}

func NewIDCardOCRRequest() (request *IDCardOCRRequest) {
    request = &IDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "IDCardOCR")
    return
}

func NewIDCardOCRResponse() (response *IDCardOCRResponse) {
    response = &IDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持中国大陆居民二代身份证正反面所有字段的识别，包括姓名、性别、民族、出生日期、住址、公民身份证号、签发机关、有效期限；具备身份证照片、人像照片的裁剪功能和翻拍、PS、复印件告警功能，以及边框和框内遮挡告警、临时身份证告警和身份证有效期不合法告警等扩展功能。
func (c *Client) IDCardOCR(request *IDCardOCRRequest) (response *IDCardOCRResponse, err error) {
    if request == nil {
        request = NewIDCardOCRRequest()
    }
    response = NewIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewInstitutionOCRRequest() (request *InstitutionOCRRequest) {
    request = &InstitutionOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "InstitutionOCR")
    return
}

func NewInstitutionOCRResponse() (response *InstitutionOCRResponse) {
    response = &InstitutionOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持事业单位法人证书关键字段识别，包括注册号、有效期、住所、名称、法定代表人等。
func (c *Client) InstitutionOCR(request *InstitutionOCRRequest) (response *InstitutionOCRResponse, err error) {
    if request == nil {
        request = NewInstitutionOCRRequest()
    }
    response = NewInstitutionOCRResponse()
    err = c.Send(request, response)
    return
}

func NewInvoiceGeneralOCRRequest() (request *InvoiceGeneralOCRRequest) {
    request = &InvoiceGeneralOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "InvoiceGeneralOCR")
    return
}

func NewInvoiceGeneralOCRResponse() (response *InvoiceGeneralOCRResponse) {
    response = &InvoiceGeneralOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对通用机打发票的发票代码、发票号码、日期、购买方识别号、销售方识别号、校验码、小写金额等关键字段的识别。
func (c *Client) InvoiceGeneralOCR(request *InvoiceGeneralOCRRequest) (response *InvoiceGeneralOCRResponse, err error) {
    if request == nil {
        request = NewInvoiceGeneralOCRRequest()
    }
    response = NewInvoiceGeneralOCRResponse()
    err = c.Send(request, response)
    return
}

func NewLicensePlateOCRRequest() (request *LicensePlateOCRRequest) {
    request = &LicensePlateOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "LicensePlateOCR")
    return
}

func NewLicensePlateOCRResponse() (response *LicensePlateOCRResponse) {
    response = &LicensePlateOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对中国大陆机动车车牌的自动定位和识别，返回地域编号和车牌号信息。
func (c *Client) LicensePlateOCR(request *LicensePlateOCRRequest) (response *LicensePlateOCRResponse, err error) {
    if request == nil {
        request = NewLicensePlateOCRRequest()
    }
    response = NewLicensePlateOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDCardOCRRequest() (request *MLIDCardOCRRequest) {
    request = &MLIDCardOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDCardOCR")
    return
}

func NewMLIDCardOCRResponse() (response *MLIDCardOCRResponse) {
    response = &MLIDCardOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持马来西亚身份证识别，识别字段包括身份证号、姓名、性别、地址；具备身份证人像照片的裁剪功能和翻拍、复印件告警功能。
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
func (c *Client) MLIDCardOCR(request *MLIDCardOCRRequest) (response *MLIDCardOCRResponse, err error) {
    if request == nil {
        request = NewMLIDCardOCRRequest()
    }
    response = NewMLIDCardOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMLIDPassportOCRRequest() (request *MLIDPassportOCRRequest) {
    request = &MLIDPassportOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "MLIDPassportOCR")
    return
}

func NewMLIDPassportOCRResponse() (response *MLIDPassportOCRResponse) {
    response = &MLIDPassportOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持马来西亚身护照识别，识别字段包括护照ID、姓名、出生日期、性别、有效期、发行国、国籍；具备护照人像照片的裁剪功能和翻拍、复印件告警功能。
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
func (c *Client) MLIDPassportOCR(request *MLIDPassportOCRRequest) (response *MLIDPassportOCRResponse, err error) {
    if request == nil {
        request = NewMLIDPassportOCRRequest()
    }
    response = NewMLIDPassportOCRResponse()
    err = c.Send(request, response)
    return
}

func NewMixedInvoiceDetectRequest() (request *MixedInvoiceDetectRequest) {
    request = &MixedInvoiceDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "MixedInvoiceDetect")
    return
}

func NewMixedInvoiceDetectResponse() (response *MixedInvoiceDetectResponse) {
    response = &MixedInvoiceDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持多张、多类型票据的混合检测和自动分类，返回对应票据类型。目前已支持增值税发票、增值税发票（卷票）、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票、酒店账单、客运限额发票、购物小票、完税证明共15种票据。
func (c *Client) MixedInvoiceDetect(request *MixedInvoiceDetectRequest) (response *MixedInvoiceDetectResponse, err error) {
    if request == nil {
        request = NewMixedInvoiceDetectRequest()
    }
    response = NewMixedInvoiceDetectResponse()
    err = c.Send(request, response)
    return
}

func NewMixedInvoiceOCRRequest() (request *MixedInvoiceOCRRequest) {
    request = &MixedInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "MixedInvoiceOCR")
    return
}

func NewMixedInvoiceOCRResponse() (response *MixedInvoiceOCRResponse) {
    response = &MixedInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持多张、多类型票据的混合识别，系统自动实现分割、分类和识别。目前已支持增值税发票、增值税发票（卷票）、定额发票、通用机打发票、购车发票、火车票、出租车发票、机票行程单、汽车票、轮船票、过路过桥费发票共11种票据。
func (c *Client) MixedInvoiceOCR(request *MixedInvoiceOCRRequest) (response *MixedInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewMixedInvoiceOCRRequest()
    }
    response = NewMixedInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewOrgCodeCertOCRRequest() (request *OrgCodeCertOCRRequest) {
    request = &OrgCodeCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "OrgCodeCertOCR")
    return
}

func NewOrgCodeCertOCRResponse() (response *OrgCodeCertOCRResponse) {
    response = &OrgCodeCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持组织机构代码证关键字段的识别，包括代码、有效期、地址、机构名称等。
func (c *Client) OrgCodeCertOCR(request *OrgCodeCertOCRRequest) (response *OrgCodeCertOCRResponse, err error) {
    if request == nil {
        request = NewOrgCodeCertOCRRequest()
    }
    response = NewOrgCodeCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPassportOCRRequest() (request *PassportOCRRequest) {
    request = &PassportOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "PassportOCR")
    return
}

func NewPassportOCRResponse() (response *PassportOCRResponse) {
    response = &PassportOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持中国大陆护照、中国香港护照、泰国护照及其他国外护照个人资料页多个字段的检测与识别。其中中国大陆居民护照识别，已支持字段包括英文姓名、中文姓名、国家码、护照号、出生地、出生日期、国籍英文、性别英文、有效期、签发地点英文、签发日期、持证人签名、护照机读码（MRZ码）等。中国香港护照、泰国护照及其他国外护照识别，已支持字段包括英文姓名、国籍、签发日期、性别、护照号码等。
func (c *Client) PassportOCR(request *PassportOCRRequest) (response *PassportOCRResponse, err error) {
    if request == nil {
        request = NewPassportOCRRequest()
    }
    response = NewPassportOCRResponse()
    err = c.Send(request, response)
    return
}

func NewPermitOCRRequest() (request *PermitOCRRequest) {
    request = &PermitOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "PermitOCR")
    return
}

func NewPermitOCRResponse() (response *PermitOCRResponse) {
    response = &PermitOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对卡式港澳台通行证的识别，包括签发地点、签发机关、有效期限、性别、出生日期、英文姓名、姓名、证件号等字段。
func (c *Client) PermitOCR(request *PermitOCRRequest) (response *PermitOCRResponse, err error) {
    if request == nil {
        request = NewPermitOCRRequest()
    }
    response = NewPermitOCRResponse()
    err = c.Send(request, response)
    return
}

func NewQrcodeOCRRequest() (request *QrcodeOCRRequest) {
    request = &QrcodeOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "QrcodeOCR")
    return
}

func NewQrcodeOCRResponse() (response *QrcodeOCRResponse) {
    response = &QrcodeOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持条形码和二维码的识别（包括 DataMatrix 和 PDF417）。
// 本接口暂未完全对外开放，如需咨询，请[联系商务](https://cloud.tencent.com/about/connect)
func (c *Client) QrcodeOCR(request *QrcodeOCRRequest) (response *QrcodeOCRResponse, err error) {
    if request == nil {
        request = NewQrcodeOCRRequest()
    }
    response = NewQrcodeOCRResponse()
    err = c.Send(request, response)
    return
}

func NewQuotaInvoiceOCRRequest() (request *QuotaInvoiceOCRRequest) {
    request = &QuotaInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "QuotaInvoiceOCR")
    return
}

func NewQuotaInvoiceOCRResponse() (response *QuotaInvoiceOCRResponse) {
    response = &QuotaInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持定额发票的发票号码、发票代码及金额等关键字段的识别。
func (c *Client) QuotaInvoiceOCR(request *QuotaInvoiceOCRRequest) (response *QuotaInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewQuotaInvoiceOCRRequest()
    }
    response = NewQuotaInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewShipInvoiceOCRRequest() (request *ShipInvoiceOCRRequest) {
    request = &ShipInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "ShipInvoiceOCR")
    return
}

func NewShipInvoiceOCRResponse() (response *ShipInvoiceOCRResponse) {
    response = &ShipInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持识别轮船票的发票代码、发票号码、日期、姓名、票价等字段。
func (c *Client) ShipInvoiceOCR(request *ShipInvoiceOCRRequest) (response *ShipInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewShipInvoiceOCRRequest()
    }
    response = NewShipInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTableOCRRequest() (request *TableOCRRequest) {
    request = &TableOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TableOCR")
    return
}

func NewTableOCRResponse() (response *TableOCRResponse) {
    response = &TableOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内表格文档的检测和识别，返回每个单元格的文字内容，支持将识别结果保存为 Excel 格式。
func (c *Client) TableOCR(request *TableOCRRequest) (response *TableOCRResponse, err error) {
    if request == nil {
        request = NewTableOCRRequest()
    }
    response = NewTableOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTaxiInvoiceOCRRequest() (request *TaxiInvoiceOCRRequest) {
    request = &TaxiInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TaxiInvoiceOCR")
    return
}

func NewTaxiInvoiceOCRResponse() (response *TaxiInvoiceOCRResponse) {
    response = &TaxiInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持出租车发票关键字段的识别，包括发票号码、发票代码、金额、日期等字段。
func (c *Client) TaxiInvoiceOCR(request *TaxiInvoiceOCRRequest) (response *TaxiInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewTaxiInvoiceOCRRequest()
    }
    response = NewTaxiInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTextDetectRequest() (request *TextDetectRequest) {
    request = &TextDetectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TextDetect")
    return
}

func NewTextDetectResponse() (response *TextDetectResponse) {
    response = &TextDetectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口通过检测图片中的文字信息特征，快速判断图片中有无文字并返回判断结果，帮助用户过滤无文字的图片。
func (c *Client) TextDetect(request *TextDetectRequest) (response *TextDetectResponse, err error) {
    if request == nil {
        request = NewTextDetectRequest()
    }
    response = NewTextDetectResponse()
    err = c.Send(request, response)
    return
}

func NewTollInvoiceOCRRequest() (request *TollInvoiceOCRRequest) {
    request = &TollInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TollInvoiceOCR")
    return
}

func NewTollInvoiceOCRResponse() (response *TollInvoiceOCRResponse) {
    response = &TollInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对过路过桥费发票的发票代码、发票号码、日期、小写金额等关键字段的识别。
func (c *Client) TollInvoiceOCR(request *TollInvoiceOCRRequest) (response *TollInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewTollInvoiceOCRRequest()
    }
    response = NewTollInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewTrainTicketOCRRequest() (request *TrainTicketOCRRequest) {
    request = &TrainTicketOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "TrainTicketOCR")
    return
}

func NewTrainTicketOCRResponse() (response *TrainTicketOCRResponse) {
    response = &TrainTicketOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持火车票全字段的识别，包括编号、票价、姓名、座位号、出发时间、出发站、到达站、车次、席别等。
func (c *Client) TrainTicketOCR(request *TrainTicketOCRRequest) (response *TrainTicketOCRResponse, err error) {
    if request == nil {
        request = NewTrainTicketOCRRequest()
    }
    response = NewTrainTicketOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVatInvoiceOCRRequest() (request *VatInvoiceOCRRequest) {
    request = &VatInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VatInvoiceOCR")
    return
}

func NewVatInvoiceOCRResponse() (response *VatInvoiceOCRResponse) {
    response = &VatInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持增值税专用发票、增值税普通发票、增值税电子发票全字段的内容检测和识别，包括发票代码、发票号码、开票日期、合计金额、校验码、税率、合计税额、价税合计、购买方识别号、复核、销售方识别号、开票人、密码区1、密码区2、密码区3、密码区4、发票名称、购买方名称、销售方名称、服务名称、备注、规格型号、数量、单价、金额、税额、收款人等字段。
func (c *Client) VatInvoiceOCR(request *VatInvoiceOCRRequest) (response *VatInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVatInvoiceOCRRequest()
    }
    response = NewVatInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVatRollInvoiceOCRRequest() (request *VatRollInvoiceOCRRequest) {
    request = &VatRollInvoiceOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VatRollInvoiceOCR")
    return
}

func NewVatRollInvoiceOCRResponse() (response *VatRollInvoiceOCRResponse) {
    response = &VatRollInvoiceOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持对增值税发票（卷票）的发票代码、发票号码、日期、校验码、合计金额（小写）等关键字段的识别。
func (c *Client) VatRollInvoiceOCR(request *VatRollInvoiceOCRRequest) (response *VatRollInvoiceOCRResponse, err error) {
    if request == nil {
        request = NewVatRollInvoiceOCRRequest()
    }
    response = NewVatRollInvoiceOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVehicleLicenseOCRRequest() (request *VehicleLicenseOCRRequest) {
    request = &VehicleLicenseOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VehicleLicenseOCR")
    return
}

func NewVehicleLicenseOCRResponse() (response *VehicleLicenseOCRResponse) {
    response = &VehicleLicenseOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持行驶证主页和副页所有字段的自动定位与识别，包含车牌号码、车辆类型、所有人、住址、使用性质、品牌型号、车辆识别代码、发动机号、注册日期、发证日期等。
func (c *Client) VehicleLicenseOCR(request *VehicleLicenseOCRRequest) (response *VehicleLicenseOCRResponse, err error) {
    if request == nil {
        request = NewVehicleLicenseOCRRequest()
    }
    response = NewVehicleLicenseOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVehicleRegCertOCRRequest() (request *VehicleRegCertOCRRequest) {
    request = &VehicleRegCertOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VehicleRegCertOCR")
    return
}

func NewVehicleRegCertOCRResponse() (response *VehicleRegCertOCRResponse) {
    response = &VehicleRegCertOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持国内机动车登记证书主要字段的结构化识别，包括机动车所有人、身份证明名称、号码、车辆型号、车辆识别代号、发动机号、制造厂名称等。
func (c *Client) VehicleRegCertOCR(request *VehicleRegCertOCRRequest) (response *VehicleRegCertOCRResponse, err error) {
    if request == nil {
        request = NewVehicleRegCertOCRRequest()
    }
    response = NewVehicleRegCertOCRResponse()
    err = c.Send(request, response)
    return
}

func NewVinOCRRequest() (request *VinOCRRequest) {
    request = &VinOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "VinOCR")
    return
}

func NewVinOCRResponse() (response *VinOCRResponse) {
    response = &VinOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持图片内车辆识别代号（VIN）的检测和识别。
func (c *Client) VinOCR(request *VinOCRRequest) (response *VinOCRResponse, err error) {
    if request == nil {
        request = NewVinOCRRequest()
    }
    response = NewVinOCRResponse()
    err = c.Send(request, response)
    return
}

func NewWaybillOCRRequest() (request *WaybillOCRRequest) {
    request = &WaybillOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ocr", APIVersion, "WaybillOCR")
    return
}

func NewWaybillOCRResponse() (response *WaybillOCRResponse) {
    response = &WaybillOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口支持市面上主流版式电子运单的识别，包括收件人和寄件人的姓名、电话、地址以及运单号等字段。
func (c *Client) WaybillOCR(request *WaybillOCRRequest) (response *WaybillOCRResponse, err error) {
    if request == nil {
        request = NewWaybillOCRRequest()
    }
    response = NewWaybillOCRResponse()
    err = c.Send(request, response)
    return
}
