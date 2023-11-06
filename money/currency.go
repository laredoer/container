package money

type Currency string

// Constants for active currency codes according to the ISO 4217 standard.
const (
	AED Currency = "AED"
	AFN Currency = "AFN"
	ALL Currency = "ALL"
	AMD Currency = "AMD"
	ANG Currency = "ANG"
	AOA Currency = "AOA"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AWG Currency = "AWG"
	AZN Currency = "AZN"
	BAM Currency = "BAM"
	BBD Currency = "BBD"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	BHD Currency = "BHD"
	BIF Currency = "BIF"
	BMD Currency = "BMD"
	BND Currency = "BND"
	BOB Currency = "BOB"
	BRL Currency = "BRL"
	BSD Currency = "BSD"
	BTN Currency = "BTN"
	BWP Currency = "BWP"
	BYN Currency = "BYN"
	BYR Currency = "BYR"
	BZD Currency = "BZD"
	CAD Currency = "CAD"
	CDF Currency = "CDF"
	CHF Currency = "CHF"
	CLF Currency = "CLF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	CNH Currency = "CNH"
	COP Currency = "COP"
	CRC Currency = "CRC"
	CUC Currency = "CUC"
	CUP Currency = "CUP"
	CVE Currency = "CVE"
	CZK Currency = "CZK"
	DJF Currency = "DJF"
	DKK Currency = "DKK"
	DOP Currency = "DOP"
	DZD Currency = "DZD"
	EEK Currency = "EEK"
	EGP Currency = "EGP"
	ERN Currency = "ERN"
	ETB Currency = "ETB"
	EUR Currency = "EUR"
	FJD Currency = "FJD"
	FKP Currency = "FKP"
	GBP Currency = "GBP"
	GEL Currency = "GEL"
	GGP Currency = "GGP"
	GHC Currency = "GHC"
	GHS Currency = "GHS"
	GIP Currency = "GIP"
	GMD Currency = "GMD"
	GNF Currency = "GNF"
	GTQ Currency = "GTQ"
	GYD Currency = "GYD"
	HKD Currency = "HKD"
	HNL Currency = "HNL"
	HRK Currency = "HRK"
	HTG Currency = "HTG"
	HUF Currency = "HUF"
	IDR Currency = "IDR"
	ILS Currency = "ILS"
	IMP Currency = "IMP"
	INR Currency = "INR"
	IQD Currency = "IQD"
	IRR Currency = "IRR"
	ISK Currency = "ISK"
	JEP Currency = "JEP"
	JMD Currency = "JMD"
	JOD Currency = "JOD"
	JPY Currency = "JPY"
	KES Currency = "KES"
	KGS Currency = "KGS"
	KHR Currency = "KHR"
	KMF Currency = "KMF"
	KPW Currency = "KPW"
	KRW Currency = "KRW"
	KWD Currency = "KWD"
	KYD Currency = "KYD"
	KZT Currency = "KZT"
	LAK Currency = "LAK"
	LBP Currency = "LBP"
	LKR Currency = "LKR"
	LRD Currency = "LRD"
	LSL Currency = "LSL"
	LTL Currency = "LTL"
	LVL Currency = "LVL"
	LYD Currency = "LYD"
	MAD Currency = "MAD"
	MDL Currency = "MDL"
	MGA Currency = "MGA"
	MKD Currency = "MKD"
	MMK Currency = "MMK"
	MNT Currency = "MNT"
	MOP Currency = "MOP"
	MUR Currency = "MUR"
	MVR Currency = "MVR"
	MWK Currency = "MWK"
	MXN Currency = "MXN"
	MYR Currency = "MYR"
	MZN Currency = "MZN"
	NAD Currency = "NAD"
	NGN Currency = "NGN"
	NIO Currency = "NIO"
	NOK Currency = "NOK"
	NPR Currency = "NPR"
	NZD Currency = "NZD"
	OMR Currency = "OMR"
	PAB Currency = "PAB"
	PEN Currency = "PEN"
	PGK Currency = "PGK"
	PHP Currency = "PHP"
	PKR Currency = "PKR"
	PLN Currency = "PLN"
	PYG Currency = "PYG"
	QAR Currency = "QAR"
	RON Currency = "RON"
	RSD Currency = "RSD"
	RUB Currency = "RUB"
	RUR Currency = "RUR"
	RWF Currency = "RWF"
	SAR Currency = "SAR"
	SBD Currency = "SBD"
	SCR Currency = "SCR"
	SDG Currency = "SDG"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	SHP Currency = "SHP"
	SKK Currency = "SKK"
	SLL Currency = "SLL"
	SOS Currency = "SOS"
	SRD Currency = "SRD"
	SSP Currency = "SSP"
	STD Currency = "STD"
	SVC Currency = "SVC"
	SYP Currency = "SYP"
	SZL Currency = "SZL"
	THB Currency = "THB"
	TJS Currency = "TJS"
	TMT Currency = "TMT"
	TND Currency = "TND"
	TOP Currency = "TOP"
	TRL Currency = "TRL"
	TRY Currency = "TRY"
	TTD Currency = "TTD"
	TWD Currency = "TWD"
	TZS Currency = "TZS"
	UAH Currency = "UAH"
	UGX Currency = "UGX"
	USD Currency = "USD"
	UYU Currency = "UYU"
	UZS Currency = "UZS"
	VEF Currency = "VEF"
	VND Currency = "VND"
	VUV Currency = "VUV"
	WST Currency = "WST"
	XAF Currency = "XAF"
	XAG Currency = "XAG"
	XAU Currency = "XAU"
	XCD Currency = "XCD"
	XDR Currency = "XDR"
	XOF Currency = "XOF"
	XPF Currency = "XPF"
	YER Currency = "YER"
	ZAR Currency = "ZAR"
	ZMW Currency = "ZMW"
	ZWD Currency = "ZWD"
)

