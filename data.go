package main

type IsoEntry struct {
	enName  string
	frName  string
	alph2   string
	alph3   string
	numCode string
}

var (
	isocodes = map[string]IsoEntry{
		"AF": {
			enName:  "Afghanistan",
			frName:  "Afghanistan (l')",
			alph2:   "AF",
			alph3:   "AFG",
			numCode: "004",
		}, "AL": {
			enName:  "Albania",
			frName:  "Albanie (l')",
			alph2:   "AL",
			alph3:   "ALB",
			numCode: "008",
		}, "DZ": {
			enName:  "Algeria",
			frName:  "Algérie (l')",
			alph2:   "DZ",
			alph3:   "DZA",
			numCode: "012",
		}, "AS": {
			enName:  "American Samoa",
			frName:  "Samoa américaines (les)",
			alph2:   "AS",
			alph3:   "ASM",
			numCode: "016",
		}, "AD": {
			enName:  "Andorra",
			frName:  "Andorre (l')",
			alph2:   "AD",
			alph3:   "AND",
			numCode: "020",
		}, "AO": {
			enName:  "Angola",
			frName:  "Angola (l')",
			alph2:   "AO",
			alph3:   "AGO",
			numCode: "024",
		}, "AI": {
			enName:  "Anguilla",
			frName:  "Anguilla",
			alph2:   "AI",
			alph3:   "AIA",
			numCode: "660",
		}, "AQ": {
			enName:  "Antarctica",
			frName:  "Antarctique (l')",
			alph2:   "AQ",
			alph3:   "ATA",
			numCode: "010",
		}, "AG": {
			enName:  "Antigua and Barbuda",
			frName:  "Antigua-et-Barbuda",
			alph2:   "AG",
			alph3:   "ATG",
			numCode: "028",
		}, "AR": {
			enName:  "Argentina",
			frName:  "Argentine (l')",
			alph2:   "AR",
			alph3:   "ARG",
			numCode: "032",
		}, "AM": {
			enName:  "Armenia",
			frName:  "Arménie (l')",
			alph2:   "AM",
			alph3:   "ARM",
			numCode: "051",
		}, "AW": {
			enName:  "Aruba",
			frName:  "Aruba",
			alph2:   "AW",
			alph3:   "ABW",
			numCode: "533",
		}, "AU": {
			enName:  "Australia",
			frName:  "Australie (l')",
			alph2:   "AU",
			alph3:   "AUS",
			numCode: "036",
		}, "AT": {
			enName:  "Austria",
			frName:  "Autriche (l')",
			alph2:   "AT",
			alph3:   "AUT",
			numCode: "040",
		}, "AZ": {
			enName:  "Azerbaijan",
			frName:  "Azerbaïdjan (l')",
			alph2:   "AZ",
			alph3:   "AZE",
			numCode: "031",
		}, "BS": {
			enName:  "Bahamas (the)",
			frName:  "Bahamas (les)",
			alph2:   "BS",
			alph3:   "BHS",
			numCode: "044",
		}, "BH": {
			enName:  "Bahrain",
			frName:  "Bahreïn",
			alph2:   "BH",
			alph3:   "BHR",
			numCode: "048",
		}, "BD": {
			enName:  "Bangladesh",
			frName:  "Bangladesh (le)",
			alph2:   "BD",
			alph3:   "BGD",
			numCode: "050",
		}, "BB": {
			enName:  "Barbados",
			frName:  "Barbade (la)",
			alph2:   "BB",
			alph3:   "BRB",
			numCode: "052",
		}, "BY": {
			enName:  "Belarus",
			frName:  "Bélarus (le)",
			alph2:   "BY",
			alph3:   "BLR",
			numCode: "112",
		}, "BE": {
			enName:  "Belgium",
			frName:  "Belgique (la)",
			alph2:   "BE",
			alph3:   "BEL",
			numCode: "056",
		}, "BZ": {
			enName:  "Belize",
			frName:  "Belize (le)",
			alph2:   "BZ",
			alph3:   "BLZ",
			numCode: "084",
		}, "BJ": {
			enName:  "Benin",
			frName:  "Bénin (le)",
			alph2:   "BJ",
			alph3:   "BEN",
			numCode: "204",
		}, "BM": {
			enName:  "Bermuda",
			frName:  "Bermudes (les)",
			alph2:   "BM",
			alph3:   "BMU",
			numCode: "060",
		}, "AX": {
			enName:  "Åland Islands", // U+00C5
			frName:  "Åland(les Îles)",
			alph2:   "AX",
			alph3:   "ALA",
			numCode: "248",
		}, "BT": {
			enName:  "Bhutan",
			frName:  "Bhoutan (le)",
			alph2:   "BT",
			alph3:   "BTN",
			numCode: "064",
		}, "BO": {
			enName:  "Bolivia (Plurinational State of)",
			frName:  "Bolivie (État plurinational de)",
			alph2:   "BO",
			alph3:   "BOL",
			numCode: "068",
		}, "BQ": {
			enName:  "Bonaire, Sint Eustatius and SabaBonaire",
			frName:  " Saint-Eustache et Saba",
			alph2:   "BQ",
			alph3:   "BES",
			numCode: "535",
		}, "BA": {
			enName:  "Bosnia and Herzegovina",
			frName:  "Bosnie-Herzégovine (la)",
			alph2:   "BA",
			alph3:   "BIH",
			numCode: "070",
		}, "BW": {
			enName:  "Botswana",
			frName:  "Botswana (le)",
			alph2:   "BW",
			alph3:   "BWA",
			numCode: "072",
		}, "BV": {
			enName:  "Bouvet Island",
			frName:  "Bouvet (l'Île)",
			alph2:   "BV",
			alph3:   "BVT",
			numCode: "074",
		}, "BR": {
			enName:  "Brazil",
			frName:  "Brésil (le)",
			alph2:   "BR",
			alph3:   "BRA",
			numCode: "076",
		}, "IO": {
			enName:  "British Indian Ocean Territory (the)",
			frName:  "Indien (le Territoire britannique de l'océan)",
			alph2:   "IO",
			alph3:   "IOT",
			numCode: "086",
		}, "BN": {
			enName:  "Brunei Darussalam",
			frName:  "Brunéi Darussalam (le)",
			alph2:   "BN",
			alph3:   "BRN",
			numCode: "096",
		}, "BG": {
			enName:  "Bulgaria",
			frName:  "Bulgarie (la)",
			alph2:   "BG",
			alph3:   "BGR",
			numCode: "100",
		}, "BF": {
			enName:  "Burkina Faso",
			frName:  "Burkina Faso (le)",
			alph2:   "BF",
			alph3:   "BFA",
			numCode: "854",
		}, "BI": {
			enName:  "Burundi",
			frName:  "Burundi (le)",
			alph2:   "BI",
			alph3:   "BDI",
			numCode: "108",
		}, "CV": {
			enName:  "Cabo Verde",
			frName:  "Cabo Verde",
			alph2:   "CV",
			alph3:   "CPV",
			numCode: "132",
		}, "KH": {
			enName:  "Cambodia",
			frName:  "Cambodge (le)",
			alph2:   "KH",
			alph3:   "KHM",
			numCode: "116",
		}, "CM": {
			enName:  "Cameroon",
			frName:  "Cameroun (le)",
			alph2:   "CM",
			alph3:   "CMR",
			numCode: "120",
		}, "CA": {
			enName:  "Canada",
			frName:  "Canada (le)",
			alph2:   "CA",
			alph3:   "CAN",
			numCode: "124",
		}, "KY": {
			enName:  "Cayman Islands (the)",
			frName:  "Caïmans (les Îles)",
			alph2:   "KY",
			alph3:   "CYM",
			numCode: "136",
		}, "CF": {
			enName:  "Central African Republic (the)",
			frName:  "République centrafricaine (la)",
			alph2:   "CF",
			alph3:   "CAF",
			numCode: "140",
		}, "TD": {
			enName:  "Chad",
			frName:  "Tchad (le)",
			alph2:   "TD",
			alph3:   "TCD",
			numCode: "148",
		}, "CL": {
			enName:  "Chile",
			frName:  "Chili (le)",
			alph2:   "CL",
			alph3:   "CHL",
			numCode: "152",
		}, "CN": {
			enName:  "China",
			frName:  "Chine (la)",
			alph2:   "CN",
			alph3:   "CHN",
			numCode: "156",
		}, "CX": {
			enName:  "Christmas Island",
			frName:  "Christmas (l'Île)",
			alph2:   "CX",
			alph3:   "CXR",
			numCode: "162",
		}, "CC": {
			enName:  "Cocos (Keeling) Islands (the)",
			frName:  "Cocos (les Îles)/ Keeling (les Îles)",
			alph2:   "CC",
			alph3:   "CCK",
			numCode: "166",
		}, "CO": {
			enName:  "Colombia",
			frName:  "Colombie (la)",
			alph2:   "CO",
			alph3:   "COL",
			numCode: "170",
		}, "KM": {
			enName:  "Comoros (the)",
			frName:  "Comores (les)",
			alph2:   "KM",
			alph3:   "COM",
			numCode: "174",
		}, "CD": {
			enName:  "Congo (the Democratic Republic of the)",
			frName:  "Congo (la République démocratique du)",
			alph2:   "CD",
			alph3:   "COD",
			numCode: "180",
		}, "CG": {
			enName:  "Congo (the)",
			frName:  "Congo (le)",
			alph2:   "CG",
			alph3:   "COG",
			numCode: "178",
		}, "CK": {
			enName:  "Cook Islands (the)",
			frName:  "Cook (les Îles)",
			alph2:   "CK",
			alph3:   "COK",
			numCode: "184",
		}, "CR": {
			enName:  "Costa Rica",
			frName:  "Costa Rica (le)",
			alph2:   "CR",
			alph3:   "CRI",
			numCode: "188",
		}, "HR": {
			enName:  "Croatia",
			frName:  "Croatie (la)",
			alph2:   "HR",
			alph3:   "HRV",
			numCode: "191",
		}, "CU": {
			enName:  "Cuba",
			frName:  "Cuba",
			alph2:   "CU",
			alph3:   "CUB",
			numCode: "192",
		}, "CW": {
			enName:  "Curaçao",
			frName:  "Curaçao",
			alph2:   "CW",
			alph3:   "CUW",
			numCode: "531",
		}, "CY": {
			enName:  "Cyprus",
			frName:  "Chypre",
			alph2:   "CY",
			alph3:   "CYP",
			numCode: "196",
		}, "CZ": {
			enName:  "Czechia",
			frName:  "Tchéquie (la)",
			alph2:   "CZ",
			alph3:   "CZE",
			numCode: "203",
		}, "CI": {
			enName:  "Côte d'Ivoire",
			frName:  "Côte d'Ivoire (la)",
			alph2:   "CI",
			alph3:   "CIV",
			numCode: "384",
		}, "DK": {
			enName:  "Denmark",
			frName:  "Danemark (le)",
			alph2:   "DK",
			alph3:   "DNK",
			numCode: "208",
		}, "DJ": {
			enName:  "Djibouti",
			frName:  "Djibouti",
			alph2:   "DJ",
			alph3:   "DJI",
			numCode: "262",
		}, "DM": {
			enName:  "Dominica",
			frName:  "Dominique (la)",
			alph2:   "DM",
			alph3:   "DMA",
			numCode: "212",
		}, "DO": {
			enName:  "Dominican Republic (the)",
			frName:  "dominicaine (la République)",
			alph2:   "DO",
			alph3:   "DOM",
			numCode: "214",
		}, "EC": {
			enName:  "Ecuador",
			frName:  "Équateur (l')",
			alph2:   "EC",
			alph3:   "ECU",
			numCode: "218",
		}, "EG": {
			enName:  "Egypt",
			frName:  "Égypte (l')",
			alph2:   "EG",
			alph3:   "EGY",
			numCode: "818",
		}, "SV": {
			enName:  "El Salvador",
			frName:  "El Salvador",
			alph2:   "SV",
			alph3:   "SLV",
			numCode: "222",
		}, "GQ": {
			enName:  "Equatorial Guinea",
			frName:  "Guinée équatoriale (la)",
			alph2:   "GQ",
			alph3:   "GNQ",
			numCode: "226",
		}, "ER": {
			enName:  "Eritrea",
			frName:  "Érythrée (l')",
			alph2:   "ER",
			alph3:   "ERI",
			numCode: "232",
		}, "EE": {
			enName:  "Estonia",
			frName:  "Estonie (l')",
			alph2:   "EE",
			alph3:   "EST",
			numCode: "233",
		}, "SZ": {
			enName:  "Eswatini",
			frName:  "Eswatini (l')",
			alph2:   "SZ",
			alph3:   "SWZ",
			numCode: "748",
		}, "ET": {
			enName:  "Ethiopia",
			frName:  "Éthiopie (l')",
			alph2:   "ET",
			alph3:   "ETH",
			numCode: "231",
		}, "FK": {
			enName:  "Falkland Islands (the) [Malvinas]",
			frName:  "Falkland (les Îles)/Malouines (les Îles)",
			alph2:   "FK",
			alph3:   "FLK",
			numCode: "238",
		}, "FO": {
			enName:  "Faroe Islands (the)",
			frName:  "Féroé (les Îles)",
			alph2:   "FO",
			alph3:   "FRO",
			numCode: "234",
		}, "FJ": {
			enName:  "Fiji",
			frName:  "Fidji (les)",
			alph2:   "FJ",
			alph3:   "FJI",
			numCode: "242",
		}, "FI": {
			enName:  "Finland",
			frName:  "Finlande (la)",
			alph2:   "FI",
			alph3:   "FIN",
			numCode: "246",
		}, "FR": {
			enName:  "France",
			frName:  "France (la)",
			alph2:   "FR",
			alph3:   "FRA",
			numCode: "250",
		}, "GF": {
			enName:  "French Guiana",
			frName:  "Guyane française (la )",
			alph2:   "GF",
			alph3:   "GUF",
			numCode: "254",
		}, "PF": {
			enName:  "French Polynesia",
			frName:  "Polynésie française (la)",
			alph2:   "PF",
			alph3:   "PYF",
			numCode: "258",
		}, "TF": {
			enName:  "French Southern Territories (the)",
			frName:  "Terres australes françaises (les)",
			alph2:   "TF",
			alph3:   "ATF",
			numCode: "260",
		}, "GA": {
			enName:  "Gabon",
			frName:  "Gabon (le)",
			alph2:   "GA",
			alph3:   "GAB",
			numCode: "266",
		}, "GM": {
			enName:  "Gambia (the)",
			frName:  "Gambie (la)",
			alph2:   "GM",
			alph3:   "GMB",
			numCode: "270",
		}, "GE": {
			enName:  "Georgia",
			frName:  "Géorgie (la)",
			alph2:   "GE",
			alph3:   "GEO",
			numCode: "268",
		}, "DE": {
			enName:  "Germany",
			frName:  "Allemagne (l')",
			alph2:   "DE",
			alph3:   "DEU",
			numCode: "276",
		}, "GH": {
			enName:  "Ghana",
			frName:  "Ghana (le)",
			alph2:   "GH",
			alph3:   "GHA",
			numCode: "288",
		}, "GI": {
			enName:  "Gibraltar",
			frName:  "Gibraltar",
			alph2:   "GI",
			alph3:   "GIB",
			numCode: "292",
		}, "GR": {
			enName:  "Greece",
			frName:  "Grèce (la)",
			alph2:   "GR",
			alph3:   "GRC",
			numCode: "300",
		}, "GL": {
			enName:  "Greenland",
			frName:  "Groenland (le)",
			alph2:   "GL",
			alph3:   "GRL",
			numCode: "304",
		}, "GD": {
			enName:  "Grenada",
			frName:  "Grenade (la)",
			alph2:   "GD",
			alph3:   "GRD",
			numCode: "308",
		}, "GP": {
			enName:  "Guadeloupe",
			frName:  "Guadeloupe (la)",
			alph2:   "GP",
			alph3:   "GLP",
			numCode: "312",
		}, "GU": {
			enName:  "Guam",
			frName:  "Guam",
			alph2:   "GU",
			alph3:   "GUM",
			numCode: "316",
		}, "GT": {
			enName:  "Guatemala",
			frName:  "Guatemala (le)",
			alph2:   "GT",
			alph3:   "GTM",
			numCode: "320",
		}, "GG": {
			enName:  "Guernsey",
			frName:  "Guernesey",
			alph2:   "GG",
			alph3:   "GGY",
			numCode: "831",
		}, "GN": {
			enName:  "Guinea",
			frName:  "Guinée (la)",
			alph2:   "GN",
			alph3:   "GIN",
			numCode: "324",
		}, "GW": {
			enName:  "Guinea-Bissau",
			frName:  "Guinée-Bissau (la)",
			alph2:   "GW",
			alph3:   "GNB",
			numCode: "624",
		}, "GY": {
			enName:  "Guyana",
			frName:  "Guyana (le)",
			alph2:   "GY",
			alph3:   "GUY",
			numCode: "328",
		}, "HT": {
			enName:  "Haiti",
			frName:  "Haïti",
			alph2:   "HT",
			alph3:   "HTI",
			numCode: "332",
		}, "HM": {
			enName:  "Heard Island and McDonald Islands",
			frName:  "Heard-et-Îles MacDonald (l'Île)",
			alph2:   "HM",
			alph3:   "HMD",
			numCode: "334",
		}, "VA": {
			enName:  "Holy See (the)",
			frName:  "Saint-Siège (le)",
			alph2:   "VA",
			alph3:   "VAT",
			numCode: "336",
		}, "HN": {
			enName:  "Honduras",
			frName:  "Honduras (le)",
			alph2:   "HN",
			alph3:   "HND",
			numCode: "340",
		}, "HK": {
			enName:  "Hong Kong",
			frName:  "Hong Kong",
			alph2:   "HK",
			alph3:   "HKG",
			numCode: "344",
		}, "HU": {
			enName:  "Hungary",
			frName:  "Hongrie (la)",
			alph2:   "HU",
			alph3:   "HUN",
			numCode: "348",
		}, "IS": {
			enName:  "Iceland",
			frName:  "Islande (l')",
			alph2:   "IS",
			alph3:   "ISL",
			numCode: "352",
		}, "IN": {
			enName:  "India",
			frName:  "Inde (l')",
			alph2:   "IN",
			alph3:   "IND",
			numCode: "356",
		}, "ID": {
			enName:  "Indonesia",
			frName:  "Indonésie (l')",
			alph2:   "ID",
			alph3:   "IDN",
			numCode: "360",
		}, "IR": {
			enName:  "Iran (Islamic Republic of)",
			frName:  "Iran (République Islamique d')",
			alph2:   "IR",
			alph3:   "IRN",
			numCode: "364",
		}, "IQ": {
			enName:  "Iraq",
			frName:  "Iraq (l')",
			alph2:   "IQ",
			alph3:   "IRQ",
			numCode: "368",
		}, "IE": {
			enName:  "Ireland",
			frName:  "Irlande (l')",
			alph2:   "IE",
			alph3:   "IRL",
			numCode: "372",
		}, "IM": {
			enName:  "Isle of Man",
			frName:  "Île de Man",
			alph2:   "IM",
			alph3:   "IMN",
			numCode: "833",
		}, "IL": {
			enName:  "Israel",
			frName:  "Israël",
			alph2:   "IL",
			alph3:   "ISR",
			numCode: "376",
		}, "IT": {
			enName:  "Italy",
			frName:  "Italie (l')",
			alph2:   "IT",
			alph3:   "ITA",
			numCode: "380",
		}, "JM": {
			enName:  "Jamaica",
			frName:  "Jamaïque (la)",
			alph2:   "JM",
			alph3:   "JAM",
			numCode: "388",
		}, "JP": {
			enName:  "Japan",
			frName:  "Japon (le)",
			alph2:   "JP",
			alph3:   "JPN",
			numCode: "392",
		}, "JE": {
			enName:  "Jersey",
			frName:  "Jersey",
			alph2:   "JE",
			alph3:   "JEY",
			numCode: "832",
		}, "JO": {
			enName:  "Jordan",
			frName:  "Jordanie (la)",
			alph2:   "JO",
			alph3:   "JOR",
			numCode: "400",
		}, "KZ": {
			enName:  "Kazakhstan",
			frName:  "Kazakhstan (le)",
			alph2:   "KZ",
			alph3:   "KAZ",
			numCode: "398",
		}, "KE": {
			enName:  "Kenya",
			frName:  "Kenya (le)",
			alph2:   "KE",
			alph3:   "KEN",
			numCode: "404",
		}, "KI": {
			enName:  "Kiribati",
			frName:  "Kiribati",
			alph2:   "KI",
			alph3:   "KIR",
			numCode: "296",
		}, "KP": {
			enName:  "Korea (the Democratic People's Republic of)",
			frName:  "Corée (la République populaire démocratique de)",
			alph2:   "KP",
			alph3:   "PRK",
			numCode: "408",
		}, "KR": {
			enName:  "Korea (the Republic of)",
			frName:  "Corée (la République de)",
			alph2:   "KR",
			alph3:   "KOR",
			numCode: "410",
		}, "KW": {
			enName:  "Kuwait",
			frName:  "Koweït (le)",
			alph2:   "KW",
			alph3:   "KWT",
			numCode: "414",
		}, "KG": {
			enName:  "Kyrgyzstan",
			frName:  "Kirghizistan (le)",
			alph2:   "KG",
			alph3:   "KGZ",
			numCode: "417",
		}, "LA": {
			enName:  "Lao People's Democratic Republic (the)",
			frName:  "Lao (la République démocratique populaire)",
			alph2:   "LA",
			alph3:   "LAO",
			numCode: "418",
		}, "LV": {
			enName:  "Latvia",
			frName:  "Lettonie (la)",
			alph2:   "LV",
			alph3:   "LVA",
			numCode: "428",
		}, "LB": {
			enName:  "Lebanon",
			frName:  "Liban (le)",
			alph2:   "LB",
			alph3:   "LBN",
			numCode: "422",
		}, "LS": {
			enName:  "Lesotho",
			frName:  "Lesotho (le)",
			alph2:   "LS",
			alph3:   "LSO",
			numCode: "426",
		}, "LR": {
			enName:  "Liberia",
			frName:  "Libéria (le)",
			alph2:   "LR",
			alph3:   "LBR",
			numCode: "430",
		}, "LY": {
			enName:  "Libya",
			frName:  "Libye (la)",
			alph2:   "LY",
			alph3:   "LBY",
			numCode: "434",
		}, "LI": {
			enName:  "Liechtenstein",
			frName:  "Liechtenstein (le)",
			alph2:   "LI",
			alph3:   "LIE",
			numCode: "438",
		}, "LT": {
			enName:  "Lithuania",
			frName:  "Lituanie (la)",
			alph2:   "LT",
			alph3:   "LTU",
			numCode: "440",
		}, "LU": {
			enName:  "Luxembourg",
			frName:  "Luxembourg (le)",
			alph2:   "LU",
			alph3:   "LUX",
			numCode: "442",
		}, "MO": {
			enName:  "Macao",
			frName:  "Macao",
			alph2:   "MO",
			alph3:   "MAC",
			numCode: "446",
		}, "MG": {
			enName:  "Madagascar",
			frName:  "Madagascar",
			alph2:   "MG",
			alph3:   "MDG",
			numCode: "450",
		}, "MW": {
			enName:  "Malawi",
			frName:  "Malawi (le)",
			alph2:   "MW",
			alph3:   "MWI",
			numCode: "454",
		}, "MY": {
			enName:  "Malaysia",
			frName:  "Malaisie (la)",
			alph2:   "MY",
			alph3:   "MYS",
			numCode: "458",
		}, "MV": {
			enName:  "Maldives",
			frName:  "Maldives (les)",
			alph2:   "MV",
			alph3:   "MDV",
			numCode: "462",
		}, "ML": {
			enName:  "Mali",
			frName:  "Mali (le)",
			alph2:   "ML",
			alph3:   "MLI",
			numCode: "466",
		}, "MT": {
			enName:  "Malta",
			frName:  "Malte",
			alph2:   "MT",
			alph3:   "MLT",
			numCode: "470",
		}, "MH": {
			enName:  "Marshall Islands (the)",
			frName:  "Marshall (les Îles)",
			alph2:   "MH",
			alph3:   "MHL",
			numCode: "584",
		}, "MQ": {
			enName:  "Martinique",
			frName:  "Martinique (la)",
			alph2:   "MQ",
			alph3:   "MTQ",
			numCode: "474",
		}, "MR": {
			enName:  "Mauritania",
			frName:  "Mauritanie (la)",
			alph2:   "MR",
			alph3:   "MRT",
			numCode: "478",
		}, "MU": {
			enName:  "Mauritius",
			frName:  "Maurice",
			alph2:   "MU",
			alph3:   "MUS",
			numCode: "480",
		}, "YT": {
			enName:  "Mayotte",
			frName:  "Mayotte",
			alph2:   "YT",
			alph3:   "MYT",
			numCode: "175",
		}, "MX": {
			enName:  "Mexico",
			frName:  "Mexique (le)",
			alph2:   "MX",
			alph3:   "MEX",
			numCode: "484",
		}, "FM": {
			enName:  "Micronesia (Federated States of)",
			frName:  "Micronésie (États fédérés de)",
			alph2:   "FM",
			alph3:   "FSM",
			numCode: "583",
		}, "MD": {
			enName:  "Moldova (the Republic of)",
			frName:  "Moldova (la République de)",
			alph2:   "MD",
			alph3:   "MDA",
			numCode: "498",
		}, "MC": {
			enName:  "Monaco",
			frName:  "Monaco",
			alph2:   "MC",
			alph3:   "MCO",
			numCode: "492",
		}, "MN": {
			enName:  "Mongolia",
			frName:  "Mongolie (la)",
			alph2:   "MN",
			alph3:   "MNG",
			numCode: "496",
		}, "ME": {
			enName:  "Montenegro",
			frName:  "Monténégro (le)",
			alph2:   "ME",
			alph3:   "MNE",
			numCode: "499",
		}, "MS": {
			enName:  "Montserrat",
			frName:  "Montserrat",
			alph2:   "MS",
			alph3:   "MSR",
			numCode: "500",
		}, "MA": {
			enName:  "Morocco",
			frName:  "Maroc (le)",
			alph2:   "MA",
			alph3:   "MAR",
			numCode: "504",
		}, "MZ": {
			enName:  "Mozambique",
			frName:  "Mozambique (le)",
			alph2:   "MZ",
			alph3:   "MOZ",
			numCode: "508",
		}, "MM": {
			enName:  "Myanmar",
			frName:  "Myanmar (le)",
			alph2:   "MM",
			alph3:   "MMR",
			numCode: "104",
		}, "NA": {
			enName:  "Namibia",
			frName:  "Namibie (la)",
			alph2:   "NA",
			alph3:   "NAM",
			numCode: "516",
		}, "NR": {
			enName:  "Nauru",
			frName:  "Nauru",
			alph2:   "NR",
			alph3:   "NRU",
			numCode: "520",
		}, "NP": {
			enName:  "Nepal",
			frName:  "Népal (le)",
			alph2:   "NP",
			alph3:   "NPL",
			numCode: "524",
		}, "NL": {
			enName:  "Netherlands (the)",
			frName:  "Pays-Bas (les)",
			alph2:   "NL",
			alph3:   "NLD",
			numCode: "528",
		}, "NC": {
			enName:  "New Caledonia",
			frName:  "Nouvelle-Calédonie (la)",
			alph2:   "NC",
			alph3:   "NCL",
			numCode: "540",
		}, "NZ": {
			enName:  "New Zealand",
			frName:  "Nouvelle-Zélande (la)",
			alph2:   "NZ",
			alph3:   "NZL",
			numCode: "554",
		}, "NI": {
			enName:  "Nicaragua",
			frName:  "Nicaragua (le)",
			alph2:   "NI",
			alph3:   "NIC",
			numCode: "558",
		}, "NE": {
			enName:  "Niger (the)",
			frName:  "Niger (le)",
			alph2:   "NE",
			alph3:   "NER",
			numCode: "562",
		}, "NG": {
			enName:  "Nigeria",
			frName:  "Nigéria (le)",
			alph2:   "NG",
			alph3:   "NGA",
			numCode: "566",
		}, "NU": {
			enName:  "Niue",
			frName:  "Niue",
			alph2:   "NU",
			alph3:   "NIU",
			numCode: "570",
		}, "NF": {
			enName:  "Norfolk Island",
			frName:  "Norfolk (l'Île)",
			alph2:   "NF",
			alph3:   "NFK",
			numCode: "574",
		}, "MK": {
			enName:  "North Macedonia",
			frName:  "Macédoine du Nord (la)",
			alph2:   "MK",
			alph3:   "MKD",
			numCode: "807",
		}, "MP": {
			enName:  "Northern Mariana Islands (the)",
			frName:  "Mariannes du Nord (les Îles)",
			alph2:   "MP",
			alph3:   "MNP",
			numCode: "580",
		}, "NO": {
			enName:  "Norway",
			frName:  "Norvège (la)",
			alph2:   "NO",
			alph3:   "NOR",
			numCode: "578",
		}, "OM": {
			enName:  "Oman",
			frName:  "Oman",
			alph2:   "OM",
			alph3:   "OMN",
			numCode: "512",
		}, "PK": {
			enName:  "Pakistan",
			frName:  "Pakistan (le)",
			alph2:   "PK",
			alph3:   "PAK",
			numCode: "586",
		}, "PW": {
			enName:  "Palau",
			frName:  "Palaos (les)",
			alph2:   "PW",
			alph3:   "PLW",
			numCode: "585",
		}, "PS": {
			enName:  "Palestine, State of",
			frName:  "Palestine, État de",
			alph2:   "PS",
			alph3:   "PSE",
			numCode: "275",
		}, "PA": {
			enName:  "Panama",
			frName:  "Panama (le)",
			alph2:   "PA",
			alph3:   "PAN",
			numCode: "591",
		}, "PG": {
			enName:  "Papua New Guinea",
			frName:  "Papouasie-Nouvelle-Guinée (la)",
			alph2:   "PG",
			alph3:   "PNG",
			numCode: "598",
		}, "PY": {
			enName:  "Paraguay",
			frName:  "Paraguay (le)",
			alph2:   "PY",
			alph3:   "PRY",
			numCode: "600",
		}, "PE": {
			enName:  "Peru",
			frName:  "Pérou (le)",
			alph2:   "PE",
			alph3:   "PER",
			numCode: "604",
		}, "PH": {
			enName:  "Philippines (the)",
			frName:  "Philippines (les)",
			alph2:   "PH",
			alph3:   "PHL",
			numCode: "608",
		}, "PN": {
			enName:  "Pitcairn",
			frName:  "Pitcairn",
			alph2:   "PN",
			alph3:   "PCN",
			numCode: "612",
		}, "PL": {
			enName:  "Poland",
			frName:  "Pologne (la)",
			alph2:   "PL",
			alph3:   "POL",
			numCode: "616",
		}, "PT": {
			enName:  "Portugal",
			frName:  "Portugal (le)",
			alph2:   "PT",
			alph3:   "PRT",
			numCode: "620",
		}, "PR": {
			enName:  "Puerto Rico",
			frName:  "Porto Rico",
			alph2:   "PR",
			alph3:   "PRI",
			numCode: "630",
		}, "QA": {
			enName:  "Qatar",
			frName:  "Qatar (le)",
			alph2:   "QA",
			alph3:   "QAT",
			numCode: "634",
		}, "RO": {
			enName:  "Romania",
			frName:  "Roumanie (la)",
			alph2:   "RO",
			alph3:   "ROU",
			numCode: "642",
		}, "RU": {
			enName:  "Russian Federation (the)",
			frName:  "Russie (la Fédération de)",
			alph2:   "RU",
			alph3:   "RUS",
			numCode: "643",
		}, "RW": {
			enName:  "Rwanda",
			frName:  "Rwanda (le)",
			alph2:   "RW",
			alph3:   "RWA",
			numCode: "646",
		}, "RE": {
			enName:  "Réunion",
			frName:  "Réunion (La)",
			alph2:   "RE",
			alph3:   "REU",
			numCode: "638",
		}, "BL": {
			enName:  "Saint Barthélemy",
			frName:  "Saint-Barthélemy",
			alph2:   "BL",
			alph3:   "BLM",
			numCode: "652",
		}, "SH": {
			enName:  "Saint Helena, Ascension and Tristan da Cunha",
			frName:  "Sainte-Hélène, Ascension et Tristan da Cunha",
			alph2:   "SH",
			alph3:   "SHN",
			numCode: "654",
		}, "KN": {
			enName:  "Saint Kitts and Nevis",
			frName:  "Saint-Kitts-et-Nevis",
			alph2:   "KN",
			alph3:   "KNA",
			numCode: "659",
		}, "LC": {
			enName:  "Saint Lucia",
			frName:  "Sainte-Lucie",
			alph2:   "LC",
			alph3:   "LCA",
			numCode: "662",
		}, "MF": {
			enName:  "Saint Martin (French part)",
			frName:  "Saint-Martin (partie française)",
			alph2:   "MF",
			alph3:   "MAF",
			numCode: "663",
		}, "PM": {
			enName:  "Saint Pierre and Miquelon",
			frName:  "Saint-Pierre-et-Miquelon",
			alph2:   "PM",
			alph3:   "SPM",
			numCode: "666",
		}, "VC": {
			enName:  "Saint Vincent and the Grenadines",
			frName:  "Saint-Vincent-et-les Grenadines",
			alph2:   "VC",
			alph3:   "VCT",
			numCode: "670",
		}, "WS": {
			enName:  "Samoa",
			frName:  "Samoa (le)",
			alph2:   "WS",
			alph3:   "WSM",
			numCode: "882",
		}, "SM": {
			enName:  "San Marino",
			frName:  "Saint-Marin",
			alph2:   "SM",
			alph3:   "SMR",
			numCode: "674",
		}, "ST": {
			enName:  "Sao Tome and Principe",
			frName:  "Sao Tomé-et-Principe",
			alph2:   "ST",
			alph3:   "STP",
			numCode: "678",
		}, "SA": {
			enName:  "Saudi Arabia",
			frName:  "Arabie saoudite (l')",
			alph2:   "SA",
			alph3:   "SAU",
			numCode: "682",
		}, "SN": {
			enName:  "Senegal",
			frName:  "Sénégal (le)",
			alph2:   "SN",
			alph3:   "SEN",
			numCode: "686",
		}, "RS": {
			enName:  "Serbia",
			frName:  "Serbie (la)",
			alph2:   "RS",
			alph3:   "SRB",
			numCode: "688",
		}, "SC": {
			enName:  "Seychelles",
			frName:  "Seychelles (les)",
			alph2:   "SC",
			alph3:   "SYC",
			numCode: "690",
		}, "SL": {
			enName:  "Sierra Leone",
			frName:  "Sierra Leone (la)",
			alph2:   "SL",
			alph3:   "SLE",
			numCode: "694",
		}, "SG": {
			enName:  "Singapore",
			frName:  "Singapour",
			alph2:   "SG",
			alph3:   "SGP",
			numCode: "702",
		}, "SX": {
			enName:  "Sint Maarten (Dutch part)",
			frName:  "Saint-Martin (partie néerlandaise)",
			alph2:   "SX",
			alph3:   "SXM",
			numCode: "534",
		}, "SK": {
			enName:  "Slovakia",
			frName:  "Slovaquie (la)",
			alph2:   "SK",
			alph3:   "SVK",
			numCode: "703",
		}, "SI": {
			enName:  "Slovenia",
			frName:  "Slovénie (la)",
			alph2:   "SI",
			alph3:   "SVN",
			numCode: "705",
		}, "SB": {
			enName:  "Solomon Islands",
			frName:  "Salomon (les Îles)",
			alph2:   "SB",
			alph3:   "SLB",
			numCode: "090",
		}, "SO": {
			enName:  "Somalia",
			frName:  "Somalie (la)",
			alph2:   "SO",
			alph3:   "SOM",
			numCode: "706",
		}, "ZA": {
			enName:  "South Africa",
			frName:  "Afrique du Sud (l')",
			alph2:   "ZA",
			alph3:   "ZAF",
			numCode: "710",
		}, "GS": {
			enName:  "South Georgia and the South Sandwich Islands",
			frName:  "Géorgie du Sud-et-les Îles Sandwich du Sud (la)",
			alph2:   "GS",
			alph3:   "SGS",
			numCode: "239",
		}, "SS": {
			enName:  "South Sudan",
			frName:  "Soudan du Sud (le)",
			alph2:   "SS",
			alph3:   "SSD",
			numCode: "728",
		}, "ES": {
			enName:  "Spain",
			frName:  "Espagne (l')",
			alph2:   "ES",
			alph3:   "ESP",
			numCode: "724",
		}, "LK": {
			enName:  "Sri Lanka",
			frName:  "Sri Lanka",
			alph2:   "LK",
			alph3:   "LKA",
			numCode: "144",
		}, "SD": {
			enName:  "Sudan (the)",
			frName:  "Soudan (le)",
			alph2:   "SD",
			alph3:   "SDN",
			numCode: "729",
		}, "SR": {
			enName:  "Suriname",
			frName:  "Suriname (le)",
			alph2:   "SR",
			alph3:   "SUR",
			numCode: "740",
		}, "SJ": {
			enName:  "Svalbard and Jan Mayen",
			frName:  "Svalbard et l'Île Jan Mayen (le)",
			alph2:   "SJ",
			alph3:   "SJM",
			numCode: "744",
		}, "SE": {
			enName:  "Sweden",
			frName:  "Suède (la)",
			alph2:   "SE",
			alph3:   "SWE",
			numCode: "752",
		}, "CH": {
			enName:  "Switzerland",
			frName:  "Suisse (la)",
			alph2:   "CH",
			alph3:   "CHE",
			numCode: "756",
		}, "SY": {
			enName:  "Syrian Arab Republic (the)",
			frName:  "République arabe syrienne (la)",
			alph2:   "SY",
			alph3:   "SYR",
			numCode: "760",
		}, "TW": {
			enName:  "Taiwan (Province of China)",
			frName:  "Taïwan (Province de Chine)",
			alph2:   "TW",
			alph3:   "TWN",
			numCode: "158",
		}, "TJ": {
			enName:  "Tajikistan",
			frName:  "Tadjikistan (le)",
			alph2:   "TJ",
			alph3:   "TJK",
			numCode: "762",
		}, "TZ": {
			enName:  "Tanzania, the United Republic of",
			frName:  "Tanzanie (la République-Unie de)",
			alph2:   "TZ",
			alph3:   "TZA",
			numCode: "834",
		}, "TH": {
			enName:  "Thailand",
			frName:  "Thaïlande (la)",
			alph2:   "TH",
			alph3:   "THA",
			numCode: "764",
		}, "TL": {
			enName:  "Timor-Leste",
			frName:  "Timor-Leste (le)",
			alph2:   "TL",
			alph3:   "TLS",
			numCode: "626",
		}, "TG": {
			enName:  "Togo",
			frName:  "Togo (le)",
			alph2:   "TG",
			alph3:   "TGO",
			numCode: "768",
		}, "TK": {
			enName:  "Tokelau",
			frName:  "Tokelau (les)",
			alph2:   "TK",
			alph3:   "TKL",
			numCode: "772",
		}, "TO": {
			enName:  "Tonga",
			frName:  "Tonga (les)",
			alph2:   "TO",
			alph3:   "TON",
			numCode: "776",
		}, "TT": {
			enName:  "Trinidad and Tobago",
			frName:  "Trinité-et-Tobago (la)",
			alph2:   "TT",
			alph3:   "TTO",
			numCode: "780",
		}, "TN": {
			enName:  "Tunisia",
			frName:  "Tunisie (la)",
			alph2:   "TN",
			alph3:   "TUN",
			numCode: "788",
		}, "TR": {
			enName:  "Turkey",
			frName:  "Turquie (la)",
			alph2:   "TR",
			alph3:   "TUR",
			numCode: "792",
		}, "TM": {
			enName:  "Turkmenistan",
			frName:  "Turkménistan (le)",
			alph2:   "TM",
			alph3:   "TKM",
			numCode: "795",
		}, "TC": {
			enName:  "Turks and Caicos Islands (the)",
			frName:  "Turks-et-Caïcos (les Îles)",
			alph2:   "TC",
			alph3:   "TCA",
			numCode: "796",
		}, "TV": {
			enName:  "Tuvalu",
			frName:  "Tuvalu (les)",
			alph2:   "TV",
			alph3:   "TUV",
			numCode: "798",
		}, "UG": {
			enName:  "Uganda",
			frName:  "Ouganda (l')",
			alph2:   "UG",
			alph3:   "UGA",
			numCode: "800",
		}, "UA": {
			enName:  "Ukraine",
			frName:  "Ukraine (l')",
			alph2:   "UA",
			alph3:   "UKR",
			numCode: "804",
		}, "AE": {
			enName:  "United Arab Emirates (the)",
			frName:  "Émirats arabes unis (les)",
			alph2:   "AE",
			alph3:   "ARE",
			numCode: "784",
		}, "GB": {
			enName:  "United Kingdom of Great Britain and Northern Ireland (the)",
			frName:  "Royaume-Uni de Grande-Bretagne et d'Irlande du Nord (le)",
			alph2:   "GB",
			alph3:   "GBR",
			numCode: "826",
		}, "UM": {
			enName:  "United States Minor Outlying Islands (the)",
			frName:  "Îles mineures éloignées des États-Unis (les)",
			alph2:   "UM",
			alph3:   "UMI",
			numCode: "581",
		}, "US": {
			enName:  "United States of America (the)",
			frName:  "États-Unis d'Amérique (les)",
			alph2:   "US",
			alph3:   "USA",
			numCode: "840",
		}, "UY": {
			enName:  "Uruguay",
			frName:  "Uruguay (l')",
			alph2:   "UY",
			alph3:   "URY",
			numCode: "858",
		}, "UZ": {
			enName:  "Uzbekistan",
			frName:  "Ouzbékistan (l')",
			alph2:   "UZ",
			alph3:   "UZB",
			numCode: "860",
		}, "VU": {
			enName:  "Vanuatu",
			frName:  "Vanuatu (le)",
			alph2:   "VU",
			alph3:   "VUT",
			numCode: "548",
		}, "VE": {
			enName:  "Venezuela (Bolivarian Republic of)",
			frName:  "Venezuela (République bolivarienne du)",
			alph2:   "VE",
			alph3:   "VEN",
			numCode: "862",
		}, "VN": {
			enName:  "Viet Nam",
			frName:  "Viet Nam (le)",
			alph2:   "VN",
			alph3:   "VNM",
			numCode: "704",
		}, "VG": {
			enName:  "Virgin Islands (British)",
			frName:  "Vierges britanniques (les Îles)",
			alph2:   "VG",
			alph3:   "VGB",
			numCode: "092",
		}, "VI": {
			enName:  "Virgin Islands (U.S.)",
			frName:  "Vierges des États-Unis (les Îles)",
			alph2:   "VI",
			alph3:   "VIR",
			numCode: "850",
		}, "WF": {
			enName:  "Wallis and Futuna",
			frName:  "Wallis-et-Futuna",
			alph2:   "WF",
			alph3:   "WLF",
			numCode: "876",
		}, "EH": {
			enName:  "Western Sahara*",
			frName:  "Sahara occidental (le)*",
			alph2:   "EH",
			alph3:   "ESH",
			numCode: "732",
		}, "YE": {
			enName:  "Yemen",
			frName:  "Yémen (le)",
			alph2:   "YE",
			alph3:   "YEM",
			numCode: "887",
		}, "ZM": {
			enName:  "Zambia",
			frName:  "Zambie (la)",
			alph2:   "ZM",
			alph3:   "ZMB",
			numCode: "894",
		}, "ZW": {
			enName:  "Zimbabwe",
			frName:  "Zimbabwe (le)",
			alph2:   "ZW",
			alph3:   "ZWE",
			numCode: "716",
		},
	}
	_ = isocodes // TODO: delete in production
)
