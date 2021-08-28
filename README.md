# Go Programming Language - Examples - Simple HTTP Server & HTTP Client

The Go programming language is an open-source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

https://golang.org/doc/

## Used technology stack in this example

1. go1.17.windows-amd64.msi - Download - https://golang.org/dl/
2. MS Visual Studio Code
3. Chrome Browser
4. POSTMAN

## Run the following Go example one by one by topic wise to learn the Go programming

### Call by Value	
$ go run .\Go_1_Call_by_Value.go
### Call by Reference	
$ go run .\Go_2_Call_by_Reference.go
### Function	
$ go run .\Go_3_Function.go
### Short Variable Declaration	
$ go run .\Go_4_Short_Variable_Declaration.go
### Type Of	
$ go run .\Go_5_Type_Of.go
### Data Types	
$ go run .\Go_6_Data_Types.go
### Variables	
$ go run .\Go_7_Variables.go
### Functions as Values	
$ go run .\Go_8_Functions_as_Values.go
### Function Closure	
$ go run .\Go_9_Function_Closure.go
### Method	
$ go run .\Go_10_Method.go
### Local and Global Variables	
$ go run .\Go_11_Local_and_Global_Variables.go
### Array and For Loop	
$ go run .\Go_12_Array_and_For_Loop.go
### Slice in Array	
$ go run .\Go_13_Slice_in_Array.go
### Pointers	
$ go run .\Go_14_Pointers.go
### Swap Pointers	
$ go run .\Go_15_Swap_Pointers.go
### Pointers to Pointers	
$ go run .\Go_16_Pointers_to_Pointers.go
### Pointers in Array	
$ go run .\Go_17_Pointers_in_Array.go
### Range in Array	
$ go run .\Go_18_Range_in_Array.go
### Map	
$ go run .\Go_19_Map.go
### Map	
$ go run .\Go_20_Map.go
### Recursion	
$ go run .\Go_21_Recursion.go
### Recursion	
$ go run .\Go_22_Recursion.go
### Type Cast	
$ go run .\Go_23_Type_Cast.go
### Interface	
$ go run .\Go_24_Interface.go
### Error Handling	
$ go run .\Go_25_Error_Handling.go
### Reading Input in Console 1	
$ go run .\Go_26_Reading_Input_in_Console_1.go
### Reading Input in Console 2	
$ go run .\Go_27_Reading_Input_in_Console_2.go
### Read File Line by Line	
$ go run .\Go_28_Read_File_Line_by_Line.go
### Read File into String	
$ go run .\Go_29_Read_File_into_String.go
### Write File	
$ go run .\Go_30_Write_File.go
### Delete File	
$ go run .\Go_31_Delete_File.go
### Run GO HTTP Server
$ go run .\Go_32_Http_Server.go
### Access the following POST URL from POSTMAN
POST http://localhost:8090/hello
Req JSON: 
{
    "name": "Rahamath S",
    "address": "India"
}
### Also, access the following GET URL from POSTMAN or Chrome browser
http://localhost:8090/
http://localhost:8090/headers
### Also, you can run the following http client go program
$ go run .\Go_33_Http_Client.go
