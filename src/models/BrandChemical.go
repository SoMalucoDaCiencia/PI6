package models

type (
	ChemicalBrand string

	Brand struct {
		BrandId   uint64
		BrandName string
	}
)

const (
	Farma1Farma        ChemicalBrand = "1Farma"
	Abbott             ChemicalBrand = "Abbott"
	Abbvie             ChemicalBrand = "Abbvie"
	Accord             ChemicalBrand = "Accord"
	Ache               ChemicalBrand = "Ache"
	Actavis            ChemicalBrand = "Actavis"
	Addera             ChemicalBrand = "Addera"
	Alcon              ChemicalBrand = "Alcon"
	Allergan           ChemicalBrand = "Allergan"
	Althaia            ChemicalBrand = "Althaia"
	Amgen              ChemicalBrand = "Amgen"
	Apsen              ChemicalBrand = "Apsen"
	Arese              ChemicalBrand = "Arese"
	ArrowPharma        ChemicalBrand = "Arrow Pharma"
	AspenPharma        ChemicalBrand = "Aspen Pharma"
	Astellas           ChemicalBrand = "Astellas"
	Astrazeneca        ChemicalBrand = "Astrazeneca"
	Avert              ChemicalBrand = "Avert"
	Bago               ChemicalBrand = "Bago"
	Baldacci           ChemicalBrand = "Baldacci"
	BauschLomb         ChemicalBrand = "Bausch Lomb"
	Baxter             ChemicalBrand = "Baxter"
	Bayer              ChemicalBrand = "Bayer"
	Beiersdorf         ChemicalBrand = "Beiersdorf"
	Belcher            ChemicalBrand = "Belcher"
	Bergamo            ChemicalBrand = "Bergamo"
	BesinsHealthcare   ChemicalBrand = "Besins Healthcare"
	Biochimico         ChemicalBrand = "Biochimico"
	Biodinamica        ChemicalBrand = "Biodinamica"
	Biogen             ChemicalBrand = "Biogen"
	Biolab             ChemicalBrand = "Biolab"
	Biomm              ChemicalBrand = "Biomm"
	Bionatus           ChemicalBrand = "Bionatus"
	Biopas             ChemicalBrand = "Biopas"
	Biosintetica       ChemicalBrand = "Biosintetica"
	Blanver            ChemicalBrand = "Blanver"
	BlauFarmaceutica   ChemicalBrand = "Blau Farmaceutica"
	Boehringer         ChemicalBrand = "Boehringer"
	Bonifik            ChemicalBrand = "Bonifik"
	BracePharma        ChemicalBrand = "Brace Pharma"
	Brasterapica       ChemicalBrand = "Brasterapica"
	Bristol            ChemicalBrand = "Bristol"
	Cannten            ChemicalBrand = "Cannten"
	Carnot             ChemicalBrand = "Carnot"
	Catarinense        ChemicalBrand = "Catarinense"
	Cazi               ChemicalBrand = "Cazi"
	CelleraFarma       ChemicalBrand = "Cellera Farma"
	Chiesi             ChemicalBrand = "Chiesi"
	Cifarma            ChemicalBrand = "Cifarma"
	Cimed              ChemicalBrand = "Cimed"
	Cosmed             ChemicalBrand = "Cosmed"
	Cristalia          ChemicalBrand = "Cristalia"
	DaiichiSankyo      ChemicalBrand = "Daiichi Sankyo"
	Daudt              ChemicalBrand = "Daudt"
	Delta              ChemicalBrand = "Delta"
	Diffucap           ChemicalBrand = "Diffucap"
	Dismed             ChemicalBrand = "Dismed"
	DKT                ChemicalBrand = "DKT"
	DoctorReddys       ChemicalBrand = "Doctor Reddys"
	Dprev              ChemicalBrand = "Dprev"
	EaseLabs           ChemicalBrand = "Ease Labs"
	EliLilly           ChemicalBrand = "Eli Lilly"
	EliteDistribuidora ChemicalBrand = "Elite Distribuidora"
	Elofar             ChemicalBrand = "Elofar"
	EMS                ChemicalBrand = "EMS"
	Erimax             ChemicalBrand = "Erimax"
	Eurofarma          ChemicalBrand = "Eurofarma"
	Exeltis            ChemicalBrand = "Exeltis"
	Farmoquimica       ChemicalBrand = "Farmoquimica"
	Ferrer             ChemicalBrand = "Ferrer"
	Ferring            ChemicalBrand = "Ferring"
	FQMMelora          ChemicalBrand = "FQM Melora"
	Galderma           ChemicalBrand = "Galderma"
	Gallia             ChemicalBrand = "Gallia"
	Gbio               ChemicalBrand = "Gbio"
	Genom              ChemicalBrand = "Genom"
	Genzyme            ChemicalBrand = "Genzyme"
	Geolab             ChemicalBrand = "Geolab"
	Germed             ChemicalBrand = "Germed"
	Geyer              ChemicalBrand = "Geyer"
	GileadSciences     ChemicalBrand = "Gilead Sciences"
	Glenmark           ChemicalBrand = "Glenmark"
	GPZ                ChemicalBrand = "GPZ"
	GreenCare          ChemicalBrand = "GreenCare"
	Gross              ChemicalBrand = "Gross"
	Grunenthal         ChemicalBrand = "Grunenthal"
	GSK                ChemicalBrand = "GSK"
	Hebron             ChemicalBrand = "Hebron"
	Heel               ChemicalBrand = "Heel"
	Hemafarma          ChemicalBrand = "Hemafarma"
	Herbarium          ChemicalBrand = "Herbarium"
	HyperaPharma       ChemicalBrand = "Hypera Pharma"
	Hypermarcas        ChemicalBrand = "Hypermarcas"
	Ipsen              ChemicalBrand = "Ipsen"
	Janssen            ChemicalBrand = "Janssen"
	JpFarma            ChemicalBrand = "Jp Farma"
	KleyHertz          ChemicalBrand = "Kley Hertz"
	LabBrasBio         ChemicalBrand = "Lab BrasBio"
	LaboratorioGlobo   ChemicalBrand = "Laboratorio Globo"
	Latinofarma        ChemicalBrand = "Latinofarma"
	LegrandPharma      ChemicalBrand = "Legrand Pharma"
	LeoPharma          ChemicalBrand = "Leo Pharma"
	Libbs              ChemicalBrand = "Libbs"
	LOKELMA            ChemicalBrand = "LOKELMA"
	Lundbeck           ChemicalBrand = "Lundbeck"
	MabraFarmaceutica  ChemicalBrand = "Mabra Farmaceutica"
	MantecorpFarmasa   ChemicalBrand = "Mantecorp Farmasa"
	MantecorpSkincare  ChemicalBrand = "Mantecorp Skincare"
	Maquira            ChemicalBrand = "Maquira"
	Marjan             ChemicalBrand = "Marjan"
	Medley             ChemicalBrand = "Medley"
	Medquimica         ChemicalBrand = "Medquimica"
	Melcon             ChemicalBrand = "Melcon"
	Mepha              ChemicalBrand = "Mepha"
	Merck              ChemicalBrand = "Merck"
	Miorrelax          ChemicalBrand = "Miorrelax"
	Moksha8            ChemicalBrand = "Moksha8"
	Momenta            ChemicalBrand = "Momenta"
	MSD                ChemicalBrand = "MSD"
	Multilab           ChemicalBrand = "Multilab"
	Mundipharma        ChemicalBrand = "Mundipharma"
	Mylan              ChemicalBrand = "Mylan"
	Myralis            ChemicalBrand = "Myralis"
	Natcofarma         ChemicalBrand = "Natcofarma"
	Natulab            ChemicalBrand = "Natulab"
	NaturesPlus        ChemicalBrand = "Natures Plus"
	NDS                ChemicalBrand = "NDS"
	NeoQuimica         ChemicalBrand = "Neo Quimica"
	NovaQuimica        ChemicalBrand = "Nova Quimica"
	Novartis           ChemicalBrand = "Novartis"
	NovoNordisk        ChemicalBrand = "Novo Nordisk"
	Nunature           ChemicalBrand = "Nunature"
	Ophthalmos         ChemicalBrand = "Ophthalmos"
	Organon            ChemicalBrand = "Organon"
	PG                 ChemicalBrand = "P&G"
	Pfizer             ChemicalBrand = "Pfizer"
	Pharlab            ChemicalBrand = "Pharlab"
	Pharmascience      ChemicalBrand = "Pharmascience"
	PratiDonaduzzi     ChemicalBrand = "Prati Donaduzzi"
	Ranbaxy            ChemicalBrand = "Ranbaxy"
	Roche              ChemicalBrand = "Roche"
	SaborAlternativo   ChemicalBrand = "Sabor Alternativo"
	SamsungBioepis     ChemicalBrand = "Samsung Bioepis"
	Sandoz             ChemicalBrand = "Sandoz"
	Sanofi             ChemicalBrand = "Sanofi"
	Sanval             ChemicalBrand = "Sanval"
	Sauad              ChemicalBrand = "Sauad"
	SCJohnson          ChemicalBrand = "SC Johnson"
	ScheringPlough     ChemicalBrand = "Schering-Plough"
	Servier            ChemicalBrand = "Servier"
	Shire              ChemicalBrand = "Shire"
	Stiefel            ChemicalBrand = "Stiefel"
	SunPharma          ChemicalBrand = "Sun Pharma"
	SuperaRX           ChemicalBrand = "Supera RX"
	Suprafarma         ChemicalBrand = "Suprafarma"
	Takeda             ChemicalBrand = "Takeda"
	Teuto              ChemicalBrand = "Teuto"
	Teva               ChemicalBrand = "Teva"
	Theramex           ChemicalBrand = "Theramex"
	Theraskin          ChemicalBrand = "Theraskin"
	Torrent            ChemicalBrand = "Torrent"
	TRBPharma          ChemicalBrand = "TRB Pharma"
	Tylenol            ChemicalBrand = "Tylenol"
	UcbBiopharma       ChemicalBrand = "Ucb Biopharma"
	UciFarma           ChemicalBrand = "Uci-Farma"
	UniaoQuimica       ChemicalBrand = "Uniao Quimica"
	UnitedMedical      ChemicalBrand = "United Medical"
	Valeant            ChemicalBrand = "Valeant"
	Verdemed           ChemicalBrand = "Verdemed"
	Viatris            ChemicalBrand = "Viatris"
	Vitamedic          ChemicalBrand = "Vitamedic"
	VITAPAN            ChemicalBrand = "VITAPAN"
	Wyeth              ChemicalBrand = "Wyeth"
	Zambon             ChemicalBrand = "Zambon"
	Zion               ChemicalBrand = "Zion"
	Zodiac             ChemicalBrand = "Zodiac"
	ZydusNikkho        ChemicalBrand = "Zydus Nikkho"
)

