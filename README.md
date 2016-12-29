## postalcodes

Given a 6 character Canadian postal code and a radius in kilometers, postalcodes will ouput all canadian postal codes within the given radius.

### System requirements

postalcodes requires the Go programming language. Please visit the Go project website for installation instructions: https://golang.org/doc/install

### Installation

```
$ git clone https://github.com/ccnixon/postalcodes
$ cd postalcodes
$ go install
```

### Usage

postalcodes takes a 6 character Canadian postal code and a radius in kilometers as arguments. The arguments can be passed in as flags as such:

`$ postalcodes -p=<POSTAL_CODE> -r=<RADIUS>`

The default arguments are M4W2L4 and 3 respectively. Running postalcodes without any arguments would thus be equivalent to:

`$ postalcodes -p=M4W2L4 -r=3`

Run `$ postalcodes -h` for help.

### Data Source

The program relies on this csv:

[CSV of Canadian Postal Codes](https://fusiontables.google.com/DataSource?docid=1H_cl-oyeG4FDwqJUTeI_aGKmmkJdPDzRNccp96M&hl=en_US&pli=1)





