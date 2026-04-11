# BlockIt

A command-line tool for managing domain whitelists and blacklists based on TLD data from IANA.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/JakubKyhos/Blockit.git
   cd Blockit
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o blockit main.go
   ```

4. Set up environment variables:
   - Create a `.env` file in the root directory.
   - Add your PostgreSQL database URL: `DB_URL=your_postgresql_connection_string`

5. Ensure PostgreSQL is running and accessible.

6. Run gencert.sh:
 - make sure created files are inside goproxy folder.
 - add the certificate to your certificate manager
   - MacOS: `/Applications/Utilities/Keychain\ Access.app.`
   - Windows: `certmgr.msc`

7. Setup proxy server in OS settings:
 - **Windows:** settings -> network and internet -> proxy server
 - **MacOS:** settings -> network -> click on desired network and details -> proxy (I don't own MacOS so this is untested)
 - for both: manual setting -> use proxy server: *enabled*, IP adress: `http://localhost`, port: `8080` -> *save*
   - **if you are not using BlockIt, make sure to disable the proxy server**

## Usage

Run `./main.go` to configure the application:
- The application provides an interactive command-line interface. Enter commands to manage domains.

Run `./server/main.go` to start the proxy server:

### Available Commands

- `help`: Display help information
- `setup`: Initialize the database with TLD data from IANA
- `add whitelist <domain>`: Add a domain to the whitelist
- `add whitelisttemp <domain>`: Add a temporary domain to the whitelist
- `add blacklist <domain>`: Add a domain to the blacklist
- `delete whitelist <domain>`: Remove a domain from the whitelist
- `delete whitelist temp`: Remove all expired temporary domains from the whitelist
- `delete blacklist <domain>`: Remove a domain from the blacklist
- `list whitelist`: Display all whitelisted domains
- `list blacklist`: Display all blacklisted domains
- `list tld`: Display all TLD domains
- `reset tld`: Reset the TLD database
- `reset whitelist`: Reset the whitelist database
- `reset blacklist`: Reset the blacklist database
- `blockstate <true/false> <tld/global>`: Change the block state for TLDs or globally
- `quit`: Exit the application

### Future plans
 add an easy way to populate blacklist with gambling, adult, malware ridden domains.

### Credits
 **elazarl:** for goproxy library `https://github.com/elazarl/goproxy`
 - *license:* `https://github.com/elazarl/goproxy/blob/master/LICENSE`