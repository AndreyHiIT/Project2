package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"    
)

func main() {
   reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример")
	text1, _ := reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)
  text:=strings.ReplaceAll(text1, `"-"`, `" - "`)
  operator, err := findOp(text)
		if err != nil {
		panic(err)
    }
  arguments := strings.Split(text, operator)
  arg01:=strings.TrimSpace(arguments[0])
  arg02:=strings.TrimSpace(arguments[1])
  arg1:=strings.ReplaceAll(arg01,`"`, "")
  arg2:=strings.ReplaceAll(arg02,`"`, "")
  if len(arg1)>10||len(arg2)>10 {
    panic("Для ввода допустимо не более 10 символов для каждого аргумента")
  }
  boolcheck1:=boolcheck(arg1)  //если false то arg - int, если true то arg string
  boolcheck2:=boolcheck(arg2)
  argInt2,_ :=strconv.Atoi(arg2)
   if boolcheck1==boolcheck2 {
    if operator=="+" { 
      l:=fmt.Sprintf(`"%s%s"`,arg1,arg2)
      fmt.Printf(formatResult(l))
      } else if operator==`" - "`{
      if strings.Contains(arg1,arg2) {
      l:=fmt.Sprintf("%q",strings.Replace(arg1,arg2,"",-1))
      fmt.Printf(formatResult(l))
      } else {
      l:=fmt.Sprintf("%q",arg1)
      fmt.Printf(formatResult(l))
      }
    } else if operator=="/"{
      panic("Деление строку на строку запрещено")
    }
  } else if boolcheck1!=boolcheck2 {
      if boolcheck1==false {panic("Первый аргумент является числом и отличен от второго аргумента")
    } else  {
      if argInt2<0||argInt2>10 {
       panic("Недопустимое число, допустимы числа от 1 до 10")
       } else if operator=="*" {
        l:=fmt.Sprintf("%q",strings.Repeat(arg1,argInt2))
        fmt.Printf(formatResult(l))
      } else if operator=="/"{
        num:=len(arg1)
        num2:=num/argInt2
        l:=fmt.Sprintf("%q",arg1[0:num2])
        fmt.Printf(formatResult(l))
      }
    }
  } else {
    panic("Выражение введено неверно")
  }
}


func checkfloat(arg2 string){
 if _, err := strconv.ParseFloat(arg2, 64); err != nil {
      panic("Дробные числа недопустимы") 
    }
}
  
  
func boolcheck(arg string) bool{
  _,err := strconv.Atoi(arg)
  if err != nil {
        return true
    }
    return false
}

func findOp(text string) (string, error) {
 	switch {
	  case strings.Contains(text, "+"):
		return "+", nil
	  case strings.Contains(text, `" - "`):
		return `" - "`, nil
	  case strings.Contains(text, "*"):
		return "*", nil
  	case strings.Contains(text, "/"):
		return "/", nil
  	default:
		return "", fmt.Errorf("Неверная операция.Доступные операции: +, -, /, *.")
	}
}
func formatResult(result string) string {
	if len(result) > 40 {
		return result[:37] + "..."
	}
	return result
}

