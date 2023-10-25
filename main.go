package main

import (
	"github.com/moqsien/gogpt/pkgs/tui"
)

// type ItemList struct {
// 	List []map[string]string `json:"list"`
// }

// func GetPrompts() {
// 	itemList := ItemList{List: []map[string]string{}}
// 	d := `C:\Users\moqsien\data\projects\md\ChatGPT-System-Prompts\prompts`
// 	dList, _ := os.ReadDir(d)
// 	for _, item := range dList {
// 		if item.IsDir() {
// 			dd := filepath.Join(d, item.Name())
// 			l, _ := os.ReadDir(dd)
// 			for _, entry := range l {
// 				if !entry.IsDir() {
// 					fPath := filepath.Join(dd, entry.Name())
// 					content, _ := os.ReadFile(fPath)
// 					if len(content) > 0 {
// 						sList := strings.Split(string(content), "\n")
// 						act := strings.ReplaceAll(entry.Name(), ".md", "")
// 						act = strings.ReplaceAll(act, "-", "_")
// 						for _, s := range sList {
// 							if len(s) > 40 {
// 								itemList.List = append(itemList.List, map[string]string{
// 									"act":    act,
// 									"prompt": strings.TrimSuffix(s, "\r"),
// 								})
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	result, _ := json.MarshalIndent(itemList, "", "    ")
// 	os.WriteFile("result.json", result, os.ModePerm)
// }

func main() {
	// GetPrompts()

	cnf := tui.GetDefaultConfig()
	ui := tui.NewGPTUI(cnf)
	ui.Run()

	// g := gpt.NewGPT(cnf)
	// conv := gpt.NewConversation(cnf)
	// conv.AddQuestion("write a quick sort in go, please!")
	// ctx := conv.GetMessages()
	// // fmt.Printf("%+v", ctx)
	// msg, err := g.SendMsg(ctx)
	// if err == nil {
	// 	fmt.Print(msg)
	// 	for {
	// 		msg, err = g.RecvMsg()
	// 		fmt.Printf(msg)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			break
	// 		}
	// 	}
	// }

	// tui.TextTest()

	// p := gpt.NewGPTPrompt(cnf)
	// p.ChoosePrompt()
	// fmt.Println(p.PromptStr())
}
