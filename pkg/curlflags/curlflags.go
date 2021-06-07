/*
Copyright © 2021 Mingfei Huang <himax1023@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package curlflags

// binaryCurlFlags are all curl flags that does not accept additional arguments
// acquired by `curl -h | grep -v '<'`
var binaryCurlFlags = []string{
	"--anyauth",
	"-a",
	"--append",
	"--basic",
	"--cert-status",
	"--compressed",
	"--compressed-ssh",
	"--create-dirs",
	"--crlf",
	"--digest",
	"-q",
	"--disable",
	"--disable-eprt",
	"--disable-epsv",
	"--disallow-username-in-url",
	"-f",
	"--fail",
	"--fail-early",
	"--false-start",
	"--ftp-create-dirs",
	"--ftp-pasv",
	"--ftp-pret",
	"--ftp-skip-pasv-ip",
	"--ftp-ssl-ccc",
	"--ftp-ssl-control",
	"-G",
	"--get",
	"-g",
	"--globoff",
	"--haproxy-protocol",
	"-I",
	"--head",
	"-h",
	"--help",
	"--http0.9",
	"-0",
	"--http1.0",
	"--http1.1",
	"--http2",
	"--http2-prior-knowledge",
	"--ignore-content-length",
	"-i",
	"--include",
	"-k",
	"--insecure",
	"-4",
	"--ipv4",
	"-6",
	"--ipv6",
	"-j",
	"--junk-session-cookies",
	"-l",
	"--list-only",
	"-L",
	"--location",
	"--location-trusted",
	"-M",
	"--manual",
	"--metalink",
	"--negotiate",
	"-n",
	"--netrc",
	"--netrc-optional",
	"-:",
	"--next",
	"--no-alpn",
	"-N",
	"--no-buffer",
	"--no-keepalive",
	"--no-npn",
	"--no-sessionid",
	"--ntlm",
	"--ntlm-wb",
	"--path-as-is",
	"--post301",
	"--post302",
	"--post303",
	"--preproxy",
	"-#",
	"--progress-bar",
	"-x",
	"--proxy",
	"--proxy-anyauth",
	"--proxy-basic",
	"--proxy-digest",
	"--proxy-insecure",
	"--proxy-negotiate",
	"--proxy-ntlm",
	"--proxy-ssl-allow-beast",
	"--proxy-tlsv1",
	"-p",
	"--proxytunnel",
	"-Q",
	"--quote",
	"--raw",
	"-J",
	"--remote-header-name",
	"-O",
	"--remote-name",
	"--remote-name-all",
	"-R",
	"--remote-time",
	"--request-target",
	"--retry-connrefused",
	"--sasl-ir",
	"-S",
	"--show-error",
	"-s",
	"--silent",
	"--socks5-basic",
	"--socks5-gssapi",
	"--socks5-gssapi-nec",
	"--ssl",
	"--ssl-allow-beast",
	"--ssl-no-revoke",
	"--ssl-reqd",
	"-2",
	"--sslv2",
	"-3",
	"--sslv3",
	"--stderr",
	"--styled-output",
	"--suppress-connect-headers",
	"--tcp-fastopen",
	"--tcp-nodelay",
	"--tftp-no-options",
	"--tlspassword",
	"-1",
	"--tlsv1",
	"--tlsv1.0",
	"--tlsv1.1",
	"--tlsv1.2",
	"--tlsv1.3",
	"--tr-encoding",
	"--trace-time",
	"-B",
	"--use-ascii",
	"-v",
	"--verbose",
	"-V",
	"--version",
	"--xattr",
}

// argumentCurlFlags are curl flags that require an additional argument
var argumentCurlFlags = []string{
	"--abstract-unix-socket",
	"--alt-svc",
	"--cacert",
	"--capath",
	"-E",
	"--cert",
	"--cert-type",
	"--ciphers",
	"-K",
	"--config",
	"--connect-timeout",
	"--connect-to",
	"-C",
	"--continue-at",
	"-b",
	"--cookie",
	"-c",
	"--cookie-jar",
	"--crlfile",
	"-d",
	"--data",
	"--data-ascii",
	"--data-binary",
	"--data-raw",
	"--data-urlencode",
	"--delegation",
	"--dns-interface",
	"--dns-ipv4-addr",
	"--dns-ipv6-addr",
	"--dns-servers",
	"--doh-url",
	"-D",
	"--dump-header",
	"--egd-file",
	"--engine",
	"--expect100-timeout",
	"-F",
	"--form",
	"--form-string",
	"--ftp-account",
	"--ftp-alternative-to-user",
	"--ftp-method",
	"-P",
	"--ftp-port",
	"--ftp-ssl-ccc-mode",
	"--happy-eyeballs-timeout-ms",
	"-H",
	"--header",
	"--hostpubmd5",
	"--interface",
	"--keepalive-time",
	"--key",
	"--key-type",
	"--krb",
	"--libcurl",
	"--limit-rate",
	"--local-port",
	"--login-options",
	"--mail-auth",
	"--mail-from",
	"--mail-rcpt",
	"--max-filesize",
	"--max-redirs",
	"-m",
	"--max-time",
	"--netrc-file",
	"--noproxy",
	"--oauth2-bearer",
	"-o",
	"--output",
	"--pass",
	"--pinnedpubkey",
	"--proto",
	"--proto-default",
	"--proto-redir",
	"--proxy-cacert",
	"--proxy-capath",
	"--proxy-cert",
	"--proxy-cert-type",
	"--proxy-ciphers",
	"--proxy-crlfile",
	"--proxy-header",
	"--proxy-key",
	"--proxy-key-type",
	"--proxy-pass",
	"--proxy-pinnedpubkey",
	"--proxy-service-name",
	"--proxy-tls13-ciphers",
	"--proxy-tlsauthtype",
	"--proxy-tlspassword",
	"--proxy-tlsuser",
	"-U",
	"--proxy-user",
	"--proxy1.0",
	"--pubkey",
	"--random-file",
	"-r",
	"--range",
	"-e",
	"--referer",
	"-X",
	"--request",
	"--resolve",
	"--retry",
	"--retry-delay",
	"--retry-max-time",
	"--service-name",
	"--socks4",
	"--socks4a",
	"--socks5",
	"--socks5-gssapi-service",
	"--socks5-hostname",
	"-Y",
	"--speed-limit",
	"-y",
	"--speed-time",
	"-t",
	"--telnet-option",
	"--tftp-blksize",
	"-z",
	"--time-cond",
	"--tls-max",
	"--tls13-ciphers",
	"--tlsauthtype",
	"--tlsuser",
	"--trace",
	"--trace-ascii",
	"--unix-socket",
	"-T",
	"--upload-file",
	"--url",
	"-u",
	"--user",
	"-A",
	"--user-agent",
	"-w",
	"--write-out",
}

var binaryCurlFlagsSet = (func() map[string]struct{} {
	set := map[string]struct{}{}
	for _, flag := range binaryCurlFlags {
		set[flag] = struct{}{}
	}
	return set
})()

var argumentCurlFlagsSet = (func() map[string]struct{} {
	set := map[string]struct{}{}
	for _, flag := range argumentCurlFlags {
		set[flag] = struct{}{}
	}
	return set
})()

// IsBinary returns whether given flag is a binary curl flag.
func IsBinary(flag string) bool {
	_, ok := binaryCurlFlagsSet[flag]
	return ok
}

// IsArgument returns whether given flag is a curl flag that requires argument.
func IsArgument(flag string) bool {
	_, ok := argumentCurlFlagsSet[flag]
	return ok
}

// IndexOfURLArg finds index the `url` argument in args.
// url argument is defined as the argument that is not an optional flag and is not an argument of an option flag.
// It returns -1 if url arg is not found.
func IndexOfURLArg(args []string) int {
	idx := 0
	for idx < len(args) {
		if IsBinary(args[idx]) {
			idx += 1
			continue
		}
		if IsArgument(args[idx]) {
			idx += 2
			continue
		}
		return idx
	}
	return -1
}
