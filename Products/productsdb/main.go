package productsdb

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

/*
CREATE TABLE productlink(
xid char(12) NOT NULL,
productName varchar(255) NOT NULL,
barboraProductName varchar(255) NOT NULL,
barboraProductURL  varchar(255) NOT NULL,
barboraProductPrice decimal(5,2),
rimiProductName     varchar(255) NOT NULL,
rimiProductUrl      varchar(255) NOT NULL,
rimiProductPrice    decimal(5,2)
)

INSERT INTO productlink (xid, productName, barboraProductName, barboraProductURL, rimiProductName, rimiProductURL) VALUES
('b50vl5e54p1000fo3gh0', 'Baltmaize Zeltene 350 g', 'Baltmaize ZELTENE 350g', 'https://www.barbora.lv/produkti/baltmaize-zeltene-350-g',
'Baltmaize Zeltene saldā sagriezta 350g', 'https://www.rimi.lv/e-veikals/lv/produkti/maize-un-konditoreja/maize/baltmaize/baltmaize-zeltene-salda-sagriezta-350g/p/285134aq2q')

ALTER TABLE productlink ADD PRIMARY KEY (xid);
*/

type ProductLink struct {
	xid                string
	productName        string
	barboraProductName string
	barboraProductURL  string
	barboraProductPrice float32
	rimiProductName     string
	rimiProductUrl      string
	rimiProductPrice    float32
}
