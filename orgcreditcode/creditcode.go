package orgcreditcode

import (
	"errors"
	"strconv"
	"strings"
)

var ORGWEIGHT = []int32 {3,7,9,10,5,8,4,2}

var CREDITWEIGHT = []int32 {1,3,9,27,19,26,16,17,20,29,25,13,8,24,10,30,28}

var ORGMAP_REVERSE = map[string]int32{
	"0":0 ,
	"1":1 ,
	"2":2 ,
	"3":3 ,
	"4":4 ,
	"5":5 ,
	"6":6 ,
	"7":7 ,
	"8":8 ,
	"9":9 ,
	"A":10,
	"B":11,
	"C":12,
	"D":13,
	"E":14,
	"F":15,
	"G":16,
	"H":17,
	"I":18,
	"J":19,
	"K":20,
	"L":21,
	"M":22,
	"N":23,
	"O":24,
	"P":25,
	"Q":26,
	"R":27,
	"S":28,
	"T":29,
	"U":30,
	"V":31,
	"W":32,
	"X":33,
	"Y":34,
	"Z":35,
}

var CREDITMAP_REVERSE = map[string]int32{
	"0":0 ,
	"1":1 ,
	"2":2 ,
	"3":3 ,
	"4":4 ,
	"5":5 ,
	"6":6 ,
	"7":7 ,
	"8":8 ,
	"9":9 ,
	"A":10,
	"B":11,
	"C":12,
	"D":13,
	"E":14,
	"F":15,
	"G":16,
	"H":17,
	"J":18,
	"K":19,
	"L":20,
	"M":21,
	"N":22,
	"P":23,
	"Q":24,
	"R":25,
	"T":26,
	"U":27,
	"W":28,
	"X":29,
	"Y":30,
}
var CREDITMAP = map[int32]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
	16: "G",
	17: "H",
	18: "J",
	19: "K",
	20: "L",
	21: "M",
	22: "N",
	23: "P",
	24: "Q",
	25: "R",
	26: "T",
	27: "U",
	28: "W",
	29: "X",
	30: "Y",
}

// 根据社会信用码的前17位生成第18位校验码
func GenCreditCheckCode(code string) (string, error){
	if len(code) != 17 {
		return "", errors.New("code must be 17 character")
	}
	var originWeightSum int32
	for i,c:=  range code{
		originWeightSum += CREDITMAP_REVERSE[string(c)] * CREDITWEIGHT[i]
	}
	checkCodeNumber := 31 - originWeightSum % 31
	if checkCodeNumber == 31 {
		checkCodeNumber = 0
	}
	return code + CREDITMAP[checkCodeNumber], nil
}

// 根据８位机构组织编码生成其校验码，作为社会信用编码的第17位
func GenOrgCheckCode(code string) (string, error){
	if len(code) != 8 {
		return "", errors.New("code must be 8 character")
	}
	var originWeightSum int32
	for i,c:=  range code{
		originWeightSum += ORGMAP_REVERSE[string(c)]  * ORGWEIGHT[i]
	}
	checkCodeNumber := 11 - originWeightSum % 11
	switch checkCodeNumber {
	case 10:
		return code +"X",nil
	case 11:
		return code + "0", nil
	default:
		return code + strconv.Itoa(int(checkCodeNumber)), nil
	}
}
//manageType 　登记管理部门代码１位　工商为９
//orgType　　　机构类别代码１位　　企业为１　个体工商户为２
//areaCode　　 6位机关行政区划码
//orgCode      8位组织机构代码(不包含校验码，实际组织机构代码为９位)
func GenCreditCode(manageType string, orgType string, areaCode string, orgCode string) (string, error){
	var err error
	orgCodeAll, err := GenOrgCheckCode(orgCode)
	if err != nil {
		return "", err
	}

	creditCodeAll, err := GenCreditCheckCode(strings.Join([]string{manageType, orgType, areaCode, orgCodeAll},""))
	if err != nil {
		return "", err
	}

	return creditCodeAll, nil
}
