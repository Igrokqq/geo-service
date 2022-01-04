package country

import (
	"github.com/vfilipovsky/geo-service/internal/modules/continent"
	"github.com/vfilipovsky/geo-service/internal/modules/region"
	"github.com/vfilipovsky/geo-service/internal/modules/timezone"
)

type Country struct {
	ID uint `gorm:"primaryKey"`

	Languages  []Language          `gorm:"many2many:country_languages"`
	Currencies []Currency          `gorm:"many2many:country_currencies"`
	Timezones  []timezone.Timezone `gorm:"many2many:country_timezones"`
	Neighbours []Country           `gorm:"many2many:country_neighbours"`

	Translations []Translation      `gorm:"foreignKey:CountryId"`
	Maps         []Map              `gorm:"foreignKey:CountryId"`
	AltSpelling  []AlternativeSpell `gorm:"foreignKey:CountryId"`

	PostalCode PostalCode `gorm:"not null"`
	Name       CName      `gorm:"not null"`
	DialId     DialId     `gorm:"not null"`
	Position   Position   `gorm:"not null"`

	Continent   continent.Continent
	ContinentId uint `gorm:"not null"`

	Region   region.Region
	RegionId uint `gorm:"not null"`

	Iso2Code    string  `gorm:"unique;not null;size:2;"`
	Iso3Code    string  `gorm:"unique;not null;size:3;"`
	Area        float64 `gorm:"default:0.0"`
	Population  int     `gorm:"default:0"`
	UnicodeFlag string
}

type DialId struct {
	ID uint `gorm:"primaryKey"`

	Root      string
	Suffixes  []DialIdSuffix `gorm:"foreignKey:DialId"`
	CountryId uint           `gorm:"not null"`
}

type DialIdSuffix struct {
	ID uint `gorm:"primaryKey"`

	Suffix string `gorm:"not null;size:50"`
	DialId uint   `gorm:"not null"`
}

type AlternativeSpell struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:150"`
	CountryId uint   `gorm:"not null"`
}

type CName struct {
	ID        uint   `gorm:"primaryKey"`
	Common    string `gorm:"not null;size:150"`
	Official  string `gorm:"not null;size:150"`
	CountryId uint   `gorm:"not null"`
}

type Currency struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique;not null;size:150"`
	Symbol string `gorm:"unique;not null;size:10"`
	Code   string `gorm:"unique;not null;size:10"`
}

type Language struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null;size:25"`
	Name string `gorm:"unique;not null;size:150"`
}

type Translation struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"not null;size:25"`
	Official  string `gorm:"not null;size:150"`
	Common    string `gorm:"size:150"`
	CountryId uint   `gorm:"not null"`
}

type Position struct {
	ID        uint    `gorm:"primaryKey"`
	Longitude float64 `gorm:"default:0.0"`
	Latitude  float64 `gorm:"default:0.0"`
	CountryId uint    `gorm:"not null"`
}

type Map struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Url       string `gorm:"size:255"`
	CountryId uint   `gorm:"not null"`
}

type PostalCode struct {
	ID        uint   `gorm:"primaryKey"`
	Format    string `gorm:"size:150;not null"`
	Regex     string `gorm:"size:150;not null"`
	CountryId uint   `gorm:"not null"`
}
