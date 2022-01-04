package country

import (
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

	PostalCode PostalCode
	Name       CName
	DialId     DialId
	Position   Position

	Region   region.Region
	RegionId uint

	Iso2Code    string `gorm:"unique;not null"`
	Iso3Code    string `gorm:"unique;not null"`
	Area        float64
	UnicodeFlag string
	Population  int
}

type DialId struct {
	ID uint `gorm:"primaryKey"`

	Root      string
	Suffixes  []DialIdSuffix `gorm:"foreignKey:DialId"`
	CountryId uint
}

type DialIdSuffix struct {
	ID uint `gorm:"primaryKey"`

	Suffix string
	DialId uint
}

type AlternativeSpell struct {
	ID uint `gorm:"primaryKey"`

	Name      string
	CountryId uint
}

type CName struct {
	ID uint `gorm:"primaryKey"`

	Common    string
	Official  string
	CountryId uint
}

type Currency struct {
	ID uint `gorm:"primaryKey"`

	Name   string `gorm:"unique;not null"`
	Symbol string `gorm:"unique;not null"`
	Code   string `gorm:"unique;not null"`
}

type Language struct {
	ID uint `gorm:"primaryKey"`

	Code string `gorm:"unique;not null"`
	Name string `gorm:"unique;not null"`
}

type Translation struct {
	ID uint `gorm:"primaryKey"`

	Code      string
	Official  string
	Common    string
	CountryId uint
}

type Position struct {
	ID uint `gorm:"primaryKey"`

	Longitude float64
	Latitude  float64
	CountryId uint
}

type Map struct {
	ID uint `gorm:"primaryKey"`

	Name      string
	Url       string
	CountryId uint
}

type PostalCode struct {
	ID uint `gorm:"primaryKey"`

	Format    string
	Regex     string
	CountryId uint
}
