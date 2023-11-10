package dto

type Base struct {
	Query  string `json:"query"`
	Result Result `json:"result"`
}

type Result struct {
	Description           string                 `json:"description"`
	Business_line_data    BusinessLineData       `json:"business_line_data"`
	Coverage_data         []CoverageData         `json:"coverage_data"`
	Tecnical_product_data []TechnicalProductData `json:"technical_product_data"`
}

type BusinessLineData struct {
	Business_line                BusinessLine                 `json:"business_line"`
	Product_payment_method       []ProductPaymentMethod       `json:"product_payment_method"`
	Product_renewal_cycle        []ProductRenewalCycle        `json:"product_renewal_cycle"`
	Commercial_network_attribute []CommercialNetworkAttribute `json:"commercial_network_attribute"`
	Renewal_parameter            []RenewalParameter           `json:"renewal_parameter"`
}

type CoverageData struct {
	Coverage                    Coverage                    `json:"coverage"`
	Capital                     []Capital                   `json:"capital"`
	Coverage_concept            []CoverageConcept           `json:"coverage_concept"`
	Franchise                   []Franchise                 `json:"franchise"`
	Subcoverage                 []Subcoverage               `json:"subcoverage"`
	Temporary_premium_reduction []TemporaryPremiumReduction `json:"temporary_premium_reduction"`
}

type BusinessLine struct {
	Ciaascod    string `json:"ciaascod"`
	Producto    string `json:"producto"`
	Ramopcod    string `json:"ramopcod"`
	Descripcion string `json:"descripcion"`
	Flujofor    string `json:"flujofor"`
	Antidias    int    `json:"antidias"`
	Cicloren    string `json:"cicloren"`
	Duracann    int    `json:"duracann"`
	Descridab   string `json:"descridab"`
	Duractip    string `json:"duractip"`
	Fechaefe    string `json:"fechaefe"`
	Fmprocol    string `json:"fmprocol"`
	Lsfljfor    string `json:"lsfljfor"`
	Ndiasdps    int    `json:"ndiasdps"`
	Prodbase    string `json:"prodbase"`
	Producer    string `json:"producer"`
	Productip   string `json:"productip"`
	Swactivo    bool   `json:"swactivo"`
	Swctlcb     bool   `json:"swxtlcb"`
	Swreacep    bool   `json:"swreacep"`
	Swsusbpm    bool   `json:"swsusbpm"`
	Swtramed    bool   `json:"swtramed"`
	Swtranot    bool   `json:"swtranot"`
	Swunicue    bool   `json:"swunicue"`
	Tablabas    string `json:"tablabas"`
	Tablacom    string `json:"tablacom"`
	Tipocole    string `json:"tipocole"`
	Tipoprod    string `json:"tipoprod"`
}

type ProductPaymentMethod struct {
	Mediopag string  `json:"mediopag"`
	Medireci string  `json:"medireci"`
	Mediexto string  `json:"mediexto"`
	Swexcole bool    `json:"swexcole"`
	Recimini float32 `json:"recimini"`
}

type ProductRenewalCycle struct {
	Cicloren string `json:"cicloren"`
}

type CommercialNetworkAttribute struct {
	Fieldcod    string `json:"fieldcod"`
	Descripcion string `json:"descripcion"`
	Orden       int    `json:"orden"`
	Swobliga    bool   `json:"swobliga"`
	Length      int    `json:"lenght"`
	Precision   int    `json:"Precision"`
	Tipodato    string `json:"tipodato"`
	Classhelp   string `json:"classhelp"`
}

type RenewalParameter struct {
	Swrenova bool   `json:"swrenova"`
	Proceren string `json:"proceren"`
	Swpreren bool   `json:"swpreren"`
	Ndiasren int    `json:"ndiasren"`
	Swvigrie bool   `json:"swvigrie"`
	Ndiaspre int    `json:"ndiaspre"`
	Ndiascnr int    `json:"ndiascnr"`
	Ndiascre int    `json:"ndiascre"`
	NdiasVri int    `json:"ndiasvri"`
}

type Coverage struct {
	Garancod    string `json:"garancod"`
	Gramodgs    string `json:"gramodgs"`
	Cobernum    int    `json:"cobernum"`
	Agrupacion  string `json:"agrupacion"`
	Garappal    string `json:"garappal"`
	Fechaefe    string `json:"fechaefe"`
	Fechbaja    string `json:"fechbaja"`
	Edadmini    int    `json:"edadmini"`
	Edadmaxi    int    `json:"edadmaxi"`
	Swpriniv    bool   `json:"swpriniv"`
	Swtranot    bool   `json:"swtranot"`
	Swopcdto    bool   `json:"swopcdto"`
	Swtarman    bool   `json:"swtarman"`
	Swbonif     bool   `json:"swbonif"`
	Swobliga    bool   `json:"swobliga"`
	Swfran      bool   `json:"swfran"`
	Swcapita    bool   `json:"swcapita"`
	Swcalpro    bool   `json:"swcalpro"`
	Swsimpro    bool   `json:"swsimpro"`
	Cobecmin    int    `json:"cobecmin"`
	Cobecmax    int    `json:"cobecmax"`
	Primamin    int    `json:"primamin"`
	Dedumaxi    int    `json:"dedumaxi"`
	Recamaxi    int    `json:"recamaxi"`
	Cobepmin    int    `json:"cobepmin"`
	Cobepmax    int    `json:"cobepmax"`
	Porcreva    int    `json:"porcreva"`
	Basestec    string `json:"basestec"`
	Swips       bool   `json:"swips"`
	Swclea      bool   `json:"swclea"`
	Edadsali    int    `json:"edadsali"`
	Swcons      bool   `json:"swcons"`
	Tipocapcons string `json:"tipocapcons"`
	Tiporecasir string `json:"tiporecasir"`
}

