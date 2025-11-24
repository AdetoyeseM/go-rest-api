
package models


type UserModel struct {

	ID 					int 			`json:"id"`
	Email 				string 			`json:"email"`
	Password 			string			`json:"password"`
}