func (c Currency) String() string {
	return string(c)
}

// 获取币种多语言
func (c Currency) Lang() map[string]string {

	res := make(map[string]string)

	switch c {
	case AED:
		res["en"] = "United Arab Emirates Dirham"
		res["zh-CN"] = "阿联酋迪拉姆"
		res["zh-HK"] = "阿聯酋迪拉姆"
	case AFN:
		res["en"] = "Afghan Afghani"
		res["zh-CN"] = "阿富汗阿尔及利亚法郎"
		res["zh-HK"] = "阿富汗阿爾及利亞法郎"
	case ALL:
		res["en"] = "Albanian Lek"
		res["zh-CN"] = "阿尔巴尼亚列克"
		res["zh-HK"] = "阿爾巴尼亞列克"
	case AMD:
		res["en"] = "Armenian Dram"
		res["zh-CN"] = "亚美尼亚德拉姆"
		res["zh-HK"] = "亞美尼亞德拉姆"
	case ANG:
		res["en"] = "Netherlands Antillean Guilder"
		res["zh-CN"] = "荷属安的列斯盾"
		res["zh-HK"] = "荷屬安的列斯盾"
	case AOA:
		res["en"] = "Angolan Kwanza"
		res["zh-CN"] = "安哥拉宽扎"
		res["zh-HK"] = "安哥拉寬紮"
	case ARS:
		res["en"] = "Argentine Peso"
		res["zh-CN"] = "阿根廷比索"
		res["zh-HK"] = "阿根廷比索"
	case AUD:
		res["en"] = "Australian Dollar"
		res["zh-CN"] = "澳大利亚元"
		res["zh-HK"] = "澳大利亞元"
	case AWG:
		res["en"] = "Aruban Florin"
		res["zh-CN"] = "阿鲁巴弗罗林"
		res["zh-HK"] = "阿魯巴弗羅林"
	case AZN:
		res["en"] = "Azerbaijani Manat"
		res["zh-CN"] = "阿塞拜疆马纳特"
		res["zh-HK"] = "阿塞拜疆馬納特"
	case BAM:
		res["en"] = "Bosnia-Herzegovina Convertible Mark"
		res["zh-CN"] = "波斯尼亚和黑塞哥维那元"
		res["zh-HK"] = "波斯尼亞和黑塞哥維那元"
	case BBD:
		res["en"] = "Barbados Dollar"
		res["zh-CN"] = "巴巴多斯元"
		res["zh-HK"] = "巴巴多斯元"
	case BDT:
		res["en"] = "Bangladeshi Taka"
		res["zh-CN"] = "孟加拉塔卡"
		res["zh-HK"] = "孟加拉塔卡"
	case BGN:
		res["en"] = "Bulgarian Lev"
		res["zh-CN"] = "保加利亚列弗"
		res["zh-HK"] = "保加利亞列弗"
	case BHD:
		res["en"] = "Bahraini Dinar"
		res["zh-CN"] = "巴林第纳尔"
		res["zh-HK"] = "巴林第納爾"
	case BIF:
		res["en"] = "Burundian Franc"
		res["zh-CN"] = "布隆迪法郎"
		res["zh-HK"] = "佈隆迪法郎"
	case BMD:
		res["en"] = "Bermudian Dollar"
		res["zh-CN"] = "百慕大元"
		res["zh-HK"] = "百慕大元"
	case BND:
		res["en"] = "Brunei Dollar"
		res["zh-CN"] = "文莱元"
		res["zh-HK"] = "文萊元"
	case BOB:
		res["en"] = "Bolivian Boliviano"
		res["zh-CN"] = "玻利维亚诺"
		res["zh-HK"] = "玻利維亞諾"
	case BRL:
		res["en"] = "Brazilian Real"
		res["zh-CN"] = "巴西雷亚尔"
		res["zh-HK"] = "巴西雷亞爾"
	case BSD:
		res["en"] = "Bahamian Dollar"
		res["zh-CN"] = "巴哈马元"
		res["zh-HK"] = "巴哈馬元"
	case BTN:
		res["en"] = "Bhutanese Ngultrum"
		res["zh-CN"] = "不丹尼斯特鲁姆"
		res["zh-HK"] = "不丹尼斯特魯姆"
	case BWP:
		res["en"] = "Botswanan Pula"
		res["zh-CN"] = "博茨瓦纳普拉"
		res["zh-HK"] = "博茨瓦納普拉"
	case BYR:
		res["en"] = "Belarusian Ruble"
		res["zh-CN"] = "白俄罗斯卢布"
		res["zh-HK"] = "白俄羅斯盧佈"
	case BZD:
		res["en"] = "Belize Dollar"
		res["zh-CN"] = "伯利兹元"
		res["zh-HK"] = "伯利茲元"
	case CAD:
		res["en"] = "Canadian Dollar"
		res["zh-CN"] = "加拿大元"
		res["zh-HK"] = "加拿大元"
	case CDF:
		res["en"] = "Congolese Franc"
		res["zh-CN"] = "刚果法郎"
		res["zh-HK"] = "剛果法郎"
	case CHF:
		res["en"] = "Swiss Franc"
		res["zh-CN"] = "瑞士法郎"
		res["zh-HK"] = "瑞士法郎"
	case CLP:
		res["en"] = "Chilean Peso"
		res["zh-CN"] = "智利比索"
		res["zh-HK"] = "智利比索"
	case CNY:
		res["en"] = "Chinese Yuan"
		res["zh-CN"] = "人民币"
		res["zh-HK"] = "人民幣"
	case CNH:
		res["en"] = "Chinese Yuan"
		res["zh-CN"] = "离岸人民币"
		res["zh-HK"] = "離岸人民幣"
	case COP:
		res["en"] = "Colombian Peso"
		res["zh-CN"] = "哥伦比亚比索"
		res["zh-HK"] = "哥倫比亞比索"
	case CRC:
		res["en"] = "Costa Rican Colon"
		res["zh-CN"] = "哥斯达黎加科朗"
		res["zh-HK"] = "哥斯達黎加科朗"
	case CUP:
		res["en"] = "Cuban Peso"
		res["zh-CN"] = "古巴比索"
		res["zh-HK"] = "古巴比索"
	case CVE:
		res["en"] = "Cape Verde Escudo"
		res["zh-CN"] = "佛得角埃斯库多"
		res["zh-HK"] = "佛得角埃斯庫多"
	case CZK:
		res["en"] = "Czech Koruna"
		res["zh-CN"] = "捷克克朗"
		res["zh-HK"] = "捷克克朗"
	case DJF:
		res["en"] = "Djiboutian Franc"
		res["zh-CN"] = "吉布提法郎"
		res["zh-HK"] = "吉佈提法郎"
	case DKK:
		res["en"] = "Danish Krone"
		res["zh-CN"] = "丹麦克朗"
		res["zh-HK"] = "丹麥克朗"
	case DOP:
		res["en"] = "Dominican Peso"
		res["zh-CN"] = "多明尼加比索"
		res["zh-HK"] = "多明尼加比索"
	case DZD:
		res["en"] = "Algerian Dinar"
		res["zh-CN"] = "阿尔及利亚第纳尔"
		res["zh-HK"] = "阿爾及利亞第納爾"
	case EGP:
		res["en"] = "Egyptian Pound"
		res["zh-CN"] = "埃及镑"
		res["zh-HK"] = "埃及鎊"
	case EUR:
		res["en"] = "Euro"
		res["zh-CN"] = "欧元"
		res["zh-HK"] = "歐元"
	case FJD:
		res["en"] = "Fijian Dollar"
		res["zh-CN"] = "斐济元"
		res["zh-HK"] = "斐濟元"
	case FKP:
		res["en"] = "Falkland Islands Pound"
		res["zh-CN"] = "福克兰群岛镑"
		res["zh-HK"] = "福克蘭群島鎊"
	case GBP:
		res["en"] = "Pound Sterling"
		res["zh-CN"] = "英镑"
		res["zh-HK"] = "英鎊"
	case GEL:
		res["en"] = "Georgian Lari"
		res["zh-CN"] = "格鲁吉亚拉里"
		res["zh-HK"] = "格魯吉亞拉裏"
	case GHS:
		res["en"] = "Ghanaian Cedi"
		res["zh-CN"] = "加纳塞地"
		res["zh-HK"] = "加納塞地"
	case GIP:
		res["en"] = "Gibraltar Pound"
		res["zh-CN"] = "直布罗陀镑"
		res["zh-HK"] = "直佈羅陀鎊"
	case GMD:
		res["en"] = "Gambian Dalasi"
		res["zh-CN"] = "冈比亚达拉西亚"
		res["zh-HK"] = "岡比亞達拉西亞"
	case GTQ:
		res["en"] = "Guatemalan Quetzal"
		res["zh-CN"] = "危地马拉格查尔"
		res["zh-HK"] = "危地馬拉格查爾"
	case GYD:
		res["en"] = "Guyanese Dollar"
		res["zh-CN"] = "圭亚那元"
		res["zh-HK"] = "圭亞那元"
	case HKD:
		res["en"] = "Hong Kong Dollar"
		res["zh-CN"] = "港币"
		res["zh-HK"] = "港幣"
	case HNL:
		res["en"] = "Honduran Lempira"
		res["zh-CN"] = "洪都拉斯伦皮拉"
		res["zh-HK"] = "洪都拉斯倫皮拉"
	case HRK:
		res["en"] = "Croatian Kuna"
		res["zh-CN"] = "克罗地亚库纳"
		res["zh-HK"] = "克羅地亞庫納"
	case HUF:
		res["en"] = "Hungarian Forint"
		res["zh-CN"] = "匈牙利福林"
		res["zh-HK"] = "匈牙利福林"
	case IDR:
		res["en"] = "Indonesian Rupiah"
		res["zh-CN"] = "印尼盾"
		res["zh-HK"] = "印尼盾"
	case ILS:
		res["en"] = "Israeli Shekel"
		res["zh-CN"] = "以色列谢克尔"
		res["zh-HK"] = "以色列謝克爾"
	case INR:
		res["en"] = "Indian Rupee"
		res["zh-CN"] = "印度卢比"
		res["zh-HK"] = "印度盧比"
	case IQD:
		res["en"] = "Iraqi Dinar"
		res["zh-CN"] = "伊拉克第纳尔"
		res["zh-HK"] = "伊拉克第納爾"
	case IRR:
		res["en"] = "Iranian Rial"
		res["zh-CN"] = "伊朗里亚尔"
		res["zh-HK"] = "伊朗裏亞爾"
	case JMD:
		res["en"] = "Jamaican Dollar"
		res["zh-CN"] = "牙买加元"
		res["zh-HK"] = "牙買加元"
	case JOD:
		res["en"] = "Jordanian Dinar"
		res["zh-CN"] = "约旦第纳尔"
		res["zh-HK"] = "約旦第納爾"
	case JPY:
		res["en"] = "Japanese Yen"
		res["zh-CN"] = "日元"
		res["zh-HK"] = "日元"
	case KES:
		res["en"] = "Kenyan Shilling"
		res["zh-CN"] = "肯尼亚先令"
		res["zh-HK"] = "肯尼亞先令"
	case KGS:
		res["en"] = "Kyrgyzstani Som"
		res["zh-CN"] = "吉尔吉斯斯坦索姆"
		res["zh-HK"] = "吉爾吉斯斯坦索姆"
	case KHR:
		res["en"] = "Cambodian Riel"
		res["zh-CN"] = "柬埔寨卢比"
		res["zh-HK"] = "柬埔寨盧比"
	case KMF:
		res["en"] = "Comorian Franc"
		res["zh-CN"] = "科摩罗法郎"
		res["zh-HK"] = "科摩羅法郎"
	case KRW:
		res["en"] = "South Korean Won"
		res["zh-CN"] = "韩元"
		res["zh-HK"] = "韓元"
	case KZT:
		res["en"] = "Kazakhstani Tenge"
		res["zh-CN"] = "哈萨克斯坦坚戈"
		res["zh-HK"] = "哈薩克斯坦堅戈"
	case LAK:
		res["en"] = "Lao Kip"
		res["zh-CN"] = "老挝基普"
		res["zh-HK"] = "老撾基普"
	case LBP:
		res["en"] = "Lebanese Pound"
		res["zh-CN"] = "黎巴嫩镑"
		res["zh-HK"] = "黎巴嫩鎊"
	case LKR:
		res["en"] = "Sri Lanka Rupee"
		res["zh-CN"] = "斯里兰卡卢比"
		res["zh-HK"] = "斯裏蘭卡盧比"
	case LRD:
		res["en"] = "Liberian Dollar"
		res["zh-CN"] = "利比里亚元"
		res["zh-HK"] = "利比裏亞元"
	case LSL:
		res["en"] = "Lesotho Loti"
		res["zh-CN"] = "莱索托里亚"
		res["zh-HK"] = "萊索托裏亞"
	case LYD:
		res["en"] = "Libyan Dinar"
		res["zh-CN"] = "利比亚第纳尔"
		res["zh-HK"] = "利比亞第納爾"
	case MAD:
		res["en"] = "Moroccan Dirham"
		res["zh-CN"] = "摩洛哥迪拉姆"
		res["zh-HK"] = "摩洛哥迪拉姆"
	case MDL:
		res["en"] = "Moldovan Leu"
		res["zh-CN"] = "摩尔多瓦列伊"
		res["zh-HK"] = "摩爾多瓦列伊"
	case MGA:
		res["en"] = "Malagasy Ariary"
		res["zh-CN"] = "马达加斯加阿里亚里"
		res["zh-HK"] = "馬達加斯加阿裏亞裏"
	case MKD:
		res["en"] = "Macedonian Denar"
		res["zh-CN"] = "马其顿第纳尔"
		res["zh-HK"] = "馬其頓第納爾"
	case MMK:
		res["en"] = "Myanma Kyat"
		res["zh-CN"] = "缅甸元"
		res["zh-HK"] = "緬甸元"
	case MNT:
		res["en"] = "Mongolian Tugrik"
		res["zh-CN"] = "蒙古图格里克"
		res["zh-HK"] = "蒙古圖格裏克"
	case MOP:
		res["en"] = "Macanese Pataca"
		res["zh-CN"] = "澳门元"
		res["zh-HK"] = "澳門元"

	case MUR:
		res["en"] = "Mauritius Rupee"
		res["zh-CN"] = "毛里求斯卢比"
		res["zh-HK"] = "毛裏求斯盧比"
	case MVR:
		res["en"] = "Maldivian Rufiyaa"
		res["zh-CN"] = "马尔代夫瓜拉"
		res["zh-HK"] = "馬爾代夫瓜拉"
	case MWK:
		res["en"] = "Malawian Kwacha"
		res["zh-CN"] = "马拉维克劳"
		res["zh-HK"] = "馬拉維克勞"
	case MXN:
		res["en"] = "Mexican Peso"
		res["zh-CN"] = "墨西哥比索"
		res["zh-HK"] = "墨西哥比索"
	case MYR:
		res["en"] = "Malaysian Ringgit"
		res["zh-CN"] = "马来西亚林吉"
		res["zh-HK"] = "馬來西亞林吉"
	case MZN:
		res["en"] = "Mozambican Metical"
		res["zh-CN"] = "莫桑比克元"
		res["zh-HK"] = "莫桑比克元"
	case NAD:
		res["en"] = "Namibian Dollar"
		res["zh-CN"] = "纳米比亚元"
		res["zh-HK"] = "納米比亞元"
	case NGN:
		res["en"] = "Nigerian Naira"
		res["zh-CN"] = "尼日利亚奈拉"
		res["zh-HK"] = "尼日利亞奈拉"
	case NIO:
		res["en"] = "Nicaraguan Cordoba"
		res["zh-CN"] = "尼加拉瓜科多巴"
		res["zh-HK"] = "尼加拉瓜科多巴"
	case NOK:
		res["en"] = "Norwegian Krone"
		res["zh-CN"] = "挪威克朗"
		res["zh-HK"] = "挪威克朗"
	case NPR:
		res["en"] = "Nepalese Rupee"
		res["zh-CN"] = "尼泊尔卢比"
		res["zh-HK"] = "尼泊爾盧比"
	case NZD:
		res["en"] = "New Zealand Dollar"
		res["zh-CN"] = "新西兰元"
		res["zh-HK"] = "新西蘭元"
	case PHP:
		res["en"] = "Philippine Peso"
		res["zh-CN"] = "菲律宾比索"
		res["zh-HK"] = "菲律賓比索"
	case PKR:
		res["en"] = "Pakistani Rupee"
		res["zh-CN"] = "巴基斯坦卢比"
		res["zh-HK"] = "巴基斯坦盧比"
	case PLN:
		res["en"] = "Polish Zloty"
		res["zh-CN"] = "波兰兹罗提"
		res["zh-HK"] = "波蘭茲羅提"
	case PYG:
		res["en"] = "Paraguayan Guarani"
		res["zh-CN"] = "巴拉圭瓜拉尼"
		res["zh-HK"] = "巴拉圭瓜拉尼"
	case QAR:
		res["en"] = "Qatari Rial"
		res["zh-CN"] = "卡塔尔里亚尔"
		res["zh-HK"] = "卡塔爾裏亞爾"
	case RON:
		res["en"] = "Romanian Leu"
		res["zh-CN"] = "罗马尼亚列伊"
		res["zh-HK"] = "羅馬尼亞列伊"
	case RUB:
		res["en"] = "Russian Ruble"
		res["zh-CN"] = "俄罗斯卢布"
		res["zh-HK"] = "俄羅斯盧佈"
	case RWF:
		res["en"] = "Rwandan Franc"
		res["zh-CN"] = "卢旺达法郎"
		res["zh-HK"] = "盧旺達法郎"
	case SAR:
		res["en"] = "Saudi Riyal"
		res["zh-CN"] = "沙特里亚尔"
		res["zh-HK"] = "沙特裏亞爾"
	case SBD:
		res["en"] = "Solomon Islands Dollar"
		res["zh-CN"] = "所罗门群岛元"
		res["zh-HK"] = "所羅門群島元"
	case SCR:
		res["en"] = "Seychelles Rupee"
		res["zh-CN"] = "塞舌尔卢比"
		res["zh-HK"] = "塞舌爾盧比"
	case SEK:
		res["en"] = "Swedish Krona"
		res["zh-CN"] = "瑞典克朗"
		res["zh-HK"] = "瑞典克朗"
	case SGD:
		res["en"] = "Singapore Dollar"
		res["zh-CN"] = "新加坡元"
		res["zh-HK"] = "新加坡元"
	case SHP:
		res["en"] = "Saint Helena Pound"
		res["zh-CN"] = "圣赫勒拿镑"
		res["zh-HK"] = "聖赫勒拿鎊"
	case SLL:
		res["en"] = "Sierra Leonean Leone"
		res["zh-CN"] = "塞拉利昂利昂"
		res["zh-HK"] = "塞拉利昂利昂"
	case SOS:
		res["en"] = "Somali Shilling"
		res["zh-CN"] = "索马里先令"
		res["zh-HK"] = "索馬裏先令"
	case SRD:
		res["en"] = "Surinamese Dollar"
		res["zh-CN"] = "苏里南元"
		res["zh-HK"] = "蘇裏南元"
	case STD:
		res["en"] = "São Tomé and Príncipe Dobra"
		res["zh-CN"] = "圣多美和普林西比多布拉"
		res["zh-HK"] = "聖多美和普林西比多佈拉"
	case SYP:
		res["en"] = "Syrian Pound"
		res["zh-CN"] = "叙利亚镑"
		res["zh-HK"] = "敘利亞鎊"
	case SZL:
		res["en"] = "Swazi Lilangeni"
		res["zh-CN"] = "斯威士兰先令"
		res["zh-HK"] = "斯威士蘭先令"
	case THB:
		res["en"] = "Thai Baht"
		res["zh-CN"] = "泰国铢"
		res["zh-HK"] = "泰國銖"
	case TJS:
		res["en"] = "Tajikistani Somoni"
		res["zh-CN"] = "塔吉克斯坦索莫尼"
		res["zh-HK"] = "塔吉克斯坦索莫尼"
	case TMT:
		res["en"] = "Turkmenistani Manat"
		res["zh-CN"] = "土库曼斯坦马纳特"
		res["zh-HK"] = "土庫曼斯坦馬納特"
	case TND:
		res["en"] = "Tunisian Dinar"
		res["zh-CN"] = "突尼斯第纳尔"
		res["zh-HK"] = "突尼斯第納爾"
	case TOP:
		res["en"] = "Tongan Paʻanga"
		res["zh-CN"] = "汤加邦"
		res["zh-HK"] = "湯加邦"
	case TRY:
		res["en"] = "Turkish Lira"
		res["zh-CN"] = "土耳其里拉"
		res["zh-HK"] = "土耳其裏拉"
	case TTD:
		res["en"] = "Trinidad and Tobago Dollar"
		res["zh-CN"] = "特立尼达和多巴哥元"
		res["zh-HK"] = "特立尼達和多巴哥元"
	case TWD:
		res["en"] = "New Taiwan Dollar"
		res["zh-CN"] = "新台币"
		res["zh-HK"] = "新臺幣"
	case TZS:
		res["en"] = "Tanzanian Shilling"
		res["zh-CN"] = "坦桑尼亚先令"
		res["zh-HK"] = "坦桑尼亞先令"
	case UAH:
		res["en"] = "Ukrainian Hryvnia"
		res["zh-CN"] = "乌克兰格里夫纳"
		res["zh-HK"] = "烏克蘭格裏夫納"
	case UGX:
		res["en"] = "Ugandan Shilling"
		res["zh-CN"] = "乌干达先令"
		res["zh-HK"] = "烏幹達先令"
	case USD:
		res["en"] = "United States Dollar"
		res["zh-CN"] = "美元"
		res["zh-HK"] = "美元"
	case UYU:
		res["en"] = "Uruguayan Peso"
		res["zh-CN"] = "乌拉圭比索"
		res["zh-HK"] = "烏拉圭比索"
	case UZS:
		res["en"] = "Uzbekistani Som"
		res["zh-CN"] = "乌兹别克斯坦索姆"
		res["zh-HK"] = "烏茲別克斯坦索姆"
	case VND:
		res["en"] = "Vietnamese Dong"
		res["zh-CN"] = "越南盾"
		res["zh-HK"] = "越南盾"
	case VUV:
		res["en"] = "Vanuatu Vatu"
		res["zh-CN"] = "瓦努阿图瓦图"
		res["zh-HK"] = "瓦努阿圖瓦圖"
	case WST:
		res["en"] = "Samoan Tala"
		res["zh-CN"] = "萨摩亚塔拉"
		res["zh-HK"] = "薩摩亞塔拉"
	case XAF:
		res["en"] = "CFA Franc BEAC"
		res["zh-CN"] = "非洲坦桑尼亚法郎"
		res["zh-HK"] = "非洲坦桑尼亞法郎"
	case XCD:
		res["en"] = "East Caribbean Dollar"
		res["zh-CN"] = "东加勒比元"
		res["zh-HK"] = "東加勒比元"
	case XPF:
		res["en"] = "CFP Franc"
		res["zh-CN"] = "CFP 法郎"
		res["zh-HK"] = "CFP 法郎"
	case YER:
		res["en"] = "Yemeni Rial"
		res["zh-CN"] = "也门里亚尔"
		res["zh-HK"] = "也門裏亞爾"
	case ZAR:
		res["en"] = "South African Rand"
		res["zh-CN"] = "南非兰特"
		res["zh-HK"] = "南非蘭特"
	case ZMW:
		res["en"] = "Zambian Kwacha"
		res["zh-CN"] = "赞比亚克瓦查"
		res["zh-HK"] = "贊比亞克瓦查"
	case ZWD:
		res["en"] = "Zimbabwean Dollar"
		res["zh-CN"] = "津巴布韦元"
		res["zh-HK"] = "津巴佈韋元"
	default:
		res["en"] = "Unknown"
		res["zh-CN"] = "未知"
		res["zh-HK"] = "未知"
	}

	return res
}