type Capital struct {
	Capitdat    string `json:"capitdat"`
	Descripcion string `json:"descripcion"`
	Swforpor    bool   `json:"swforpor"`
	Swforcap    bool   `json:"swforcap"`
}

type CoverageConcept struct {
	Conceimp  string `json:"conceimp"`
	Tipoimpu  string `json:"tipoimpu"`
	Swliqimpu bool   `json:"swliqimpu"`
	Gastocob  int    `json:"gastocob"`
}

type Franchise struct {
	Codifran    string `json:"codifran"`
	Descripcion string `json:"descripcion"`
	Diasfran    string `json:"diasfran"`
	Impofran    int    `json:"impofran"`
}

type Subcoverage struct {
	Cobernum int    `json:"cobernum"`
	Capiaseg int    `json:"capiaseg"`
	Impofran int    `json:"impofran"`
	Codsubco string `json:"codsubco"`
	Porccapi int    `json:"porccapi"`
	Porcrepa int    `json:"porcrepa"`
}

type TemporaryPremiumReduction struct {
	Diasanti int `json:"diasanti"`
	Porcprim int `json:"porcprim"`
}

type TechnicalProductData struct {
	Technical_product          TechnicalProduct           `json:"technical_product"`
	Technical_product_coverage []TechnicalProductCoverage `json:"technical_product_coverage"`
	Commercial_offer           []CommercialOffer          `json:"commercial_offer"`
}

type TechnicalProduct struct {
	Modalcod     string `json:"modalcod"`
	Cuessalu     string `json:"cuessalu"`
	Cicloren     string `json:"cicloren"`
	Swunica      bool   `json:"swunica"`
	Basestec     string `json:"basestec"`
	Ffincome     string `json:"ffincome"`
	Fechalta     string `json:"fechalta"`
	Duractip     string `json:"duractip"`
	Ramopcod     string `json:"ramopcod"`
	Swmensual    bool   `json:"swmensual"`
	Swanual      bool   `json:"swanual"`
	Swsemestral  bool   `json:"swsemestral"`
	Swtrimestral bool   `json:"swtrimestral"`
	Tipomoda     string `json:"tipomoda"`
	Orden        int    `json:"orden"`
	Descripcion  string `json:"descripcion"`
}

type TechnicalProductCoverage struct {
	Garancod string `json:"garancod"`
	Orden    int    `json:"orden"`
	Fechaefe string `json:"fechaefe"`
	Fechbaja string `json:"fechbaja"`
	Swobliga bool   `json:"swobliga"`
	Basestec string `json:"basestec"`
	Modalcob string `json:"modalcob"`
	Modulcob string `json:"modulcob"`
	Garappal string `json:"garappal"`
	Porcappi int    `json:"porcappi"`
	Tipocapi string `json:"tipocapi"`
	Tipodfra string `json:"tipodfra"`
	Capitmax int    `json:"capitmax"`
	Capitmin int    `json:"capitmin"`
	Tiporeva string `json:"tiporeva"`
	Revalpor int    `json:"revalpor"`
	Listacap string `json:"listacap"`
	Listafra string `json:"listafra"`
	Edadmini int    `json:"edadmini"`
	Edadsali int    `json:"edadsali"`
	Edadmaxi int    `json:"edadmaxi"`
	Capiaseg int    `json:"capiaseg"`
	Swambpol bool   `json:"swambpol"`
}
type CommercialOffer struct {
	Swcobper                  bool                      `json:"swcobper"`
	Fechaven                  string                    `json:"fechaven"`
	Basestec                  string                    `json:"basestec"`
	Fechaefe                  string                    `json:"fechaefe"`
	Puntomax                  int                       `json:"puntomax"`
	Tablaren                  string                    `json:"tablaren"`
	Swpolmar                  bool                      `json:"swpolmar"`
	Orden                     int                       `json:"orden"`
	Reembolso                 bool                      `json:"reembolso"`
	Swanual                   bool                      `json:"swanual"`
	Swsemestral               bool                      `json:"swsemestral"`
	Swtrimestral              bool                      `json:"swtrimestral"`
	Swunicue                  bool                      `json:"swunicue"`
	Cuessalu                  string                    `json:"cuessalu"`
	Swunica                   bool                      `json:"swunica"`
	Impofran                  int                       `json:"impofran"`
	Cobeprsa                  string                    `json:"cobeprsa"`
	Swbimestral               bool                      `json:"swbimestral"`
	Swactivo                  bool                      `json:"swactivo"`
	Swmensual                 bool                      `json:"swmensual"`
	Combicod                  string                    `json:"combicod"`
	Tipococo                  string                    `json:"tipococo"`
	Descripcion               string                    `json:"descripcion"`
	Ramopcod                  string                    `json:"ramopcod"`
	Swtramed                  bool                      `json:"swtramed"`
	Prefinum                  string                    `json:"prefinum"`
	Cgestora                  string                    `json:"cgestora"`
	Cccgesto                  string                    `json:"cccgesto"`
	Distribution_channel      []Distribution            `json:"distribution_channel"`
	Indexed_link_payment_date []IndexedLinkPaymentDate  `json:"indexed_link_payment_date"`
	InvestmentFund            []InvestmentFund          `json:"investment_fund"`
	Venture_capital           []VentureCapital          `json:"venture_capital"`
	Variable_data             []VariableData            `json:"variable_data"`
	Commercial_offer_coverage []CommercialOfferCoverage `json:"commercial_offer_coverage"`
}

