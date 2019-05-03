package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "Paris"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"
	fmt.Println(countryCapitalMap)
	fmt.Println(countryCapitalMap)
	 /*使用键输出地图值 */ 
	 for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }
}

// Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。