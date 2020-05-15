package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
	MyHtml    []temp
	OpTitle   string
}

type temp struct {
	OpTitless     string
	Content       string
	Img           string
	Short_content string
}

func main() {
	//assets
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//asset html
	tmpl := template.Must(template.ParseFiles("./layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			OpTitle: "It's my life and I try",
			MyHtml: []temp{
				{
					OpTitless:     "Fight with monster in my body.",
					Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
					Img:           "https://www.imgworlds.com/wp-content/uploads/2015/12/18-CONTACTUS-HEADER.jpg",
					Short_content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				},
				{
					OpTitless:     "Fight with monster in my body.",
					Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
					Img:           "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSwHlSQzp5UFHhZSbi9ZgGoq_sDQtBqcn_PsqDnJbJaqn7xCe2x&usqp=CAU",
					Short_content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				},
				{
					OpTitless:     "Fight with monster in my body.",
					Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
					Img:           "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcR9_qHPgWsGQGzkGd1vVBLlAYI-8BdVP1l_3Tv5oiGqarULLrsn&usqp=CAU",
					Short_content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				},
				{
					OpTitless:     "Fight with monster in my body.",
					Content:       "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
					Img:           "https://www.cartoonbrew.com/wp-content/uploads/2019/01/gumball_fell-580x326.jpg",
					Short_content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8081", nil)
}
