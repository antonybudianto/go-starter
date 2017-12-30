package main

func main() {
	a := App{}
	a.Initialize("root", "hello", "rest_api_example")

	a.Run(":8000")
}
