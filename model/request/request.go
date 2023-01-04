/*跟前端溝通的地方 */
package request



type RegisterRequest struct{
	Username	string `json:"username"`
	Password  	string `json:"password"`
	Email		string `json:"Email"`
	Birthday  	string `json:"Birthday"`
}

type LoginRequest struct{
	Username		string `json:"username"`
	IP				string `json:"ip"`
	Token			string `json:"token"`
}




