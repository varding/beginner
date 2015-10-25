package conf

import (
	"fmt"
	"github.com/alecthomas/log4go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/*
when all yml files loaded and appended in one big file then call yml.UnMarshal,only the last key will be saved if key is not unique
so yml files will be parsed respetively and merged by hand
*/

func load_all_locals(path string) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		log4go.Error("read dir err:%v", err)
	}

	/*
		example
		"en":
		  errors:
		    messages:
	*/

	for _, i := range info {
		if i.IsDir() {
			continue
		}
		d, err2 := ioutil.ReadFile(fmt.Sprintf("%s%s", path, i.Name()))
		if err2 != nil {
			log4go.Error("load yml file err:%v", err2)
			continue
		}

		//parse
		out := make(map[interface{}]map[interface{}]interface{})
		if err3 := yaml.Unmarshal(d, out); err3 != nil {
			log4go.Error("yml parse err:%v", err3)
		}
		//en:map...
		for lang, data := range out {
			//error:...
			for k, v := range data {
				//add k,v to the child map when lang exist
				if c, ok := locals[lang]; ok {
					c[k] = v
				} else {
					locals[lang] = data
				}
			}

		}
	}

	//print length
	// fmt.Println(len(content["en"]))
	// fmt.Println(len(content["zh-CN"]))
	// fmt.Println(len(content["zh-TW"]))

	//pretty print the result
	// fmt.Println("en")
	// pretty_print(locals["en"], 1)

	// fmt.Println("zh-CN")
	// pretty_print(locals["zh-CN"], 1)

	// fmt.Println("zh-TW")
	// pretty_print(locals["zh-TW"], 1)
}

func pretty_print(m map[interface{}]interface{}, depth int) {
	for k, v := range m {
		switch v.(type) {
		case map[interface{}]interface{}:
			for i := 0; i < depth; i++ {
				fmt.Print("\t")
			}
			fmt.Printf("%v:\r\n", k)
			pretty_print(v.(map[interface{}]interface{}), depth+1)
		case string:
			for i := 0; i < depth; i++ {
				fmt.Print("\t")
			}
			fmt.Printf("%s:%s\r\n", k, v)
		}
	}
}

func init() {
	load_all_locals("conf/locales/")
}

var locals = make(map[interface{}]map[interface{}]interface{})

func Local(lang string) map[interface{}]interface{} {
	return locals[lang]
}