func GetChemicalBrand() []ChemicalBrand {
	return []ChemicalBrand{
		Farma1Farma,
		Abbott,
		Abbvie,
		Accord,
		Ache,
		Actavis,
		Addera,
		Alcon,
		Allergan,
		Althaia,
		Amgen,
		Apsen,
		Arese,
		ArrowPharma,
		AspenPharma,
		Astellas,
		Astrazeneca,
		Avert,
		Bago,
		Baldacci,
		BauschLomb,
		Baxter,
		Bayer,
		Beiersdorf,
		Belcher,
		Bergamo,
		BesinsHealthcare,
		Biochimico,
		Biodinamica,
		Biogen,
		Biolab,
		Biomm,
		Bionatus,
		Biopas,
		Biosintetica,
		Blanver,
		BlauFarmaceutica,
		Boehringer,
		Bonifik,
		BracePharma,
		Brasterapica,
		Bristol,
		Cannten,
		Carnot,
		Catarinense,
		Cazi,
		CelleraFarma,
		Chiesi,
		Cifarma,
		Cimed,
		Cosmed,
		Cristalia,
		DaiichiSankyo,
		Daudt,
		Delta,
		Diffucap,
		Dismed,
		DKT,
		DoctorReddys,
		Dprev,
		EaseLabs,
		EliLilly,
		EliteDistribuidora,
		Elofar,
		EMS,
		Erimax,
		Eurofarma,
		Exeltis,
		Farmoquimica,
		Ferrer,
		Ferring,
		FQMMelora,
		Galderma,
		Gallia,
		Gbio,
		Genom,
		Genzyme,
		Geolab,
		Germed,
		Geyer,
		GileadSciences,
		Glenmark,
		GPZ,
		GreenCare,
		Gross,
		Grunenthal,
		GSK,
		Hebron,
		Heel,
		Hemafarma,
		Herbarium,
		HyperaPharma,
		Hypermarcas,
		Ipsen,
		Janssen,
		JpFarma,
		KleyHertz,
		LabBrasBio,
		LaboratorioGlobo,
		Latinofarma,
		LegrandPharma,
		LeoPharma,
		Libbs,
		LOKELMA,
		Lundbeck,
		MabraFarmaceutica,
		MantecorpFarmasa,
		MantecorpSkincare,
		Maquira,
		Marjan,
		Medley,
		Medquimica,
		Melcon,
		Mepha,
		Merck,
		Miorrelax,
		Moksha8,
		Momenta,
		MSD,
		Multilab,
		Mundipharma,
		Mylan,
		Myralis,
		Natcofarma,
		Natulab,
		NaturesPlus,
		NDS,
		NeoQuimica,
		NovaQuimica,
		Novartis,
		NovoNordisk,
		Nunature,
		Ophthalmos,
		Organon,
		PG,
		Pfizer,
		Pharlab,
		Pharmascience,
		PratiDonaduzzi,
		Ranbaxy,
		Roche,
		SaborAlternativo,
		SamsungBioepis,
		Sandoz,
		Sanofi,
		Sanval,
		Sauad,
		SCJohnson,
		ScheringPlough,
		Servier,
		Shire,
		Stiefel,
		SunPharma,
		SuperaRX,
		Suprafarma,
		Takeda,
		Teuto,
		Teva,
		Theramex,
		Theraskin,
		Torrent,
		TRBPharma,
		Tylenol,
		UcbBiopharma,
		UciFarma,
		UniaoQuimica,
		UnitedMedical,
		Valeant,
		Verdemed,
		Viatris,
		Vitamedic,
		VITAPAN,
		Wyeth,
		Zambon,
		Zion,
		Zodiac,
		ZydusNikkho,
	}
}
