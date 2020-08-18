package main

import (
	"gopl-practise/ch7/7-16/eval"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./ch7/7-16/index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func calc(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./ch7/7-16/index.html"))
	exprStr := req.PostFormValue("expr")
	envStr := req.PostFormValue("env")
	expr, err := eval.Parse(exprStr)
	if err != nil {
		log.Fatal(err)
	}
	env, err := parseEnv(envStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = tmpl.Execute(w, expr.Eval(env)); err != nil {
		log.Fatal(err)
	}
}

func parseEnv(envStr string) (eval.Env, error) {
	env := make(eval.Env)
	fields := strings.FieldsFunc(envStr,
		func(r rune) bool {
			return strings.ContainsRune(`:=,{}\"`, r) ||
				unicode.IsSpace(r)
		})
	for i := 0; i+1 < len(fields); i += 2 {
		k := strings.TrimSpace(fields[i])
		v := strings.TrimSpace(fields[i+1])
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}

		env[eval.Var(k)] = val
	}
	return env, nil
}
