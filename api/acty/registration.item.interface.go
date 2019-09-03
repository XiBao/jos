package acty

type RegistrationItem struct {
	InformationChannel     string `json:"informationChannel,omitempty" codec:"informationChannel,omitempty"`
	Birthday               string `json:"birthday,omitempty" codec:"birthday,omitempty"`
	Sex                    string `json:"sex,omitempty" codec:"sex,omitempty"`
	Job                    string `json:"job,omitempty" codec:"job,omitempty"`
	SkuName                string `json:"skuName,omitempty" codec:"skuName,omitempty"`
	AddressDetail          string `json:"addressDetail,omitempty" codec:"addressDetail,omitempty"`
	EmergencyContact       string `json:"emergencyContact,omitempty" codec:"emergencyContact,omitempty"`
	Nationality            string `json:"nationality,omitempty" codec:"nationality,omitempty"`
	EmergencyContactNumber string `json:"emergencyContactNumber,omitempty" codec:"emergencyContactNumber,omitempty"`
	PhoneNumber            string `json:"phoneNumber,omitempty" codec:"phoneNumber,omitempty"`
	Email                  string `json:"email,omitempty" codec:"email,omitempty"`
	HomeAddress            string `json:"homeAddress,omitempty" codec:"homeAddress,omitempty"`
	Name                   string `json:"name,omitempty" codec:"name,omitempty"`
	BeastResult            string `json:"beastResult,omitempty" codec:"beastResult,omitempty"`
	SkuId                  string `json:"skuId,omitempty" codec:"skuId,omitempty"`
	ClothingSize           string `json:"clothingSize,omitempty" codec:"clothingSize,omitempty"`
	CertificatePictureUrl  string `json:"certificatePictureUrl,omitempty" codec:"certificatePictureUrl,omitempty"`
	OrderId                string `json:"orderId,omitempty" codec:"orderId,omitempty"`
	IdNumber               string `json:"idNumber,omitempty" codec:"idNumber,omitempty"`
}
