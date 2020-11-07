package gokit

import (
	"net/url"
)

//ParseURLQuery 获取url query中的参数
func ParseURLQuery(rawurl, field string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	// fmt.Println(u.Scheme)
	// fmt.Println(u.User)
	// fmt.Println(u.User.Username())
	// p, _ := u.User.Password()
	// fmt.Println(p)
	// fmt.Println(u.Host)
	// fmt.Println(u.Path)
	// fmt.Println(u.Fragment)
	// fmt.Println(u.RawQuery)
	v, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", err
	}
	return v.Get(field), nil
}