type Distribution struct {
	Canalmed string `json:"canalmed"`
}

type IndexedLinkPaymentDate struct {
	Fechpago string `json:"fechpago"`
	Fechaob1 string `json:"fechaob1"`
	Fechaob2 string `json:"fechaob2"`
	Fechaob3 string `json:"fechaob3"`
	Fechaob4 string `json:"fechaob4"`
	Cupon    int    `json:"cupon"`
}

type InvestmentFund struct {
	Fondoinv    string `json:"fondoinv"`
	Descripcion string `json:"descripcion"`
}

type VentureCapital struct {
	Capiries   string `json:"capiries"`
	Orden      int    `json:"orden"`
	Swobliga   bool   `json:"swobliga"`
	Swreadonly bool   `json:"swreadonly"`
	Eonblur    string `json:"eonblur"`
}

type VariableData struct {
	Nombdato string `json:"nombdato"`
	Orden    int    `json:"orden"`
	Swobliga bool   `json:"swobliga"`
}

type CommercialOfferCoverage struct {
	Numevers       int                `json:"numevers"`
	Garancod       string             `json:"garancod"`
	Orden          int                `json:"orden"`
	Swmodifi       bool               `json:"swmodifi"`
	Swhide         bool               `json:"swhide"`
	Swcapita       bool               `json:"swcapita"`
	Swhidden       bool               `json:"swhidden"`
	Swoption       bool               `json:"swoption"`
	Cobepmax       int                `json:"cobepmax"`
	Cobepmin       int                `json:"cobepmin"`
	Basegral       string             `json:"basegral"`
	Compotar       string             `json:"compotar"`
	Swsercob       bool               `json:"swsercob"`
	Swtarifa       bool               `json:"swtarifa"`
	Cotarifa       string             `json:"cotarifa"`
	Fechaefe       string             `json:"fechaefe"`
	Fechbaja       string             `json:"fechbaja"`
	Swobliga       bool               `json:"swobliga"`
	Basestec       string             `json:"basestec"`
	Modalcob       string             `json:"modalcob"`
	Modulcob       string             `json:"modulcob"`
	Garappal       string             `json:"garappal"`
	Porcappi       int                `json:"porcappi"`
	Tipocapi       string             `json:"tipocapi"`
	Tipodfra       string             `json:"tipodfra"`
	Capitmax       int                `json:"capitmax"`
	Capitmin       int                `json:"capitmin"`
	Tiporeva       string             `json:"tiporeva"`
	Revalpor       int                `json:"revalpor"`
	Listacap       string             `json:"listacap"`
	Listafra       string             `json:"listafra"`
	Edadmini       int                `json:"edadmini"`
	Edadsali       int                `json:"edadsali"`
	Edadmaxi       int                `json:"edadmaxi"`
	Capiaseg       int                `json:"capiaseg"`
	Swambpol       bool               `json:"swambpol"`
	Limit          []Limit            `json:"limit"`
	Waiting_period []WaitingPeriod    `json:"waiting_period"`
	Partial_value  []PartialValue     `json:"partial_value"`
	Capital        []CapitalComercial `json:"capital"`
}

type Limit struct {
	Numevers int `json:"numevers"`
	Capitmax int `json:"capitmax"`
	Capitmin int `json:"capitmin"`
}

type WaitingPeriod struct {
	Numevers  int    `json:"numevers"`
	Carencia  string `json:"carencia"`
	Swdefault bool   `json:"swdefault"`
}

type PartialValue struct {
	Numevers    int    `json:"numevers"`
	Codigo      string `json:"codigo"`
	Orden       int    `json:"orden"`
	Valorpar    int    `json:"valorpar"`
	Descripcion string `json:"descripcion"`
}

type CapitalComercial struct {
	Numevers    int    `json:"numevers"`
	Capital     int    `json:"capital"`
	Descripcion string `json:"descripcion"`
	Tcapgara    string `json:"tcapgara"`
	Swdefault   bool   `json:"swdefault"`
	Formcapi    string `json:"formcapi"`
}
