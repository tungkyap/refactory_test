package main

import (
    "encoding/json"
    "fmt"
)

type profileList struct {
    id int `json:"id"`
    username string `json:"username"`
    profile profile `json:"profile"`
    // articles []articles `json:articles`
}

type profile struct {
    full_name string `json:"full_name"`
    birthday string `json:"birthday"`
    phones string `json:phones`
}

// type articles []struct {
//
// }

func main() {
    fmt.Println("Test json manipulation")
    profile := profile{full_name: "Adi Pradana", birthday: "1995-12-21", phones: "021"}
    profileList := profileList{id: 123, username: "admin", profile: profile}

    fmt.Printf("%+v\n",profileList)
    var jsonData []byte
    jsonData, err := json.MarshalIndent(profileList, "", "  ")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(jsonData))
}
