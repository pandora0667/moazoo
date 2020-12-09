package start

import "fmt"

func Logo()  {

	license := `
The MIT License (MIT)

Copyright (c) <2020> <moazoo>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
`
	fmt.Println(license)

	fmt.Println("Multi protocol IoT server platform MOAZOO Version 0.1 ALPHA")
	fmt.Println("Hanbat National University ")

	// http://patorjk.com/software/taag/#p=testall&f=Graffiti&t=Code%20Generator
	platform :=
		`
        .__              _____  __                                              
__  _  _|__| ___________/ ____\/  |_    _____   _________  ____________   ____  
\ \/ \/ /  |/  ___/  _ \   __\\   __\  /     \ /  _ \__  \ \___   /  _ \ /  _ \ 
 \     /|  |\___ (  <_> )  |   |  |   |  Y Y  (  <_> ) __ \_/    (  <_> |  <_> )
  \/\_/ |__/____  >____/|__|   |__|   |__|_|  /\____(____  /_____ \____/ \____/ 
                \/                          \/           \/      \/             

`
	fmt.Println(platform)
	fmt.Println("name : seongwon lee , email : seongwon@edu.hanbat.ac.kr")

}