/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: utils.go
 * File Name: vs.go
 * Last Modified: 15/01/2021, 20:33
 */

package httputil

var macVs = []string{
	"10.0.0",
	"10.0.1",
	"10.0.2",
	"10.0.3",
	"10.0.4",
	"10.1.0",
	"10.1.1",
	"10.1.2",
	"10.1.3",
	"10.1.4",
	"10.1.5",
	"10.2.0",
	"10.2.1",
	"10.2.2",
	"10.2.3",
	"10.2.4",
	"10.2.5",
	"10.2.6",
	"10.2.7",
	"10.2.8",
	"10.3.0",
	"10.3.1",
	"10.3.2",
	"10.3.3",
	"10.3.4",
	"10.3.5",
	"10.3.6",
	"10.3.7",
	"10.3.8",
	"10.3.9",
	"10.4.0",
	"10.4.1",
	"10.4.2",
	"10.4.3",
	"10.4.4",
	"10.4.5",
	"10.4.6",
	"10.4.7",
	"10.4.8",
	"10.4.9",
	"10.4.10",
	"10.4.11",
	"10.5.0",
	"10.5.1",
	"10.5.2",
	"10.5.3",
	"10.5.4",
	"10.5.5",
	"10.5.6",
	"10.5.7",
	"10.5.8",
	"10.6.0",
	"10.6.1",
	"10.6.2",
	"10.6.3",
	"10.6.4",
	"10.6.5",
	"10.6.6",
	"10.6.7",
	"10.6.8",
	"10.7.0",
	"10.7.1",
	"10.7.2",
	"10.7.3",
	"10.7.4",
	"10.7.5",
	"10.8.0",
	"10.8.1",
	"10.8.2",
	"10.8.3",
	"10.8.4",
	"10.8.5",
	"10.9.0",
	"10.9.1",
	"10.9.2",
	"10.9.3",
	"10.9.4",
	"10.9.5",
	"10.10",
	"10.10.0",
	"10.10.1",
	"10.10.2",
	"10.10.3",
	"10.10.4",
	"10.10.5",
	"10.11",
	"10.11.0",
	"10.11.1",
	"10.11.2",
	"10.11.3",
	"10.11.4",
	"10.11.5",
	"10.11.6",
	"10.12.0",
	"10.12.1",
	"10.12.2",
	"10.12.3",
	"10.12.4",
	"10.12.5",
	"10.12.6",
	"10.13.0",
	"10.13.1",
	"10.13.2",
	"10.13.3",
	"10.13.4",
	"10.13.5",
	"10.13.6",
	"10.14.0",
	"10.14.1",
	"10.14.2",
	"10.14.3",
	"10.14.4",
	"10.14.5",
	"10.14.6",
	"10.15.0",
	"10.15.1",
	"10.15.2",
	"10.15.3",
	"10.15.4",
	"10.15.5",
	"10.15.6",
	"10.15.7",
	"11.0",
	"11.0.1",
	"11.1",
}

