package bver

// A VersionInfo has basic information about a Bible version
type VersionInfo struct {
	ID        string
	Name      string
	Aliases   []string
	Copyright string
	Publisher string
}

// AvailableLanguages contains the codes for the languages that are available in the Versions variable
var AvailableLanguages = []string{"en"}

// Versions contains information about different Bible versions
var Versions = map[string][]VersionInfo{
	"en": {
		{
			ID:        "ACV",
			Name:      "A Conservative Version",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "AMP",
			Name:      "Amplified Bible",
			Aliases:   []string{"Amplified Bible Classic Edition", "AMPC", "Amplified Bible, Classic Edition"},
			Copyright: "Copyright 1954, 1958, 1962, 1964, 1965, 1987, 2015 Zondervan and the Lockman Foundation",
			Publisher: "Zondervan Publishing House",
		},
		{
			ID:        "ASV",
			Name:      "American Standard Version",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "DARBY",
			Name:      "Darby Translation",
			Aliases:   []string{"Darby Bible", "Darby Bible Translation", "DBY"},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "RHE",
			Name:      "Douay-Rheims Bible",
			Aliases:   []string{"Douay Rheims Bible", "Douay Rheims", "Douay-Rheims Catholic Bible", "DRBO", "DOUR", "Douay"},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "ESV",
			Name:      "English Standard Version",
			Aliases:   []string{},
			Copyright: "Copyright 2001 by Crossway Bibles, a publishing ministry of Good News Publishers",
			Publisher: "Crossway",
		},
		{
			ID:        "GNT",
			Name:      "Good News Translation",
			Aliases:   []string{"Good News Bible"},
			Copyright: "Copyright American Bible Society 1966, 1967, 1970, 1971, 1976, 1979",
			Publisher: "HarperCollins",
		},
		{
			ID:        "GW",
			Name:      "God's Word Translation",
			Aliases:   []string{"Gods Word Translation", "God's Word", "Gods Word"},
			Copyright: "Copyright 1995 by God's Word to the Nations",
			Publisher: "Baker Publishing Group",
		},
		{
			ID:        "HNV",
			Name:      "Hebrew Names Version",
			Aliases:   []string{"Hebrew Names"},
			Copyright: "",
			Publisher: "",
		},
		{
			ID:        "HCSB",
			Name:      "Holman Christian Study Bible",
			Aliases:   []string{"CSB", "HCS", "Holman", "Holman Christian Standard Bible"},
			Copyright: "Copyright 2004 Holman Bible Publishers",
			Publisher: "Holman Bible Publishers",
		},
		{
			ID:        "JUB2000",
			Name:      "Jubilee Bible 2000",
			Aliases:   []string{"English Jubilee 2000 Bible", "JUB", "Jubilee Bible"},
			Copyright: "Copyright 2000, 2001, 2010, 2013 by Russell M. Stendal",
			Publisher: "Russell M. Stendal",
		},
		{
			ID:        "KJV",
			Name:      "King James Version",
			Aliases:   []string{"Authorised Version", "Authorised King James Version", "Authorised (King James) Version", "Authorized Version", "Authorized King James Version", "Authorized (King James) Version", "King James Bible", "AV", "AKJV", "KJB"},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "KJ2000",
			Name:      "King James 2000",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "LEB",
			Name:      "Lexham English Bible",
			Aliases:   []string{},
			Copyright: "Copyright 2010 Logos Bible Software",
			Publisher: "Logos Bible Software",
		},
		{
			ID:        "MSG",
			Name:      "The Message Bible",
			Aliases:   []string{"The Message", "Message"},
			Copyright: "Copyright 2002 Eugene H. Peterson",
			Publisher: "Eugene H. Peterson",
		},
		{
			ID:        "NASB",
			Name:      "New American Standard Bible",
			Aliases:   []string{"NAS"},
			Copyright: "Copyright 1960, 1962, 1963, 1968, 1971, 1972, 1973, 1975, 1977, 1995 by The Lockman Foundation",
			Publisher: "The Lockman Foundation",
		},
		{
			ID:        "NCV",
			Name:      "New Century Version",
			Aliases:   []string{},
			Copyright: "unknown",
			Publisher: "Thomas Nelson",
		},
		{
			ID:        "NET",
			Name:      "New English Translation",
			Aliases:   []string{"NET Bible"},
			Copyright: "Copyright 2005 Biblical Studies Press, L.L.C.",
			Publisher: "Biblical Studies Press, L.L.C.",
		},
		{
			ID:        "NHEB",
			Name:      "New Heart English Bible",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "NIV",
			Name:      "New International Version",
			Aliases:   []string{},
			Copyright: "Copyright 1973, 1978, 1984, 2011 Biblica",
			Publisher: "Biblica (Worldwide), Zondervan (US)",
		},
		{
			ID:        "NKJV",
			Name:      "New King James Version",
			Aliases:   []string{},
			Copyright: "Copyright 1979, 1980, 1982 Thomas Nelson",
			Publisher: "Thomas Nelson",
		},
		{
			ID:        "NLT",
			Name:      "New Living Translation",
			Aliases:   []string{},
			Copyright: "Copyright 1996, 2004, 2015 by Tyndale House Foundation",
			Publisher: "Tyndale House Publishers",
		},
		{
			ID:        "NRSV",
			Name:      "New Revised Standard Version",
			Aliases:   []string{"New Revised Standard"},
			Copyright: "Copyright 1989 by the Division of Christian Education of the National Council of the Churches of Christ in the USA",
			Publisher: "National Council of Churches",
		},
		{
			ID:        "RSV",
			Name:      "Revised Standard Version",
			Aliases:   []string{},
			Copyright: "Copyright 1946, 1952, 1971 (Apocrypha 1957, 1977) by the Division of Christian Education of the National Council of the Churches of Christ in the USA",
			Publisher: "",
		},
		{
			ID:        "WEB",
			Name:      "World English Bible",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "WMB",
			Name:      "World Messianic Bible",
			Aliases:   []string{},
			Copyright: "",
			Publisher: "",
		},
		{
			ID:        "WYC",
			Name:      "Wycliffe",
			Aliases:   []string{},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "WBT",
			Name:      "The Webster Bible",
			Aliases:   []string{"Webster"},
			Copyright: "Public Domain",
			Publisher: "",
		},
		{
			ID:        "YLT",
			Name:      "Young's Literal Translation",
			Aliases:   []string{"Youngs Literal Translation"},
			Copyright: "Public Domain",
			Publisher: "",
		},
	},
}
