# BlockIt

A command-line tool for managing domain whitelists and blacklists based on TLD data from IANA utilizing elazarl's goproxy library.

## Motivation
- About 3.4 billion phishing emails are sent daily.
- Over 90% of all cyberattacks begin with a phishing email.
- 94% of malware is delivered via email.

- How many of us suffer from gambling social media or adult content addiction.

- This project seeks to solve these problems by allowing you to block domains based on just the TLD or the whole domain name.



## Quick Start

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
 - **MacOS:** settings -> network -> click on desired network and details -> proxy
 - for both: manual setting -> use proxy server: *enabled*, IP address: `http://localhost`, port: `8080` -> *save*
   - **if you are not using BlockIt, make sure to disable the proxy server**

## Usage

Run `go run .` to configure the application:
- The application provides an interactive command-line interface. Enter commands to manage domains.

Run `go run ./server/main.go` to start the proxy server:

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

## Contributing

Follow the **quick start** section.

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

### Future plans
 add an easy way to populate blacklist with gambling, adult, malware ridden domains.

### Credits
 **elazarl:** for goproxy library `https://github.com/elazarl/goproxy`
 - *license:* `https://github.com/elazarl/goproxy/blob/master/LICENSE`