var iosVs = []string{
	"3.0",
	"3.0.1",
	"3.1 ",
	"3.1.2",
	"3.1.3",
	"4.0",
	"4.0.1",
	"4.0.2",
	"4.1",
	"4.2.1",
	"4.2.5",
	"4.2.6",
	"4.2.7",
	"4.2.8",
	"4.2.9",
	"4.2.10",
	"4.3",
	"4.3.1",
	"4.3.2",
	"4.3.3",
	"4.3.4",
	"4.3.5",
	"5.0",
	"5.0.1",
	"5.1",
	"5.1.1",
	"6.0",
	"6.0.1",
	"6.0.2",
	"6.1",
	"6.1.1",
	"6.1.2",
	"6.1.3",
	"6.1.4",
	"6.1.6",
	"7.0",
	"7.0.1",
	"7.0.2",
	"7.0.3",
	"7.0.4",
	"7.0.5",
	"7.0.6",
	"7.1",
	"7.1.1",
	"7.1.2",
	"8.0",
	"8.0.1",
	"8.0.2",
	"8.1",
	"8.1.1",
	"8.1.2",
	"8.1.3",
	"8.2",
	"8.3",
	"8.4",
	"8.4.1",
	"9.0",
	"9.0.1",
	"9.0.2",
	"9.1",
	"9.2",
	"9.2.1",
	"9.3",
	"9.3.1",
	"9.3.2",
	"9.3.3",
	"9.3.4",
	"9.3.5",
	"9.3.6",
	"10.0",
	"10.0.2",
	"10.0.3",
	"10.1",
	"10.1.1",
	"10.2",
	"10.2.1",
	"10.3",
	"10.3.1",
	"10.3.2",
	"10.3.3",
	"10.3.4",
	"11.0",
	"11.0.1",
	"11.0.2",
	"11.0.3",
	"11.1",
	"11.1.1",
	"11.1.2",
	"11.2",
	"11.2.1",
	"11.2.2",
	"11.2.5",
	"11.2.6",
	"11.3",
	"11.3.1",
	"11.4",
	"11.4.1",
	"12.0",
	"12.0.1",
	"12.1",
	"12.1.1",
	"12.1.2",
	"12.1.3",
	"12.1.4",
	"12.2",
	"12.3",
	"12.3.1",
	"12.3.2",
	"12.4",
	"12.4.1",
	"12.4.2",
	"12.4.3",
	"12.4.4",
	"12.4.5",
	"12.4.6",
	"12.4.7",
	"12.4.8",
	"12.4.9",
	"12.5",
	"13.0",
	"13.1",
	"13.1.1",
	"13.1.2",
	"13.1.3",
	"13.2",
	"13.2.1",
	"13.2.2",
	"13.2.3",
	"13.3",
	"13.3.1",
	"13.4",
	"13.4.1",
	"13.5",
	"13.5.1",
	"13.6",
	"13.6.1",
	"13.7",
	"14.0",
	"14.0.1",
	"14.1",
	"14.2",
	"14.2.1",
	"14.3",
}

var webkitVs = map[string][]string{
	"ios": {
		"419.3",
		"420.1",
		"525.18.1",
		"525.20",
		"528.18",
		"528.16",
		"531.21.10",
		"532.9",
		"6531.22.7",
		"533.17.9",
		"6533.18.5",
		"534.46",
		"536.26",
		"537.51.1",
		"9537.53",
		"537.51.2",
		"600.1.3",
		"600.1.4",
		"600.1.4",
		"601.1.46",
		"601.1",
		"602.1.38",
		"602.1",
		"602.1.50",
		"602.2.14",
		"602.3.12",
		"602.4.6",
		"603.1.30",
		"603.2.4",
		"603.2.4",
		"604.1.38",
		"604.1",
		"604.3.5",
		"604.4.7",
		"604.5.6",
		"605.1.15",
		"606.1.36",
		"604.1",
		"607.1.40",
		"608.2.11",
		"604.1",
	},
	"mac": {
		"85.8.5",
		"312.3",
		"312.5",
		"312.6",
		"416.11",
		"419.3",
		"522.11",
		"522.12",
		"522.12.1",
		"523.10",
		"525.13",
		"525.17",
		"525.20",
		"525.21",
		"525.26",
		"525.27",
		"525.28",
		"526.11.2",
		"528.16",
		"528.17",
		"530.17",
		"530.18",
		"530.19",
		"531.9",
		"531.21.10",
		"531.22.7",
		"533.16",
		"533.17.8",
		"533.18.5",
		"533.19.4",
		"533.16",
		"533.17.8",
		"533.18.5",
		"533.19.4",
		"533.20.27",
		"533.21.1",
		"533.22.3",
		"534.48.3",
		"534.51.22",
		"534.52.7",
		"534.53.10",
		"534.54.16",
		"534.55.3",
		"534.56.5",
		"534.57.2",
		"534.58.2",
		"534.59.8",
		"534.59.10",
		"536.25",
		"536.26",
		"536.26.17",
		"536.28.10",
		"536.29.13",
		"536.30.1",
		"537.43.58",
		"537.73.11",
		"537.78.2",
		"537.85.17",
		"537.71",
		"537.73.11",
		"537.75.14",
		"537.76.4",
		"537.77.4",
		"537.78.2",
		"600.3.18",
		"600.8.9",
		"538.35.8",
		"600.6.3",
		"600.7.12",
		"601.1.56",
		"601.2.7",
		"601.3.9",
		"601.4.4",
		"601.5.17",
		"601.6.17",
		"601.7.1",
		"601.7.8",
		"602.1.50",
		"602.2.14",
		"602.3.12",
		"602.4.8",
		"603.1.30",
		"603.2.4",
		"603.3.8",
		"604.2.4",
		"605.1.33",
		"606.1.36",
		"607.1.40",
		"608.2.11",
		"610.2.11",
		"610.3.7.1.9",
	},
	"windows": {
		"522.11.3",
		"522.12.2",
		"522.13.1",
		"522.15.5",
		"523.12.9",
		"523.13",
		"523.15",
		"525.13",
		"525.17",
		"525.21",
		"525.26.13",
		"525.27.1",
		"525.28.1",
		"525.29.1",
		"526.12.2",
		"528.1.1",
		"528.16",
		"528.17",
		"530.17",
		"530.19.1",
		"531.9.1",
		"531.21.10",
		"531.22.7",
		"533.16",
		"533.17.8",
		"533.18.5",
		"533.19.4",
		"533.20.27",
		"533.21.1",
		"534.50",
		"534.51.22",
		"534.52.7",
		"534.54.16",
		"534.55.3",
		"534.57.2",
	},
}

