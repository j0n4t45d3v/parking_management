CREATE TABLE IF NOT EXISTS vehicles (
  id SERIAL NOT NULL PRIMARY KEY,
  band VARCHAR(30) NOT NULL,
  model VARCHAR(30) NOT NULL,
  plate VARCHAR(7) NOT NULL,
  type CHAR(1) NOT NULL DEFAULT 'C' CHECK (type = 'C' OR type = 'M'),
  withdrawn CHAR(1) NOT NULL DEFAULT 'N' CHECK (withdrawn = 'Y' OR withdrawn = 'N')
)
