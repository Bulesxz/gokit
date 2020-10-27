/*
通过google 服务账号 操作google sheet接口
*/

package third

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
)

// NewService new 一个google sheet 服务接口
func NewService(email, privateKey, privateKeyID, tokenURL string, Scopes []string) (*sheets.Service, error) {
	// Create a JWT configurations object for the Google service account
	conf := &jwt.Config{
		Email:        email,
		PrivateKey:   []byte(privateKey),
		PrivateKeyID: privateKeyID,
		TokenURL:     tokenURL,
		Scopes:       Scopes,
	}

	client := conf.Client(oauth2.NoContext)

	// Create a service object for Google sheets
	return sheets.New(client)
}

// GetListTitle 获取所有子表标题
func GetListTitle(srv *sheets.Service, spreadsheetID string) ([]string, error) {
	shs, err := srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, err
	}
	titles := []string{}
	for i := 0; i < len(shs.Sheets); i++ {
		titles = append(titles, shs.Sheets[i].Properties.Title)
	}
	return titles, nil
}

//ReadRows 读取所有行
func ReadRows(srv *sheets.Service, spreadsheetID string, readRange string) ([][]interface{}, error) {
	//readRange := "Sheet1!A:A"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

//WriteRow append一行
func WriteRow(srv *sheets.Service, spreadsheetID, writeRange string, value []interface{}) error {
	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	var rb sheets.ValueRange
	rb.Values = append(rb.Values, value)
	_, err := srv.Spreadsheets.Values.Append(spreadsheetID, writeRange, &rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Do()
	return err
}