var chromeVs = []string{
	"0.2.149",
	"0.3.154",
	"0.4.154",
	"1.0.154",
	"2.0.172",
	"3.0.195",
	"4.0.249",
	"4.1.249",
	"5.0.375",
	"6.0.472",
	"7.0.517",
	"8.0.552",
	"9.0.597",
	"10.0.648",
	"11.0.696",
	"12.0.742",
	"13.0.782",
	"14.0.835",
	"15.0.874",
	"16.0.912",
	"17.0.963",
	"18.0.1025",
	"19.0.1084",
	"20.0.1132",
	"21.0.1180",
	"22.0.1229",
	"23.0.1271",
	"24.0.1312",
	"25.0.1364",
	"26.0.1410",
	"27.0.1453",
	"28.0.1500",
	"29.0.1547",
	"30.0.1599",
	"31.0.1650",
	"32.0.1700",
	"33.0.1750",
	"34.0.1847",
	"35.0.1916",
	"36.0.1985",
	"37.0.2062",
	"38.0.2125",
	"39.0.2171",
	"40.0.2214",
	"41.0.2272",
	"42.0.2311",
	"43.0.2357",
	"44.0.2403",
	"45.0.2454",
	"46.0.2490",
	"47.0.2526",
	"48.0.2564",
	"49.0.2623",
	"50.0.2661",
	"51.0.2704",
	"52.0.2743",
	"53.0.2785",
	"54.0.2840",
	"55.0.2883",
	"56.0.2924",
	"57.0.2987",
	"58.0.3029",
	"59.0.3071",
	"60.0.3112",
	"61.0.3163",
	"62.0.3202",
	"63.0.3239",
	"64.0.3282",
	"65.0.3325",
	"66.0.3359",
	"67.0.3396",
	"68.0.3440",
	"69.0.3497",
	"70.0.3538",
	"71.0.3578",
	"72.0.3626",
	"73.0.3683",
	"74.0.3729",
	"75.0.3770",
	"76.0.3809",
	"77.0.3865",
	"78.0.3904",
	"79.0.3945",
	"80.0.3987",
	"81.0.4044",
	"83.0.4103",
	"84.0.4147",
	"85.0.4183",
	"86.0.4240",
	"87.0.4280",
	"88.0.4324",
	"89.0",
	"89.0",
}

var edgeHTMLVs = []string{
	"12.0",
	"12.10049",
	"12.10166",
	"12.10240",
	"12.10525",
	"12.10532",
	"13.10547",
	"13.10565",
	"13.10586",
	"13.11099",
	"14.14267",
	"14.14279",
	"14.14291",
	"14.14316",
	"14.14327",
	"14.14342",
	"14.14352",
	"14.14356",
	"14.14361",
	"14.14366",
	"14.14367",
	"14.14376",
	"14.14393",
	"14.14901",
	"14.14915",
	"14.14926",
	"15.14942",
	"15.14959",
	"15.14986",
	"15.15063",
	"16",
}
